package request

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
