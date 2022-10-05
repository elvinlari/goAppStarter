package controllers

import (
  "net/http"
  "io/ioutil"
  "os"
  "bytes"
  
  "github.com/mygoapp/api/responses"
  "github.com/mygoapp/api/models"



  "encoding/json"
	"errors"
	"fmt"
	
	"log"
	"math"
	"net/http"
	
	"sync"
	"time"
	"unicode/utf8"

	// "bufio"
	
	"encoding/gob"
	"strings"

	"github.com/elvnneinlari/synqSMS/api/auth"
	"github.com/elvnneinlari/synqSMS/api/comms"
	"github.com/elvnneinlari/synqSMS/api/models"
	"github.com/elvnneinlari/synqSMS/api/responses"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"

)

type WorkRequest struct {
	Send      models.Send  `json:"send"`
}

func (server *Server) GetParams(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	uSr := keys.Get("user")
	from := keys.Get("source")
	to := keys.Get("dest")
	message := keys.Get("message")
	msId := keys.Get("msgID")
	meta := keys.Get("METADATA")

	send := &models.Send{
		User:    uSr,
		From:    from,
		To:      to,
		Message: message,
		MssId:   msId,
		MetaDt:  meta,
	}

	send.Prepare()
	err = send.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

// Sending params to function
	work := WorkRequest{
		Send: *send,
		User: usr}
	err = work.SendTo()
	if err != nil {
		log.Println("Error sending message to 'destination':", err)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path))
	responses.JSON(w, http.StatusOK, "Message sent to server")

}

func (server *Server) PostParams(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	send := models.Send{}
	err = json.Unmarshal(body, &send)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	send.Prepare()
	err = send.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}


  work := WorkRequest{
		Send: send}

  responseCode, err = work.SendTo()
	if err != nil {
		log.Println("Error sending message to 'destination': ", err)
	}

  log.Println("Successfully sent message to 'destination': ", responseCode)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path))
	responses.JSON(w, http.StatusOK, "Message sent to server")
}


func (work *WorkRequest) SendTo() (int, error) {
	byteKey3, err := json.Marshal(work)
	if err != nil {
		log.Fatalf("Error marshalling message: %s", err)
		return err
	}

  log.Println("Message bytes: ", byteKey3)
  respond := 0

// // env variables  
//   url := os.Getenv("URL")
// 	user := os.Getenv("USER")
// 	password := os.Getenv("PASS")

// // get request
//   req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return 0, err
// 	}
//   q := req.URL.Query()
// 	q.Add("param1", work.Send.user)
// 	q.Add("param2", work.Send.pass)
// 	req.URL.RawQuery = q.Encode()
// 	log.Println(req.URL.String())
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Println("errror:", err)
// 		return 0, err
// 	}
// 	defer res.Body.Close()

// // post request
//   reqbody := bytes.NewBuffer(byteKey3)
// 	req, err := http.NewRequest("POST", url, reqbody)
// 	if err != nil {
// 		log.Println("failed to create client request")
// 		return 0, err
// 	}
//   req.Header.Set("Content-Type", "application/json")
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Println("errror:", err)
// 		return 0, err
// 	}
// 	defer res.Body.Close()

//   resond = res.StatusCode

  return resond, nil
}