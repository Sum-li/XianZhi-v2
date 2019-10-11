package main

import (
	"XianZhi/clog"
	"XianZhi/controller"
	"github.com/gin-gonic/gin"
)

func main()  {
	_ = clog.InitLogger("console",clog.Config{Log_level:"debug"})
	r := gin.New()

	goods := r.Group("/goods")
	{
		goods.GET("/category",controller.CategoryList)
		goods.GET("/index",controller.GoodsIndex)
		goods.GET("/goodscategory",controller.GoodsCategory)
		goods.GET("/detail",controller.Goodsdetail)
		goods.GET("/wantbuy",controller.WantbuyGoodsList)
		goods.GET("/mypub",controller.PubGoodsList)
		goods.GET("/havebuy",controller.HavebuyGoodsList)
		goods.GET("/sold",controller.SoldGoodsList)
		goods.GET("/active",controller.ActiveGoodsList)
	}

	_ = r.Run(":8081")
}
