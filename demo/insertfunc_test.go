package main_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"testing"
)

type People struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
}

type Trash struct {
	TrashOne string `json:"IamATrash"`
	TrashTwo string `json:"IamAnotherTrash"`
}

func TestInsertFunctionWithExistID(t *testing.T) {
	postItem := People{"12365885", "Tele", "Homo", "pxx@126.com"}
	reqJson, _ := json.Marshal(postItem)
	resp, err := http.Post("http://localhost:3000/insert", "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err := errors.New("you should get a error")
		if err != nil {
			return
		}
	}
}

func TestInsertFunctionWithLetterID(t *testing.T) {
	postItem := People{"hello-world", "Tell", "Homo", "pxx@126.com"}
	reqJson, _ := json.Marshal(postItem)
	resp, err := http.Post("http://localhost:3000/insert", "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err := errors.New("you should get a error")
		if err != nil {
			return
		}
	}
}

func TestInsertFunctionWithNoID(t *testing.T) {
	postItem := People{" ", "Tell", "Homo", "pxx@126.com"}
	reqJson, _ := json.Marshal(postItem)
	resp, err := http.Post("http://localhost:3000/insert", "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err := errors.New("you should get a error")
		if err != nil {
			return
		}
	}
}

func TestInsertFunctionWithTrashMessage(t *testing.T) {
	postItem := Trash{"here is trash", "here is another trash"}
	reqJson, _ := json.Marshal(postItem)
	resp, err := http.Post("http://localhost:3000/insert", "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err := errors.New("you should get a error")
		if err != nil {
			return
		}
	}
}

func TestInsertFunctionWithRandomID(t *testing.T) {
	id := rand.Intn(10000)
	postItem := People{string(rune(id)), "Tele", "Homo", "pxx@126.com"}
	reqJson, _ := json.Marshal(postItem)
	resp, err := http.Post("http://localhost:3000/insert", "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode <= 200 && resp.StatusCode >= 299 {
		err := errors.New("you should get a valid replay")
		if err != nil {
			return
		}
	}
	getresp, err := http.Get("http://localhost:3000/query/" + string(rune(id)))
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(getresp.Body)
	// b, err := util.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}
	res := bytes.Compare(b, reqJson)

	if res != 0 {
		err := errors.New("you should get a same replay")
		if err != nil {
			return
		}
	}

}
