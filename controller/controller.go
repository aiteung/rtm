package controller

import (
	"context"
	"fmt"
	"net/http"
	"rtm/config"

	pkg "github.com/MSyahidAlFajri/packagertm"
	"github.com/gofiber/fiber/v2"
	jobdesk "github.com/harisriyoni3/rtmpackage"
	rtpkg "github.com/rofinafiin/rtm-package"
	"go.mongodb.org/mongo-driver/bson"
	accpkg "github.com/daffaaudyapramana/packagertm"
	lpkg "github.com/GilarYa/packagertm"
	monitor "github.com/xshazks/rtmbdsaw"
)

var usercol = "data_user"
var datartm = "data_rtm"

//func WsWhatsAuthQR(c *websocket.Conn) {
//	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
//}
//
//func PostWhatsAuthRequest(c *fiber.Ctx) error {
//	if string(c.Request().Host()) == config.Internalhost {
//		var req whatsauth.WhatsauthRequest
//		err := c.BodyParser(&req)
//		if err != nil {
//			return err
//		}
//		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
//		return c.JSON(ntfbtn)
//	} else {
//		var ws whatsauth.WhatsauthStatus
//		ws.Status = string(c.Request().Host())
//		return c.JSON(ws)
//	}
//
//}

func GetHome(c *fiber.Ctx) error {
	//getip := musik.GetIPaddress()
	getip := "Hello guys"
	return c.JSON(getip)
}

func Getdatauser(c *fiber.Ctx) error {
	id := "cc2"
	getstats := rtpkg.GetDatauser(id, config.MongoConn, usercol)
	fmt.Println(getstats)
	return c.JSON(getstats)
}

func InsertData(c *fiber.Ctx) error {
	database := config.MongoConn
	var jumlah rtpkg.User
	if err := c.BodyParser(&jumlah); err != nil {
		return err
	}
	Inserted := rtpkg.InsertDataUser(database,
		jumlah.Iduser,
		jumlah.Nama,
		jumlah.Email,
		jumlah.Handphone,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": Inserted,
	})
}

func GetDataUserbyPhone(c *fiber.Ctx) error {
	hp := c.Params("handphone")
	data := rtpkg.GetDataUserFromPhone(hp, config.MongoConn, "data_user")
	fmt.Println(data)
	return c.JSON(data)
}

func DeleteDataUser(c *fiber.Ctx) error {
	hp := c.Params("handphone")
	data := rtpkg.DeleteData(hp, config.MongoConn, "data_user")
	return c.JSON(data)
}

func Getdatartm(c *fiber.Ctx) error {
	namarapat := "Rapat Akreditasi"
	getstats := pkg.GetDataRtm(namarapat, config.MongoConn, datartm)
	fmt.Println(getstats)
	return c.JSON(getstats)
}

func GetDataRtmByAgenda(c *fiber.Ctx) error {
	hp := c.Params("agendarapat")
	data := pkg.GetDataRtmFromAgenda(hp, config.MongoConn, "data_rtm")
	fmt.Println(data)
	return c.JSON(data)
}

func DeleteDataRtmFromLokasi(c *fiber.Ctx) error {
	hp := c.Params("lokasirapat")
	data := pkg.DeleteDataRtm(hp, config.MongoConn, "data_rtm")
	return c.JSON(data)
}

func InsertDataRapat(c *fiber.Ctx) error {
	database := config.MongoConn
	var tambah pkg.DataRTM
	if err := c.BodyParser(&tambah); err != nil {
		return err
	}
	Inserted := rtpkg.InsertDataUser(database,
		tambah.NamaRapat,
		tambah.TanggalRapat,
		tambah.LokasiRapat,
		tambah.AgendaRapat,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil Tersimpan.",
		"inserted_id": Inserted,
	})
}

// JOB
func InsertDataJob(c *fiber.Ctx) error {
	database := config.MongoConn
	var job jobdesk.Job
	if err := c.BodyParser(&job); err != nil {
		return err
	}
	Inserted := jobdesk.InsertDataJob(database,
		job.Job_title,
		job.Deskripsi,
		job.Deadline,
		job.Priority,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Job berhasil disimpan.",
		"inserted_id": Inserted,
	})
}
func GetDataJob(c *fiber.Ctx) error {
	hp := c.Params("priority")
	data := jobdesk.GetDataJob(hp, config.MongoConn, "data_job")
	fmt.Println(data)
	return c.JSON(data)
}
func GetDataJobtitle(c *fiber.Ctx) error {
	hp := c.Params("job_title")
	data := jobdesk.GetDataJobtitle(hp, config.MongoConn, "data_job")
	fmt.Println(data)
	return c.JSON(data)
}
func DeleteDataJob(c *fiber.Ctx) error {
	hp := c.Params("priority")
	data := jobdesk.DeleteDataJob(hp, config.MongoConn, "DeleteDataJob")
	return c.JSON(data)
}
func DeleteDataJobtitle(c *fiber.Ctx) error {
	hp := c.Params("job_title")
	data := jobdesk.DeleteDataJobtitle(hp, config.MongoConn, "DeleteDataJob")
	return c.JSON(data)
}

// ACCOUNTS
func InsertDataAccounts(c *fiber.Ctx) error {
	database := config.MongoConn
	var accounts accpkg.Accounts
	if err := c.BodyParser(&accounts); err != nil {
		return err
	}
	Inserted := accpkg.InsertDataAccounts(database,
		accounts.Nama,
		accounts.Email,
		accounts.Sosial,
		accounts.Perusahaan,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Accounts berhasil disimpan.",
		"inserted_id": Inserted,
	})
}
func GetDataAccounts(c *fiber.Ctx) error {
	hp := c.Params("perusahaan")
	data := accpkg.GetDataAccounts(hp, config.MongoConn, "data_accounts")
	fmt.Println(data)
	return c.JSON(data)
}
func GetDataNama(c *fiber.Ctx) error {
	hp := c.Params("nama")
	data := accpkg.GetDataNama(hp, config.MongoConn, "data_accounts")
	fmt.Println(data)
	return c.JSON(data)
}
func DeleteDataAccounts(c *fiber.Ctx) error {
	hp := c.Params("perusahaan")
	data := accpkg.DeleteDataAccounts(hp, config.MongoConn, "DeleteDataAccounts")
	return c.JSON(data)
}
func DeleteDataNama(c *fiber.Ctx) error {
	hp := c.Params("nama")
	data := accpkg.DeleteDataAccounts(hp, config.MongoConn, "DeleteDataAccounts")
	return c.JSON(data)
}

// ABSENSI
func InsertAbsensi(nip string, date string, status string, employee string, positions string) (InsertedID interface{}) {
	var absensi Absensi
	absensi.NIP = nip
	absensi.Date = date
	absensi.Status = status
	absensi.Employee = employee
	absensi.Positions = positions

	return InsertOneDoc("176892865", "absensi", absensi)
}
func InsertPosisi(nama_posisi string, bidang string) (InsertedID interface{}) {
	var posisi Posisi
	posisi.Nama_Posisi = nama_posisi
	posisi.Bidang = bidang

	return InsertOneDoc("Engineering Manager", "posisi", posisi)
}
func GetDataAbsensi(nip string) (data []Absensi) {
	user := MongoConnect("rtmtugas").Collection("absen")
	filter := bson.M{"nip": nip}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAbsensi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetDataPosisi(nama_posisi string) (data []Posisi) {
	user := MongoConnect("rtmtugas").Collection("Posisi")
	filter := bson.M{"nama_posisi": nama_posisi}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPosisi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// CS
func GetDataNamacs(c *fiber.Ctx) error {
	namacs := "Gilar"
	getstats := lpkg.GetDataNamacs(namacs, config.MongoConn, "data_DataCS")
	fmt.Println(getstats)
	return c.JSON(getstats)
}
func GetDataNegaracs(c *fiber.Ctx) error {
	hpcs := c.Params("negaracs")
	data := lpkg.GetDataNegaracs(hpcs, config.MongoConn, "data_DataCS")
	fmt.Println(data)
	return c.JSON(data)
}
func DeleteDataNamacs(c *fiber.Ctx) error {
	hpcs := c.Params("namacs")
	data := lpkg.DeleteDataNamacs(hpcs, config.MongoConn, "data_DataCS")
	return c.JSON(data)
}
func DeleteDataNegaracs(c *fiber.Ctx) error {
	hp := c.Params("negara")
	data := lpkg.DeleteDataNegaracs(hp, config.MongoConn, "data_DataCS")
	return c.JSON(data)
}
func InsertDataCS(c *fiber.Ctx) error {
	database := config.MongoConn
	var tambah lpkg.DataCS
	if err := c.BodyParser(&tambah); err != nil {
		return err
	}
	Inserted := lpkg.InsertDataCS(database,
		tambah.Namacs,
		tambah.Emailcs,
		tambah.Nohpcs,
		tambah.Negaracs,
		tambah.Desccs,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil Tersimpan.",
		"inserted_id": Inserted,
	})
}
// monitoring
func Insertmonitor(c *fiber.Ctx) error {
	var rtmdb monitor.Monitor
	if err := c.BodyParser(&rtmdb); err != nil {
		return err
	}
	Inserted := monitor.Insertmonitor(
		rtmdb.Proker,
		rtmdb.Status,
		rtmdb.About,
		rtmdb.Karyawan,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Job berhasil disimpan.",
		"inserted_id": Inserted,
	})
}

func GetDatamonitor(c *fiber.Ctx) error {
	status := c.Params("status")
	data := monitor.GetDatamonitor(status)
	fmt.Println(data)
	return c.JSON(data)
}

func GetDatakaryawan(c *fiber.Ctx) error {
	karyawan := c.Params("karyawan")
	data := monitor.GetDatakaryawan(karyawan)
	fmt.Println(data)
	return c.JSON(data)
}

func GetDataproker(c *fiber.Ctx) error {
	proker := c.Params("proker")
	data := monitor.GetDataproker(proker)
	fmt.Println(data)
	return c.JSON(data)
}
