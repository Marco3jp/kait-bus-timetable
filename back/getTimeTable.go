package main

import (
  "fmt"
  "html"
  "log"
  "net/http"

  "github.com/mattn/go-sqlite3"
)

type go struct {
	Fast []int `json:"fast"`
	From []int `json:"from"`
	To   []int `json:"to"`
}

type return struct {
	Fast []int `json:"fast"`
	From []int `json:"from"`
}

func main() {
    goDb, goErr := sql.Open("sqlite3", "./go.db");
    returnDb, returnErr := sql.Open("sqlite3", "./retrun.db");

    if goErr != nil || returnErr != nil{
      panic(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func getGoTime(fromto,id)  {
    //曜日はtime.Weekday()で受け取れて、0が日曜日、1が……という仕組み。祝日どうするか考えてないけど未実装のままいこう……
    //ちょっと一旦寝て考え直します//
    var row
    switch fromto {
    case 0: //from
        switch id {
        case 0:
            if time.Weekday()==0 {
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=0 AND dayType=2
                    ORDER BY time ASC LIMIT 1`
                    )
            }else if time.Weekday()<6{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=0 AND dayType=0
                    ORDER BY time ASC LIMIT 1`
                    )
            }else{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=0 AND dayType=1
                    ORDER BY time ASC LIMIT 1`
                    )
            }
        case 1:
            if time.Weekday()==0 {
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=1 AND dayType=2
                    ORDER BY time ASC LIMIT 1`
                    )
            }else if time.Weekday()<7{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=1 AND dayType=0
                    ORDER BY time ASC LIMIT 1`
                    )
            }else{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=1 AND dayType=1
                    ORDER BY time ASC LIMIT 1`
                    )
            }
        }
    case 1: //to
        switch id {
        case 0:
            return "No Data."
            /*
            if time.Weekday()==0 {
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=0 AND dayType=2
                    ORDER BY time ASC LIMIT 1`
                    )
            */
            }else if time.Weekday()<6{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=0 AND dayType=0
                    ORDER BY time ASC LIMIT 1`
                    )
            }else{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=0 AND dayType=1
                    ORDER BY time ASC LIMIT 1`
                    )
            }
        case 1:
            if time.Weekday()==0 {
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=1 AND dayType=2
                    ORDER BY time ASC LIMIT 1`
                    )
            }else if time.Weekday()<6 {
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=1 AND dayType=0
                    ORDER BY time ASC LIMIT 1`
                    )
            }else{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=1 AND dayType=1
                    ORDER BY time ASC LIMIT 1`
                    )
            }
        case 2:
            if time.Weekday()==0 {
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=2 AND dayType=2
                    ORDER BY time ASC LIMIT 1`
                    )
            }else if time.Weekday()<6{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=2 AND dayType=0
                    ORDER BY time ASC LIMIT 1`
                    )
            }else{
                row := goDb.QueryRow(
                    `SELECT * FROM "GO"
                    WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND to=2 AND dayType=1
                    ORDER BY time ASC LIMIT 1`
                    )
            }
        }
    }
}

func getReturnTime(id)  {
    switch id {
    case 0:
        if time.Weekday()==0 {
            return "No Data."
            /*
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=0 AND dayType=2
                ORDER BY time ASC LIMIT 1`
                )
            */
        }else if time.Weekday()<6{
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=0 AND dayType=0
                ORDER BY time ASC LIMIT 1`
                )
        }else{
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=0 AND dayType=1
                ORDER BY time ASC LIMIT 1`
                )
        }
    case 1:
        if time.Weekday()==0 {
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=1 AND dayType=2
                ORDER BY time ASC LIMIT 1`
                )
        }else if time.Weekday()<6{
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=1 AND dayType=0
                ORDER BY time ASC LIMIT 1`
                )
        }else{
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=1 AND dayType=1
                ORDER BY time ASC LIMIT 1`
                )
        }
    }
    case 2:
        if time.Weekday()==0 {
            //存在しない
            return "No Data."
            /*
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=2 AND dayType=2
                ORDER BY time ASC LIMIT 1`
                )
            */
        }else if time.Weekday()<6{
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=2 AND dayType=0
                ORDER BY time ASC LIMIT 1`
                )
        }else{
            //存在しない
            return "No Data."
            /*
            row := returnDb.QueryRow(
                `SELECT * FROM "RETURN"
                WHERE strftime('%H','now','localtime')*60+strftime('%M','now','localtime')<time AND from=2 AND dayType=1
                ORDER BY time ASC LIMIT 1`
                )
            */
        }
    }

}
