package routerfunc

import (
	"demo/mysqlconnection"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type People struct {
	User_id  string
	Username string
	Sex      string
	Email    string
} //对数据库内的struct进行了定义

type resErr struct {
	ErrCode int    `json:"err_code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type normalRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    People `json:"data"`
}

type deleteRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func SearchPeopleWithID(w http.ResponseWriter, r *http.Request) {
	db := mysqlconnection.DatebaseInit()
	vars := mux.Vars(r)
	id := vars["id"]
	queryThing := "SELECT * FROM person WHERE user_id =" + id
	var p People
	err := db.Get(&p, queryThing)
	fmt.Println(err)

	if err != nil {
		w.WriteHeader(500)
		errRes := resErr{1, "can not find this person", "None"}
		resError, err := json.Marshal(errRes)
		if err != nil {
			resError = nil
			fmt.Println("err:", err.Error())
		}
		fmt.Println(string(resError))
		w.Write(resError)
		return
	}
	//error throw

	w.WriteHeader(300)
	res := normalRes{0, "OK", p}
	corRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error")
	}

	w.Write(corRes)
} //通过id对人进行查找,返回的值以json形式存入body中,如果查无此人或者查找的id不符合格式,那么将返回一个500 error.

func InsertPeople(w http.ResponseWriter, r *http.Request) {

	temp := json.NewDecoder(r.Body)
	var postContent map[string]string
	temp.Decode(&postContent)
	fmt.Println()
	db := mysqlconnection.DatebaseInit()
	newPeopleSql := `INSERT INTO person(user_id ,username ,sex ,email)VALUES(?,?,?,?)`
	result, err := db.Exec(newPeopleSql, postContent["user_id"], postContent["username"], postContent["sex"], postContent["email"])
	if err != nil {
		w.WriteHeader(500)
		errRes := resErr{1, "can not insert this person", "None"}
		resErr, err := json.Marshal(errRes)
		if err != nil {
			resErr = nil
		}
		w.Write(resErr)
		return
	}
	id2, _ := result.LastInsertId()
	fmt.Println(id2)

	p := People{postContent["user_id"], postContent["username"], postContent["sex"], postContent["email"]}
	w.WriteHeader(300)
	res := normalRes{0, "OK", p}
	corRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error")
	}

	w.Write(corRes)
} //插入一个数据,如果数据不合法就返回一个error

func UpdatePeople(w http.ResponseWriter, r *http.Request) {

	temp := json.NewDecoder(r.Body)
	var postContent map[string]string
	temp.Decode(&postContent)
	vars := mux.Vars(r)
	id := vars["id"]
	idInJson, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error")
	}
	updataStuct := People{"", "", "", ""}
	json.Unmarshal(idInJson, updataStuct)

	if updataStuct.User_id != id {
		w.WriteHeader(500)
		fmt.Println(err)
		errRes := resErr{1, "you can not change user_id", "None"}
		resError, errNew := json.Marshal(errRes)
		if errNew != nil {
			resError = nil
		}
		w.Write(resError)
		return
	}

	db := mysqlconnection.DatebaseInit()
	newPeopleSql := `UPDATE person SET username =?, sex=?,email=? WHERE user_id=?`
	db.MustExec(newPeopleSql, postContent["username"], postContent["sex"], postContent["email"], id)

	queryThing := ("SELECT * FROM person WHERE user_id=" + id)
	var p People
	fmt.Println(queryThing)
	errorBro := db.Get(&p, queryThing)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(errorBro)
		errRes := resErr{1, "can not updata this person", "None"}
		resError, errNew := json.Marshal(errRes)
		if errNew != nil {
			resError = nil
		}
		w.Write(resError)
		return
	}

	w.WriteHeader(300)
	fmt.Println(p)
	res := normalRes{0, "OK", p}
	corRes, errorSec := json.Marshal(res)
	if errorSec != nil {
		fmt.Println("error")
		return
	}
	w.Write(corRes)
	return

} //更新一个人的数据,如果缺少数据或者数据不合法,就返回一个error,传入的新人数据必须包含username,sex和email

func DeletePeopleUseID(w http.ResponseWriter, r *http.Request) {
	db := mysqlconnection.DatebaseInit()
	vars := mux.Vars(r)
	id := vars["id"]
	deleteSQL := "DELETE FROM person WHERE user_id=?"
	queryThing := "SELECT * FROM person WHERE user_id=" + id
	var p People
	fmt.Println(queryThing)
	err := db.Get(&p, queryThing)
	if err != nil {
		w.WriteHeader(500)
		errRes := resErr{1, "please provide a valid person", "None"}
		resErr, err := json.Marshal(errRes)
		if err != nil {
			resErr = nil
		}
		w.Write(resErr)
		return
	} //error throw

	db.Exec(deleteSQL, id)
	w.WriteHeader(300)
	res := deleteRes{0, "OK", "None"}
	corRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error")
	}
	w.Write(corRes)

} //删除一个人的数据,如果缺少数据或者数据不合法,就返回一个error.
