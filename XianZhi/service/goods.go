package service

import (
	"XianZhi/clog"
	"XianZhi/dbops"
	"XianZhi/models"
	"fmt"
	"strconv"
)

func GetCategoryList() (res []*models.GoodsCategory,err error) {
	res,err = dbops.SelectAllCategory()
	if err != nil {
		fmt.Printf("in service get categorylist failed,err:%V\n",err)
	}
	return
}

func GetIndexGoodsList(pnStr string) (goodses []*models.Goods, err error) {
	pn,err := strconv.Atoi(pnStr)
	if err !=nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	skip := (pn-1)*15
	goodses,err = dbops.SelectIndexGoods(skip)
	if err != nil {
		clog.Error("get goodses failed,err:%v",err)
		return
	}
	return
}

func GetCategoryGoodsList(cidStr,pnStr string) (goodses []*models.Goods, err error) {
	pn,err := strconv.Atoi(pnStr)
	if err !=nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	skip := (pn-1)*15
	cid,err := strconv.Atoi(cidStr)
	if err !=nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	goodses,err = dbops.SelectCategoryGoods(cid,skip)
	if err != nil {
		clog.Error("get goodses failed,err:%v",err)
		return
	}
	return
}

func GetGoodsInfo(uidStr,gidStr string) (goodsInfo *models.GoodsInfo,err error) {
	uid,err := strconv.Atoi(uidStr)
	if err != nil {
		clog.Error("get uid failed,err:%v",err)
		return
	}
	gid,err := strconv.Atoi(gidStr)
	if err != nil {
		clog.Error("get gid failed,err:%v",err)
		return
	}

	goodsInfo,err = dbops.SelectGoodsInfo(uid,gid)
	if err != nil {
		clog.Error("get goodsInfo failed,err:%v",err)
		return
	}
	return
}

func GetHavebuyGoodsList(uidStr,pnStr string) (goodsList []*models.HavebuyGoods,err  error) {
	uid,err := strconv.Atoi(uidStr)
	if err != nil {
		clog.Error("get uid failed,err:%v",err)
		return
	}
	pn,err := strconv.Atoi(pnStr)
	if err != nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	skip := (pn-1)*15
	goodsList,err = dbops.SelectHavebuyGoodsList(uid,skip)
	if err != nil {
		clog.Error("get goodsList falied,err:%v",err)
		return
	}
	return
}

func GetWantbuyGoodsList(uidStr,pnStr string) (goodsList []*models.WantbuyGoods,err error) {
	uid,err := strconv.Atoi(uidStr)
	if err != nil {
		clog.Error("get uid failed,err:%v",err)
		return
	}
	pn,err := strconv.Atoi(pnStr)
	if err != nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	skip := (pn-1)*15
	goodsList,err = dbops.SelectWantbuyGoodsList(uid,skip)
	if err != nil {
		clog.Error("get goodsList falied,err:%v",err)
		return
	}
	return
}

func GetPubGoodsList(uidStr,pnStr string) (goodsList []*models.PubGoods,err error) {
	uid,err := strconv.Atoi(uidStr)
	if err != nil {
		clog.Error("get uid failed,err:%v",err)
		return
	}
	pn,err := strconv.Atoi(pnStr)
	if err != nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	skip := (pn-1)*15
	goodsList,err = dbops.SelectPubGoodsList(uid,skip)
	if err != nil {
		clog.Error("get goodsList falied,err:%v",err)
		return
	}
	return
}

func GetSoldGoodsList(uidStr,pnStr string) (goodsList []*models.SoldGoods,err error) {
	uid,err := strconv.Atoi(uidStr)
	if err != nil {
		clog.Error("get uid failed,err:%v",err)
		return
	}
	pn,err := strconv.Atoi(pnStr)
	if err != nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	skip := (pn-1)*15
	goodsList,err = dbops.SelectSoldGoodsList(uid,skip)
	if err != nil {
		clog.Error("get goodsList falied,err:%v",err)
		return
	}
	return
}

func GetActiveGoodsList(active,pnStr string) (goodsList []*models.Goods,err error) {
	pn,err := strconv.Atoi(pnStr)
	if err != nil {
		clog.Error("get pn failed,err:%v",err)
		return
	}
	skip := (pn-1)*15
	goodsList,err = dbops.SelectActiveGoodsList(active,skip)
	if err != nil {
		clog.Error("get goodsList failed,err:%v",err)
		return
	}
	return
}
