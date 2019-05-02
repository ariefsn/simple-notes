package core

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"
	_ "git.eaciitapp.com/sebar/dbflex/drivers/mongodb"
)

var dbConnection dbflex.IConnection

// Database = Get Database
func Database() dbflex.IConnection {
	if dbConnection != nil {
		return dbConnection
	}

	log.Println("===> Connecting to database.")

	// Direct
	// dbConnection, err := NewDatabaseConnection()
	// if err != nil {
	// 	log.Println("===> Failed to connect to database. Error : ", err.Error())
	// 	os.Exit(0)
	// }

	// With Pool
	pooling := PoolDatabaseConnection()
	// defer pooling.Close()

	pool, err := pooling.Get()

	if err != nil {
		log.Println("===> Failed to get pool connection. Error : ", err.Error())
		os.Exit(0)
	}

	defer pool.Release()

	toolkit.Println("Pool Size", pooling.Size())
	toolkit.Println("Pool Created", pooling.Count())
	toolkit.Println("Pool Free", pooling.FreeCount())

	dbConnection = pool.Connection()

	return dbConnection
}

// NewDatabaseConnection = New Database Connection
func NewDatabaseConnection() (dbflex.IConnection, error) {
	dbHost := Configuration().GetString("database.host")
	dbUser := Configuration().GetString("database.username")
	dbPass := Configuration().GetString("database.password")
	dbName := Configuration().GetString("database.name")

	connectionString := "mongodb://"
	if dbHost != "" && dbName != "" && dbUser != "" && dbPass != "" {
		connectionString = fmt.Sprintf("%s%s:%s@%s/%s", connectionString, dbUser, dbPass, dbHost, dbName)
	} else if dbHost != "" && dbName != "" && dbUser != "" {
		connectionString = fmt.Sprintf("%s%s@%s/%s", connectionString, dbUser, dbHost, dbName)
	} else if dbHost != "" && dbName != "" {
		connectionString = fmt.Sprintf("%s%s/%s", connectionString, dbHost, dbName)
	} else {
		return nil, fmt.Errorf("Unable to connect to the database server. Please check the configuration")
	}

	conn, err := dbflex.NewConnectionFromURI(connectionString, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	err = conn.Connect()
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	return conn, nil
}

// PoolDatabaseConnection = Create connection pool
func PoolDatabaseConnection() *dbflex.DbPooling {
	pooling := dbflex.NewDbPooling(10, NewDatabaseConnection)

	pooling.Timeout = 30 * time.Second

	// defer pooling.Close()

	return pooling
}
