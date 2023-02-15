package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tools-go/httptools"
	"tools-go/httptools/common"
)

func StartHttpServer() {

	routers := make(map[string][]httptools.RouterInfo)

	routers["GET"] = []httptools.RouterInfo{
		{Uri: "/", Handlers: []gin.HandlerFunc{a}},
		{Uri: "/abc", Handlers: []gin.HandlerFunc{b}},
	}

	httptools.StartServer(routers, ":8080")
}

func a(context *gin.Context) {
	fmt.Println(1)
	common.Success(context, nil)

}

func b(context *gin.Context) {
	context.String(200, "Hello, b")
	fmt.Println(2)
}
