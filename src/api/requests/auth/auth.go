package request_auth

import "github.com/nilavu-backend/src/db/model"

type SignUP_Request struct {
	First_Name string
	Last_Name string
    Email string
	Password string
	Role model.Role
}

type Login_Request struct {
    Email string
	Password string
}