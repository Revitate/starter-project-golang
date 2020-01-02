package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"log"
	"starter-project/connector"
	"starter-project/controller"
	"starter-project/repository"
	"starter-project/router"

	_ "starter-project/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time: ${time_rfc3339}, remote_ip: ${remote_ip}, method: ${method}, uri: ${uri}, status: ${status}, latency: ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	binding(e)

	if err := e.Start(":5000"); err != nil {
		log.Fatal(err)
	}
}

func binding(e *echo.Echo) {
	mockDb := connector.NewInMemoryKeyValue()

	mongoConnectorOption := connector.MongoConnectorOption{
		Database: "john",
		URI:      "mongodb://john:john@localhost:27017/admin",
	}
	mongoDB, err := connector.NewMongoConnector(mongoConnectorOption)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = mongoDB.Client().Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	sqlConnectorOption := connector.SqlConnectorOption{
		DBName:   "postgres",
		Type:     "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "john",
		Password: "john",
		SSLMode:  "disable",
	}
	sqlDb, err := connector.NewSqlConnector(sqlConnectorOption)
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDb.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(mockDb)
	userController := controller.NewUserController(userRepo)
	router.NewUserRouter(e, userController)

	pingService := controller.NewPingService()
	router.NewPingRouter(e, pingService)
}
