package main

import (
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/api/server"
	"github.com/smallretardedfish/go-chat/internal/connector"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"github.com/smallretardedfish/go-chat/logging"
)

func main() {

	cfg, err := configs.NewConfig()
	if err != nil {
		panic(err)
	}
	log := logging.NewLogger()

	db, err := configs.NewDB(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user_repo.NewUserRepo(db)
	roomRepo := room_repo.NewRoomRepo(db)
	//messageRepo := message_repo.NewMessageRepo(db)
	credsRepo := user_cred_repo.NewUserCredentialsRepo(db)

	//messageSvc := chat.NewMessageServiceImpl(messageRepo)
	roomSvc := chat.NewRoomServiceImpl(roomRepo)
	authSvc := user.NewAuthServiceImpl(credsRepo, userRepo)
	userSvc := user.NewUserServiceImpl(userRepo)

	conn := connector.NewConnector(log)

	httpServer := server.NewHTTPServer(log, roomSvc, userSvc, authSvc, conn)
	if err := httpServer.Start(cfg.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
