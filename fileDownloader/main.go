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

func DownloadFile(url, destDir string) error {
	filename := filepath.Base(url)
	filePath := filepath.Join(destDir, filename)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	fmt.Println("Downloading", url)
	start := time.Now()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		_, _ = os.Readlink(filePath)
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Download %s took %s\n", filename, time.Since(start))
	return nil
}

func SequentialDownloader(urls []string, destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	start := time.Now()
	for _, url := range urls {
		if err := DownloadFile(url, destDir); err != nil {
			fmt.Printf("Error downloading %s: %s\n", url, err)
			continue
		}
	}

	fmt.Printf("Download took %s\n", time.Since(start))
	return nil
}

type Result struct {
	URL      string
	Filename string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	results := make(chan Result)

	var wg sync.WaitGroup

	limiter := make(chan struct{}, maxConcurrent)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}
			defer func() { <-limiter }()

			start := time.Now()
			filename := filepath.Base(url)
			filePath := filepath.Join(destDir, filename)

			out, err := os.Create(filePath)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer func(out *os.File) {
				err := out.Close()
				if err != nil {

				}
			}(out)
			resp, err := http.Get(url)
			if err != nil {
				results <- Result{URL: url, Error: err}
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(resp.Body)
			if resp.StatusCode != http.StatusOK {
				results <- Result{URL: url, Error: fmt.Errorf("bad status: %s", resp.Status)}
				return
			}
			size, err := io.Copy(out, resp.Body)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			timeSinceStart := time.Since(start)
			results <- Result{URL: url, Size: size, Duration: timeSinceStart}
			return
		}(url)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	start := time.Now()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error downloading %s: %s\n", result.URL, result.Error)
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("Downloaded %s (%d bytes) in %s\n", result.Filename, result.Size, result.Duration)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("All downloads completed in %s\n, Total: %d bytes\n", elapsed, totalSize)
	if len(errors) > 0 {
		return fmt.Errorf("%d errors occurred", len(errors))
	}
	return nil
}

func main() {

	urls := []string{"https://download.samplelib.com/mp3/sample-3s.mp3",
		"https://download.samplelib.com/mp4/sample-5s.mp4"}

	err := ConcurrentDownloader(urls, "./", 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Downloaded", urls)

}
