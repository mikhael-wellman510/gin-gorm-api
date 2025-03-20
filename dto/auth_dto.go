package dto

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)
