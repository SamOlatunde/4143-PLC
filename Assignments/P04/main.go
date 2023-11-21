package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	// TODO: Implement sequential download logic

	for index, value := range urls {

		downloadImage(value, fmt.Sprintf("image%d.jpg", index), nil, nil)
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	// TODO: Implement concurrent download logic
	// Declare a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a buffered channel to collect errors from goroutines
	errCh := make(chan error, len(urls))

	// Iterate over the URLs and launch a goroutine for each download
	for index, value := range urls {
		// Increment the WaitGroup counter to indicate a new goroutine
		wg.Add(1)

		// Launch a goroutine to download the image
		go downloadImage(value, fmt.Sprintf("image%d.jpg", index), &wg, errCh)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the error channel after all downloads are completed
	close(errCh)

	// Process errors from goroutines
	for err := range errCh {
		// Print the error if it is not nil
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",

		"https://images.unsplash.com/photo-1590068561151-2aa0b87cda13?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1610035974356-3e9f2c818347?q=80&w=1935&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://cdn.stocksnap.io/img-thumbs/960w/people-man_2W0L5IENXQ.jpg",
		"https://images.pexels.com/photos/751696/pexels-photo-751696.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://cdn.pixabay.com/photo/2017/05/11/12/35/girl-2304038_1280.jpg"}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
/*func downloadImage(url, filename string) error {
	// TODO: Implement download logic
	imageUrl := url
	// Create an HTTP GET request
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return err
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return nil
	}

	// Create a new file to save the image
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return err
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return err
	}

	fmt.Println("Image downloaded and saved as ", filename)
	return nil
}*/

func downloadImage(url, filename string, wg *sync.WaitGroup, ch chan<- error) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	imageURL := url

	// Create an HTTP GET request
	response, err := http.Get(imageURL)
	if err != nil {
		if ch != nil {
			ch <- fmt.Errorf("Error making the request for %s: %v", imageURL, err)
		}
		return
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		if ch != nil {
			ch <- fmt.Errorf("Error downloading %s: Status code %d", imageURL, response.StatusCode)
		}
		return
	}

	// Create a new file to save the image
	outputFile, err := os.Create(filename)
	if err != nil {
		if ch != nil {
			ch <- fmt.Errorf("Error creating the file %s: %v", filename, err)
		}
		return
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		if ch != nil {
			ch <- fmt.Errorf("Error saving the image %s: %v", filename, err)
		}
		return
	}

	fmt.Printf("Image downloaded and saved as %s\n", filename)
	if ch != nil {
		ch <- nil
	}
}
