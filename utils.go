package main

// Importing required packages
import (
	"strings"
	"io/ioutil"
	"log"
	"fmt"
)

// Takes filename as an argument
// return boolean based on
// if the file exists or not
func isFileValid(filename string)bool {
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Takes filename as an argument
// return boolean based on
// if the file format is mov or mp4 (currently)
func isFormatValid(filename string)bool {
	avalfmts := []string{"mov", "mp4"}
	flext := strings.Split(filename, ".")
	if len(flext) == 1 {
		return false
	}
	ext := flext[1]
	return contains(avalfmts, strings.ToLower(ext))
}

// Takes filename as an argument
// return formatted directory name
// i.e filename_frames
func generateOutputTemplate(filename string)string {
	fl := strings.Split(filename, ".")
	outputDir := strings.ToLower(fl[0])
	return fmt.Sprintf("%s_frames", outputDir)
}

// helper function to check if
// elem x exists in the given []string
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
