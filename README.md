# Storganizer Workspace

> *Fast, native duplicate finder for media-aware cleanup.*

This workspace contains **Storganizer**, a high-concurrency media toolkit designed for finding duplicates and near-duplicates across messy media folders.

## What it does

- **Concurrent Scanning** - Native Go routines for high-speed file crawling
- **Perceptual Analysis** - Visual fingerprinting (pHash/dHash) to find non-exact matches
- **Automated Reporting** - GitHub-native web dashboard generation
- **Cross-Platform CLI** - Single-binary execution for Linux and Windows

## Project structure

```text
/
├── Storganizer/      # Main application core
│   ├── main.go       # Go scanning engine
│   ├── index.html    # Dashboard template
│   └── Makefile      # Build orchestration
├── .github/          # GitHub Automation
│   └── workflows/    # CI/CD report pipelines
└── README.md
```

## Stack

| Layer | Tooling |
|---|---|
| Engine | Go (Golang) |
| Web | HTML5 / Vanilla JS |
| Automation | GitHub Actions |
| Deployment | GitHub Pages |

## Requirements

- **Go 1.20+** (for local builds)
- **Linux** or **Windows** (for binary execution)
- **ffmpeg** (optional, for video support)

## Running

```bash
# Build and run locally
cd Storganizer
go build -o storganizer main.go
./storganizer -path /path/to/media -html
```

## Typical workflows

### Automated GitHub Audit
1. Push media to your repository.
2. The GitHub Action builds the Go engine.
3. A scan is performed and results are baked into `index.html`.
4. The report is published to your GitHub Pages site automatically.

### Manual Local Scan
1. Download or build the `storganizer` binary.
2. Run it against a target folder with the `-html` flag.
3. Open the resulting `index.html` to review matches.

## Notes

- Storganizer is optimized for large libraries where byte-level SHA-1 hashes are insufficient.
- The web dashboard is static and self-contained, making it ideal for hosting on GitHub or sharing as a single file.
