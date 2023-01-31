package bo

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID        string `json:"user_id"`
	CharacterType int    `json:"character_type"`
	Token         string `json:"token"`
}

type ApplyRegisterRequest struct {
	Email string `json:"email" binding:"required"`
}

type ApplyRegisterResponse struct{}

type ActiveUserRequest struct {
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	CharacterType int    `json:"character_type" binding:"required"`
	ActiveCode    string `json:"active_code" binding:"required"`
}

type ActiveUserResponse struct{}

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordResponse struct{}

type ActiveResetPasswordRequest struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
}

type ActiveResetPasswordResponse struct{}

type ApplyChangeUserEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type ApplyChangeUserEmailResponse struct{}

type ActiveChangeUserEmailRequest struct {
	Email      string `json:"email" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
}

type ActiveChangeUserEmailResponse struct{}
