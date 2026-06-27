package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Downloader  Downloader  `json:"downloader"`
	FileChecker FileChecker `json:"file_checker"`
	RenameFiles RenameFiles `json:"rename_files"`
}

type Downloader struct {
	CookiesFilePath   string `json:"cookies_file_path"`
	IgOutputPath      string `json:"ig_output_path"`
	YtOutputPath      string `json:"yt_output_path"`
	YtDownloadQuality string `json:"yt_download_quality"`
	YtDownloadFormat  string `json:"yt_download_format,omitempty"`
}

type FileChecker struct {
	CheckFolderPath string `json:"check_folder_path"`
}

type RenameFiles struct {
	InputPath  string `json:"input_path"`
	OutputPath string `json:"output_path"`
}

const configFileName = "config.json"

func LoadConfig() (*Config, error) {
	var cfg Config

	path, err := FindConfigFilepath()
	if err != nil {
		return nil, err
	}

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(configFile, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FindConfigFilepath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		configPath := filepath.Join(dir, configFileName)

		if fileExists(configPath) {
			return configPath, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", os.ErrNotExist
}
