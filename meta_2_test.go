package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"text/tabwriter"
)

// Read data from an API call that gives structured data in JSON format.
// Arrange that data in the requested format.
// Use https://jsonplaceholder.typicode.com/posts
// Post: userid, id, title, body
func TestMeta2_0(t *testing.T) {
	type Post struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}
	const url = "https://jsonplaceholder.typicode.com/posts"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}
}

func TestMeta2_1(t *testing.T) {
	// Post represents the structure of each post from the JSONPlaceholder API
	type Post struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	// Define the API URL
	apiURL := "https://jsonplaceholder.typicode.com/posts"

	// Make the HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check if the response status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: API returned status code %d\n", resp.StatusCode)
		os.Exit(1)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		os.Exit(1)
	}

	// Parse the JSON response into a slice of Post structs
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Print the number of posts fetched
	fmt.Printf("Fetched %d posts from the API\n\n", len(posts))

	// Create a tabwriter for formatted output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)

	// Print header
	fmt.Fprintln(w, "ID\tUserID\tTitle\tBody Preview")
	fmt.Fprintln(w, "---\t------\t-----\t------------")

	// Print data in tabular format
	for _, post := range posts {
		// Get a preview of the body (first 30 characters)
		bodyPreview := post.Body
		if len(bodyPreview) > 30 {
			bodyPreview = bodyPreview[:30] + "..."
		}

		fmt.Fprintf(w, "%d\t%d\t%s\t%s\n", post.ID, post.UserID, post.Title, bodyPreview)
	}

	// Flush the tabwriter output
	w.Flush()
}
