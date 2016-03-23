// CREATE TABLE userinfo
// (
//     ID INT PRIMARY KEY     NOT NULL,
//     first_name TEXT NOT NULL,
//     last_name TEXT NOT NULL,
//     company_name TEXT NOT NULL,
//     address TEXT NOT NULL,
//     city TEXT NOT NULL,
//     county TEXT NOT NULL,
//     postal TEXT NOT NULL,
//     phone1 TEXT NOT NULL,
//     phone2 TEXT NOT NULL,
//     email TEXT NOT NULL,
//     web TEXT NOT NULL
// );

// CREATE TABLE userinfo (
//     uid INTEGER PRIMARY KEY AUTOINCREMENT,
//     username VARCHAR(64) NULL,
//     departname VARCHAR(64) NULL,
//     created DATE NULL
// );

package main
import (
  "fmt"
  "time"
  "log"
  "os"
  "encoding/csv"
  "database/sql"
  _ "github.com/lib/pq"
)

const (
  DB_USER     = "pguser"
  DB_PASSWORD = "kien@uit"
  DB_NAME     = "golang"
)

func main() {
  fmt.Println("### - Read file CSV and import data to PostgreSQL")

  dbinfo := fmt.Sprintf("user=%s host=localhost password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("StartTime: %v\n", time.Now())

  csvFile, err := os.Open("data.csv")

  if err != nil {
    fmt.Println(err)
  }

  defer csvFile.Close()

  reader := csv.NewReader(csvFile)
  reader.FieldsPerRecord = -1
  csvData, err := reader.ReadAll()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for _, each := range csvData {
    // Thu hien luu xuong DB postgre
    fmt.Println(each)
  }

  stmt, err := db.Exec("INSERT INTO userinfo(uid, username, departname, created) values(2, 'astaxie', '研发部门', '2012-12-09')")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(stmt)


  db.Close()
  fmt.Printf("StopTime: %v\n", time.Now())
}


