package interfaces

type IF01_01_IN struct {
	UserID   string `json:"user_id"`
	PassHash string `json:"pass_hash"`
}

type IF01_01_OUT struct {
	Result    string `json:"result"`
	SessionID string `json:"session_id"`
	Detail    string `json:"detail"`
}
