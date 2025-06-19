package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Talent    Role = "talent"
	Employer  Role = "employer"
	Admin 	  Role = "admin"

)

type User struct {
	ID 				primitive.ObjectID 		`bson:"_id,omitempty" json:"id"`
	First_Name 		string 					`bson:"first_name" json:"first_name"`
	Last_Name 		string 					`bson:"last_name" json:"last_name"`
	Ful_Name 		string 					`bson:"full_name" json:"full_name"`
	Email 			string 					`bson:"email" json:"email"`
	Password 		string 					`bson:"password" json:"password"`
	Refresh_token 	string 					`bson:"refresh_token" json:"refresh_token"`
	Role     		Role            		`bson:"role" json:"role"`
	CreatedAt 		time.Time         		`bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt 		time.Time         		`bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

