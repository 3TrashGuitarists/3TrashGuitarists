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
	ShoppingCart *ShoppingCart `orm:"reverse(one)"`
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
	Description string
	Price float64
	StockNumber int `orm:"default(0)"`
	ShoppingCart *ShoppingCart `orm:"reverse(one)"`
}

type ShoppingCart struct{
	Id int `orm:"pk;auto"`
	UserInformation3TG *UserInformation3TG `orm:"rel(one)"`
	Guitars *Guitars `orm:"rel(one)"`
	Amount int
	PurchaseStatus bool `orm:"default(false)"` //true for successfully purchased, false for still in the cart
	Order *Order `orm:"rel(fk)"`
}

type Order struct{
	Id int `orm:"pk;auto"`
	Address string
	OrderNumber string
	ShippingStatus int `orm:"default(0)"`  // 0 for in stock, -1 for on the way, 1 for delivered
	ShoppingCarts []*ShoppingCart `orm:"reverse(many)"`
}

func (u *UserInformation3TG)TableName() string{
	return "user_information"
}

func (g *Guitars)TableName() string{
	return "guitars"
}

func (s *ShoppingCart)TableName() string{
	return "shopping_cart"
}

func (o *Order)TableName() string{
	return "order"
}

func init() {
	orm.RegisterModel(new(UserInformation3TG))
	orm.RegisterModel(new(Guitars))
	orm.RegisterModel(new(ShoppingCart))
	orm.RegisterModel(new(Order))
}