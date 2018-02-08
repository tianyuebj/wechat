package web

import (
	"github.com/gin-gonic/gin"
)

func Strap(router gin.IRouter) {
	router.GET("/", handlerWeChatGet)
}
