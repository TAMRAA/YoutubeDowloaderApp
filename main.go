package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <YouTube Video URL>")
		return
	}

	videoURL := os.Args[1]

	cmd := exec.Command("yt-dlp", videoURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error downloading video: %v", err)
	}

	fmt.Println("Video downloaded successfully!")
}