package main

import (
	"gwi/handler"
	"gwi/repository"
	"gwi/repository/sqlite"
	"gwi/router"
	"log"
)

func initSQLiteDB(userDB string, assetDb string) (repository.UserRepo, repository.AssetRepo) {
	dbu, err := sqlite.UserDBConnect(userDB)
	if err != nil {
		log.Fatal(err)
	}
	err = dbu.UserDBInitialize(true)
	if err != nil {
		log.Fatal(err)
	}
	dba, err := sqlite.AssetDBConnect(assetDb)
	if err != nil {
		log.Fatal(err)
	}
	err = dba.AssetDBInitialize(true)
	if err != nil {
		log.Fatal(err)
	}
	return dbu, dba
}

func initDB(dbname string, userDB string, assetDb string) (repository.UserRepo, repository.AssetRepo) {
	switch dbname {
	case "sqlite":
		return initSQLiteDB(userDB, userDB)
	default:
		panic("Unknown db")
	}
}

func main() {
	r := router.New()
	v1 := r.Group("/api")
	dbu, dba := initDB("sqlite", "sqlite.db", "sqlite.db")
	h := handler.NewHandler(dbu, dba)
	defer h.Close()
	h.Register(v1)
	r.Logger.Fatal(r.Start(":1323"))
}
