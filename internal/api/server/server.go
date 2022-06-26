package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/internal/api/handlers/auth_handlers"
	"github.com/smallretardedfish/go-chat/internal/api/handlers/room_handlers"
	"github.com/smallretardedfish/go-chat/internal/api/handlers/user_handlers"
	"github.com/smallretardedfish/go-chat/internal/api/handlers/ws_handler"
	"github.com/smallretardedfish/go-chat/internal/api/middleware"
	"github.com/smallretardedfish/go-chat/internal/connector"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/logging"
)

type Server interface {
	Start(serverAddress string) error
}

type HTTPServer struct {
	log         logging.Logger
	roomService chat.RoomService
	userService user.UserService
	authService user.AuthService
	connector   connector.Connector
}

func NewHTTPServer(log logging.Logger, roomService chat.RoomService,
	userService user.UserService, authService user.AuthService,
	connector connector.Connector) *HTTPServer {
	return &HTTPServer{
		log:         log,
		roomService: roomService,
		userService: userService,
		authService: authService,
		connector:   connector,
	}
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
	privateGroup.Delete("rooms/:room_id/remove/:user_id",
		room_handlers.RemoveCurrentUserFromRoomHandler(
			s.log,
			s.roomService))

	privateGroup.Post("rooms/remove", room_handlers.RemoveUsersFromRoomHandler(s.log, s.roomService))
	privateGroup.Put("/rooms/:id", room_handlers.UpdateRoomHandler(s.log, s.roomService))
	privateGroup.Delete("/rooms/:id", room_handlers.DeleteRoomHandler(s.log, s.roomService))

	privateGroup.Get("/users", user_handlers.GetUsersHandler(s.log, s.userService))
	privateGroup.Get("/users/:id", user_handlers.GetUserHandler(s.log, s.userService))
	privateGroup.Patch("/users", user_handlers.UpdateUserHandler(s.log, s.userService))
	privateGroup.Delete("users/:id", user_handlers.DeleteUserHandler(s.log, s.userService))

	app.All("/ws", middleware.AuthMiddleware(s.log, s.userService), middleware.WsMiddleware, ws_handler.WsHandler(s.log, s.connector))

	return app.Listen(serverAddress)
}
