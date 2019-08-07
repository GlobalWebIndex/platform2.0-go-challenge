package repository

import "gwi/model"

// UserRepo is an interface that needs to be implemented by all the user storage mechanisms
type UserRepo interface {
	Create(us *model.User) (*model.User, error)
	Retrieve(us *model.User) (*model.User, error)
	Delete(us *model.User) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	GetUserFavourites(us *model.User, startIndex *uint, pageSize uint, ar AssetRepo) (*model.User, error)
	AddUserFavourite(us *model.User, ar AssetRepo, as *model.Asset) (*model.User, error)
	RemoveUserFavourite(us *model.User, ar AssetRepo, as *model.Asset) (*model.User, error)
	Close() error
	// DBConnect(atts ...string) (*interface{}, error)
}
