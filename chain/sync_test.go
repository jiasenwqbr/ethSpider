package chain

import (
	"context"
	"ethSpider/conf"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestSync(t *testing.T) {
	//spider, err := New(conf.WSS_ETH_RPC)
	//if err != nil {
	//	fmt.Printf("Initial error%s\n", err)
	//}
	//
	//// ETH client
	//err = spider.Sync()
	//if err != nil {
	//	fmt.Printf("Syncing error%s\n", err)
	//}

	r := gin.Default()
	r.GET("/block", func(c *gin.Context) {
		client, err := ethclient.Dial(conf.WSS_ETH_RPC)
		if err != nil {
			fmt.Println(err)
		}

		ctx := context.Background()

		data, err := ParsingBlock(ctx, client, uint64(16089709))
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"message": data,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
