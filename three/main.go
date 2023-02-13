package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{"https://www.google.com", "https://www.yahoo.com", "https://www.varzesh3.com"}
	responses, err := GetResponses(urls, 1*time.Second)
	if err != nil {
		fmt.Printf("Error getting responses: %s\n", err)
		return
	}

	fmt.Println(responses)
}

func GetResponses(urls []string, timeout time.Duration) ([]string, error) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	responses := make([]string, 0, len(urls))
	errChan := make(chan error, len(urls))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				errChan <- fmt.Errorf("Error creating request for %s: %s\n", url, err)
				return
			}

			client := http.DefaultClient
			resp, err := client.Do(req)
			if err != nil {
				errChan <- fmt.Errorf("Error making request to %s: %s\n", url, err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Response received from %s\n", url)
			responses = append(responses, url)
		}(url)
	}

	wg.Wait()
	select {
	case err := <-errChan:
		return nil, err
	default:
		return responses, nil
	}
}
