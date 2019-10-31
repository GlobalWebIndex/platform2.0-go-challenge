package users

// User structure
type User struct {
	userID    string
	favorites map[string]struct{}
}

// CreateNewUser creates a new user
func CreateNewUser(userID string) User {
	return User{
		userID:    userID,
		favorites: make(map[string]struct{})}
}

// GetUserFavorites returns the favorite list of a specific user
func (user *User) GetUserFavorites() map[string]struct{} {
	return user.favorites
}
