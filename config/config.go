package config

import (
	"fmt"
	"os"

	userRepo "github.com/Jiran03/mailku/user/repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
}

var Conf Config

func Init() {
	Conf = Config{
		DBNAME: os.Getenv("DB_NAME"),
		DBUSER: os.Getenv("DB_USER"),
		DBPASS: os.Getenv("DB_PASS"),
		DBHOST: os.Getenv("DB_HOST"),
		DBPORT: os.Getenv("DB_PORT"),
	}
}

func DBInit() (DB *gorm.DB) {

	//mysql
	if os.Getenv("DB") == "mysql" {
		DB, _ = gorm.Open(
			mysql.Open(
				fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
					Conf.DBUSER,
					Conf.DBPASS,
					Conf.DBHOST,
					Conf.DBPORT,
					Conf.DBNAME,
				),
			),
		)

		return
	}

	//postgre
	dbURL := fmt.Sprintf(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			Conf.DBHOST,
			Conf.DBUSER,
			Conf.DBPASS,
			Conf.DBNAME,
			Conf.DBPORT,
		),
	)

	DB, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURL, // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,  // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	return
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&userRepo.User{},
	)
}
