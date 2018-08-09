package main

import (
	"flag"
	"log"
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
		"encoding/json"
	"time"
)

type CreationTime struct {
	CreationTime time.Time `json:"creation_time"`
}

type Tags struct {
	Tags *CreationTime `json:"tags"`
	FileName string `json:"filename"`
	Duration string `json:"duration"`
}

type FileInfo struct {
	Format *Tags `json:"format"`
	FrameCount int `json:"frame_count"`
}

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
	extractFileInfo(input, output)
}

func extractFileInfo(filename string, outDir string) {
	args := []string{"-v", "quiet", filename, "-print_format", "json", "-show_format"}
	cmd, _ := exec.Command("ffprobe", args...).Output()
	log.Print(fmt.Sprintf("Extracting infromation from %s", filename))
	var i FileInfo
	// Getting video creation time
	json.Unmarshal(cmd, &i)
	// Getting frame count
	files, _ := ioutil.ReadDir(outDir)
	i.FrameCount = len(files)
	data, _ := json.Marshal(i)
	ioutil.WriteFile(outDir+"/info.json", data, 0644)
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