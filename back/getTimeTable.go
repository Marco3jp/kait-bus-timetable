package main

import (
  "fmt"
  "html"
  "log"
  "time"
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

//最速を返す関数
func getGoFastJson(w http.ResponseWriter, r *http.Request){
    dayType := getDayType()
    nowTime := getNowTime()

    row := goDb.QueryRow(
        `SELECT * FROM "GO"
        WHERE ?<time AND dayType=?
        ORDER BY time ASC LIMIT 1`,
        nowTime,
        dayType
    )

    ans := goOne{}
    var tmp int //読み捨て
    row.Scan(&ans.time,&ans.from,&ans.to,&tmp);
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

//特定種の最速を返す
func getGoJson(w http.ResponseWriter, r *http.Request){
    q := r.URL.Query()
    row:=getGoTime(q.fromto,q.id);

    ans := goOne{}
    var tmp int
    row.Scan(&ans.time,&ans.from,&ans.to,&tmp);
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

//行き各種の最速をすべて返す
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

    ans.fast = checkGoFastTime(from,to)
    ans.from = from
    ans.to = to

    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

//検索に使う部分、返すのは時間//
func getGoTime(fromto int,id int) int{
    dayType := getDayType()
    nowTime := getNowTime()

    var mode string

    if fromto==0{
        mode="from"
    }else{
        mode="to"
    }

    row := goDb.QueryRow(
        `SELECT * FROM "GO"
        WHERE ?<time AND ?=? AND dayType=?
        ORDER BY time ASC LIMIT 1`,
        nowTime,
        mode,
        id,
        dayType,
    )

    var time int
    var [3]tmp int //読み捨て
    row.Scan(&time,tmp[0],tmp[1],tmp[2])
    return time
}

func getReturnFastJson(w http.ResponseWriter, r *http.Request){
    dayType := getDayType()
    nowTime := getNowTime()

    row := goDb.QueryRow(
        `SELECT * FROM "RETURN"
        WHERE ?<time AND dayType=?
        ORDER BY time ASC LIMIT 1`,
        nowTime,
        dayType
    )

    ans := returnOne{}
    var tmp int //読み捨て
    row.Scan(&ans.time,&ans.from,&tmp);
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

//特定種の最速を返す
func getReturnJson(w http.ResponseWriter, r *http.Request){
    q := r.URL.Query()
    row:=getGoTime(q.id);

    ans := returnOne{}
    var tmp int
    row.Scan(&ans.time,&ans.from,&tmp);
    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

//帰り各種の最速をすべて返す
func getReturnFullJson(w http.ResponseWriter, r *http.Request)  {
    ans := returnFull{}
    var [3]from int

    for i := 0; i < 3; i++ {
        from[i]=getReturnTime(i)
    }

    fast := [2]int{from[0],0}

    for i := 1; i < 3; i++ {
        if(fast[0]>from[i]){
            fast[0]==from[i]
            fast[1]==i
        }
    }

    ans.fast = fast
    ans.from = from

    outJson, err := json.Marshal(&ans)
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(outJson))
}

func getReturnTime(id int) int{
    dayType := getDayType()
    nowTime := getNowTime()
    row := returnDb.QueryRow(
        `SELECT * FROM "RETURN"
        WHERE ?<time AND from=? AND dayType=?
        ORDER BY time ASC LIMIT 1`,
        nowTime,
        id,
        dayType
    )

    var time int
    var [2]tmp int //読み捨て
    row.Scan(&time,tmp[0],tmp[1])
    return time
}

func getDayType()  {
    var n int
    t := time.Now()
    if t.Weekday()==0{
        n = 2
    }else if t.Weekday()<6{
        n = 0
    }else{
        n = 1
    }
    return n
}

func getNowTime(){
    t := time.Now();
    return t.Hour()*60+t.Minute()
}

func checkGoFastTime([2]a [3]b) {
    fromFast := [2]int{a[0],0}
    if(a[1]<fromFast[0]){
        fromFast[0]=a[1]
        fromFast[1]=1
    }

    toFast := [2]int{b[0],0}
    for i := 1; i < 3; i++ {
        if b[i]<toFast[0]{
            toFast[0]=b[i]
            toFast[1]=i
        }
    }
    return [fromFast[0],fromFast[1],toFast[1]]
}
