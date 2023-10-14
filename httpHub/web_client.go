package httpHub

import (
	"ethSpider/httpHub/router"
	"ethSpider/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type Server struct {
	HttpServer   *http.Server
	Handler      *gin.Engine
	Logger       logging.StandardLogger
	IpAddr       string
	Port         uint16
	WriteTimeout time.Duration
}

func NewServer(ipAddr string, port uint16, writeTimeout time.Duration) (*Server, error) {
	var log = logger.SetupLog("web_server")

	// 关闭gin日志打印输出（适用于生产环境）
	mode := gin.ReleaseMode
	gin.SetMode(mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Warnf("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	handler := router.Router()

	server := &Server{
		Handler:      handler,
		Logger:       log,
		IpAddr:       ipAddr,
		Port:         port,
		WriteTimeout: writeTimeout,
	}

	return server, nil
}

func (s *Server) Run() error {
	s.Logger.Infof("Listening and serving HTTP on %s:%d", s.IpAddr, s.Port)
	s.HttpServer = createServer(s.Handler, s.IpAddr, s.Port, s.WriteTimeout)
	err := s.HttpServer.ListenAndServe()
	return errors.Errorf("Faild to run HTTP server:%v", err)
}

func createServer(handler *gin.Engine, ipAddr string, port uint16, writeTimeout time.Duration) *http.Server {
	url := fmt.Sprintf("%s:%d", ipAddr, port)

	s := &http.Server{
		Addr:           url,
		Handler:        handler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   writeTimeout,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s
}

// RespMsgStr 返回一条信息
func RespMsgStr(c *gin.Context, msg string, code uint) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":   code,
			"result": msg,
		},
	)
}

// RespMsgStruct 返回消息结构体
func RespMsgStruct(c *gin.Context, msg any, code uint) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":   code,
			"result": msg,
		},
	)
}
