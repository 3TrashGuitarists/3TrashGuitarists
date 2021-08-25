package controllers

type CheckoutController struct {
	//beego.Controller
	BaseController
}

func (c *CheckoutController) Get() {
	c.TplName = "Checkout.html"

}
