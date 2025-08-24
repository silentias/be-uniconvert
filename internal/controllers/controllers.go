package controllers

import (
	"be-uniconvert/internal/services"
	"be-uniconvert/pkg"

	"github.com/gin-gonic/gin"
)

// @Summary Audio files convert
// @Accept multipart/form-data
// @Param file formData file true "audiofile to convert"
// @Param ext formData string true "ext"
// @Produce json
// @Router /convert/audio [post]
func Audio(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	ext := c.PostForm("ext")
	if len(ext) == 0 {
		c.JSON(400, gin.H{"msg": "ext required"})
	}
	filename := file.Filename
	if !pkg.IsAllowedExtension(filename) || !pkg.IsAllowedExtension(ext) {
		c.JSON(400, gin.H{"msg": "ext must be audio format"})
		return
	}

	err = services.UploadFile(file)
	if err != nil {
		c.JSON(500, gin.H{"msg": "Internal server error to upload file"})
		return
	}

	pathToFile, err := services.ConvertAudio(file, ext)
	if err != nil {
		c.JSON(500, gin.H{"msg": "Internal server error to upload file"})
		return
	}

	c.JSON(200, gin.H{
		"url": pathToFile,
	})
}
