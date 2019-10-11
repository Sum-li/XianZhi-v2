package dbops

import (
	"XianZhi/clog"
	"fmt"
	"testing"
)

func TestMain(m *testing.M)  {
	_ = clog.InitLogger("console",clog.Config{Log_level:"debug"})
	//clearTables()
	m.Run()
	//clearTables  ()
}

func clearTables() {
	//清除数据库数据
	//_,_ = dbConn.Exec("truncate table")
}

func TestRun(t *testing.T)  {
	t.Run("SelectAllCategory",testSelectAllCategory)
	t.Run("SelectIndexGoods",testSelectIndexGoods)
	t.Run("SelectCategoryGoods",testSelectCategoryGoods)
	t.Run("SelectGoodsInfo",testSelectGoodsInfo)
	t.Run("SelectHavebuyGoodsList",testSelectHavebuyGoodsList)
	t.Run("SelectWantbuyGoodsList",testSelectWantbuyGoodsList)
	t.Run("SelectPubGoodsList",testSelectPubGoodsList)
	t.Run("SelectSoldGoodsList",testSelectSoldGoodsList)
	t.Run("SelectActiveGoodsList",testSelectActiveGoodsList)
}

func testSelectAllCategory(t *testing.T) {
	res,err := SelectAllCategory()
	if err != nil {
		t.Errorf("select category error ,err:%#v",err)
	}
	fmt.Printf("categorylist:%#v",res)
}

func testSelectIndexGoods(t *testing.T) {
	res,err := SelectIndexGoods(10)
	if err != nil {
		t.Errorf("select category error ,err:%#v",err)
	}
	fmt.Printf("categorylist:%#v",res)
}

func testSelectCategoryGoods(t *testing.T) {
	res,err := SelectCategoryGoods(10,10)
	if err != nil {
		t.Errorf("select category error ,err:%#v",err)
	}
	fmt.Printf("categorylist:%#v",res)
}

func testSelectGoodsInfo(t *testing.T) {
	res,err := SelectGoodsInfo(10,18)
	if err != nil {
		t.Errorf("select goods info error,err:%v",err)
	}
	fmt.Printf("goods info:%#v",res)
}

func testSelectHavebuyGoodsList(t *testing.T) {
	res,err := SelectHavebuyGoodsList(587,0)
	if err != nil {
		t.Errorf("select goods info error,err:%v",err)
	}
	fmt.Printf("goods info:%#v",res)
}

func testSelectWantbuyGoodsList(t *testing.T) {
	res,err := SelectWantbuyGoodsList(775,1)
	if err != nil {
		t.Errorf("select wantbuy failed,err:%v",err)
	}
	fmt.Printf("goodsList:%#v",res)
}

func testSelectPubGoodsList(t *testing.T) {
	res,err := SelectPubGoodsList(0,1)
	if err != nil {
		t.Errorf("select pubGoods failed,err:%v",err)
	}
	fmt.Printf("goodsList:%#v",res)
}

func testSelectSoldGoodsList(t *testing.T) {
	res,err := SelectSoldGoodsList(2,1)
	if err != nil {
		t.Errorf("select soldGoods failed,err:%v",err)
	}
	fmt.Printf("goodsList:%#v",res)
}

func testSelectActiveGoodsList(t *testing.T) {
	res,err := SelectActiveGoodsList("gd",0)
	if err != nil {
		t.Errorf("select activeGoods failed,err:%v",err)
	}
	fmt.Printf("goodsList:%#v",res)
}