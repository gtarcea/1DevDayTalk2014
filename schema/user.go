package schema

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}
