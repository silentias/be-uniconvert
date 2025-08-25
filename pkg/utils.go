package pkg

import (
	"crypto/rand"
	"math/big"
	"path/filepath"
	"strings"
)

var AllowedExtensions = map[string]bool{
	"mp3":  true,
	"wav":  true,
	"flac": true,
	"aac":  true,
}

func IsAllowedExtension(filename string) bool {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	return AllowedExtensions[ext]
}


func RandomString(n int) (string, error) {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, n)
    for i := 0; i < n; i++ {
        num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
        if err != nil {
            return "", err
        }
        result[i] = letters[num.Int64()]
    }
    return string(result), nil
}
