package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/api/handlers/auth_handlers"
	"github.com/smallretardedfish/go-chat/internal/api/handlers/room_handlers"
	"github.com/smallretardedfish/go-chat/internal/api/middleware"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
)

type Server interface {
	Start(serverAddress string) error
}

type HTTPServer struct {
	log         configs.Logger
	roomService chat.RoomService
	userService user.UserService
	authService user.AuthService
}

func NewHTTPServer(log configs.Logger, roomService chat.RoomService, userService user.UserService, authService user.AuthService) *HTTPServer {
	return &HTTPServer{log: log, roomService: roomService, userService: userService, authService: authService}
}

func (s *HTTPServer) Start(serverAddress string) error {
	app := fiber.New()
	publicGroup := app.Group("/api/v1")
	publicGroup.Post("/sign-up", auth_handlers.RegisterHandler(s.log, s.authService))
	publicGroup.Post("/sign-in", auth_handlers.SignInHandler(s.log, s.authService))

	privateGroup := app.Group("/api/v1", middleware.AuthMiddleware(s.log, s.userService))
	privateGroup.Get("/rooms", room_handlers.GetRoomsHandler(s.log, s.roomService))
	privateGroup.Get("/rooms/:id", room_handlers.GetRoomHandler(s.log, s.roomService))
	privateGroup.Post("/rooms", room_handlers.CreateRoomHandler(s.log, s.roomService))
	privateGroup.Post("rooms/add", room_handlers.AddUserToRoomHandler(s.log, s.roomService))
	privateGroup.Delete("rooms/add", room_handlers.RemoveUserFromRoomHandler(s.log, s.roomService))

	//privateGroup.Put("/rooms", nil)
	//privateGroup.Delete("/rooms/:id", nil)

	//TODO implement other endpoints
	// messages are processed by WebSockets
	//GetMessage are only needed here
	return app.Listen(serverAddress)
}
