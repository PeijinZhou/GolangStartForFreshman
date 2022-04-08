package main_test

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestQueryFunctionWithExistID(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/query/12365885")
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
}

func TestQueryFunctionWithNoneExistID(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/query/456123856")
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

func TestQueryFunctionWithStringID(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/query/AreUFuckingSerious?")
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

func TestQueryFunctionWithEmptyID(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/query/ ?")
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

func TestQueryFunctionWithNegativeID(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/query/-114514")
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
