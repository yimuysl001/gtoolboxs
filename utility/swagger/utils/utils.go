package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"strconv"
)

func GetJs(router *ghttp.RouterGroup, relativePath string, hexContent string) {
	rs, err := hex.DecodeString(hexContent)
	if nil != err {
		fmt.Println("err:", err)
		return
	}
	router.GET(relativePath, func(r *ghttp.Request) {

		r.Response.Status = http.StatusOK
		r.Response.Header().Set("content-type", "application/javascript;charset=UTF-8")
		r.Response.Header().Set("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Write(rs)
	})
}

func GetCss(router *ghttp.RouterGroup, relativePath string, content string) {
	router.GET(relativePath, func(r *ghttp.Request) {
		rs := []byte(content)

		r.Response.Status = http.StatusOK
		r.Response.Header().Set("content-type", "text/css;charset=UTF-8")
		r.Response.Header().Set("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Write(rs)
	})
}

func GetOther(router *ghttp.RouterGroup, relativePath string, hexContent string) {
	rs, err := hex.DecodeString(hexContent)
	if nil != err {
		fmt.Println("err:", err)
		return
	}
	router.GET(relativePath, func(r *ghttp.Request) {
		r.Response.Status = http.StatusOK
		r.Response.Header().Set("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Write(rs)
	})

}

func GetJson(router *ghttp.RouterGroup, relativePath string, json string) {
	router.GET(relativePath, func(r *ghttp.Request) {
		rs := []byte(json)
		r.Response.Status = http.StatusOK
		r.Response.Header().Set("content-type", "application/json;charset=UTF-8")
		r.Response.Header().Set("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Write(rs)
	})
}

func GetHtml(router *ghttp.RouterGroup, relativePath string, content string) {
	rs := []byte(content)
	router.GET(relativePath, func(r *ghttp.Request) {
		r.Response.Status = http.StatusOK
		r.Response.Header().Set("content-type", "text/html;charset=UTF-8")
		r.Response.Header().Set("content-length", strconv.Itoa(len(rs)))
		r.Response.Header().Set("connection", "keep-alive")
		r.Response.Write(rs)
	})

}
