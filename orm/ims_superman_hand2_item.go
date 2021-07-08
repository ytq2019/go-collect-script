package orm

import "gorm.io/gorm"

type ImsSupermanHand2Item struct {
	ID                   int64   `gorm:"column:id" json:"id" form:"id"`
	Uniacid              int64   `gorm:"column:uniacid" json:"uniacid" form:"uniacid"`
	Cid                  int64   `gorm:"column:cid" json:"cid" form:"cid"`
	CreditCid            int64   `gorm:"column:credit_cid" json:"credit_cid" form:"credit_cid"`
	SellerUid            int64   `gorm:"column:seller_uid" json:"seller_uid" form:"seller_uid"`
	Title                string  `gorm:"column:title" json:"title" form:"title"`
	Type                 int64   `gorm:"column:type" json:"type" form:"type"`
	Description          string  `gorm:"column:description" json:"description" form:"description"`
	Tags                 string  `gorm:"column:tags" json:"tags" form:"tags"`
	Summary              string  `gorm:"column:summary" json:"summary" form:"summary"`
	Album                string  `gorm:"column:album" json:"album" form:"album"`
	Thumb                string  `gorm:"column:thumb" json:"thumb" form:"thumb"`
	Video                string  `gorm:"column:video" json:"video" form:"video"`
	VideoThumb           string  `gorm:"column:video_thumb" json:"video_thumb" form:"video_thumb"`
	Avatar               string  `gorm:"column:avatar" json:"avatar" form:"avatar"`
	Nickname             string  `gorm:"column:nickname" json:"nickname" form:"nickname"`
	Phone                string  `gorm:"column:phone" json:"phone" form:"phone"`
	Wechat               string  `gorm:"column:wechat" json:"wechat" form:"wechat"`
	Appid                string  `gorm:"column:appid" json:"appid" form:"appid"`
	Url                  string  `gorm:"column:url" json:"url" form:"url"`
	ItemType             int64   `gorm:"column:item_type" json:"item_type" form:"item_type"`
	BuyType              int64   `gorm:"column:buy_type" json:"buy_type" form:"buy_type"`
	Wechatpay            int64   `gorm:"column:wechatpay" json:"wechatpay" form:"wechatpay"`
	Stock                int64   `gorm:"column:stock" json:"stock" form:"stock"`
	SellType1            int64   `gorm:"column:sell_type1" json:"sell_type1" form:"sell_type1"`
	SellType2            int64   `gorm:"column:sell_type2" json:"sell_type2" form:"sell_type2"`
	SellType3            int64   `gorm:"column:sell_type3" json:"sell_type3" form:"sell_type3"`
	WholesaleNumber      int64   `gorm:"column:wholesale_number" json:"wholesale_number" form:"wholesale_number"`
	WholesaleSinglePrice float64 `gorm:"column:wholesale_single_price" json:"wholesale_single_price" form:"wholesale_single_price"`
	WholesaleEmptyPrice  float64 `gorm:"column:wholesale_empty_price" json:"wholesale_empty_price" form:"wholesale_empty_price"`
	Price                float64 `gorm:"column:price" json:"price" form:"price"`
	Commission           float64 `gorm:"column:commission" json:"commission" form:"commission"`
	Credit               float64 `gorm:"column:credit" json:"credit" form:"credit"`
	OriginPrice          float64 `gorm:"column:origin_price" json:"origin_price" form:"origin_price"`
	ExpressPrice         float64 `gorm:"column:express_price" json:"express_price" form:"express_price"`
	DepositPrice         float64 `gorm:"column:deposit_price" json:"deposit_price" form:"deposit_price"`
	FreeShip             int64   `gorm:"column:free_ship" json:"free_ship" form:"free_ship"`
	TradeType1           int64   `gorm:"column:trade_type1" json:"trade_type1" form:"trade_type1"`
	TradeType2           int64   `gorm:"column:trade_type2" json:"trade_type2" form:"trade_type2"`
	TradeType3           int64   `gorm:"column:trade_type3" json:"trade_type3" form:"trade_type3"`
	FetchAddress         string  `gorm:"column:fetch_address" json:"fetch_address" form:"fetch_address"`
	Isbn                 string  `gorm:"column:isbn" json:"isbn" form:"isbn"`
	BookFields           string  `gorm:"column:book_fields" json:"book_fields" form:"book_fields"`
	SetTopFields         string  `gorm:"column:set_top_fields" json:"set_top_fields" form:"set_top_fields"`
	AreaPoints           string  `gorm:"column:area_points" json:"area_points" form:"area_points"`
	Lng                  string  `gorm:"column:lng" json:"lng" form:"lng"`
	Lat                  string  `gorm:"column:lat" json:"lat" form:"lat"`
	Province             string  `gorm:"column:province" json:"province" form:"province"`
	City                 string  `gorm:"column:city" json:"city" form:"city"`
	Address              string  `gorm:"column:address" json:"address" form:"address"`
	Ip                   string  `gorm:"column:ip" json:"ip" form:"ip"`
	Status               int64   `gorm:"column:status" json:"status" form:"status"`
	IsShop               int64   `gorm:"column:is_shop" json:"is_shop" form:"is_shop"`
	CreditTip            int64   `gorm:"column:credit_tip" json:"credit_tip" form:"credit_tip"`
	PayPosition          int64   `gorm:"column:pay_position" json:"pay_position" form:"pay_position"`
	Reason               string  `gorm:"column:reason" json:"reason" form:"reason"`
	AddFields            string  `gorm:"column:add_fields" json:"add_fields" form:"add_fields"`
	PageView             int64   `gorm:"column:page_view" json:"page_view" form:"page_view"`
	RealPageView         int64   `gorm:"column:real_page_view" json:"real_page_view" form:"real_page_view"`
	SellRegion           int64   `gorm:"column:sell_region" json:"sell_region" form:"sell_region"`
	IsCredit             int64   `gorm:"column:is_credit" json:"is_credit" form:"is_credit"`
	ExchangeTotal        int64   `gorm:"column:exchange_total" json:"exchange_total" form:"exchange_total"`
	Createtime           int64   `gorm:"column:createtime" json:"createtime" form:"createtime"`
	Updatetime           int64   `gorm:"column:updatetime" json:"updatetime" form:"updatetime"`
	Expiretime           int64   `gorm:"column:expiretime" json:"expiretime" form:"expiretime"`
	Refreshtime          int64   `gorm:"column:refreshtime" json:"refreshtime" form:"refreshtime"`
}

func (i *ImsSupermanHand2Item) TableName() string {
	return "ims_superman_hand2_item"
}

func (i *ImsSupermanHand2Item) Create(db *gorm.DB) error {
	return db.Table(i.TableName()).Create(i).Error
}

func (i *ImsSupermanHand2Item) SelectList(db *gorm.DB) ([]ImsSupermanHand2Item, error) {
	var items []ImsSupermanHand2Item
	//处理到了 236463
	if err := db.Model(i).Where("id <= 236463").Order("id desc").Limit(10000).Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (c *ImsSupermanHand2Item) Updates(tx *gorm.DB, ups map[string]interface{}) error {
	return tx.Model(c).Updates(ups).Error
}
