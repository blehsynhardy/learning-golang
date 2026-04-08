package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// downloader performs a synchronous download of a single file from a URL into a destination directory.
// It cleans up the partial file on any error.
func downloader(url, destDir string) (int64, error) {
	filename := filepath.Base(url)
	filePath := filepath.Join(destDir, filename)

	output, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer output.Close()

	fmt.Println("Downloading", url)
	start := time.Now()

	// Execute the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		_ = os.Remove(filePath)
		return 0, err
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(filePath)
		return 0, fmt.Errorf("bad status: %s", resp.Status)
	}

	// Stream the response body directly to the file
	size, err := io.Copy(output, resp.Body)
	if err != nil {
		_ = os.Remove(filePath) // clean up partial file
		return 0, err
	}

	fmt.Printf("downloaded %s (%d bytes) in %v\n", filename, size, time.Since(start))
	return size, nil
}

// sequentialDownloader downloads files one at a time.
func sequentialDownloader(urls []string, destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	start := time.Now()
	var errs []error

	for _, url := range urls {
		if _, err := downloader(url, destDir); err != nil {
			fmt.Printf("error downloading %s: %v\n", url, err)
			errs = append(errs, err)
		}
	}

	fmt.Printf("all sequential downloads completed in %v\n", time.Since(start))

	if len(errs) > 0 {
		return fmt.Errorf("%d download(s) failed: %v", len(errs), errs)
	}
	return nil
}

// Result holds the outcome of a single concurrent download.
type Result struct {
	URL      string
	FileName string
	Size     int64 // actual byte count, not string length
	Duration time.Duration
	Error    error
}

// concurrentDownloader downloads files in parallel, capped at maxConcurrent.
func concurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	// results channel collects outcomes from goroutines
	results := make(chan Result)
	// limiter channel acts as a semaphore to control concurrency
	limiter := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup
	
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			// acquire slot
			limiter <- struct{}{}
			defer func() { <-limiter }()

			start := time.Now()
			filename := filepath.Base(url)
			filePath := filepath.Join(destDir, filename)

			output, err := os.Create(filePath)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer output.Close()

			resp, err := http.Get(url)
			if err != nil {
				_ = os.Remove(filePath)
				results <- Result{URL: url, Error: err}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				_ = os.Remove(filePath)
				results <- Result{URL: url, Error: fmt.Errorf("bad status: %s", resp.Status)}
				return
			}

			size, err := io.Copy(output, resp.Body)
			if err != nil {
				_ = os.Remove(filePath) // FIX: clean up partial file on copy error
				results <- Result{URL: url, Error: err}
				return
			}

			results <- Result{
				URL:      url,
				FileName: filename,
				Size:     size, // FIX: store actual int64 byte count
				Duration: time.Since(start),
			}
		}(url)
	}

	// Close results channel once all goroutines finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errs []error // FIX: renamed from `error` to avoid shadowing the built-in
	start := time.Now()

	// Process results as they arrive from the channel
	for res := range results {
		if res.Error != nil {
			fmt.Printf("error downloading %s: %v\n", res.URL, res.Error)
			errs = append(errs, res.Error)
			continue
		}

		totalSize += res.Size // FIX: accumulate actual byte count
		fmt.Printf("downloaded %s (%d bytes) in %v\n", res.FileName, res.Size, res.Duration)
	}

	fmt.Printf("all downloads completed in %v — total size: %d bytes\n", time.Since(start), totalSize)

	if len(errs) > 0 {
		return fmt.Errorf("%d download(s) failed: %v", len(errs), errs)
	}
	return nil
}

func main() {
	urls := []string{
		"https://cdn.pixabay.com/photo/2024/09/15/13/03/cow-9049115_1280.jpg",
		"https://cdn.pixabay.com/photo/2026/03/23/07/20/salofoto-dog-10187835_1280.jpg",
		"https://cdn.pixabay.com/photo/2025/04/23/20/38/cupcake-9553838_1280.jpg",
	}

	// Start concurrent download with a limit of 3 simultaneous connections
	if err := concurrentDownloader(urls, "./downloads", 3); err != nil {
		log.Println("completed with errors:", err)
		return
	}

	log.Println("done")
}