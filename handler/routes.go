package handler

import (
	"github.com/labstack/echo/v4"
)

// Register the api routes
func (h *Handler) Register(v1 *echo.Group) {
	users := v1.Group("/users")
	users.GET("", h.GetAllUsers)
	users.POST("/create", h.CreateUser)
	// users.POST("/get", h.RetrieveUser)
	users.GET("/get/:id", h.RetrieveUser)
	users.DELETE("", h.DeleteUser)
	users.POST("/favourites", h.GetUserFavourites)
	users.PUT("/favourites", h.AddUserFavourite)
	users.DELETE("/favourites", h.RemoveUserFavourite)

	assets := v1.Group("/assets")
	assets.POST("/create", h.CreateAsset)
	assets.GET("/get/:id", h.RetrieveAsset)
	assets.PUT("", h.UpdateAsset)
	assets.DELETE("", h.DeleteAsset)

}
