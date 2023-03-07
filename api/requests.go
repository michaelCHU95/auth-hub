package api

type LoginRequest struct {
	Email    string
	Password string
}

type SignupRequest struct {
	LoginRequest
	First_Name string
	Last_Name  string
}
