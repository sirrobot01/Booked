package user

import "errors"

type UserRegister struct {
	Username        string `json:"username"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (u *UserRegister) Validate() error {
	if u.Password != u.ConfirmPassword {
		return errors.New("password and confirm password must be the same")
	}
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	return nil
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLogin) Validate() error {
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	return nil
}
