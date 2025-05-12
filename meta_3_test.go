package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

// Reading from STDIN and receiving two parameters
// which was host and port that I need to reach and print status.
//
// What if this input is the file with 1tb
// (here interaction with large inputs & optimization)

func TestMeta3aStdin(t *testing.T) {
	host, port := "google.com", "80"
	target := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.DialTimeout("tcp", target, time.Second)
	if err != nil {
		fmt.Println("unreachable", target, err)
	} else {
		fmt.Println("reachable", target)
	}
	defer conn.Close()
}

// Read from a large 1TB input file.
// For all host port pairs:
// * Connect to host and port
// * Print status
func TestMeta3aLargeFile(t *testing.T) {

}

// Reading from STDIN and receiving two parameters
// which was host and port that I need to reach and print status.
//
// What if this input is the file with 1tb
// (here interaction with large inputs & optimization)

const (
	maxWorkers = 100             // Limit concurrent connections
	timeout    = 5 * time.Second // Connection timeout
	batchSize  = 10000           // Process inputs in batches
)

// checkHost attempts to connect to host:port and returns the status
func checkHost(host string, port string) string {
	target := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		return fmt.Sprintf("%s:%s - Unreachable: %v", host, port, err)
	}

	defer conn.Close()
	return fmt.Sprintf("%s:%s - Reachable", host, port)
}

func TestMeta3(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)

	// For extremely large files, we might need to increase the buffer size
	const maxScanTokenSize = 1024 * 1024 // 1MB
	buf := make([]byte, maxScanTokenSize)
	scanner.Buffer(buf, maxScanTokenSize)

	// Create a worker pool using semaphore pattern
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxWorkers)
	results := make(chan string, maxWorkers) // Buffer for results

	// Start a goroutine to print results to avoid blocking workers
	done := make(chan struct{})
	go func() {
		for result := range results {
			fmt.Println(result)
		}
		close(done)
	}()

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) < 2 {
			continue // Skip invalid lines
		}

		host := strings.TrimSpace(parts[0])
		port := strings.TrimSpace(parts[1])

		lineCount++

		// Throttle with semaphore
		wg.Add(1)
		semaphore <- struct{}{} // Acquire semaphore

		go func(h, p string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release semaphore

			result := checkHost(h, p)
			results <- result
		}(host, port)

		// Process in batches to avoid memory consumption
		if lineCount%batchSize == 0 {
			wg.Wait() // Wait for current batch to complete
		}
	}

	// Wait for all remaining workers to complete
	wg.Wait()
	close(results)
	<-done // Wait for all results to be printed

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
