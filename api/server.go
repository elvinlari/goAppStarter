package api 

import (
   "fmt"
   "log"
   "os"

   "github.com/joho/godotenv"
   "github.com/elvnneinlari/vodatz/api/controllers"
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
   server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
   port := os.Getenv("GO_PORT")
   server.Run(port)
}


