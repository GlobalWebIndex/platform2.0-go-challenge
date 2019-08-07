package handler

import "gwi/repository"

// Handler model containing the repositories for users and assets
type Handler struct {
	userRepo  repository.UserRepo
	assetRepo repository.AssetRepo
}

// NewHandler instantiates a handler object
func NewHandler(ur repository.UserRepo, ar repository.AssetRepo) *Handler {
	return &Handler{
		userRepo:  ur,
		assetRepo: ar,
	}
}

// Close closes the open handler resources
func (h *Handler) Close() error {
	err := h.userRepo.Close()
	if err != nil {
		panic(err)
	}
	err = h.assetRepo.Close()
	if err != nil {
		panic(err)
	}
	return nil
}
