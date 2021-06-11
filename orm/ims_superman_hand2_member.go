package orm

import "gorm.io/gorm"

type ImsSupermanHand2Member struct {
	ID         int64   `gorm:"column:id" json:"id" form:"id"`
	Uniacid    int64   `gorm:"column:uniacid" json:"uniacid" form:"uniacid"`
	Uid        int64   `gorm:"column:uid" json:"uid" form:"uid"`
	Getcash    int64   `gorm:"column:getcash" json:"getcash" form:"getcash"`
	Balance    float64 `gorm:"column:balance" json:"balance" form:"balance"`
	Createtime int64   `gorm:"column:createtime" json:"createtime" form:"createtime"`
	Updatetime int64   `gorm:"column:updatetime" json:"updatetime" form:"updatetime"`
}

func (i *ImsSupermanHand2Member) TableName() string {
	return "ims_superman_hand2_member"
}

func (i *ImsSupermanHand2Member) Create(db *gorm.DB) error {
	return db.Table(i.TableName()).Create(i).Error
}

func (i *ImsSupermanHand2Member) LoadLast(db *gorm.DB) error {
	return db.Table(i.TableName()).Where("uid < 5000").Order("uid desc").Limit(1).Scan(i).Error
}

func (i *ImsSupermanHand2Member) LoadByUid(db *gorm.DB) error {
	return db.Table(i.TableName()).Where("uid = ?", i.Uid).First(i).Error
}

func (i *ImsSupermanHand2Member) Updates(db *gorm.DB) error {
	return db.Model(i).Updates(i).Error
}
