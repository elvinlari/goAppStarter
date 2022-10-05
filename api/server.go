package api 

import (
   "fmt"
   "log"
   "os"

   "github.com/joho/godotenv"
   "github.com/mygoapp/api/controllers"
)

var server = controllers.Server{}

func Run() {
   var err error
   err = godotenv.Load()
   if err != nil {
   	log.Fatalf("Error getting env, not coming through %v", err)
   } else {
   	fmt.Println("We are getting the env values")
   }
   server.Initialize()
   port := os.Getenv("GO_PORT")
   server.Run(port)
}


