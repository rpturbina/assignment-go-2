package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/assigment-go-2/config/gin"
	"github.com/rpturbina/assigment-go-2/config/postgres"
	"github.com/rpturbina/assigment-go-2/pkg/domain/message"
	userrepo "github.com/rpturbina/assigment-go-2/pkg/repository/user"
	userhandler "github.com/rpturbina/assigment-go-2/pkg/server/http/handler/user"
	userusecase "github.com/rpturbina/assigment-go-2/pkg/usecase/user"

	orderrepo "github.com/rpturbina/assigment-go-2/pkg/repository/order"
	orderhandler "github.com/rpturbina/assigment-go-2/pkg/server/http/handler/order"
	orderusecase "github.com/rpturbina/assigment-go-2/pkg/usecase/order"

	router "github.com/rpturbina/assigment-go-2/pkg/server/http/router/v1"
)

// ASSESSMENT
// buat API
// - get user
// sebelum membuat order
//	- table dengan relasi order -> user (FOREIGN KEY)
// 			ref:https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-create-table/
// 	- code base untuk repo, usecase, dll
// - create order
// - get order by user

func main() {
	// generate postgres config and connect to postgres
	// this postgres client, will be used in repository layer
	postgresCln := postgres.NewPostgresConnection(postgres.Config{Host: "localhost", Port: "5432", User: "postgres", Password: "mysecretpassword", DatabaseName: "postgres"})

	// gin engine
	ginEngine := engine.NewGinHttp(engine.Config{Port: ":8080"})
	ginEngine.GetGin().Use(gin.Recovery(), gin.Logger())

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		resMap := map[string]any{
			"code":       0,
			"message":    "server up and running",
			"start_time": startTime,
		}

		var respStruct message.Response

		resByte, err := json.Marshal(resMap)
		if err != nil {
			log.Panic(err)
		}

		err = json.Unmarshal(resByte, &respStruct)
		if err != nil {
			log.Panic(err)
		}

		ctx.JSON(http.StatusOK, respStruct)
	})

	// generate user repository
	userRepo := userrepo.NewUserRepo(postgresCln)

	// initiate use case
	userUsecase := userusecase.NewUserUsecase(userRepo)

	// initiate handler
	userHandler := userhandler.NewUserHandler(userUsecase)

	// initiate router
	router.NewUserRouter(ginEngine, userHandler).Routers()

	// generate order repository
	orderRepo := orderrepo.NewOrderRepo(postgresCln)

	// initiate use case
	orderUsecase := orderusecase.NewOrderUsecase(orderRepo)

	// initiate handler
	orderHandler := orderhandler.NewOrderHandler(orderUsecase)

	// initiate router
	router.NewOrderRouter(ginEngine, orderHandler).Routers()

	ginEngine.Serve()
}
