package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple_golang/common"
	"simple_golang/component"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err = c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
	}
}
