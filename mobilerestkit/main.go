package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/alamyudi/echo-app/echokit"
	"github.com/alamyudi/echo-app/mobilerestkit/helpers"

	ctrl "github.com/alamyudi/echo-app/mobilerestkit/controllers"

	_ "github.com/go-sql-driver/mysql"
)

type (
	// CustomValidator for validate input
	CustomValidator struct {
		validator *validator.Validate
	}
)

// Validate to validate structure
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var conf helpers.Config

func initConfig() helpers.Config {
	var configFile string

	// Load configuration as a command line parameter
	flag.StringVar(&configFile, "config", "config/dev.conf", "Provide an absolute path to the configuration file")
	flag.Parse()
	conf = helpers.ReadConfig(configFile)
	return conf
}

func initRoutes(nk *echokit.EchoKit, cnf helpers.Config) *echo.Echo {
	// set model to controller
	ctrl.SetIAPManager(nk, cnf)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Pre(middleware.AddTrailingSlash())

	// echo configuration
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	// REST End poin
	apiGroup := e.Group("/v1")

	// Basic Auth
	apiGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == cnf.BaseAuth.User && password == cnf.BaseAuth.Password {
			return true, nil
		}
		return false, nil
	}))

	// CORS
	apiGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	apiGroup.GET("/product/", ctrl.GetProducts)
	apiGroup.GET("/product/:id/", ctrl.GetProductByID)
	apiGroup.PUT("/product/:id/", ctrl.PutProductByID)
	apiGroup.DELETE("/product/:id/", ctrl.DeleteProductByID)
	apiGroup.POST("/product/", ctrl.PostProduct)

	apiGroup.GET("/content/", ctrl.GetContents)
	apiGroup.GET("/content/:id/", ctrl.GetContentByID)
	apiGroup.PUT("/content/:id/", ctrl.PutContentByID)
	apiGroup.DELETE("/content/:id/", ctrl.DeleteContentByID)
	apiGroup.POST("/content/", ctrl.PostContent)

	if !conf.Debug {
		e.HideBanner = true
	}

	return e
}

func main() {
	// read config file
	conf := initConfig()

	mysqlConfig := conf.MysqlConfig

	msqlDSN := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)

	// connect to the mysql
	db, err := sql.Open("mysql", msqlDSN)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	navKit := echokit.Init(db)

	e := initRoutes(navKit, conf)
	e.Logger.Fatal(e.Start(":" + conf.Port))
}
