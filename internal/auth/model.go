package auth

type Session struct {
	SessionID string `json:"session_id"`
	UserID    int64  `json:"user_ID"`
	Username  string `json:"username"`
}
