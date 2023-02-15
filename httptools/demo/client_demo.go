package demo

import (
	"fmt"
	"net/url"
	"tools-go/httptools"
)

func StartHttpClient() {
	all, _ := httptools.SendGet("https://hanyu.cool", url.Values{"ABC": {"1"}})
	fmt.Println(all)
}
