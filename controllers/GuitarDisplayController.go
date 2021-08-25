package controllers

import (
	"3TrashGuitarists/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GuitarDisplayController struct {
	beego.Controller
}

func (c *GuitarDisplayController) Get() {
	brand := c.Ctx.Input.Param(":brand")
	series := c.Ctx.Input.Param(":series")

	o := orm.NewOrm()
	guitars := []models.Guitars{}
	o.QueryTable(new(models.Guitars)).Filter("BrandName",brand).Filter("SeriesName",series).All(&guitars)

	c.Data["guitars"] = guitars
	fmt.Println(brand + series)
	fmt.Println(c.Data["guitars"])
	c.TplName = "GuitarDisplay.html"



}
