package controller

import (
	"backend/interfaces"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService interfaces.UserService
}

func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{UserService: service}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	if err := uc.UserService.Create(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "create failed"})
	}
	return c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	fmt.Printf("RAW ID STRING: [%q]\n", idStr)

	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Atoi ERROR:", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	user, err := uc.UserService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	idStr := strings.TrimSpace(c.Param("id")) // ✅ bersihkan newline/spasi
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	user.ID = id // Set ID dari path ke objek user
	if err := uc.UserService.Update(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Update failed"})
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	fmt.Printf("RAW ID FROM PATH: [%q]\n", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	if err := uc.UserService.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "delete failed"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "user deleted"})
}
