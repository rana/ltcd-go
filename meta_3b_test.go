package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// Reading from STDIN and receiving two parameters
// which was host and port that I need to reach and print status.
//
// What if this input is the file with 1tb
// (here interaction with large inputs & optimization)

func TestMeta3bStdin(t *testing.T) {
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
