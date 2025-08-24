package pkg

import (
	"path/filepath"
	"strings"
)

var allowedExtensions = map[string]bool{
	"mp3":  true,
	"wav":  true,
	"flac": true,
	"aac":  true,
}

func IsAllowedExtension(filename string) bool {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	return allowedExtensions[ext]
}