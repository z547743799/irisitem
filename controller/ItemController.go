package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gitlab.com/z547743799/iriscommon/utils"
	"gitlab.com/z547743799/irismanager/service"
)

type PageController struct {
	Ctx     iris.Context
	Service service.TbItemService
}

func (c *PageController) GetBy(itemid int64) mvc.Result {

	velus := c.Ctx.GetCookie("token")
	Sen := utils.Manager.Start(c.Ctx)
	jsonuser := Sen.GetString(velus)
	if velus == "" || jsonuser == "" {
		c.Ctx.Redirect("http://127.0.0.1:8088")
		return nil
	} else {
		Item := c.Service.Get(itemid)
		ItemDes := c.Service.GetItemDescById(itemid)

		return mvc.View{
			Name: "item.html",
			Data: iris.Map{
				"item":    Item,
				"itemdes": ItemDes,
			},
		}

	}

}
