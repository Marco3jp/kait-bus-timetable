package main

import (
  "fmt"
  "html"
  "log"
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/mattn/go-sqlite3"
)

type gp struct{
    Time int `json:"time"`
    From int `json:"from"`
    To int `json:"to"`
}

type return struct{
    Time int `json:"fast"`
    From int `json:"from"`
}

type goFull struct {
	Fast []int `json:"fast"`
	From []int `json:"from"`
	To   []int `json:"to"`
}

type returnFull struct {
	Fast []int `json:"fast"`
	From []int `json:"from"`
}

func main() {
    goDb, goErr := sql.Open("sqlite3", "./go.db");
    returnDb, returnErr := sql.Open("sqlite3", "./retrun.db");
    r := mux.NewRouter()

    if goErr != nil || returnErr != nil{
      panic(err)
    }

    r.HandleFunc("/api/goFast", getGoFastJson)
    r.HandleFunc("/api/go", getGoTime)
    r.HandleFunc("/api/goFull", )

    r.HandleFunc("/api/returnFast", )
    r.HandleFunc("/api/return", )
    r.HandleFunc("/api/returnFull", )

    http.ListenAndServe(":7650", r)
}

func getGoFastJson(w http.ResponseWriter, r *http.Request){
    dayType := getDayType()
    row := goDb.QueryRow(
        `SELECT * FROM "GO"
        WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND dayType=?
        ORDER BY time ASC LIMIT 1`,
        dayType
    )

    ans := go{}
    row.Scan(&ans.time,&ans.from,&ans.to);
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

func QueryStringHandler() {
    q := r.URL.Query()
}

func getGoTimeJson(fromto,id)  {
    row:=getGoTime(fromto,id);
}

func getGoTime(fromto,id)  {
    dayType := getDayType()
    var mode string

    if fromto==0{
        mode="from"
    }else{
        mode="to"
    }

    row := goDb.QueryRow(
        `SELECT * FROM "GO"
        WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND ?=? AND dayType=?
        ORDER BY time ASC LIMIT 1`,
        mode,
        id,
        dayType,
    )

    return row
}


func getReturnTime(id)  {
    dayType := getDayType()
    row := returnDb.QueryRow(
        `SELECT * FROM "RETURN"
        WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=? AND dayType=?
        ORDER BY time ASC LIMIT 1`,
        id,
        dayType
    )
    return row
}


func getDayType()  {
    var t int
    if time.Weekday()==0{
        t = 2
    }else if time.Weekday()<6{
        t = 0
    }else{
        t = 1
    }
    return t
}
