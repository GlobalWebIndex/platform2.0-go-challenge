package handler

import (
	"gwi/model"
	"gwi/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAsset is a handler function for handling the asset creation API call
func (h *Handler) CreateAsset(c echo.Context) error {
	asset := new(model.Asset)
	if err := c.Bind(asset); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	as, err := h.assetRepo.Create(asset)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrAlreadyExists {
			c.JSON(409, utils.NewError(utils.ErrAlreadyExists))
		}
		return err
	}
	return c.JSON(200, as)
}

// RetrieveAsset is a handler function for handling the asset retrieval API call
func (h *Handler) RetrieveAsset(c echo.Context) error {
	asset := new(model.Asset)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	asset.ID = uint(id)
	as, err := h.assetRepo.Retrieve(asset)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound {
			c.JSON(404, utils.NewError(utils.ErrNotFound))
		}
		return err
	}
	return c.JSON(200, as)
}

// UpdateAsset is a handler function for handling the asset description update API call
func (h *Handler) UpdateAsset(c echo.Context) error {
	asset := new(model.Asset)
	if err := c.Bind(asset); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	as, err := h.assetRepo.Update(asset)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound {
			c.JSON(404, utils.NewError(utils.ErrNotFound))
		}
		return err
	}
	return c.JSON(200, as)
}

// DeleteAsset is a handler function for handling the asset deletion API call
func (h *Handler) DeleteAsset(c echo.Context) error {
	asset := new(model.Asset)
	if err := c.Bind(asset); err != nil {
		c.Logger().Debug(err.Error())
		return err
	}
	as, err := h.assetRepo.Delete(asset)
	if err != nil {
		c.Logger().Debug(err.Error())
		if err == utils.ErrNotFound {
			c.JSON(404, utils.NewError(utils.ErrNotFound))
		}
		return err
	}
	return c.JSON(200, as)
}
