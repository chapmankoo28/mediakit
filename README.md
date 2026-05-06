# mediakit

A simple tool for downloading and managing media files.

## Setup

1. Install `deno` first for handling YouTube JS challenges.

2. Copy `config.example.yaml` to `config.yaml` and update it:

```yaml
downloader:
    # required for Instagram downloads
    cookies_file_path: instagram_cookies.txt
    ig_output_path: IG
    yt_output_path: YT
    # Available values: "720" (Default), "1080", "best"
    yt_download_quality: "720"

    # Available values: "mp3", "m4a", "opus", "mp4" (Default), "webm"
    yt_download_format: mp4

file_checker:
    check_folder_path: output

rename_files:
    input_path: rename
    output_path: output
```

## Usage

### Build the project

```bash
# macOS or Linux
go build -o mediakit

# Windows
go build -o mediakit.exe
```

```powershell
# Build for linux/amd64 on Windows
$env:GOOS="linux"; $env:GOARCH="amd64"; go build; Remove-Item Env:GOOS; Remove-Item Env:GOARCH
```

### Run it

```bash
# MacOS or Linux
./mediakit

# Windows
.\mediakit.exe
```
