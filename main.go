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

	keys := []string{"1", "2", "3"}
	options := map[string]string{
		keys[0]: "Video Downloader",
		keys[1]: "Instagram Downloader",
		keys[2]: "Renamer",
	}

	fmt.Printf("\nDownload and Format Helper")
	fmt.Printf(" (using config: %s)\n", filepath.Base(cfgPath))
	for _, key := range keys {
		fmt.Printf("  %s) %s\n", key, options[key])
	}
	fmt.Println()

	var text string
	for true {
		text = reader.ReadInput("Please select an option (1..3): ")
		if options[text] == "" {
			fmt.Println("Invalid option.")
			continue
		} else {
			break
		}
	}

	var msg string
	var err2 error
	switch text {
	case keys[0]:
		msg, err2 = downloader.YtdlpDownloader()
	case keys[1]:
		msg, err2 = downloader.IgDownloader()
	case keys[2]:
		msg, err2 = renamer.RenameFiles()
	default:
		fmt.Println("\nInvalid option.")
	}

	if err2 != nil {
		fmt.Fprintf(os.Stderr, "\033[31mError:\033[0m %v\n", err2)
	} else {
		fmt.Printf("\n%v\n", msg)
	}
}
