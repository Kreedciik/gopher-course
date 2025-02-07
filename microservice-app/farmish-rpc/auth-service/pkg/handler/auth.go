package handler

import (
	"farmish/model"
)

func (h *Handler) SignUp(newUser model.CreateUserDTO, response *string) error {
	if err := h.services.SignUp(newUser); err != nil {
		return err
	}
	*response = "user registered successfully"
	return nil
}

func (h *Handler) SignIn(credentials model.SignInDTO, response *model.Token) error {
	accessToken, err := h.services.SignIn(credentials)
	if err != nil {
		return err
	}
	response.AccessToken = accessToken
	return nil
}
