package Application

import (
	"ReportBoxEGTSTelemetryService/Application/EGTSPackge"
	"fmt"
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
	defer Connection.Close()
	File, Error := os.OpenFile("message.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if Error != nil {
		fmt.Println("File open error:", Error)
		return
	}
	defer File.Close()
	Buffer := make([]byte, 2048)
	for {
		ByteCount, Error := Connection.Read(Buffer)
		if Error != nil {
			fmt.Println("Client disconnected:", Error)
			return
		}
		Data := Buffer[:ByteCount]

		File.Write([]byte(fmt.Sprintf("% x\n", Data)))
		Package := &EGTSPackge.EGTSPackge{}
		Error = Package.Decode(Data)
		if Error != nil {
			File.Write([]byte(fmt.Sprintf("Ошибка: %s\n", Error)))
		} else {
			File.Write([]byte(fmt.Sprintf("Protocol Version %d\n", Package.Header.ProtocolVersion)))
			File.Write([]byte(fmt.Sprintf("Security Key Id %d\n", Package.Header.SecurityKeyId)))
			File.Write([]byte(fmt.Sprintf("Prefix Flag %d\n", Package.Header.EGTSPackgeHeaderFlags.Prefix)))
			File.Write([]byte(fmt.Sprintf("Route Flag %t\n", Package.Header.EGTSPackgeHeaderFlags.Route)))
			File.Write([]byte(fmt.Sprintf("Encryption Algorithm Flag %d\n", Package.Header.EGTSPackgeHeaderFlags.EncryptionAlgorithm)))
			File.Write([]byte(fmt.Sprintf("Compression Flag %t\n", Package.Header.EGTSPackgeHeaderFlags.Compression)))
			File.Write([]byte(fmt.Sprintf("Priority Flag %d\n", Package.Header.EGTSPackgeHeaderFlags.Priority)))
			File.Write([]byte(fmt.Sprintf("Header Length %d\n", Package.Header.HeaderLength)))
			File.Write([]byte(fmt.Sprintf("Header Encoding %d\n", Package.Header.HeaderEncoding)))
		}

		Connection.Write([]byte("OK\n"))
	}
}
