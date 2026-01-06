package repo

type UserRepo struct {
	// Add fields if necessary
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserByID() string {
	// Dummy implementation
	return "User data"
}
