package student

type Student struct {
	FirstName string `json:"first_name" bson:"full_name" validate:"required"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func (s Student) GetFullName() string {
	return s.FirstName + " " + s.LastName
}

func (s *Student) SetEmail(e string) {
	s.Email = e
}
