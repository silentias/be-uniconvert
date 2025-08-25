package services

import (
	"be-uniconvert/pkg"
	"fmt"
	"mime/multipart"
	"os/exec"
	"path/filepath"
	"strings"
)

func ConvertAudio(file *multipart.FileHeader, ext string) (string, error) {
	filename := filepath.Base(file.Filename)
	inputPath := filepath.Join("uploads", filename)
	nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
	hash, err := pkg.RandomString(10)
	if err != nil {
		return "", err
	}
	outputPath := filepath.Join("uploads", nameWithoutExt+hash+ext)

	cmd := exec.Command("ffmpeg", "-i", inputPath, outputPath)
	stderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Ошибка FFmpeg:", err)
		fmt.Println("Подробности:", string(stderr))
		return "", err
	}

	return outputPath, nil
}
