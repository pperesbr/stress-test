package main

import (
	"flag"
	"fmt"
	"stress-test/internal"
	"sync"
	"time"
)

func main() {

	url := flag.String("url", "", "url for test")
	numRequests := flag.Int("requests", 1, "number of requests")
	concurrency := flag.Int("concurrency", 1, "number of gorotuines")

	flag.Parse()

	if *url == "" {
		panic("please, type one URL")
	}

	pool := make(chan struct{}, *concurrency)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var report internal.Report
	start := time.Now()

	for i := 1; i <= *numRequests; i++ {
		wg.Add(1)

		pool <- struct{}{}
		go func(id int) {
			defer func() { <-pool }()
			internal.Worker(&wg, &mu, *url, &report)
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Total de requests: %d\n", *numRequests)
	fmt.Printf("Total de requests simuntaneas: %d\n", *concurrency)
	fmt.Printf("Tempo total de execucao: %s\n", duration)
	fmt.Printf("Total de requests com status code 200: %d\n", report.StatusCode200)
	fmt.Printf("Total de requests com status code 404: %d\n", report.StatusCode404)
	fmt.Printf("Total de requests com status code 500: %d\n", report.StatusCode500)
}
