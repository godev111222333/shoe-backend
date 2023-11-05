package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/google/uuid"
)

type RegisterUserRequest struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *APIServer) RegisterUser(c *gin.Context) {
	req := &RegisterUserRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &model.User{
		Phone:     req.Phone,
		Name:      req.Name,
		AvatarURL: "",
		Email:     req.Email,
		Balance:   0,
		Status:    model.UserStatusPendingRegistration,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.store.UserStore.Create(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.otpService.SendOTP(model.OTPTypeRegister, req.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "registration OTP sent",
	})
}

type LoginRequest struct {
	Phone string `form:"phone"`
}

func (s *APIServer) UserInfo(c *gin.Context) {
	req := &LoginRequest{}
	if err := c.Bind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := s.store.UserStore.GetByPhone(req.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

type VerifyRegistrationRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func (s *APIServer) VerifyRegistration(c *gin.Context) {
	req := &VerifyRegistrationRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isMatched, err := s.otpService.VerifyOTP(model.OTPTypeRegister, req.Email, req.OTP)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !isMatched {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "otp not matched",
		})
		return
	}

	user, err := s.store.UserStore.GetByEmail(req.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	if err = s.store.UserStore.UpdateUser(user.ID, map[string]interface{}{
		"status": model.UserStatusActive,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type UserLoginRequest struct {
	Email string `json:"email"`
}

func (s *APIServer) UserLogin(c *gin.Context) {
	req := &UserLoginRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.otpService.SendOTP(model.OTPTypeLogin, req.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type VerifyLoginRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type UserResponse struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	AvatarURL string `json:"avatar_url"`
}

type VerifyLoginResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	*UserResponse
}

func (s *APIServer) VerifyLogin(c *gin.Context) {
	req := &VerifyLoginRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isMatch, err := s.otpService.VerifyOTP(model.OTPTypeLogin, req.Email, req.OTP)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !isMatch {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid OTP",
		})
		return
	}

	accessToken, accessPayload, err := s.tokenMaker.CreateToken(req.Email, "user", 24*time.Hour)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := s.store.UserStore.GetByEmail(req.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, &VerifyLoginResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		UserResponse: &UserResponse{
			Email:     user.Email,
			Phone:     user.Phone,
			Name:      user.Name,
			Balance:   user.Balance,
			AvatarURL: user.AvatarURL,
		},
	})
}

func (s *APIServer) UploadImage(c *gin.Context) {
	image, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := c.GetHeader("userID")
	if len(userID) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please put userID in header",
		})
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

	// save uuid to database
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
