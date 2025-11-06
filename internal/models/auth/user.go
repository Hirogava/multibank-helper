package auth

type BankUser struct {
	BankUserID string `json:"bank_user_id" binding:"required"`
}

type User struct {
	ID string `json:"id" binding:"required"`
	BankUserID string `json:"bank_user_id" binding:"required"`
	Token Tokens `json:"token" binding:"required"`
}
