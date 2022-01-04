package main

import (
	"fmt"
	"gofiber/handler"
	"gofiber/repository"
	"gofiber/service"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()
	_ = db

	// NOTE : Using Fiber Web Framework
	app := fiber.New()

	// NOTE : Official Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})

	setupRouter(app, db)

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))

	_ = nativeMux()
}

func setupRouter(app *fiber.App, db *sqlx.DB) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		// NOTE : Inline handler function
		return c.SendString("Hello World!")
	})

	app.Get("/hello/:first_name/:last_name?", func(c *fiber.Ctx) error {
		// NOTE : Get params with optional + default
		fName := c.Params("first_name", "default_first_name")
		lName := c.Params("last_name", "default_last_name")
		return c.SendString(fmt.Sprintf("fn = %v | ln = %v", fName, lName))
	})

	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	app.Post("/signup", userHandler.SignUpHandler)
	app.Post("/login", userHandler.SignInHandler)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // NOTE : หรืออาจจะเป็น json ก็ได้
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// NOTE : ดึงจาก ENV มาทับก็ได้นะ !!!
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dataSourceName := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dataSourceName)
	if err != nil {
		panic(err)
	}

	// NOTE : Set Database config on the fly
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func nativeMux() error {
	nativeApp := http.NewServeMux()

	nativeApp.HandleFunc("/native/mux/{id}", func(rw http.ResponseWriter, r *http.Request) {
		// NOTE : Manually Handle stuff....
		log.Println(r.Method)
	})

	// http.ListenAndServe(":8002", nativeApp)

	return nil
}
