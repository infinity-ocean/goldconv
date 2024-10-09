package controller

// import "time"

// type account struct {
// 	ID        int       `json:"id"`
// 	Username  string    `json:"username"`
// 	Email     string    `json:"email"`
// 	Password  string    `json:"-"`
// 	Number    int64     `json:"number"`
// 	CreatedAt time.Time `json:"createdAt"`
// 	Balance   int64     `json:"balance"`
// }

// retrieving acc from POST json
type accountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Balance  string `json:"balance"`
}
