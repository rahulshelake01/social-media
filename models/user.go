package models

type UserRegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Gender    int8   `json:"gender"`
	DOB       string `json:"dob"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetails struct {
	UID       int64  `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Gender    int8   `json:"gender"`
	DOB       string `json:"dob"`
}

type UserRegisterResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type UserLoginResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	} `json:"data"`
}
