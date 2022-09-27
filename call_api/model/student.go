package model
// Student :
type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
	Age       int    `json:"age"`
}
