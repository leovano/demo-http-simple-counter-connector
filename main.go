package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("COUNTER_SERVER_ENDPOINT")
	if addr == "" {
		log.Fatal("COUNTER_SERVER_ENDPOINT environment variable not set")
	}
	url := fmt.Sprintf("http://%s/", addr)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("server returned non-OK status: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading response: %v", err)
	}
}
