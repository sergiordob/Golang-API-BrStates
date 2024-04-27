package database

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)



var DatabaseConnection *gorm.DB

func OpenDB() {
	var err error
	configurationString := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable"

	DatabaseConnection, err = gorm.Open(postgres.Open(configurationString))

	if err != nil {
		log.Println("Error: ", err)
	} else {
		log.Println("Database opened!")
	}

	sqlDB, err := DatabaseConnection.DB()
	if err != nil {
		log.Println("Erro ao obter a instância do banco de dados: ", err)
	} else {
		err = sqlDB.Ping()
		if err != nil {
			log.Println("Erro ao conectar ao banco de dados: ", err)
		} else {
			log.Println("Conexão com o banco de dados bem sucedida!")
		}
	}
}


