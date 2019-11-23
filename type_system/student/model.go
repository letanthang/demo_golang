package student

type Student struct {
	FirstName string `json:"first_name" bson:"full_name" validate:"required"`
	LastName  string
	Age       int
	Email     string
}

func (s Student) GetFullName() string {
	return s.FirstName + " " + s.LastName
}

func (s *Student) SetEmail(e string) {
	s.Email = e
}
