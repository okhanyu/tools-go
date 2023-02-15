package httptools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var engine *gin.Engine

type RouterInfo struct {
	Uri      string
	Handlers []gin.HandlerFunc
}

// StartServer 启动服务
func StartServer(routers map[string][]RouterInfo, ipPort string) {
	engine = gin.New()
	engine.Use(gin.Recovery())
	for method, routerGroup := range routers {
		if method == http.MethodGet {
			for _, info := range routerGroup {
				engine.GET(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodPost {
			for _, info := range routerGroup {
				engine.POST(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodPut {
			for _, info := range routerGroup {
				engine.PUT(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodDelete {
			for _, info := range routerGroup {
				engine.DELETE(info.Uri, info.Handlers...)
			}
		}
		if method == http.MethodPatch {
			for _, info := range routerGroup {
				engine.PATCH(info.Uri, info.Handlers...)
			}
		}
	}

	err := engine.Run(ipPort)

	if err != nil {
		fmt.Println(err)
		return
	}
}
