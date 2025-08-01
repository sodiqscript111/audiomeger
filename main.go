package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	searchPath := filepath.Join("audio", "*.mp3")
	fmt.Println("Looking for mp3 files in:", searchPath)

	files, err := filepath.Glob(searchPath)
	if err != nil {
		fmt.Println("Error searching for files:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No mp3 files found")
		return
	}

	fmt.Println("Found the following mp3 files:")
	for _, file := range files {
		fmt.Println(" -", file)
	}

	output := "output.mp3"

	args := []string{"-y"}
	for _, file := range files {
		args = append(args, "-i", file)
	}
	args = append(args, "-filter_complex")

	concatFilter := ""
	for i := range files {
		concatFilter += fmt.Sprintf("[%d:0]", i)
	}
	concatFilter += fmt.Sprintf("concat=n=%d:v=0:a=1[out]", len(files))

	args = append(args, concatFilter, "-map", "[out]", output)

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Merging mp3 files with ffmpeg...")
	err = cmd.Run()
	if err != nil {
		fmt.Println("ffmpeg failed:", err)
		return
	}

	fmt.Println("Merge complete: output.mp3")
}
