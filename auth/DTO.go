package auth

type SignInReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResp struct {
	AccessToken  string `json:"token" validate:"required"`
	RefreshToken string `json:"refresh_token"`
}

type SignOutDTO struct {
	UserID string `json:"user_id" validate:"required"`
}

type RolesReqDTO struct {
	RoleID string `json:"role_id" validate:"required"`
}

type RolesResp struct {
	Id       string `json:"id"`
	RoleType string `json:"role_type"`
}

type UserMeResp struct {
	Id        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Role      RolesResp `json:"role"`
	Username  string    `json:"username"`
}
