package main

import (
	"strings"
	"testing"
)

// Time complexity: O(n), n is the length of the string.
// Space complexity: O(n), up to n characters stored on a stack and string builder.
// https://claude.ai/chat/ce0d47f6-4722-47ac-aad7-995ce30cf5f3
func simplifyPath71(pth string) string {
	// Simplify Path
	// Given a string pth.
	// String pth is an Unix-style absolute file path.
	// Transform pth into a simplified form.
	// Return the simplified form.
	// Conditions:
	//	* Single period (.) is current directory.
	//	* Double period (..) is parent directory.
	//	* Multiple slashes (///) are treat as single slash.
	//	* Three or more periods is a valid directory name.
	// Simplified Form:
	//	* Starts with single slash (/)
	//	* Directories separated by single slash.
	//	* Does not end with slash unless root directory.
	//	* Does not contain single or double period.

	// Split and filter string.
	var stk []string
	for _, dir := range strings.Split(pth, "/") {
		// Filter empty string and current directory.
		if dir == "" || dir == "." {
			continue
		}

		// Navigate parent directory.
		if dir == ".." {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
			continue
		}

		// Append valid directory.
		stk = append(stk, dir)
	}

	// Handle root directory case.
	if len(stk) == 0 {
		return "/"
	}

	// Build simplfied path.
	var bld strings.Builder
	for _, dir := range stk {
		bld.WriteString("/")
		bld.WriteString(dir)
	}

	return bld.String()
}

func TestSimplifyPath71(t *testing.T) {
	tests := []struct {
		name string
		pth  string
		want string
	}{
		{
			name: "basic path",
			pth:  "/home/",
			want: "/home",
		},
		{
			name: "multiple slashes",
			pth:  "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "parent directory navigation",
			pth:  "/home/user/Documents/../Pictures",
			want: "/home/user/Pictures",
		},
		{
			name: "beyond root",
			pth:  "/../",
			want: "/",
		},
		{
			name: "valid periods as directory name",
			pth:  "/.../a/../b/c/../d/./",
			want: "/.../b/d",
		},
		{
			name: "complex navigation",
			pth:  "/a/./b/../../c/",
			want: "/c",
		},
		{
			name: "root path",
			pth:  "/",
			want: "/",
		},
		{
			name: "relative current directory",
			pth:  "/./",
			want: "/",
		},
		{
			name: "multiple parent directories",
			pth:  "/a/b/c/../../d",
			want: "/a/d",
		},
		{
			name: "many dots as valid directory",
			pth:  "/...../foo",
			want: "/...../foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplifyPath71(tt.pth)
			if got != tt.want {
				t.Errorf("simplifyPath(%q) = %q, want %q", tt.pth, got, tt.want)
			}
		})
	}
}
