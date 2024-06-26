package handler

import (
	"fmt"
	resp "github.com/atauov/image-converter/internal/lib/api/response"
	"github.com/atauov/image-converter/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

func (h *Handlers) uploadImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("user id should be numeric"))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("failed to get image"))
		return
	}

	if err = checkInputImage(file); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	open, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("failed to open file"))
		return
	}
	defer open.Close()

	filename := filepath.Base(file.Filename)
	path := filepath.Join(h.cfg.UploadDir, filename)
	fmt.Println(filename)

	if err = c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("upload file err"))
		return
	}

	image := models.Image{
		UserID:   id,
		Filename: filename,
		IsDone:   false,
	}

	fmt.Println(image)

	//TODO send to MQ job

	c.JSON(http.StatusOK, resp.OK())
}

func (h *Handlers) changeImage(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("failed to get image"))
		return
	}

	if err = checkInputImage(file); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	_ = id
	//TODO change images
	//TODO resend img
}

func (h *Handlers) getAllImages(c *gin.Context) {
	//TODO return first page
	//TODO make pagination
}

func (h *Handlers) getByKey(c *gin.Context) {
	key := c.Param("key")
	_ = key
}

func (h *Handlers) deleteByURL(c *gin.Context) {
	url := c.Param("url")
	_ = url
}

func (h *Handlers) deleteByKey(c *gin.Context) {
	key := c.Param("key")
	_ = key
}
