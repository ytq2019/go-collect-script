package tests

import (
	"encoding/json"
	"fmt"
	"github.com/techoner/gophp/serialize"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
	"supply_warehouse/dto"
	"supply_warehouse/orm"
	"supply_warehouse/utils"
	"testing"
	"time"
)

func TestCollectNewGoods(t *testing.T) {
	db := utils.InitDb()

	for cate := 39; cate <= 47; cate++ {
		log.Println(fmt.Sprintf("=======开始采集分类为%d数据=======", cate))
		for page := 1; page <= 3; page++ {
			log.Println(fmt.Sprintf("=======开始采集第%d页=======", page))
			url := fmt.Sprintf(`https://w7.dapp100.cn/app/index.php?i=27&t=0&v=4.3.5&from=wxapp&c=entry&a=wxapp&do=getProducts&m=mzhk_sun&sign=09815a452384fb11466d9db08f783ca0&lat=40.106178283691406&lon=116.3189697265625&openid=oB6Tt0JxXVGsRuORPRzmoalusX7U&keyword=&brand_cate=%d&type=1&page=%d&aid=&userid=95534`, cate, page)
			bodyBytes, err := utils.HttpGet(url)
			if err != nil {
				t.Fail()
			}
			pn := &dto.GetProductsNew{}
			if err := json.Unmarshal(bodyBytes, pn); err != nil {
				t.Fail()
			}
			if err := saveOrUpdateGoods(db, pn); err != nil {
				fmt.Println(err.Error())
				t.Fail()
			}
			log.Println(fmt.Sprintf("=======第%d页采集完毕=======", page))
		}
		log.Println(fmt.Sprintf("=======分类为%d数据采集完毕=======", cate))
	}

}

func saveOrUpdateGoods(db *gorm.DB, gn *dto.GetProductsNew) error {
	for _, v := range gn.Date {
		fmt.Println(fmt.Sprintf("商品名称 %s", v.Content))
		uid, _ := strconv.ParseInt(v.Bid, 10, 64)
		item := &orm.ImsSupermanHand2Item{}
		item.Uniacid = 2
		gid, _ := strconv.ParseInt(v.Gid, 10, 64)
		item.ID = gid
		var categoryId int64
		if len(v.Facility) > 0 {
			switch v.Facility[0] {
			case "工厂直发":
				categoryId = 39
			case "大宗货源":
				categoryId = 40
			case "临期仓库":
				categoryId = 41
			case "贸易商":
				categoryId = 42
			}
		} else {
			categoryId = 41
		}

		item.Cid = categoryId
		item.SellerUid = uid
		item.Title = v.Content
		item.Type = 1
		item.Description = v.Content
		item.Wechatpay = 1
		item.Stock = 1
		item.SellType3 = 1
		item.TradeType2 = 1
		item.Lng = v.Longitude
		item.Lat = v.Latitude
		item.Address = v.Address
		item.Ip = "127.0.0.1"
		item.Status = 1
		item.IsShop = 1
		item.CreditTip = 1
		item.SellRegion = 1
		parse, _ := time.Parse("2006-01-02 15:04:05", v.Selftime)
		item.Createtime = parse.Unix()
		item.Refreshtime = parse.Unix()
		//处理图片
		if len(v.Pics) > 0 {
			pics := make([]string, 0)
			videos := make([]string, 0)
			for _, v := range v.Pics {
				//图片
				if v.Type == "1" {
					pics = append(pics, v.URL)
				} else {
					videos = append(videos, v.URL)
				}
			}
			piscByte, _ := serialize.MarshalSlice(pics)
			videosByte, _ := serialize.MarshalSlice(videos)
			item.Album = string(piscByte)
			item.Thumb = string(piscByte)
			if len(videos) > 0 {
				item.Video = string(videosByte)
			}
			item.VideoThumb = "images/2/2021/05/xHL2YRyeY8r224rrlpv22SKYPCmH27hM.jpg"
		} else {
			item.Album = `a:0:{}`
			item.Thumb = `a:0:{}`
		}

		if err := db.Create(item).Error; err != nil {
			if strings.Contains(err.Error(), "Duplicate") {
				fmt.Println(fmt.Sprintf("商品gid=%d,已存在，跳过", gid))
				continue
			} else {
				return err
			}
		}
		fmt.Println(fmt.Sprintf("成功采集一条商品记录,gid=%d", gid))
	}
	return nil
}
