package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type goOne struct {
	Time int `json:"time"`
	From int `json:"from"`
	To   int `json:"to"`
}

type returnOne struct {
	Time int `json:"time"`
	From int `json:"from"`
}

type goFull struct {
	Fast [3]int `json:"fast"`
	From [2]int `json:"from"`
	To   [3]int `json:"to"`
}

type returnFull struct {
	Fast [2]int `json:"fast"`
	From [3]int `json:"from"`
}

// Directory for test //
var goDb, goErr = sql.Open("sqlite3", "/home/bustime/DataBase/go.db")
var returnDb, returnErr = sql.Open("sqlite3", "/home/bustime/DataBase/return.db")

func main() {
	r := mux.NewRouter()
	if goErr != nil {
		panic(goErr)
	} else if returnErr != nil {
		panic(returnErr)
	}

	r.HandleFunc("/api/goFast", getGoFastJson)
	r.HandleFunc("/api/go", getGoJson) //go?fromto=XX?id=XXと送ること//
	r.HandleFunc("/api/goFull", getGoFullJson)

	r.HandleFunc("/api/returnFast", getReturnFastJson)
	r.HandleFunc("/api/return", getReturnJson) //return?id=XXと送ること//
	r.HandleFunc("/api/returnFull", getReturnFullJson)

	http.ListenAndServe(":7650", r)
}

//最速を返す関数
func getGoFastJson(w http.ResponseWriter, r *http.Request) {
	w = setACAO(w)
	dayType := getDayType()
	nowTime := getNowTime()

	row := goDb.QueryRow(
		"SELECT * FROM GO WHERE time > ? AND dayType = ? ORDER BY time ASC LIMIT 1",
		nowTime,
		dayType,
	)

	ans := goOne{}
	var tmp int //読み捨て
	dataErr := row.Scan(&ans.Time, &ans.From, &ans.To, &tmp)
	if dataErr == sql.ErrNoRows {
		returnErrCode(w, 1)
		return
	} else if dataErr != nil {
		returnErrCode(w, 10)
		return
	}

	outJson, err := json.Marshal(&ans)

	if err != nil {
		returnErrCode(w, 10)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outJson))
	return
}

//特定種の最速を返す
func getGoJson(w http.ResponseWriter, r *http.Request) {
	w = setACAO(w)
	q := r.URL.Query()
	fromtoString := q["fromto"][0]
	idString := q["id"][0]
	fromto, fromtoOk := strconv.Atoi(fromtoString)
	id, idOk := strconv.Atoi(idString)

	if fromtoOk != nil || idOk != nil {
		returnErrCode(w, 10)
		return
	}

	timeArray := getGoTime(fromto, id)
	ans := goOne{}
	ans.Time = timeArray[0]
	ans.From = timeArray[1]
	ans.To = timeArray[2]

	outJson, err := json.Marshal(&ans)
	if err != nil {
		returnErrCode(w, 10)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outJson))
	return
}

//行き各種の最速をすべて返す
func getGoFullJson(w http.ResponseWriter, r *http.Request) {
	w = setACAO(w)
	ans := goFull{}
	var from [2]int
	var to [3]int
	var tmp [3]int

	for i := 0; i < 2; i++ {
		tmp = getGoTime(0, i)
		from[i] = tmp[0]
	}
	for i := 0; i < 3; i++ {
		tmp = getGoTime(1, i)
		to[i] = tmp[0]
	}

	ans.Fast = checkGoFastTime(from, to)
	ans.From = from
	ans.To = to

	outJson, err := json.Marshal(&ans)
	if err != nil {
		returnErrCode(w, 10)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outJson))
	return
}

//検索に使う部分、返すのは配列//
func getGoTime(fromto int, id int) [3]int {
	dayType := getDayType()
	nowTime := getNowTime()
	var ans [3]int
	var tmp int //読み捨て

	switch fromto {
	case 0:
		row := goDb.QueryRow(
			"SELECT * FROM GO WHERE time > ? AND stop = ? AND dayType = ? ORDER BY time ASC LIMIT 1",
			nowTime,
			id,
			dayType,
		)

		row.Scan(&ans[0], &ans[1], &ans[2], &tmp)
		break
	case 1:
		row := goDb.QueryRow(
			"SELECT * FROM GO WHERE time > ? AND endpoint = ? AND dayType = ? ORDER BY time ASC LIMIT 1",
			nowTime,
			id,
			dayType,
		)

		row.Scan(&ans[0], &ans[1], &ans[2], &tmp)
		break
	}

	return ans
}

func getReturnFastJson(w http.ResponseWriter, r *http.Request) {
	w = setACAO(w)
	dayType := getDayType()
	nowTime := getNowTime()

	row := returnDb.QueryRow(
		"SELECT * FROM RETURN WHERE time > ? AND dayType = ? ORDER BY time ASC LIMIT 1",
		nowTime,
		dayType,
	)

	ans := returnOne{}
	var tmp int //読み捨て
	dataErr := row.Scan(&ans.Time, &ans.From, &tmp)
	if dataErr == sql.ErrNoRows {
		returnErrCode(w, 1)
		return
	} else if dataErr != nil {
		returnErrCode(w, 10)
		return
	}

	outJson, err := json.Marshal(&ans)

	if err != nil {
		returnErrCode(w, 10)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outJson))
	return
}

//特定種の最速を返す
func getReturnJson(w http.ResponseWriter, r *http.Request) {
	w = setACAO(w)
	q := r.URL.Query()
	idString := q.Get("id")
	id, idOk := strconv.Atoi(idString)

	if idOk != nil {
		returnErrCode(w, 10)
		return
	}

	timeArray := getReturnTime(id)
	ans := returnOne{}
	ans.Time = timeArray[0]
	ans.From = timeArray[1]

	outJson, err := json.Marshal(&ans)
	if err != nil {
		returnErrCode(w, 10)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outJson))
	return
}

//帰り各種の最速をすべて返す
func getReturnFullJson(w http.ResponseWriter, r *http.Request) {
	w = setACAO(w)
	ans := returnFull{}
	var tmp [2]int
	var from [3]int

	for i := 0; i < 3; i++ {
		tmp = getReturnTime(i)
		from[i] = tmp[0]
	}

	fast := [2]int{from[0], 0}

	for i := 1; i < 3; i++ {
		if fast[0] > from[i] {
			fast[0] = from[i]
			fast[1] = i
		}
	}

	ans.Fast = fast
	ans.From = from

	outJson, err := json.Marshal(&ans)
	if err != nil {
		returnErrCode(w, 10)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outJson))
	return
}

func getReturnTime(id int) [2]int {
	dayType := getDayType()
	nowTime := getNowTime()
	row := returnDb.QueryRow(
		"SELECT * FROM RETURN WHERE time > ? AND stop = ? AND dayType = ? ORDER BY time ASC LIMIT 1",
		nowTime,
		id,
		dayType,
	)

	var ans [2]int
	var tmp int //読み捨て
	row.Scan(&ans[0], &ans[1], &tmp)
	return ans
}

func getDayType() int {
	var n int
	t := time.Now()
	if t.Weekday() == 0 {
		n = 2
	} else if t.Weekday() < 6 {
		n = 0
	} else {
		n = 1
	}
	return n
}

func getNowTime() int {
	t := time.Now()
	return t.Hour()*60 + t.Minute()
}

func checkGoFastTime(a [2]int, b [3]int) [3]int {
	fromFast := [2]int{a[0], 0}
	if a[1] < fromFast[0] {
		fromFast[0] = a[1]
		fromFast[1] = 1
	}

	toFast := [2]int{b[0], 0}
	for i := 1; i < 3; i++ {
		if b[i] < toFast[0] {
			toFast[0] = b[i]
			toFast[1] = i
		}
	}
	ans := [3]int{fromFast[0], fromFast[1], toFast[1]}
	return ans
}

func returnErrCode(w http.ResponseWriter, code int) {
	w = setACAO(w)
	w.Header().Set("Content-Type", "application/json")
	var out string = "{\"error_id\":" + strconv.Itoa(code) + "}"
	outJson, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, outJson)
	return
}

func setACAO(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}
