package router

import (
	"ethSpider/handler"
	"ethSpider/logger"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"net/http"
)

var log = logger.SetupLog("router")

func Router() *gin.Engine {
	engine := gin.New()
	engine.Use(helmet.Default())
	// 使用后端跨域中间件
	engine.Use(cors())

	/*
	   接口汇总
	*/
	// 设置分组路由
	v1 := engine.Group("/v1")
	// 健康因子
	v1.GET("/test", handler.Test)

	return engine
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		//defer func() {
		//	err := recover()
		//	// err报错返回Any类型， 可以忽略
		//	if err != nil {
		//		log.Warnf("Panic info is: %v", err)
		//	}
		//}()

		c.Next()
	}
}
