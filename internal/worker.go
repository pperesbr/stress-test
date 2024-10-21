package internal

import (
	"net/http"
	"sync"
)

func Worker(wg *sync.WaitGroup, mu *sync.Mutex, url string, report *Report) {
	defer wg.Done()

	resp, err := http.Get(url)

	if err != nil {
		mu.Lock()
		report.StatusCode500++
		mu.Unlock()
		return
	}

	defer resp.Body.Close()

	statusCode := resp.StatusCode

	if statusCode == 200 {
		mu.Lock()
		report.StatusCode200++
		mu.Unlock()
	}

	if statusCode == 404 {
		mu.Lock()
		report.StatusCode404++
		mu.Unlock()
	}

}
