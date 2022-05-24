package main

import (
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/repositories/user_repo"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := .NewDb(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user_repo.NewUserRepo(db)

	// Tests
}
