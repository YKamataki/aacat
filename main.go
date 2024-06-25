package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/u2takey/ffmpeg-go"
)

func main() {
	//
	dir, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	var aacFilePaths []string
	var numberOfFile int
	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".aac") {
			aacFilePaths = append(aacFilePaths, file.Name())
		}
	}
	numberOfFile = len(aacFilePaths)

	// Make stream
	var streams []*ffmpeg_go.Stream
	for _, aacFilePath := range aacFilePaths {
		streams = append(streams, ffmpeg_go.Input(aacFilePath))
	}
	stream := ffmpeg_go.Concat(streams)

	// concat
	filterArg := fmt.Sprintf("concat=n=%d:v=0:a=1", numberOfFile)
	stream = stream.Audio().Filter("amix", ffmpeg_go.Args{filterArg}, ffmpeg_go.KwArgs{})

	// Make output
	outputFileName := fmt.Sprintf("%d-%d-%d.aac", time.Now().Year(), time.Now().Month(), time.Now().Day())
	stream.Output(outputFileName)

	// Run
	stream.Run()
}
