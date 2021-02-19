package server

import "encoding/json"

type Error struct {
	Code  int    `json:"error_code"`
	Title string `json:"error_description"`
}

func WriteErrorJSON(code int, title string) string {
	e := Error{
		Code:  code,
		Title: title,
	}

	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}
