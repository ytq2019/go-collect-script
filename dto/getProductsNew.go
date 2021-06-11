package dto

type GetProductsNew struct {
	Status int          `json:"status"`
	Msg    string       `json:"msg"`
	Date   ProductItems `json:"date"`
}

type ProductItems []struct {
	Bid             string   `json:"bid"`
	Wechat          []string `json:"wechat"`
	Phone           string   `json:"phone"`
	Longitude       string   `json:"longitude"`
	Latitude        string   `json:"latitude"`
	Img             string   `json:"img"`
	Bname           string   `json:"bname"`
	Type            string   `json:"type"`
	Address         string   `json:"address"`
	Facility        []string `json:"facility"`
	Feature         string   `json:"feature"`
	Gid             string   `json:"gid"`
	Content         string   `json:"content"`
	Status          string   `json:"status"`
	WechatReptileID string   `json:"wechat_reptile_id"`
	Pic             string   `json:"pic"`
	IndexImg        string   `json:"index_img"`
	Selftime        string   `json:"selftime"`
	Pics            []struct {
		URL  string `json:"url"`
		Type string `json:"type"`
	} `json:"pics"`
}

func (s ProductItems) Len() int {
	return len(s)
}
func (s ProductItems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ProductItems) Less(i, j int) bool {
	return s[i].Gid > s[j].Gid
}
