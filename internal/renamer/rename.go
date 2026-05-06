package renamer

import (
	"crypto/sha256"
	"fmt"
	"io"
	"mediakit/internal/config"
	"mediakit/internal/reader"
	"os"
	"path/filepath"
	"slices"
)

func RenameFiles() (string, error) {
	renewMode := reader.ConfirmAction("Enable renew mode?", false)

	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("error reading config file: %v", err)
	}

	inputDir := cfg.RenameFiles.InputPath
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		return "", fmt.Errorf("no such directory: %s", inputDir)
	}

	outputDir := cfg.RenameFiles.OutputPath
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0766); err != nil {
			return "", fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	files, err := os.ReadDir(inputDir)
	if err != nil {
		return "", fmt.Errorf("failed to read input directory: %v", err)
	}

	var fileToRename []string
	for _, file := range files {
		if !file.IsDir() {
			fileToRename = append(fileToRename, file.Name())
		}
	}

	slices.Sort(fileToRename)

	if len(fileToRename) == 0 {
		return "", fmt.Errorf("no files found in directory '%v'", inputDir)
	}

	fmt.Print("Files to be renamed: ")
	if renewMode {
		fmt.Print("(renew mode: each file will have a new hex code)")
	}
	fmt.Println()
	for i, f := range fileToRename {
		fmt.Printf("%3d. %s\n", i+1, f)
	}

	if !reader.ConfirmAction("Rename these files?", true) {
		return "Aborted.", nil
	}

	fmt.Println()

	fileHashes := make(map[string]string)
	var seriesHash string

	if renewMode {
		for _, fName := range fileToRename {
			file, err := os.Open(filepath.Join(inputDir, fName))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to open file %v: %v", fName, err)
				continue
			}

			hasher := sha256.New()
			if _, err := io.Copy(hasher, file); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to hash file %v: %v", fName, err)
			}
			file.Close()
			fileHashes[fName] = fmt.Sprintf("%x", hasher.Sum(nil))[:6]
		}
	} else {
		hasher := sha256.New()
		for _, fName := range fileToRename {
			file, err := os.Open(filepath.Join(inputDir, fName))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to open file %v: %v", fName, err)
				continue
			}

			fileHasher := sha256.New()
			writer := io.MultiWriter(hasher, fileHasher)
			if _, err := io.Copy(writer, file); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to hash file %v: %v", fName, err)
			}
			file.Close()

			fileHashes[fName] = fmt.Sprintf("%x", fileHasher.Sum(nil))[:6]
		}
		seriesHash = fmt.Sprintf("%x", hasher.Sum(nil))[:6]
	}

	count := 1
	for _, fName := range fileToRename {
		file := filepath.Join(inputDir, fName)

		fileHash := fileHashes[fName]

		suffix := filepath.Ext(fName)

		fmt.Printf("Org name: %v\n", file)

		var newFileName string
		if renewMode {
			newFileName = fmt.Sprintf("%s_1%s", fileHash, suffix)
		} else {
			newFileName = fmt.Sprintf("%s_%d_%s%s", seriesHash, count, fileHash, suffix)
		}

		fmt.Printf("New name: %v\n", newFileName)

		src, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open source file: %v", err)
			continue
		}

		dst, err := os.Create(filepath.Join(outputDir, newFileName))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create destination file: %v", err)
			src.Close()
			continue
		}

		_, err = io.Copy(dst, src)
		src.Close()
		if closeErr := dst.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to copy file %v: %v", fName, err)
			continue
		}

		fmt.Println()

		if !renewMode {
			count++
		}
	}

	return "Done.", nil
}
