package orm

import "gorm.io/gorm"

type ImsMcMembers struct {
	Uid             int64  `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	Uniacid         int    `gorm:"column:uniacid;NOT NULL"`
	Mobile          string `gorm:"column:mobile;NOT NULL"`
	Email           string `gorm:"column:email;NOT NULL"`
	Password        string `gorm:"column:password;NOT NULL"`
	Salt            string `gorm:"column:salt;NOT NULL"`
	Groupid         int    `gorm:"column:groupid;NOT NULL"`
	Credit1         string `gorm:"column:credit1;NOT NULL"`
	Credit2         string `gorm:"column:credit2;NOT NULL"`
	Credit3         string `gorm:"column:credit3;NOT NULL"`
	Credit4         string `gorm:"column:credit4;NOT NULL"`
	Credit5         string `gorm:"column:credit5;NOT NULL"`
	Credit6         string `gorm:"column:credit6;NOT NULL"`
	Createtime      uint   `gorm:"column:createtime;NOT NULL"`
	Realname        string `gorm:"column:realname;NOT NULL"`
	Nickname        string `gorm:"column:nickname;NOT NULL"`
	Avatar          string `gorm:"column:avatar;NOT NULL"`
	Qq              string `gorm:"column:qq;NOT NULL"`
	Vip             uint   `gorm:"column:vip;NOT NULL"`
	Gender          int    `gorm:"column:gender;NOT NULL"`
	Birthyear       uint   `gorm:"column:birthyear;NOT NULL"`
	Birthmonth      uint   `gorm:"column:birthmonth;NOT NULL"`
	Birthday        uint   `gorm:"column:birthday;NOT NULL"`
	Constellation   string `gorm:"column:constellation;NOT NULL"`
	Zodiac          string `gorm:"column:zodiac;NOT NULL"`
	Telephone       string `gorm:"column:telephone;NOT NULL"`
	Idcard          string `gorm:"column:idcard;NOT NULL"`
	Studentid       string `gorm:"column:studentid;NOT NULL"`
	Grade           string `gorm:"column:grade;NOT NULL"`
	Address         string `gorm:"column:address;NOT NULL"`
	Zipcode         string `gorm:"column:zipcode;NOT NULL"`
	Nationality     string `gorm:"column:nationality;NOT NULL"`
	Resideprovince  string `gorm:"column:resideprovince;NOT NULL"`
	Residecity      string `gorm:"column:residecity;NOT NULL"`
	Residedist      string `gorm:"column:residedist;NOT NULL"`
	Graduateschool  string `gorm:"column:graduateschool;NOT NULL"`
	Company         string `gorm:"column:company;NOT NULL"`
	Education       string `gorm:"column:education;NOT NULL"`
	Occupation      string `gorm:"column:occupation;NOT NULL"`
	Position        string `gorm:"column:position;NOT NULL"`
	Revenue         string `gorm:"column:revenue;NOT NULL"`
	Affectivestatus string `gorm:"column:affectivestatus;NOT NULL"`
	Lookingfor      string `gorm:"column:lookingfor;NOT NULL"`
	Bloodtype       string `gorm:"column:bloodtype;NOT NULL"`
	Height          string `gorm:"column:height;NOT NULL"`
	Weight          string `gorm:"column:weight;NOT NULL"`
	Alipay          string `gorm:"column:alipay;NOT NULL"`
	Msn             string `gorm:"column:msn;NOT NULL"`
	Taobao          string `gorm:"column:taobao;NOT NULL"`
	Site            string `gorm:"column:site;NOT NULL"`
	Bio             string `gorm:"column:bio;NOT NULL"`
	Interest        string `gorm:"column:interest;NOT NULL"`
	PayPassword     string `gorm:"column:pay_password;NOT NULL"`
	UserFrom        int    `gorm:"column:user_from;NOT NULL"`
}

func (i *ImsMcMembers) TableName() string {
	return "ims_mc_members"
}

func (i *ImsMcMembers) Create(db *gorm.DB) error {
	return db.Model(i).Create(i).Error
}

func (i *ImsMcMembers) LoadByUid(db *gorm.DB) error {
	return db.Model(i).Where("uid = ?", i.Uid).First(i).Error
}
