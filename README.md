# Storganizer

> *Fast, native duplicate finder for media-aware cleanup.*

Storganizer is a high-performance toolkit for identifying duplicate and near-duplicate media without relying solely on exact file hashes. It catches true duplicates, downsized copies, re-exports, and near-identical shots using perceptual fingerprints.

## What it does

- **Exact duplicate detection** - full-file SHA-1 matching
- **Perceptual matching** - pHash and dHash analysis to find visually similar images
- **Resolution-aware checks** - identifies same content at different pixel dimensions
- **Re-export detection** - catches visually identical files with different sizes or encoding
- **Near-identical shots** - groups similar photos from the same burst or scene
- **Interactive reporting** - generates self-contained HTML dashboards for browser-based audits
- **Video support** - optional ffmpeg integration

## Requirements

### Linux/macOS

**Build:**
- Go 1.20+ (`go version`)

**Runtime:**
- None - completely self-contained
- Optional: `ffmpeg` / `ffprobe` (for video fingerprinting)

### Windows

**Build:**
- Go 1.20+ (download from [go.dev](https://go.dev/dl/))

**Runtime:**
- None - completely self-contained
- Optional: ffmpeg (for video)

## Installation

### Linux/macOS

```bash
# Clone the repository
git clone https://github.com/Ivoryy06/storganizer.git ~/storganizer

# Build the binary
cd ~/storganizer
go build -o storganizer .

# Install globally (optional)
sudo cp storganizer /usr/local/bin/
```

### Windows

1. **Download Go** from [go.dev/dl](https://go.dev/dl/) and install

2. **Open Command Prompt** (Win+X → Terminal)

3. **Clone and build:**
```cmd
git clone https://github.com/Ivoryy06/storganizer.git
cd storganizer
go build -o storganizer.exe .
```

4. **Run from any location:**
```cmd
move storganizer.exe C:\Windows\System32\
```

## Running

### Linux/macOS

```bash
# Scan current directory
./storganizer

# Scan specific folder
./storganizer -path ~/Pictures

# Generate HTML report
./storganizer -path ~/Pictures -html

# Output JSON
./storganizer -path ~/Pictures -json results.json

# Strict matching (fewer false positives)
./storganizer -path ~/Pictures -threshold 5
```

### Windows

Open **Command Prompt**:

```cmd
storganizer.exe

storganizer.exe -path "C:\Users\YourName\Pictures"

storganizer.exe -path "C:\Users\YourName\Pictures" -html
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-path` | Directory to scan | Current directory |
| `-threshold` | Perceptual distance (lower = stricter) | 10 |
| `-json` | Output JSON file | No output |
| `-html` | Output HTML report | No output |

## Tuning Thresholds

| Threshold | Sensitivity | Typical Use Case |
|---|---|---|
| **5-8** | Stricter | True duplicates and simple resizes |
| **10-12** | Standard | Re-exports and edited copies |
| **15-20** | Loose | Near-identical shots in a burst |

## Typical Workflows

### Local Media Audit
1. Build the tool: `go build -o storganizer .`
2. Run with HTML output: `./storganizer -path ./my-media -html`
3. Open `index.html` to review visual matches side-by-side.

### GitHub-Native Reporting
1. Push your media files to a GitHub repository.
2. The included GitHub Action builds the Go CLI and runs a scan.
3. A self-contained report is automatically deployed to your **GitHub Pages**.

## Troubleshooting

- **No images found?** - Make sure path is correct and contains supported formats (JPEG, PNG, WEBP, GIF, BMP)
- **Too many groups?** - Increase threshold: `-threshold 15`
- **Not finding similar enough images?** - Lower threshold: `-threshold 5`
- **Video not supported?** - Install ffmpeg: `sudo pacman -S ffmpeg` (Linux) or download from ffmpeg.org

## Notes

- **Performance:** Scanning is concurrent; large folders are processed using all available CPU cores.
- **Safety:** Storganizer is a **reporting tool**; it identifies matches but does not delete files automatically.
- **Video:** Video matching uses heuristic sampling. For forensic-level video analysis, consider dedicated tools.

## Project Structure

```
storganizer/
├── main.go       # Native Go scanning engine and CLI logic
├── index.html   # Unified static dashboard template
├── Makefile     # Cross-platform build orchestration
├── go.mod       # Go module definition
├── README.md    # This file
└── LICENSE      # MIT License
```

## License

MIT