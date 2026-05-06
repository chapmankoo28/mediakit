package downloader

import (
	"fmt"
	"mediakit/internal/config"
	"mediakit/internal/reader"
	"os"
	"os/exec"
)

func findGallerydlPath() (string, error) {
	path, err := exec.LookPath("gallery-dl")
	if err != nil {
		if _, err := exec.LookPath("./gallery-dl"); err == nil {
			return "./gallery-dl", nil
		}
		return "", os.ErrNotExist
	}
	return path, nil
}

func IgDownloader() (string, error) {
	path, err := findGallerydlPath()
	if err != nil {
		return "", fmt.Errorf("gallery-dl not found")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("error reading config file: %v", err)
	}

	cookies := cfg.Downloader.CookiesFilePath
	if _, err := os.Stat(cookies); os.IsNotExist(err) {
		return "", fmt.Errorf("no such file: %s", cookies)
	}

	outputPath := cfg.Downloader.IgOutputPath
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		if err := os.MkdirAll(outputPath, 0766); err != nil {
			return "", fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	url := reader.ReadInput("Input the Instagram post URL: ")

	args := []string{"--cookies", cookies, "-D", outputPath, url}
	fmt.Println("args:", args)
	if !reader.ConfirmAction("Is this OK?", true) {
		return "Aborted.", nil
	}

	cmd := exec.Command(path, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error during Instagram download: %v", err)
	}

	return "Done.", nil
}
