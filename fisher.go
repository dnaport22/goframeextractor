package main

import (
	"flag"
	"log"
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
)

func extractFramesFromVideo(input string, output string) {
	args := []string{"-i", input, "-r", "1/1", output+"/frame_%06d.png"}
	cmd := exec.Command("ffmpeg", args...)
	log.Print(fmt.Sprintf("Extracting frames from %s", input))
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Finished extracting frames")
	log.Print(fmt.Sprintf("Extracted frames are in %s directory", output))
}


func main() {
	inputFile := flag.String("input", "", "Input video file")
	flag.Parse()

	filename := *inputFile

	if !isFormatValid(filename) {
		log.Fatal("Invalid file format\n\tAcceptable formats: mp4, mov")
	}

	if isFormatValid(filename) {
		outDir := generateOutputTemplate(filename)
		_, err := ioutil.ReadDir(outDir)
		if err != nil {
			os.Mkdir(outDir, 0777)
		}
		os.RemoveAll(outDir)
		os.Mkdir(outDir, 0777)
		extractFramesFromVideo(filename, outDir)
	}
}