package main

import (
	"ReportBoxEGTSTelemetryService/Application"
	"log"
)

func main() {
	/*Создаем экземпляр сервиса и считываем настройки*/
	Service := &Application.Service{}
	Error := Service.ReadSettings()
	if Error != nil {
		log.Fatal(Error)
	}
	Error = Service.InitService()
	if Error != nil {
		log.Fatal(Error)
	}
	Service.Start()

}
