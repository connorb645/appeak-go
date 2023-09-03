package route

import (
	"time"

	"github.com/connorb645/appeak-go/api/controller"
	"github.com/connorb645/appeak-go/store"
	"github.com/connorb645/appeak-go/usecase"
	"github.com/gin-gonic/gin"
)

func NewDocumentRouter(timeout time.Duration, hcp store.HelpCenterProvider, group *gin.RouterGroup) {
	ds := store.NewDocumentStore(hcp)
	dc := &controller.DocumentController{
		DocumentUsecase: usecase.NewDocumentUsecase(ds, timeout),
	}
	group.GET("/documents", dc.FetchAllDocuments)
}
