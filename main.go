package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	videoRegex, _ := regexp.Compile(".*\\.(mkv|mk3d|mp4|m4v|mov|qt|ogv|ogx|avi|wmv|webm|flv)")
	subRegex, _ := regexp.Compile(".*\\.(ass|ssa|srt|sub|gsub|aqt|ttxt|pjs|psb|rt|smi|ssf|sbv|usf|vtt)")

	directory := "."
	args := os.Args
	if len(args) > 1 {
		directory = args[1]
	}
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	videos := []os.DirEntry{}
	subtitles := []os.DirEntry{}

	for _, file := range files {
		if videoRegex.MatchString(file.Name()) {
			videos = append(videos, file)
			continue
		}
		if subRegex.MatchString(file.Name()) {
			subtitles = append(subtitles, file)
		}
	}
	if len(videos) != len(subtitles) {
		log.Fatalf("Found %d videos but %d subtitles, exiting without changing anything", len(videos), len(subtitles))
	}
	for i := range videos {
		vidName := videos[i].Name()
		vidBaseName := strings.TrimSuffix(vidName, filepath.Ext(vidName))
		subName := subtitles[i].Name()
		subExt := filepath.Ext(subName)
		os.Rename(subName, fmt.Sprintf("%s%s", vidBaseName, subExt))
		log.Printf("Moved %s to %s%s", subName, vidBaseName, subExt)
	}
}
