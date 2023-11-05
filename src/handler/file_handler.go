package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/google/uuid"
)

func (s *APIServer) UploadImage(c *gin.Context) {
	image, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	imageID := strings.ReplaceAll(uuid.New().String(), "-", "")
	imageExt := strings.Split(image.Filename, ".")[1]
	ImageWithExt := fmt.Sprintf("%s.%s", imageID, imageExt)
	filePath := filepath.Join("images", ImageWithExt)

	if err = c.SaveUploadedFile(image, filePath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// save avatar for user to database
	userID := c.GetHeader("user_id")
	if len(userID) != 0 {
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		if err = s.store.UserStore.UpdateUser(userIDInt, map[string]interface{}{
			"avatar_url": imageID,
		}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	if err = s.store.FileStore.Create(&model.File{
		UUID:      imageID,
		Extension: imageExt,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid": imageID,
	})
}

type GetImageRequest struct {
	UUID string `form:"uuid"`
}

func (s *APIServer) GetImage(c *gin.Context) {
	req := &GetImageRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Retrieve the file metadata from the database
	fileMetadata, err := s.store.FileStore.GetByUUID(req.UUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if fileMetadata == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	fileName := fmt.Sprintf("%s.%s", fileMetadata.UUID, fileMetadata.Extension)
	// Define the path of the file to be retrieved
	filePath := filepath.Join("images", fileName)
	// Open the file
	fileData, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileData.Close()
	// Read the first 512 bytes of the file to determine its content type
	fileHeader := make([]byte, 512)
	_, err = fileData.Read(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	fileContentType := http.DetectContentType(fileHeader)
	// Get the file info
	fileInfo, err := fileData.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}
	// Set the headers for the file transfer and return the file
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", fileContentType)
	c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	c.File(filePath)
}
