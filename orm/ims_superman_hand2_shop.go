package orm

import "gorm.io/gorm"

type ImsSupermanHand2Shop struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Uniacid    int64  `gorm:"column:uniacid" json:"uniacid" form:"uniacid"`
	Uid        int64  `gorm:"column:uid" json:"uid" form:"uid"`
	Title      string `gorm:"column:title" json:"title" form:"title"`
	Realname   string `gorm:"column:realname" json:"realname" form:"realname"`
	Phone      string `gorm:"column:phone" json:"phone" form:"phone"`
	Wechat     string `gorm:"column:wechat" json:"wechat" form:"wechat"`
	Address    string `gorm:"column:address" json:"address" form:"address"`
	Latitude   string `gorm:"column:latitude" json:"latitude" form:"latitude"`
	Longitude  string `gorm:"column:longitude" json:"longitude" form:"longitude"`
	CertImgs   string `gorm:"column:cert_imgs" json:"cert_imgs" form:"cert_imgs"`
	Status     int64  `gorm:"column:status" json:"status" form:"status"`
	Createtime int64  `gorm:"column:createtime" json:"createtime" form:"createtime"`
}

func (i *ImsSupermanHand2Shop) TableName() string {
	return "ims_superman_hand2_shop"
}

func (i *ImsSupermanHand2Shop) Create(db *gorm.DB) error {
	return db.Table(i.TableName()).Create(i).Error
}

func (i *ImsSupermanHand2Shop) LoadByUid(db *gorm.DB) error {
	return db.Table(i.TableName()).Where("uid = ?", i.Uid).First(i).Error
}

func (i *ImsSupermanHand2Shop) IsExist(db *gorm.DB) (bool, error) {
	var count int64
	if err := db.Table(i.TableName()).Where("uid = ?", i.Uid).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (i *ImsSupermanHand2Shop) Updates(db *gorm.DB) error {
	return db.Model(i).Updates(i).Error
}
