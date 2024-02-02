package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}
var ServerSetting = &Server{}

var cfg *ini.File
var DB *sql.DB

func Init() {
	var err error
	cfg, err = ini.Load("app.ini")

	if err != nil {
		log.Fatalf("Parse ini error")
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s", DatabaseSetting.User, DatabaseSetting.Password, DatabaseSetting.Host, DatabaseSetting.Name)

	DB, err = sql.Open(DatabaseSetting.Type, connection)

	if err != nil {
		log.Printf("%s", err.Error())
	}
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Printf("Cfg.MapTo %s err: %v", section, err)
	}
}
