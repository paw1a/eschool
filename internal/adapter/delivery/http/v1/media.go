package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/core/domain"
)

func (h *Handler) initMediaRoutes(api *gin.RouterGroup) {
	media := api.Group("/upload")
	{
		media.POST("/", h.uploadFile)
	}
}

func (h *Handler) uploadFile(context *gin.Context) {
	formFile, err := context.FormFile("file")
	if err != nil {
		h.errorResponse(context, BadRequestError)
		return
	}

	file, err := formFile.Open()
	if err != nil {
		h.errorResponse(context, err)
		return
	}
	defer file.Close()

	url, err := h.mediaService.SaveMediaFile(context.Request.Context(), domain.File{
		Name:   formFile.Filename,
		Reader: file,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, url)
}
