package main

import (
	"log"
	"quiz-pekan3-damar/controller"
	"quiz-pekan3-damar/router"
)

func main() {
	sqlCon, err := controller.KoneksiDB()

	if err != nil {
		log.Fatal("Gagal tersambung ke postgres ", err)
	}
	defer sqlCon.Close()

	controller.DBSqlConn = sqlCon

	PORT := ":8080"

	router.StartServer(sqlCon).Run(PORT)
}
