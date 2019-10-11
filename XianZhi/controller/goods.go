package controller

import (
	"XianZhi/clog"
	"XianZhi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取所有分类
func CategoryList(c *gin.Context) {
	res, err := service.GetCategoryList()
	if err != nil {
		clog.Error("get categorylist failed,err:%v", err)
		c.JSON(http.StatusInternalServerError, "get categorylist failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

//商品首页
func GoodsIndex(c *gin.Context) {
	pnStr := c.Query("page_count")
	res, err := service.GetIndexGoodsList(pnStr)
	if err != nil {
		clog.Error("get goodses failed,err:%v", err)
		c.JSON(http.StatusInternalServerError, "get goodses failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

//分类商品页面
func GoodsCategory(c *gin.Context) {
	cidStr := c.Query("category_id")
	pnStr := c.Query("page_count")
	res, err := service.GetCategoryGoodsList(cidStr, pnStr)
	if err != nil {
		clog.Error("get goodses failed,err:%v", err)
		c.JSON(http.StatusInternalServerError, "get goodses failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

//商品详情页面
func Goodsdetail(c *gin.Context) {
	gidStr := c.Query("goods_id")
	uidStr := c.Query("user_id")
	res, err := service.GetGoodsInfo(uidStr, gidStr)
	if err != nil {
		clog.Error("get goodsInfo failed,err:%v", err)
		c.JSON(http.StatusInternalServerError, "get goodsInfo failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

//买到的商品列表页面
func HavebuyGoodsList(c *gin.Context) {
	uidStr := c.Query("user_id")
	pnStr := c.Query("page_count")
	res,err := service.GetHavebuyGoodsList(uidStr,pnStr)
	if err != nil {
		clog.Error("get goodsList failed,err:%v",err)
		c.JSON(http.StatusInternalServerError, "get goodsList failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK,res)
}

//想要的商品列表页面
func WantbuyGoodsList(c *gin.Context)  {
	uidStr := c.Query("user_id")
	pnStr := c.Query("page_count")
	res,err := service.GetWantbuyGoodsList(uidStr,pnStr)
	if err != nil {
		clog.Error("get goodsList failed,err:%v",err)
		c.JSON(http.StatusInternalServerError, "get goodsList failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK,res)
}

//发布的商品列表页面
func PubGoodsList(c *gin.Context)  {
	uidStr := c.Query("user_id")
	pnStr := c.Query("page_count")
	res,err := service.GetPubGoodsList(uidStr,pnStr)
	if err != nil {
		clog.Error("get goodsList failed,err:%v",err)
		c.JSON(http.StatusInternalServerError, "get goodsList failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK,res)
}

//卖出的商品列表页面
func SoldGoodsList(c *gin.Context)  {
	uidStr := c.Query("user_id")
	pnStr := c.Query("page_count")
	res,err := service.GetSoldGoodsList(uidStr,pnStr)
	if err != nil {
		clog.Error("get goodsList failed,err:%v",err)
		c.JSON(http.StatusInternalServerError, "get goodsList failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK,res)
}

//活动商品列表页面
func ActiveGoodsList(c *gin.Context)  {
	active := c.Query("active")
	pnStr := c.Query("page_count")
	res,err := service.GetActiveGoodsList(active,pnStr)
	if err != nil {
		clog.Error("get goodsList failed,err:%v",err)
		c.JSON(http.StatusInternalServerError, "get goodsList failed,err:"+err.Error())
		return
	}
	c.JSON(http.StatusOK,res)
}
