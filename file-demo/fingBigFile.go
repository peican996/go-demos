package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	_ "strconv"
)

type FileEntry struct {
	Path string
	Size int64
}

type BySize []FileEntry

func (b BySize) Len() int           { return len(b) }
func (b BySize) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BySize) Less(i, j int) bool { return b[i].Size > b[j].Size }

func formatFileSize(size int64) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)
	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d bytes", size)
	}
}

func main() {
	root := "/"
	sizeThreshold := int64(1024 * 1024 * 100) // 100 MB

	var files []FileEntry

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}

		if info.IsDir() {
			// Skip directories
			return nil
		}

		if info.Size() > sizeThreshold {
			files = append(files, FileEntry{Path: path, Size: info.Size()})
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %q: %v", root, err)
	}

	// Sort files by size in descending order
	sort.Sort(BySize(files))

	// Create a new CSV file
	file, err := os.Create("large_files.csv")
	if err != nil {
		log.Fatalf("Error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"File", "Size"}
	err = writer.Write(header)
	if err != nil {
		log.Fatalf("Error writing CSV header: %v", err)
	}

	// Write sorted file entries to CSV
	for _, entry := range files {
		row := []string{entry.Path, formatFileSize(entry.Size)}
		err = writer.Write(row)
		if err != nil {
			log.Printf("Error writing CSV row for file %q: %v\n", entry.Path, err)
		}
	}

	fmt.Println("Large files have been written to large_files.csv")
}
