package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
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
	// accountRepository := repository.NewAccountRepositoryDB(db)

	// NOTE : ตรงนี้สามารถใช้เป็น mock repository แทนก็ได้
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	accountRepositoryMock := repository.NewAccountRepositoryMock()

	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)
	accountService := service.NewAccountService(accountRepositoryMock)
	accountHandler := handler.NewAccountHandler(accountService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	// NOTE : สามารถใช้ Regex ต่อท้ายเพื่อกำหนดให้เป็น Pattern ที่ถูกต้องได้เลย
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	logs.Info(fmt.Sprintf("Banking Service online at port %v", viper.GetInt("app.port")))
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
	logs.Info(fmt.Sprintf("Data Source Config = %v", dataSourceName))

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
