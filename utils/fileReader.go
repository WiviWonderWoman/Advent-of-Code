package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// findProjectRoot finds the project root by looking for the input directory
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree to find the input directory
	for {
		inputPath := filepath.Join(dir, "input")
		if info, err := os.Stat(inputPath); err == nil && info.IsDir() {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory
			break
		}
		dir = parent
	}

	// Fallback: return current directory
	return os.Getwd()
}

// ReadInput reads the input file for a given day
// Returns an array of lines from the input file
func ReadInput(day string) ([]string, error) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return nil, err
	}

	// Extract day number from "day01" format
	dayNumber := strings.TrimPrefix(day, "day")
	filePath := filepath.Join(projectRoot, "input", "day"+dayNumber+".txt")
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	return lines, nil
}

// ReadInputAsString reads the input file as a single string
func ReadInputAsString(day string) (string, error) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return "", err
	}

	// Extract day number from "day01" format
	dayNumber := strings.TrimPrefix(day, "day")
	filePath := filepath.Join(projectRoot, "input", "day"+dayNumber+".txt")
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}
