package dbops

import (
	"XianZhi/clog"
	"XianZhi/models"
)

//查询所有分类
func SelectAllCategory() (res []*models.GoodsCategory,err error) {
	stmt,err := dbConn.Prepare("select id, name from categories order by id")
	if err != nil {
		clog.Error("in dbops select category error,err:%v",err)
		return
	}
	rows,err := stmt.Query()
	if err != nil {
		clog.Error("in dbops select category error,err:%v",err)
		return
	}
	for rows.Next() {
		var re models.GoodsCategory
		if err = rows.Scan(&re.ID,&re.Name); err != nil {
			clog.Error("in dbops select category error,err:%v",err)
			return
		}
		res = append(res, &re)
	}
	defer stmt.Close()
	return
}

//查询首页中的商品列表
func SelectIndexGoods(skip int) (goodses []*models.Goods,err error) {
	stmt,err := dbConn.Prepare(`select goods.price,goods.browse,goods.name,goods.gphoto,goods.id,goods.created_at,
									   users.user_name,users.school,users.student_number from goods inner join users on 
									   users.id = goods.user_id where goods.is_buy = 0 and goods.is_low = 0 order by browse, created_at desc limit 15 offset ?`)
	if err != nil {
		clog.Error("in dbops select goods error,err:%v",err)
		return
	}
	rows,err := stmt.Query(skip)
	if err != nil {
		clog.Error("in dbops select category error,err:%v",err)
		return
	}
	for rows.Next() {
		var goods models.Goods
		if err = rows.Scan(&goods.Price,&goods.Browse,&goods.Name,&goods.Gphoto,&goods.ID,&goods.CreatedAt,&goods.UserName,&goods.School,&goods.StudentNumber); err != nil {
			clog.Error("in dbops select category error,err:%v",err)
			return
		}
		goodses = append(goodses,&goods)
	}
	defer stmt.Close()
	return
}

//查询分类页中的商品列表
func SelectCategoryGoods(id,skip int) (goodses []*models.Goods,err error) {
	stmt,err := dbConn.Prepare(`select goods.price,goods.browse,goods.name,goods.gphoto,goods.id,goods.created_at,
									   users.user_name,users.school,users.student_number from goods inner join users on 
									   users.id = goods.user_id where goods.category_id = ? and goods.is_buy = 0 and goods.is_low = 0 order by browse, created_at desc limit 15 offset ?`)
	if err != nil {
		clog.Error("in dbops select goods error,err:%v",err)
		return
	}
	rows,err := stmt.Query(id,skip)
	if err != nil {
		clog.Error("in dbops select category error,err:%v",err)
		return
	}
	for rows.Next() {
		var goods models.Goods
		if err = rows.Scan(&goods.Price,&goods.Browse,&goods.Name,&goods.Gphoto,&goods.ID,&goods.CreatedAt,&goods.UserName,&goods.School,&goods.StudentNumber); err != nil {
			clog.Error("in dbops select category error,err:%v",err)
			return
		}
		goodses = append(goodses,&goods)
	}
	defer stmt.Close()
	return
}

//查询商品的详细信息
func SelectGoodsInfo(uid,gid int) (goodsInfo *models.GoodsInfo,err error) {
	goodsInfo = &models.GoodsInfo{}
	stmt,err := dbConn.Prepare("select `name`, discribe, price, category_id, user_id, is_buy, gphoto, browse from goods where id = ?")
	if err != nil {
		clog.Error("get goods info failed,err:%v",err)
		return
	}
	err = stmt.QueryRow(gid).Scan(&goodsInfo.Name,&goodsInfo.Discribe,&goodsInfo.Price,&goodsInfo.CategoryId,&goodsInfo.UserId,&goodsInfo.IsBuy,&goodsInfo.Gphoto,&goodsInfo.Browse)
	if err != nil {
		clog.Error("get goods info failed,err:%v",err)
		return
	}
	stmt,err = dbConn.Prepare("select photo from photos where goods_id = ?")
	if err != nil {
		clog.Error("get images failed,err:%v",err)
		return
	}
	rows,err := stmt.Query(gid)
	if err != nil {
		clog.Error("get images failed,err:%v",err)
		return
	}
	for rows.Next() {
		var image string
		if err = rows.Scan(&image); err != nil {
			clog.Error("get images failed,err:%v",err)
			return
		}
		goodsInfo.Images = append(goodsInfo.Images, image)
	}
	stmt,err = dbConn.Prepare("select user_name, school, gphoto from users where id = ?")
	if err != nil {
		clog.Error("get user info failed,err:%v",err)
		return
	}
	err = stmt.QueryRow(goodsInfo.UserId).Scan(&goodsInfo.Username,&goodsInfo.School,&goodsInfo.Avavtar)
	if err != nil {
		clog.Error("get user info failed,err:%v",err)
		return
	}
	stmt,err = dbConn.Prepare("select count(*) from goods where user_id = ? and is_buy = 1")
	if err != nil {
		clog.Error("get sellCount failed,err:%v",err)
		return
	}
	err = stmt.QueryRow(goodsInfo.UserId).Scan(&goodsInfo.SellCount)
	if err != nil {
		clog.Error("get sellCount failed,err:%v",err)
		return
	}
	stmt,err = dbConn.Prepare("select * from collections where user_id = ? and goods_id = ?")
	if err != nil {
		clog.Error("get isnotcoll failed,err:%v",err)
		return
	}
	rows,err = stmt.Query(uid,gid)
	if err != nil {
		clog.Error("get isnotcoll failed,err:%v",err)
		return
	}
	goodsInfo.Isnotcoll = !rows.Next()
	defer stmt.Close()
	return
}

//查询买到的商品
func SelectHavebuyGoodsList(uid,skip int) (goodsList []*models.HavebuyGoods,err error) {
	stmt,err := dbConn.Prepare(`select goods.price, goods.name, goods.gphoto, goods.id, buys.order_id, buys.is_received from buys 
									   inner join goods on goods.id = buys.goods_id where buys.user_id = ? AND is_delete = 0
									    order by created_at desc limit 15 offset ?`)
	if err != nil {
		clog.Error("get haveBuyGoodsList falied,err:%v",err)
		return
	}
	rows,err := stmt.Query(uid,skip)
	if err != nil {
		clog.Error("get haveBuyGoodsList falied,err:%v",err)
		return
	}
	for rows.Next() {
		var goods models.HavebuyGoods
		if err = rows.Scan(&goods.Price,&goods.Name,&goods.Gphoto,&goods.Id,&goods.OrderId,&goods.IsReceived);err != nil {
			clog.Error("get haveBuyGoodsList falied,err:%v",err)
			return
		}
		goodsList = append(goodsList, &goods)
	}
	defer stmt.Close()
	return
}

func SelectWantbuyGoodsList(uid,skip int) (goodsList []*models.WantbuyGoods,err error) {
	stmt,err := dbConn.Prepare(`select goods.price, goods.name, goods.gphoto, goods.id from collections inner join 
										goods on goods.id = collections.goods_id where collections.user_id = ? order by 
										created_at, browse desc limit 15 offset ?`)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	rows,err := stmt.Query(uid,skip)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	for rows.Next() {
		var goods models.WantbuyGoods
		if err = rows.Scan(&goods.Price, &goods.Name, &goods.Gphoto, &goods.Id); err != nil {
			clog.Error("select goodsList failed,err:%v",err)
			return
		}
		goodsList = append(goodsList, &goods)
	}
	defer stmt.Close()
	return
}

func SelectPubGoodsList(uid,skip int) (goodsList []*models.PubGoods,err error) {
	stmt,err := dbConn.Prepare(`select id, price, gphoto, name, is_low, created_at, updated_at from goods 
										where user_id = ? AND is_buy = 0 order by created_at desc limit 15 offset ?`)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	rows,err := stmt.Query(uid,skip)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	for rows.Next() {
		var goods models.PubGoods
		if err = rows.Scan(&goods.Id,&goods.Price,&goods.Gphoto,&goods.Name,&goods.Is_low,&goods.CreatedAt,&goods.UpdatedAt);err != nil {
			clog.Error("select goodsList failed,err:%v",err)
			return
		}
		goodsList = append(goodsList, &goods)
	}
	defer stmt.Close()
	return
}

func SelectSoldGoodsList(uid,skip int) (goodsList []*models.SoldGoods,err error) {
	stmt,err := dbConn.Prepare(`select buys.telephone, buys.address, buys.user_id, goods.price, goods.name,
										goods.gphoto, goods.id, buys.order_id, goods.created_at from goods inner join
										 buys on buys.goods_id = goods.id where goods.user_id = ? AND is_buy = 1
										 order by created_at desc limit 15 offset ?`)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	rows,err := stmt.Query(uid,skip)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	for rows.Next() {
		var goods models.SoldGoods
		if err = rows.Scan(&goods.Telephone, &goods.Address, &goods.UserId, &goods.Price, &goods.Name, &goods.Gphoto, &goods.Id, &goods.OrderId, &goods.CreatedAt); err != nil {
			clog.Error("select goodsList failed,err:%v",err)
			return
		}
		goodsList = append(goodsList, &goods)
	}
	defer stmt.Close()
	return
}

func SelectActiveGoodsList(active string, skip int) (goodsList []*models.Goods, err error) {
	stmt,err := dbConn.Prepare(`select goods.price,goods.browse,goods.name,goods.gphoto,goods.id,goods.created_at,
										users.user_name,users.school,users.student_number from actives inner join goods 
										on goods.id = actives.goods_id AND is_buy = 0 AND is_low = 0 inner join users on 
										users.id = goods.user_id where active = ? order by goods.created_at, goods.browse
										desc limit 15 offset ?`)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	rows,err := stmt.Query(active,skip)
	if err != nil {
		clog.Error("select goodsList failed,err:%v",err)
		return
	}
	for rows.Next() {
		var goods models.Goods
		if err = rows.Scan(&goods.Price,&goods.Browse,&goods.Name,&goods.Gphoto,&goods.ID,&goods.CreatedAt,&goods.UserName,&goods.School,&goods.StudentNumber);err != nil {
			clog.Error("select goodsList failed,err:%v",err)
			return
		}
		goodsList = append(goodsList, &goods)
	}
	defer stmt.Close()
	return
}