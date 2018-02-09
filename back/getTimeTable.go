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

type goOne struct{
    Time int `json:"time"`
    From int `json:"from"`
    To int `json:"to"`
}

type returnOne struct{
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
    r.HandleFunc("/api/go", getGoJson) //go?fromto=XX?id=XXと送ること//
    r.HandleFunc("/api/goFull", getGoFullJson)

    r.HandleFunc("/api/returnFast", getReturnFastJson)
    r.HandleFunc("/api/return", getReturnJson) //return?id=XXと送ること//
    r.HandleFunc("/api/returnFull", getReturnFullJson)

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

    ans := goOne{}
    row.Scan(&ans.time,&ans.from,&ans.to);
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

func getGoJson(w http.ResponseWriter, r *http.Request){
    q := r.URL.Query()
    row:=getGoTime(q.fromto,q.id);

    ans := goOne{}
    row.Scan(&ans.time,&ans.from,&ans.to);
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

func getGoFullJson(w http.ResponseWriter, r *http.Request)  {
    ans := goFull{}
    var [2]from int
    var [3]to int

    for i := 0; i < 2; i++ {
        from[i]=getGoTime(0,i)
    }
    for i := 0; i < 3; i++ {
        to[i]=getGoTime(1,i)
    }

    //あとはここで最速を判定して、ans内にまとめてください
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
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

    var time int
    row.Scan(&time,)
    return
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
