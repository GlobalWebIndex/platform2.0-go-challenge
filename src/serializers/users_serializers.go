package serializers

import "gwi-challenge/data/models"

type LoginUserSerializer struct {
	*models.User
	AccessToken string
}

type LoginUserResponse struct {
	FullName    string `json:"full_name"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func NewLoginUserSerializer() LoginUserSerializer {
	return LoginUserSerializer{}
}

func (serializer *LoginUserSerializer) Response() *LoginUserResponse {
	response := LoginUserResponse{
		FullName:    serializer.FullName,
		Username:    serializer.Username,
		AccessToken: serializer.AccessToken,
	}
	return &response
}
