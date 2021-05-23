package tests

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"supply_warehouse/dto"
	"supply_warehouse/orm"
	"supply_warehouse/utils"
	"testing"
	"time"
)

func TestCollectNewMerchant(t *testing.T) {
	db := utils.InitDb()
	//获取最后一个商户id
	lastMember := &orm.ImsSupermanHand2Member{}
	if err := lastMember.LoadLast(db); err != nil {
		t.Fail()
	}
	//
	start := lastMember.Uid + 1
	end := start + 20
	//start := 330
	//end := 680
	for i := start; i <= end; i++ {
		fmt.Println(fmt.Sprintf("=========开始获取商户id=%d数据==========", i))
		url := fmt.Sprintf(`https://w7.dapp100.cn/app/index.php?i=27&t=0&v=4.3.5&from=wxapp&c=entry&a=wxapp&do=shopXq&m=mzhk_sun&sign=&id=%d&openid=oB6Tt0JxXVGsRuORPRzmoalusX7U&type=1&userid=95534`, i)
		bodyText, err := utils.HttpGet(url)
		if err != nil {
			t.Fail()
		}
		xq := &dto.SupplyInfo{}
		if err := json.Unmarshal(bodyText, xq); err != nil {
			panic(err)
		}
		if xq.Bid == "" {
			fmt.Println("商户信息为空")
			continue
		}

		pre := "https://w7.dapp100.cn/attachment/"
		xq.Img = pre + xq.Img
		xq.Cimg = pre + xq.Cimg
		for k, v := range xq.Logo {
			xq.Logo[k] = pre + v
		}
		//校验
		uid, _ := strconv.ParseInt(xq.Bid, 10, 64)
		if uid == 0 {
			fmt.Println("商户数据为空")
			break
		}
		fmt.Println("开始创建商户")
		//创建会员与商户
		if err := createUserAndMerchant(db, xq); err != nil {
			t.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("=========商户id=%d数据采集完成==========", i))
	}

}

func createUserAndMerchant(db *gorm.DB, xq *dto.SupplyInfo) error {
	uid, _ := strconv.ParseInt(xq.Bid, 10, 64)

	// ims_superman_hand2_member
	member := &orm.ImsSupermanHand2Member{}
	member.Uniacid = 2
	member.Uid = uid
	member.Createtime = time.Now().Unix()
	//
	shop := &orm.ImsSupermanHand2Shop{}
	shop.Uniacid = 2
	shop.Uid = uid
	shop.Title = xq.Bname
	shop.Realname = xq.Bname
	shop.Phone = xq.Phone
	shop.Address = xq.Address
	shop.Status = 1
	shop.Createtime = time.Now().Unix()
	if len(xq.Wechat) > 0 {
		shop.Wechat = xq.Wechat[0]
	}

	tx := db.Begin()
	if err := tx.Error; err != nil {
		panic(err)
	}
	//先查询 存在则更新  不存再则创建
	if err := member.LoadByUid(db); err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := member.Create(db); err != nil {
				s := err.Error()
				if !strings.Contains(s, "Duplicate") {
					tx.Rollback()
					return nil
				}
			}
		}
	} else {
		member.Updatetime = time.Now().Unix()
		if err := member.Updates(db); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := shop.LoadByUid(db); err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := shop.Create(db); err != nil {
				s := err.Error()
				if !strings.Contains(s, "Duplicate") {
					tx.Rollback()
					return nil
				}
			}
		}
	} else {
		//更新微信
		if len(xq.Wechat) > 0 {
			shop.Wechat = xq.Wechat[0]
		}
		if err := shop.Updates(db); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
