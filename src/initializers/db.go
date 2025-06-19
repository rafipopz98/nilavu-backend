package initializers

import "github.com/nilavu-backend/src/db"

func InitDB (){
	db.DbConnect()
}