package dto

type SupplyInfo struct {
	Bid           string   `json:"bid"`
	Bname         string   `json:"bname"`
	Logo          []string `json:"logo"`
	Content       string   `json:"content"`
	Uniacid       string   `json:"uniacid"`
	Status        string   `json:"status"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	Img           string   `json:"img"`
	Type          string   `json:"type"`
	Feature       string   `json:"feature"`
	Price         string   `json:"price"`
	Deliveryfee   string   `json:"deliveryfee"`
	Deliverytime  string   `json:"deliverytime"`
	Deliveryaway  string   `json:"deliveryaway"`
	InOpenid      string   `json:"in_openid"`
	BindOpenid    string   `json:"bind_openid"`
	Loginname     string   `json:"loginname"`
	Loginpassword string   `json:"loginpassword"`
	Uname         string   `json:"uname"`
	Starttime     string   `json:"starttime"`
	Endtime       string   `json:"endtime"`
	Coordinates   string   `json:"coordinates"`
	Longitude     string   `json:"longitude"`
	Latitude      string   `json:"latitude"`
	LtID          string   `json:"lt_id"`
	LtDay         string   `json:"lt_day"`
	Settleintime  string   `json:"settleintime"`
	Paytime       string   `json:"paytime"`
	Facility      []struct {
		ID            string `json:"id"`
		Facilityname  string `json:"facilityname"`
		Selectedimg   string `json:"selectedimg"`
		Unselectedimg string `json:"unselectedimg"`
		Sort          string `json:"sort"`
		Uniacid       string `json:"uniacid"`
	} `json:"facility"`
	StoreID          string        `json:"store_id"`
	StoreName        string        `json:"store_name"`
	Totalamount      string        `json:"totalamount"`
	Frozenamount     string        `json:"frozenamount"`
	Commission       string        `json:"commission"`
	Memdiscount      string        `json:"memdiscount"`
	Istop            string        `json:"istop"`
	Sort             string        `json:"sort"`
	EnablePrinter    string        `json:"enable_printer"`
	BackupsPrinter   string        `json:"backups_printer"`
	PrinterUser      string        `json:"printer_user"`
	PrinterUkey      string        `json:"printer_ukey"`
	PrinterSn        string        `json:"printer_sn"`
	Aid              string        `json:"aid"`
	ParentID1        string        `json:"parent_id_1"`
	ParentID2        string        `json:"parent_id_2"`
	ParentID3        string        `json:"parent_id_3"`
	Codeimg          string        `json:"codeimg"`
	Codeimgsrc       string        `json:"codeimgsrc"`
	BrandOpen        string        `json:"brand_open"`
	TimeOpen         string        `json:"time_open"`
	OpenPayment      string        `json:"open_payment"`
	Cimg             string        `json:"cimg"`
	IsDelivery       string        `json:"is_delivery"`
	DeliveryStart    string        `json:"delivery_start"`
	DeliveryFree     string        `json:"delivery_free"`
	DeliveryPrice    string        `json:"delivery_price"`
	DeliveryDistance string        `json:"delivery_distance"`
	Group            interface{}   `json:"group"`
	SubMchID         string        `json:"sub_mch_id"`
	IsCounp          string        `json:"is_counp"`
	PrestoreMoney    string        `json:"prestore_money"`
	InUserid         string        `json:"in_userid"`
	BindUserid       string        `json:"bind_userid"`
	LtEndtime        string        `json:"lt_endtime"`
	SaleWechat       string        `json:"sale_wechat"`
	SaleWxid         string        `json:"sale_wxid"`
	Createtime       string        `json:"createtime"`
	Star             int           `json:"star"`
	Cnums            int           `json:"cnums"`
	Coupons          []interface{} `json:"coupons"`
	Goods            []struct {
		Gid             string `json:"gid"`
		Gname           string `json:"gname"`
		Astime          string `json:"astime"`
		Selftime        string `json:"selftime"`
		Antime          string `json:"antime"`
		IsVip           string `json:"is_vip"`
		Buynum          int    `json:"buynum"`
		Lid             string `json:"lid"`
		Pic             string `json:"pic"`
		IndexImg        string `json:"index_img"`
		Content         string `json:"content"`
		WechatReptileID string `json:"wechat_reptile_id"`
		Fjoinnum        string `json:"fjoinnum"`
		IsDelivery      string `json:"is_delivery"`
		IsDeliveryLimit string `json:"is_delivery_limit"`
		DeliveryLimit   string `json:"delivery_limit"`
		Pics            []struct {
			URL  string `json:"url"`
			Type string `json:"type"`
		} `json:"pics"`
	} `json:"goods"`
	Wechat []string `json:"wechat"`
}
