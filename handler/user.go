package handler

import (
	"gwi/model"
	"gwi/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUser is a handler function for handling the user creation API call
func (h *Handler) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	us, err := h.userRepo.Create(user)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrAlreadyExists {
			c.JSON(409, utils.NewError(utils.ErrAlreadyExists))
		}
		return err
	}
	return c.JSON(200, us)
}

// RetrieveUser is a handler function for handling the user retrieval API call
func (h *Handler) RetrieveUser(c echo.Context) error {
	user := new(model.User)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	user.ID = uint(id)
	us, err := h.userRepo.Retrieve(user)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound {
			c.JSON(404, utils.NewError(utils.ErrNotFound))
		}
		return err
	}
	return c.JSON(200, us)
}

// DeleteUser is a handler function for handling the user deletion API call
func (h *Handler) DeleteUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	us, err := h.userRepo.Delete(user)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound {
			c.JSON(404, utils.NewError(utils.ErrNotFound))
		}
		return err
	}
	return c.JSON(200, us)
}

// GetAllUsers is a handler function for handling the retrieval of all users
func (h *Handler) GetAllUsers(c echo.Context) error {
	uss, err := h.userRepo.GetAllUsers()
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound {
			c.JSON(404, utils.NewError(utils.ErrNotFound))
		}
		return err
	}
	return c.JSON(200, uss)
}

// GetUserFavourites is a handler function for handling the retrieval of a specific user's paged favourites
func (h *Handler) GetUserFavourites(c echo.Context) error {
	req := new(UserFavouritesPaged)
	if err := c.Bind(req); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	if req.PageSize == 0 || req.PageSize > 100 {
		req.PageSize = 100
	}
	us, err := h.userRepo.GetUserFavourites(&req.User, &req.NextToken, req.PageSize, h.assetRepo)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound {
			c.JSON(404, utils.NewError(utils.ErrNotFound))
		}
		return err
	}
	req.User = *us
	return c.JSON(200, req)
}

// AddUserFavourite is a handler function for handling the addition of an asset as a Favorite to a user
func (h *Handler) AddUserFavourite(c echo.Context) error {
	req := new(CDFavourite)
	if err := c.Bind(req); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	us, err := h.userRepo.AddUserFavourite(&req.User, h.assetRepo, &req.Asset)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound || err == utils.ErrAlreadyExists {
			c.JSON(404, utils.NewError(err))
		}
		return err
	}
	return c.JSON(200, us)
}

// RemoveUserFavourite is a handler function for handling the removal of an asset from the Favorites of a user
func (h *Handler) RemoveUserFavourite(c echo.Context) error {
	req := new(CDFavourite)
	if err := c.Bind(req); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	us, err := h.userRepo.RemoveUserFavourite(&req.User, h.assetRepo, &req.Asset)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound || err == utils.ErrAlreadyExists {
			c.JSON(404, utils.NewError(err))
		}
		return err
	}
	return c.JSON(200, us)
}
