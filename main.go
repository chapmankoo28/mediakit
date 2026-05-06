package main

import (
	"fmt"
	"mediakit/internal/config"
	"mediakit/internal/downloader"
	"mediakit/internal/reader"
	"mediakit/internal/renamer"
	"os"
	"path/filepath"
)

func main() {
	cfgPath, err := config.FindConfigFilepath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading config: %v\n", err)
		return
	}

	fmt.Printf("\nDownload and Format Helper")
	fmt.Printf(" (using config: %s)\n", filepath.Base(cfgPath))
	fmt.Println("  1) Video Downloader")
	fmt.Println("  2) Instagram Downloader")
	fmt.Println("  3) Renamer")
	fmt.Println()

	text := reader.ReadInput("Please select an option (1..4): ")
	var msg string
	var err2 error
	switch text {
	case "1":
		msg, err2 = downloader.YtdlpDownloader()
	case "2":
		msg, err2 = downloader.IgDownloader()
	case "3":
		msg, err2 = renamer.RenameFiles()
	default:
		fmt.Println("\nInvalid option.")
	}

	if err2 != nil {
		fmt.Fprintf(os.Stderr, "\033[31mError\033[0m: %v\n", err2)
	} else {
		fmt.Printf("\n%v\n", msg)
	}
}
