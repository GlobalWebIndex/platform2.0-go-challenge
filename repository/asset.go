package repository

import "gwi/model"

// AssetRepo is an interface that needs to be implemented by all the asset storage mechanisms
type AssetRepo interface {
	Create(as *model.Asset) (*model.Asset, error)
	Retrieve(as *model.Asset) (*model.Asset, error)
	Update(as *model.Asset) (*model.Asset, error)
	Delete(as *model.Asset) (*model.Asset, error)
	Close() error
}
