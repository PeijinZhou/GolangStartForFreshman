package main_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
)

func TestDeleteFunctionWithExistUser(t *testing.T) {
	id := rand.Intn(10000)
	postItem := People{string(rune(id)), "Tele", "Homo", "pxx@126.com"}
	reqJson, _ := json.Marshal(postItem)
	http.Post("http://localhost:3000/insert", "application/json", bytes.NewBuffer(reqJson))
	http.NewRequest("DELETE", "http://localhost:3000/delete/"+string(rune(id)), nil)

	resp, err := http.Get("http://localhost:3000/query/" + string(rune(id)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err := errors.New("something wrong with the delete function")
		if err != nil {
			return
		}
	}
}
