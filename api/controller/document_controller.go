package controller

import (
	"net/http"

	"github.com/connorb645/appeak-go/domain"
	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	DocumentUsecase domain.DocumentUsecase
}

func (dc *DocumentController) FetchAllDocuments(c *gin.Context) {
	docs, err := dc.DocumentUsecase.FetchAllDocuments(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, docs)
}
