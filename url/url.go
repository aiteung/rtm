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
	page.Get("/userdata/:handphone", controller.GetDataUserbyPhone)
	page.Delete("delete/:handphone", controller.DeleteDataUser)
	page.Get("/rapat", controller.Getdatartm)
	page.Post("/insertrapat", controller.InsertDataRapat)
	page.Get("/datartm/:agenda", controller.GetDataRtmByAgenda)
	page.Delete("delete/:lokasi", controller.DeleteDataRtmFromLokasi)
	page.Post("/job/insert", controller.InsertDataJob)
	page.Get("/job/get/datajob", controller.GetDataJob)
	page.Get("/job/get/datajobtitle", controller.GetDataJobtitle)
	page.Delete("/job/delete/datajob", controller.DeleteDataJob)
	page.Delete("/job/delete/datajobtitle", controller.DeleteDataJobtitle)
	page.Post("/accounts/insert", controller.InsertDataAccounts)
	page.Get("/accounts/get/dataaccounts", controller.GetDataAccounts)
	page.Get("/accounts/get/datanama", controller.GetDataNama)
	page.Delete("/accounts/delete/dataaccounts", controller.DeleteDataAccounts)
	page.Delete("/accounts/delete/datanama", controller.DeleteDataNama)
	page.Get("/absensi/get/dataabsensi", controller.GetDataAbsensi)
	page.Get("/posisi", controller.GetDataPosisi)
	page.Post("/insertabsensi", controller.InsertAbsensi)
	page.Post("/insertPosisi", controller.InsertPosisi)
}
