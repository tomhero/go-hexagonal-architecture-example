package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	// _ = customerRepository

	// NOTE : ตรงนี้สามารถใช้เป็น mock repository แทนก็ได้
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	// _ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	// NOTE : สามารถใช้ Regex ต่อท้ายเพื่อกำหนดให้เป็น Pattern ที่ถูกต้องได้เลย
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	log.Printf("Banking Service online at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // NOTE : หรืออาจจะเป็น json ก็ได้
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	// NOTE : สรมารถรับ ENV เช่น APP_PORT=5000 ได้เลย
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
	log.Printf("Data Source Config = %v", dataSourceName)

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
