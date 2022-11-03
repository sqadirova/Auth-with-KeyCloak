package auth

type SignInReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResp struct {
	Token string `json:"token" validate:"required"`
}

type SignOutDTO struct {
	UserID string `json:"user_id" validate:"required"`
}

type RolesReqDTO struct {
	RoleID string `json:"role_id" validate:"required"`
}
