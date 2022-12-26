package main

import (
	Routers "go-elasticsearch/routes"
	"log"
	
)


func main(){


		router := Routers.SetupRouter()

		if err := router.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	}


