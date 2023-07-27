package main

import (
	"codebase/config"
	"codebase/handler"
	"codebase/usecase"

	"github.com/hfs1988/sdk/adapters/db"
	"github.com/hfs1988/sdk/adapters/rest"
	"github.com/hfs1988/sdk/client"

	_ "github.com/lib/pq"
)

func main() {
	config := config.GetConfig("./app.yaml")
	var sqlDB client.SQLDB = db.GetPostgresInstance(config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DBName)
	postgresDB, err := sqlDB.Connect()
	if err != nil {
		panic(err)
	}

	reqHandler := handler.UserHandler{
		UserUsecase: usecase.GetUserUsecase(postgresDB, sqlDB),
	}

	var router client.Router = rest.GetRouterInstance(3000)
	router.Post("/user/create", reqHandler.CreateUser)
	router.Post("/user/update", reqHandler.UpdateUser)
	router.Post("/user/delete", reqHandler.DeleteUser)
	router.Get("/user/getall", reqHandler.GetAllUser)
	router.Get("/user/get", reqHandler.GetByID)
	router.ListenAndServe()
}
