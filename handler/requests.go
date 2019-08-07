package handler

import "gwi/model"

// CDFavourite is a helper struct used for create or remove favourite requests
type CDFavourite struct {
	User  model.User  `json:"user"`
	Asset model.Asset `json:"asset"`
}

// UserFavouritesPaged is a helper struct used for facilitating paged GetFavourite requests
type UserFavouritesPaged struct {
	User      model.User `json:"user"`
	NextToken uint       `json:"next_token"`
	PageSize  uint       `json:"page_size"`
}
