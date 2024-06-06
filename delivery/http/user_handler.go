package http

import (
	"clean_arch/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(app *fiber.App, uc usecase.UserUseCase) {
	handler := &UserHandler{
		UserUseCase: uc,
	}

	app.Get("/users/:id", handler.GetUserByID)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	// Extract user ID from path
	idStr := c.Params("id")

	// Convert id from string to uint
	id, err := strconv.ParseUint(idStr, 10, 64) // Assuming you want uint64
	if err != nil {
		// Handle conversion error
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": "Invalid user ID format"})
	}

	// Call use case with the converted ID
	user, err := h.UserUseCase.GetUserByID(uint(id)) // Cast to uint if necessary, depending on the method signature
	if err != nil {
		// Handle error, e.g., user not found or other errors
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"error": "User not found"})
	}

	// Return successful response with user data
	return c.JSON(user)
}
