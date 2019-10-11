package models

//商品
type Goods struct {
	Price         int       `json:"price"`
	Name          string    `json:"name";orm:"Scan:name"`
	Gphoto        string    `json:"gphoto"`
	ID            int       `json:"good_id";db:"id"`
	UserName      string    `json:"user_name"`
	CreatedAt     string `json:"pub_time"`
	School        string    `json:"school"`
	StudentNumber string    `json:"student_number"`
	Browse        int       `json:"browse"`
}

//商品的详细信息
type GoodsInfo struct {
	Name       string   `json:"name"`
	Discribe   string   `json:"discribe"`
	Price      int      `json:"price"`
	CategoryId int      `json:"category_id"`
	UserId     int      `json:"user_id"`
	IsBuy      bool     `json:"is_buy"`
	Gphoto     string   `json:"gphoto"`
	Images     []string `json:"images"`
	Username   string   `json:"username"`
	School     string   `json:"school"`
	SellCount  int      `json:"sell_count"`
	Avavtar    string   `json:"avavtar"`
	Isnotcoll  bool     `json:"isnotcoll"`
	Browse     int      `json:"browse"`
}

//买到的商品
type HavebuyGoods struct {
	Price      int    `json:"price"`
	Id         int    `json:"id"`
	Gphoto     string `json:"gphoto"`
	Name       string `json:"name"`
	OrderId    string `json:"order_id"`
	IsReceived bool   `json:"is_received"`
}

//想要买的商品
type WantbuyGoods struct {
	Price  int    `json:"price"`
	Id     int    `json:"id"`
	Gphoto string `json:"gphoto"`
	Name   string `json:"name"`
}

//发布的商品
type PubGoods struct {
	Price     int    `json:"price"`
	Id        int    `json:"id"`
	Gphoto    string `json:"gphoto"`
	Name      string `json:"name"`
	Is_low    bool   `json:"is_low"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

//卖出的商品
type SoldGoods struct {
	Price     int    `json:"price"`
	Id        int    `json:"id"`
	Gphoto    string `json:"gphoto"`
	Name      string `json:"name"`
	OrderId   string `json:"order_id"`
	CreatedAt string `json:"created_at"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
	UserId    int    `json:"user_id"`
}
