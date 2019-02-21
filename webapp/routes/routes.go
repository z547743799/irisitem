package routes

import (
	"github.com/kataras/iris/mvc"
	"gitlab.com/z547743799/irisitem/bootstrap"
	"gitlab.com/z547743799/irisitem/controller"
	"gitlab.com/z547743799/irismanager/service"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	ItemService := service.NewTbItemService()
	Item := mvc.New(b.Party("/item"))
	Item.Register(ItemService)
	Item.Handle(new(controller.PageController))

	//admin := mvc.New(b.Party("/admin"))
	//admin.Router.Use(middleware.BasicAuth)
	//admin.Register(superstarService)
	//admin.Handle(new(controllers.AdminController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
