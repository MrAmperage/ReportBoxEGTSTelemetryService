package Application

import (
	"ReportBoxEGTSTelemetryService/Application/EGTSPackge"
	"bytes"
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

	for {

		Package := &EGTSPackge.EGTSPackge{}
		ByteMessage, Error := Package.ReadMessage(Connection)
		File.Write([]byte(fmt.Sprintf("% x\n", ByteMessage)))
		Reader := bytes.NewReader(ByteMessage)
		Error = Package.Decode(Reader)
		if Error != nil {
			File.Write([]byte(fmt.Sprintf("Ошибка: %s\n", Error)))
			return
		} else {
			File.Write([]byte(fmt.Sprintf("Protocol Version: %d\n", Package.Header.ProtocolVersion)))
			File.Write([]byte(fmt.Sprintf("Security Key Id: %d\n", Package.Header.SecurityKeyId)))
			File.Write([]byte(fmt.Sprintf("Prefix Flag: %d\n", Package.Header.Flags.Prefix)))
			File.Write([]byte(fmt.Sprintf("Route Flag: %t\n", Package.Header.Flags.Route)))
			File.Write([]byte(fmt.Sprintf("Encryption Algorithm Flag: %d\n", Package.Header.Flags.EncryptionAlgorithm)))
			File.Write([]byte(fmt.Sprintf("Compression Flag: %t\n", Package.Header.Flags.Compression)))
			File.Write([]byte(fmt.Sprintf("Priority Flag: %d\n", Package.Header.Flags.Priority)))
			File.Write([]byte(fmt.Sprintf("Header Length: %d\n", Package.Header.HeaderLength)))
			File.Write([]byte(fmt.Sprintf("Header Encoding: %d\n", Package.Header.HeaderEncoding)))
			File.Write([]byte(fmt.Sprintf("Frame Data Length: %d\n", Package.Header.FrameDataLength)))
			File.Write([]byte(fmt.Sprintf("Packet Identifier: %d\n", Package.Header.PacketIdentifier)))
			File.Write([]byte(fmt.Sprintf("Packet Type: %s\n", Package.Header.PacketTypeToString(Package.Header.PacketType))))
			if Package.Header.Flags.Route {
				File.Write([]byte(fmt.Sprintf("Peer Address: %d\n", Package.Header.PeerAddress)))
				File.Write([]byte(fmt.Sprintf("Recipient Address: %d\n", Package.Header.RecipientAddress)))
				File.Write([]byte(fmt.Sprintf("Time To Live: %d\n", Package.Header.TimeToLive)))

			}
			File.Write([]byte(fmt.Sprintf("Header Check Sum: %d\n", Package.Header.HeaderCheckSum)))
			for _, Record := range Package.Records {
				File.Write([]byte(fmt.Sprintf("Record Length: %d\n", Record.RecordHeader.RecordLength)))
				File.Write([]byte(fmt.Sprintf("Record Number: %d\n", Record.RecordHeader.RecordNumber)))
				File.Write([]byte(fmt.Sprintf("Source Service On Device: %t\n", Record.RecordHeader.RecordFlags.SourceServiceOnDevice)))
				File.Write([]byte(fmt.Sprintf("Recipient Service On Device: %t\n", Record.RecordHeader.RecordFlags.RecipientServiceOnDevice)))
				File.Write([]byte(fmt.Sprintf("Group: %t\n", Record.RecordHeader.RecordFlags.Group)))
				File.Write([]byte(fmt.Sprintf("Record Processing Priority: %d\n", Record.RecordHeader.RecordFlags.RecordProcessingPriority)))
				File.Write([]byte(fmt.Sprintf("Time Field Exist: %t\n", Record.RecordHeader.RecordFlags.TimeFieldExist)))
				File.Write([]byte(fmt.Sprintf("Event Id Field Exist: %t\n", Record.RecordHeader.RecordFlags.EventIdFieldExist)))
				File.Write([]byte(fmt.Sprintf("Object Id Field Exist: %t\n", Record.RecordHeader.RecordFlags.ObjectIdFieldExist)))
				if Record.RecordHeader.RecordFlags.ObjectIdFieldExist {
					File.Write([]byte(fmt.Sprintf("Object Identifier: %d\n", Record.RecordHeader.ObjectIdentifier)))
				}
				if Record.RecordHeader.RecordFlags.EventIdFieldExist {
					File.Write([]byte(fmt.Sprintf("Event Identifier: %d\n", Record.RecordHeader.EventIdentifier)))
				}
				if Record.RecordHeader.RecordFlags.TimeFieldExist {
					File.Write([]byte(fmt.Sprintf("Time: %d\n", Record.RecordHeader.Time)))
				}
				File.Write([]byte(fmt.Sprintf("Source Service Type: %s\n", Record.RecordHeader.RecordTypeToString(Record.RecordHeader.SourceServiceType))))
				File.Write([]byte(fmt.Sprintf("Recipient Service Type: %s\n", Record.RecordHeader.RecordTypeToString(Record.RecordHeader.RecipientServiceType))))

			}
		}

		Connection.Write([]byte("OK\n"))
	}
}
