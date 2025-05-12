package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testing"
)

// Read in file meta_4.csv.
// Count occurrences of email address.
// Print count.
func TestMeta4(t *testing.T) {
	file, _ := os.Open("meta_4.csv")
	defer file.Close()

	emailRegexp := regexp.MustCompile(`[a-zA-Z0-9._]+@[a-zA-Z0-9]+\.[a-zA-Z]+`)
	emailMap := make(map[string]int)
	emailTotal := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		emails := emailRegexp.FindAllString(line, -1)
		for _, email := range emails {
			emailMap[email]++
			emailTotal++
		}
	}

	fmt.Println("Total", emailTotal)
	for key, val := range emailMap {
		fmt.Println(val, key)
	}
}
