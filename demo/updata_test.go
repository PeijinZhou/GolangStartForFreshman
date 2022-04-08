package main_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"testing"
)

func TestUpdataFunctionWithExistUser(t *testing.T) {
	id := rand.Intn(10000)
	postItem := People{string(rune(id)), "Tele", "Homo", "pxx@126.com"}
	updataItem := People{string(rune(id)), "Tele", "male", "pxx@126.com"}
	reqJson, _ := json.Marshal(postItem)
	newReqJson, _ := json.Marshal(updataItem)
	http.Post("http://localhost:3000/insert", "application/json", bytes.NewBuffer(reqJson))

	firResp, err := http.Get("http://localhost:3000/query/" + string(rune(id)))
	if err != nil {
		log.Fatal(err)
	}

	http.Post("http://localhost:3000/updata/"+string(rune(id)), "application/json", bytes.NewBuffer(newReqJson))

	SecResp, err := http.Get("http://localhost:3000/query/" + string(rune(id)))
	if err != nil {
		log.Fatal(err)
	}

	firResCom, err := io.ReadAll(firResp.Body)
	secResCom, err := io.ReadAll(SecResp.Body)
	res := bytes.Compare(firResCom, secResCom)
	if res == 0 {
		err := errors.New("update filled")
		if err != nil {
			return
		}
	}
}
