package goembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestStringEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed Go-Language.png
var logo []byte

func TestByteEmbed(t *testing.T) {
	err := ioutil.WriteFile("new-logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed videoplayback.m4a

var video []byte

func TestVideoEmbed(t *testing.T) {
	err := ioutil.WriteFile("new-video.m4a", video, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
