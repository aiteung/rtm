package url

import (
	"rtm/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	//page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	//page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.GetHome)
	page.Get("/user", controller.Getdatauser)
	page.Post("/insert", controller.InsertData)
	page.Post("/job/masuk", controller.InsertDataJob)
	page.Get("/userdata/:handphone", controller.GetDataUserbyPhone)
	page.Delete("delete/:handphone", controller.DeleteDataUser)
}
