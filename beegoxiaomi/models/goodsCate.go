package models

type GoodsCate struct {
	Id            int
	Title         string
	CateImg       string
	Link          string
	Template      string
	Pid           int
	SubTitle      string
	Keywords      string
	Description   string
	Sort          int
	Status        int
	AddTime       int
	GoodsCateItem []GoodsCate `gorm:"foreignkey:Pid;references:Id"` //自关联

}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
