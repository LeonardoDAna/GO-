package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Handler {
    //初始化一个新的httprouter路由器实例
    router := httprouter.New()

    //使用HandlerFunc()函数为接口注册相关方法，URL模式和处理函数。
    //注意http.MethodGet和http.MethodPost是常量，等价于字符串"GET"和"POST"
    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
    router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
    router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

    //返回httprouter实例
    return router
}