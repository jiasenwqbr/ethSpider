package handler

import (
	"ethSpider/httpHub"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Test(c *gin.Context) {
	orderID, ok := c.Params.Get("id")
	if !ok {
		httpHub.RespMsgStr(c, "Order Id parameter must be set", 1000)
		return
	}
	id, _ := strconv.Atoi(orderID)

	httpHub.RespMsgStruct(c, id, 200)
}
