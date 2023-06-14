package types

type Activities struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int32   `json:"participants"`
	Price         float32 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float32 `json:"accessibility"`
}

type ErrorResp struct {
	Message string `json:"message"`
}

type Response struct {
	ID            int32
	Activity      string
	Category      string
	Participants  int32
	Price         float32
	Link          string
	Key           string
	Accessibility float32
}

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var JwtKey = []byte("dfjkgldskf3234sd")
