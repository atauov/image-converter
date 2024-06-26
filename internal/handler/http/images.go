package http

import (
	"fmt"
	resp "github.com/atauov/image-converter/internal/lib/api/response"
	"github.com/atauov/image-converter/internal/models"
	"github.com/atauov/image-converter/internal/worker"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

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
	filename = strings.ToLower(strings.ReplaceAll(filename, " ", "_"))

	path := filepath.Join(h.cfg.UploadDir, filename)

	if err = c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("upload file err"))
		return
	}

	image := models.Image{
		UserID:   id,
		Filename: path,
		IsDone:   false,
	}

	if err = h.services.Images.CreateImage(image); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("create image err"))
		return
	}

	task, err := worker.NewImageUploadTask(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(fmt.Sprintf("could not create task: %v", err)))
		return
	}

	info, err := h.asynqC.Client.Enqueue(task)
	if err != nil {
		log.Printf("could not enqueue task: %v", err)
		c.JSON(http.StatusInternalServerError, resp.Error(fmt.Sprintf("could not enqueue task: %v", err)))
		return
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	c.JSON(http.StatusOK, resp.OK())
}

func (h *Handlers) changeImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("image_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("image id should be numeric"))
	}

	file, err := c.FormFile("image")
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
	filename = strings.ToLower(strings.ReplaceAll(filename, " ", "_"))

	path := filepath.Join(h.cfg.UploadDir, filename)

	if err = c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("upload file err"))
		return
	}

	image := models.Image{
		UserID:   id,
		Filename: path,
		IsDone:   false,
	}

	if err = h.services.UpdateImage(id, image); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("update image err"))
		return
	}

	task, err := worker.NewImageUploadTask(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(fmt.Sprintf("could not create task: %v", err)))
		return
	}

	info, err := h.asynqC.Client.Enqueue(task)
	if err != nil {
		log.Printf("could not enqueue task: %v", err)
		c.JSON(http.StatusInternalServerError, resp.Error(fmt.Sprintf("could not enqueue task: %v", err)))
		return
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	c.JSON(http.StatusOK, resp.OK())
}

func (h *Handlers) getAllImages(c *gin.Context) {
	var pagination Pagination

	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("failed to bind query"))
		return
	}

	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}

	offset := (pagination.Page - 1) * pagination.Limit

	images, err := h.services.GetAllImages(pagination.Limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, images)
}

func (h *Handlers) getByKey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("user id should be numeric"))
		return
	}

	images, err := h.services.GetImagesByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, images)
}

func (h *Handlers) deleteByURL(c *gin.Context) {
	url := c.Query("url")

	if err := h.services.DeleteImageByURL(url); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	task, err := worker.NewImageDeleteTask(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error(fmt.Sprintf("could not create task: %v", err)))
		return
	}

	info, err := h.asynqC.Client.Enqueue(task)
	if err != nil {
		log.Printf("could not enqueue task: %v", err)
		c.JSON(http.StatusInternalServerError, resp.Error(fmt.Sprintf("could not enqueue task: %v", err)))
		return
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	c.JSON(http.StatusOK, resp.OK())
}

func (h *Handlers) deleteByKey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("user id should be numeric"))
		return
	}

	if err = h.services.DeleteImagesByUserID(id); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("delete image err"))
		return
	}

	c.JSON(http.StatusOK, resp.OK())
}
