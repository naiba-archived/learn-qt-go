package main

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func downloadFromURL(url string) (string, error) {
	now := time.Now()
	filename := downloadDir + strconv.Itoa(now.Year()) + "-" + now.Month().String() + "-" + strconv.Itoa(now.Day()) + "-" + strconv.Itoa(rand.Int()) + url[strings.LastIndex(url, "."):]
	output, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func getCurPath() string {
	file, _ := exec.LookPath(os.Args[0])

	path, _ := filepath.Abs(file)

	return filepath.Dir(path)
}
