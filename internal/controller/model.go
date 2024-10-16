package controller

// retrieving acc from POST json
type accountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Balance  string `json:"balance"`
}

type AccountLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}