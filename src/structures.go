package main

type JSONResponse struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

type RequestData struct {
	Date    string `json:"date"`
	Url     string `json:"url"`
	Method  string `json:"method"`
	Status  int    `json:"status"`
	User    string `json:"user"`
	Headers string `json:"headers"`
	Body    string `json:"body"`
	Comment string `json:"comment"`
}
