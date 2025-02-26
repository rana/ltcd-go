package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func simplifyPath71b(pth string) string {
	return ""
}

func TestSimplifyPath71b(t *testing.T) {
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
			got := simplifyPath71b(tt.pth)
			if got != tt.want {
				t.Errorf("simplifyPath(%q) = %q, want %q", tt.pth, got, tt.want)
			}
		})
	}
}
