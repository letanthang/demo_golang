package api

import "fmt"

type ApiResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Body       []byte `json:"body"`
}

func (res *ApiResponse) ToString() string {
	return fmt.Sprintf(`status:%v status_code:%v reponse:%v`, res.Status, res.StatusCode, string(res.Body))
}
