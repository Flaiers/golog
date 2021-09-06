package config

type JSONResponse struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

type JSONResponseInt struct {
	Error bool `json:"error"`
	Data  int  `json:"data"`
}

type RequestData struct {
	Date    string `json:"date"`
	Url     string `json:"url"`
	Method  string `json:"method"`
	Status  int    `json:"status"`
	UserID  int    `json:"user_id"`
	Body    string `json:"body"`
	Comment string `json:"comment"`
}
