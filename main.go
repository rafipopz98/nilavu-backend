package main

import (
	"github.com/nilavu-backend/src/api/routes"
	"github.com/nilavu-backend/src/initializers"
)


func init(){
	initializers.LoadEnvVaribles();
	initializers.InitDB()
}

func main(){
	router := routes.InitRouters();
	
	router.Run()
}