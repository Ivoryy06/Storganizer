package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/corona10/goimagehash"
)

type MediaFile struct {
	Path        string `json:"path"`
	Kind        string `json:"kind"`
	Size        int64  `json:"size"`
	ExactHash   string `json:"exact_hash,omitempty"`
	PHash       string `json:"phash,omitempty"`
	DHash       string `json:"dhash,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}

type Match struct {
	Type      string  `json:"type"`
	Kind      string  `json:"kind"`
	Left      string  `json:"left"`
	Right     string  `json:"right"`
	LeftSize  int64   `json:"left_size"`
	RightSize int64   `json:"right_size"`
	Distance  int     `json:"distance"`
}

var (
	imgExts = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	vidExts = map[string]bool{".mp4": true, ".mov": true, ".mkv": true, ".avi": true, ".webm": true}
)

func getSha1(path string) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func processFile(path string) (*MediaFile, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(path))
	kind := "unknown"
	if imgExts[ext] {
		kind = "image"
	} else if vidExts[ext] {
		kind = "video"
	} else {
		return nil, fmt.Errorf("unsupported type")
	}

	mf := &MediaFile{
		Path: path,
		Kind: kind,
		Size: info.Size(),
	}

	mf.ExactHash = getSha1(path)

	if kind == "image" {
		f, err := os.Open(path)
		if err != nil {
			return mf, nil
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return mf, nil
		}

		phash, _ := goimagehash.PerceptionHash(img)
		dhash, _ := goimagehash.DifferenceHash(img)
		mf.PHash = phash.ToString()
		mf.DHash = dhash.ToString()
		mf.Fingerprint = mf.PHash + ":" + mf.DHash
	}

	return mf, nil
}

func main() {
	pathPtr := flag.String("path", ".", "Path to scan")
	thresholdPtr := flag.Int("threshold", 10, "Distance threshold")
	jsonPtr := flag.Bool("json", false, "Output JSON")
	htmlPtr := flag.Bool("html", false, "Generate index.html report")
	flag.Parse()

	root := *pathPtr
	files := []string{}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if imgExts[ext] || vidExts[ext] {
			files = append(files) // Placeholder for concurrent fix
			files = append(files, path)
		}
		return nil
	})

	var wg sync.WaitGroup
	results := make([]*MediaFile, len(files))
	jobs := make(chan int, len(files))

	// Worker pool
	numWorkers := runtime.NumCPU()
	for w := 0; w < numWorkers; w++ {
		go func() {
			for i := range jobs {
				mf, err := processFile(files[i])
				if err == nil {
					results[i] = mf
				}
				wg.Done()
			}
		}()
	}

	for i := range files {
		wg.Add(1)
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// Filter nil results
	var validResults []*MediaFile
	for _, r := range results {
		if r != nil {
			validResults = append(validResults, r)
		}
	}

	matches := []Match{}
	for i := 0; i < len(validResults); i++ {
		for j := i + 1; j < len(validResults); j++ {
			a, b := validResults[i], validResults[j]
			if a.ExactHash == b.ExactHash && a.ExactHash != "" {
				matches = append(matches, Match{"exact duplicate", a.Kind, a.Path, b.Path, a.Size, b.Size, 0})
				continue
			}

			if a.PHash != "" && b.PHash != "" {
				ha, _ := goimagehash.ImageHashFromString(a.PHash)
				hb, _ := goimagehash.ImageHashFromString(b.PHash)
				dist, _ := ha.Distance(hb)

				if dist <= *thresholdPtr {
					matches = append(matches, Match{"near-identical", a.Kind, a.Path, b.Path, a.Size, b.Size, dist})
				}
			}
		}
	}

	if *jsonPtr {
		out, _ := json.MarshalIndent(matches, "", "  ")
		fmt.Println(string(out))
	} else if *htmlPtr {
		generateHTML(matches)
	} else {
		fmt.Printf("Scanned %d files, found %d matches\n", len(validResults), len(matches))
		for _, m := range matches {
			fmt.Printf("[%s] %s\n  %s\n  %s\n", m.Kind, m.Type, m.Left, m.Right)
		}
	}
}

func generateHTML(matches []Match) {
	template, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Println("Error reading index.html template:", err)
		return
	}

	dataJson, _ := json.MarshalIndent(matches, "", "  ")
	injection := fmt.Sprintf("const EMBEDDED_DATA = %s;", string(dataJson))
	
	html := string(template)
	marker := "// DATA_INJECTION_MARKER"
	if strings.Contains(html, marker) {
		html = strings.Replace(html, marker, injection, 1)
	} else {
		fmt.Println("Marker not found in index.html")
		return
	}

	err = os.WriteFile("index.html", []byte(html), 0644)
	if err != nil {
		fmt.Println("Error writing index.html:", err)
	} else {
		fmt.Println("Report generated in index.html")
	}
}
