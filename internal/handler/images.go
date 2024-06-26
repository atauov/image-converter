package handler

import (
	resp "github.com/atauov/image-converter/internal/lib/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	TYPE_JPEG = "image/jpeg"
	TYPE_PNG  = "image/png"
	TYPE_WEBP = "image/webp"
	MAX_SIZE  = 5 << 20 //5 MB
)

func (h *Handlers) uploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("failed to get image"))
		return
	}

	if file.Size > MAX_SIZE {
		c.JSON(http.StatusBadRequest, resp.Error("file size exceeds the maximum allowed size = 5 MB"))
		return
	}

	fileType := file.Header.Get("Content-Type")
	switch fileType {
	case TYPE_JPEG, TYPE_PNG, TYPE_WEBP:
	default:
		c.JSON(http.StatusBadRequest, resp.Error("file extension must be JPEG, PNG or WEBP"))
		return
	}

	open, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("failed to open file"))
	}
	defer open.Close()

}
