# Download videos and images

## Setup

1. Install ffmpeg if not already installed
2. Install yt-dlp with pipx:

```bash
pipx install yt-dlp
```

3. Install gallery-dl with pipx:

```bash
pipx install gallery-dl
```

## Usage

1. Download video from YouTube

```bash
yt-dlp \
-f "bestvideo[height=1080]+bestaudio/best" \
--remux-video mp4 \
-o "downloads/%(title)s.%(ext)s" \
"YOUTUBE VIDEO URL"
```

2. Download video from Instagram

```bash
yt-dlp \
--cookies-from-browser firefox \
-f "bestvideo/best+bestaudio/best" \
--remux-video mp4 \
-o "downloads/%(title)s.%(ext)s" \
"INSTAGRAM POST URL"
```

3. Download image from Instagram

```bash
gallery-dl \
--cookies instagram_cookies.txt \
-D "downloads" \
"INSTAGRAM POST URL"
```
