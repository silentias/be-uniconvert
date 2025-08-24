package services

import (
	"be-uniconvert/config"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(file *multipart.FileHeader) error {
	maxSize, err := config.GetSize()
	if err != nil {
		return err
	}
	if file.Size > maxSize {
		return errors.New("file too big")
	}

	filename := filepath.Base(file.Filename)
	dst := filepath.Join("uploads", filename)

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
			return err
		}

	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = out.ReadFrom(src)
	if err != nil {
		return err
	}

	return nil

}
