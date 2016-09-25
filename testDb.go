package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "time"
import "math/rand"

type Tips struct {
	id      int
	content string
	ctime   string
	hash    string
	auth    string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/tips?charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	stmtIn, err := db.Prepare("insert into tips (content, hash, ctime, auth) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIn.Close()
	var hash string = "http"
	//	timeNow := time.Now().String()
	//	timeNow1 := timeNow[0:19]
	result, err := stmtIn.Exec("ssdsa", hash, time.Now().String()[0:19], "asdas") // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	lastid, err := result.LastInsertId()
	fmt.Println(lastid)
	stmtOut, err := db.Prepare("SELECT * FROM tips WHERE hash = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	rows, err := stmtOut.Query(hash) // WHERE number = 13
	var tt []*Tips
	for rows.Next() {
		t := new(Tips)
		err = rows.Scan(&t.id, &t.content, &t.hash, &t.ctime, &t.auth)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		tt = append(tt, t)
	}
	index := rand.New(rand.NewSource(lastid - 1))
	fmt.Println(tt[index.Intn(10)].ctime)
}
