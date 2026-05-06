package downloader

import (
	"fmt"
	"mediakit/internal/config"
	"mediakit/internal/reader"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

func findYtdlp() (string, error) {
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		if _, err := exec.LookPath("./yt-dlp"); err == nil {
			return "./yt-dlp", nil
		}
		return "", os.ErrNotExist
	}
	return path, nil
}

func YtdlpDownloader() (string, error) {
	path, err := findYtdlp()
	if err != nil {
		return "", fmt.Errorf("yt-dlp not found")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("error reading config file: %v", err)
	}

	outPath := cfg.Downloader.YtOutputPath
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		if err := os.MkdirAll(outPath, 0755); err != nil {
			return "", fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	url := reader.ReadInput("Input the video URL: ")

	var args []string

	fileFormat := cfg.Downloader.YtDownloadFormat
	switch fileFormat {
	case "mp3", "m4a", "opus":
		args = append(args, "--embed-thumbnail", "-x", "--audio-format", fileFormat)
	case "mp4", "webm":
		quality := "720"
		validQualities := []string{"720", "1080", "best"}

		if slices.Contains(validQualities, quality) {
			quality = cfg.Downloader.YtDownloadQuality
			fmt.Println("Using quality:", quality)
		} else {
			fmt.Println("Unknown option, using default:", quality, "(valid options: 720, 1080, best)")
		}

		if fileFormat != "webm" {
			args = append(args, "--embed-thumbnail")
		}

		args = append(args, "--remux-video", fileFormat)

		if quality != "best" {
			args = append(args, "-f", fmt.Sprintf("bestvideo[height<=%v]+bestaudio/best", quality))
		}
	default:
		return "", fmt.Errorf("unknown download format: %s (valid options: mp3, m4a, opus, mp4, webm)", fileFormat)
	}

	args = append(args, "-o", filepath.Join(outPath, "%(title)s.%(ext)s"), url)
	fmt.Println("args:", args)

	if !reader.ConfirmAction("Is this OK?", true) {
		return "Aborted.", nil
	}

	cmd := exec.Command(path, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error during YouTube download: %v", err)
	}

	return "Done.", nil
}
