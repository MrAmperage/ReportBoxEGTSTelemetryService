package Application

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/spf13/viper"
)

/*Структура для приложения*/
type Service struct{ Listener net.Listener }

/*Считываем файл Settings.json с настройками*/
func (Service *Service) ReadSettings() error {
	viper.SetConfigName("Settings")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	Error := viper.ReadInConfig()
	if Error != nil {
		log.Fatal(Error)
	}
	return nil
}

func (Service *Service) Start() {

	for {
		Connection, Error := Service.Listener.Accept()
		if Error != nil {
			fmt.Println("Accept error:", Error)
			continue
		}
		go Service.HandlerConnection(Connection)
	}

}
func (Service *Service) InitService() error {
	ListenPort := viper.GetInt("ListenPort")
	Adress := fmt.Sprintf(":%d", ListenPort)
	Listener, Error := net.Listen("tcp", Adress)
	if Error != nil {
		return Error
	}
	Service.Listener = Listener
	return nil

}
func (Service *Service) HandlerConnection(Connection net.Conn) {
	File, Error := os.OpenFile("message.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if Error != nil {
		fmt.Println("File open error:", Error)
		return
	}
	for {
		Data, Error := io.ReadAll(Connection)
		if Error != nil {
			fmt.Println("Client disconnected:", Error)
			return
		}

		File.Write([]byte(hex.Dump(Data)))
		File.Write([]byte("\n"))
		Connection.Write([]byte("OK\n"))
		File.Close()
		Connection.Close()

	}
}
