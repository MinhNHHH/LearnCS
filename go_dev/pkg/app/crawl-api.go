package app

import (
	"net/http"

	"github.com/MinhNHHH/go_dev/pkg/crawl"
	"github.com/gin-gonic/gin"
)

type GetLinkPayload struct {
	URL string
}

func (app *Application) GetDocumentApi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inputData GetLinkPayload
		if err := ctx.ShouldBindBodyWithJSON(&inputData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, error := crawl.GetDoc(inputData.URL)
		if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})
	}
}
