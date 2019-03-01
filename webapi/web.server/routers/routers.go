package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	)

func init() {
	util.RegisterValidation()
}

func GinReady() *gin.Engine {
	router := gin.New()
	router.Use(ginx.Recovery())

	// Version
	router.GET("/version", func(c *gin.Context) {
		ginx.RenderResult(c, metadata.Info())
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(
			200,
			map[string]interface{}{
				"status":"UP",
			},
		)
	})

	router.GET("/callJava", func(c *gin.Context) {
		var  respStatus int
		respStatus, resBody, errs := webutil.PostOrGetJSON("GET", "http://127.0.0.1:8081/callbygo", nil)
		if ok := util.HanderRespStatusAndErr(c, respStatus, errs, resBody); !ok {
			fmt.Printf("调用java出错")
			return
		}
		c.JSON(
			respStatus,
			string(resBody),
		)
	})

	
	return router
}
