# Storganizer

> *Fast, native duplicate finder for media-aware cleanup.*

`Storganizer` is a high-performance toolkit for identifying duplicate and near-duplicate media without relying solely on exact file hashes. It catches true duplicates, downsized copies, re-exports, and near-identical shots using perceptual fingerprints.

## What it does

- **Exact duplicate detection** - full-file SHA-1 matching
- **Perceptual matching** - pHash and dHash analysis to find visually similar images
- **Resolution-aware checks** - identifies same content at different pixel dimensions
- **Re-export detection** - catches visually identical files with different sizes or encoding
- **Near-identical shots** - groups similar photos from the same burst or scene
- **Interactive reporting** - generates self-contained HTML dashboards for browser-based audits

## Project structure

```text
.
├── .github/      # GitHub Automation (CI/CD report pipelines)
├── main.go       # Native Go scanning engine and CLI logic
├── index.html    # Unified static dashboard template
├── Makefile      # Cross-platform build orchestration
├── go.mod        # Go module definition
└── README.md
```

## Stack

| Layer | Tooling |
|---|---|
| Core engine | Go (Golang) |
| Hashing | `goimagehash` (pHash, dHash) |
| Web UI | HTML5, CSS3, Vanilla JS |
| CI/CD | GitHub Actions |
| Video | `ffmpeg` + `ffprobe` (optional) |

## Requirements

### Build
- **Go 1.20+** is required to build from source.

### Runtime
Storganizer runs as a single native binary with no external dependencies:
- **Linux** (amd64)
- **Windows** (amd64)

*Note: `ffmpeg` is recommended for advanced video fingerprinting support.*

## Running

```bash
# Basic scan of the current directory
./storganizer

# Scan a specific media folder
./storganizer -path ~/Pictures

# Generate a web-based audit report
./storganizer -path ~/Pictures -html

# Output raw JSON for pipeline integration
./storganizer -path ~/Pictures -json
```

Windows:
```powershell
# Scan a disk path
.\storganizer.exe -path D:\Photos -html
```

## `storganizer` Flags

- `-path`       - Directory to scan (default: `.`)
- `-threshold`  - Perceptual distance limit (default: `10`, lower is stricter)
- `-json`       - Output results as JSON
- `-html`       - Bake results into `index.html` report

## Typical workflows

### Local Media Audit
1. Build the tool: `go build -o storganizer main.go`
2. Run with HTML output: `./storganizer -path ./my-media -html`
3. Open `index.html` to review visual matches side-by-side.

### GitHub-Native Reporting
1. Push your media files to a GitHub repository.
2. The included GitHub Action builds the Go CLI and runs a scan.
3. A self-contained report is automatically deployed to your **GitHub Pages**.

## Tuning Thresholds

| Threshold | Sensitivity | Typical Use Case |
|---|---|---|
| **5-8** | Stricter | True duplicates and simple resizes |
| **10-12** | Standard | Re-exports and edited copies |
| **15-20** | Loose | Near-identical shots in a burst |

## Notes

- **Performance:** Scanning is concurrent; large folders are processed using all available CPU cores.
- **Safety:** Storganizer is a **reporting tool**; it identifies matches but does not delete files automatically.
- **Video:** Video matching uses heuristic sampling. For forensic-level video analysis, consider dedicated tools.
