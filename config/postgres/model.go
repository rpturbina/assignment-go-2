package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init config struct
type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `jsong:"database_name"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

type PostgresClient interface {
	GetClient() *gorm.DB
}

type PostgresClientImpl struct {
	cln    *gorm.DB
	config Config
}

func (p *PostgresClientImpl) GetClient() *gorm.DB {
	return p.cln
}

func NewPostgresConnection(config Config) PostgresClient {
	connectionString := fmt.Sprintf(`
	host=%s 
	port=%s 
	user=%s 
	password=%s 
	dbname=%s`,
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DatabaseName)

	// open gorm connection to postgres
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	// check error connection
	if err != nil {
		log.Panic("error connection to database :", err)
	}

	return &PostgresClientImpl{cln: db, config: config}
}
