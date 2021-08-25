package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserInformation3TG struct {
	Id int `orm:"pk;auto"`
	UserName string `orm:"index;unique"`
	Password string
	Email string `orm:"index;unique"`
	CreatTime time.Time `orm:"auto_now_add;type(datetime)"`
	MoneyRemain float64
	Address1 string
	Address2 string
	Address3 string
}

type Guitars struct {
	Id int `orm:"pk;auto"`
	GuitarName string
	ModelNumber string
	BrandName string
	SeriesName string
	NeckWood string
	FingerBoardWood string
	BodyWood string
	FretsNumber int
	TopWood string
	DisplayCover string
	DisplayPicture1 string
	DisplayPicture2 string
	DisplayPicture3 string
	DisplayPicture4 string
	DisplayPicture5 string
	StockNumber int `orm:"default(0)"`
}


func (u *UserInformation3TG)TableName() string{
	return "user_information"
}

func (g *Guitars)TableName() string{
	return "guitars"
}

func init() {
	orm.RegisterModel(new(UserInformation3TG))
	orm.RegisterModel(new(Guitars))
}