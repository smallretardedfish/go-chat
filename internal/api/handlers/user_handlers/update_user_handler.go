package user_handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"net/http"
	"strconv"
)

type UserUpdateData struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

func updateUser(data UserUpdateData, usr user.User) *user.User {
	if data.Name != nil {
		usr.Name = *data.Name
	}
	if data.Email != nil {
		usr.Email = *data.Email
	}
	return &usr
}

func UpdateUserHandler(log configs.Logger, userService user.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		userToUpdate, err := userService.GetUser(int64(id))
		if err != nil {
			return err
		}
		if userToUpdate == nil {
			return errors.New("no such user")
		}

		updateData := &UserUpdateData{}
		if err := c.BodyParser(updateData); err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		updatedUser := updateUser(*updateData, *userToUpdate)
		_, err = userService.UpdateUser(*updatedUser)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}
}
