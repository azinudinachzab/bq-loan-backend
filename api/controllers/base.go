package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	//_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	//_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type DBConnection struct {
	Driver, User, Password, Port, Host, Name string
}

func (server *Server) Initialize(d DBConnection) error {

	if d.Driver != "mysql" {
		return errors.New("DB Dialect is not recognized")
	}

	DbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.Name)
	gDB, err := gorm.Open(d.Driver, DbURL)
	if err != nil {
		return err
	}

	server.DB = gDB
	//server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration
	server.Router = mux.NewRouter()
	server.initializeRoutes()
	return nil
}

func (server *Server) Run(addr string) {
	log.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
