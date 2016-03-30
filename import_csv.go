// CREATE TABLE userinfo
// (
//     ID SERIAL PRIMARY KEY NOT NULL,
//     first_name TEXT NOT NULL,
//     last_name TEXT NOT NULL,
//     company_name TEXT NOT NULL,
//     address TEXT NOT NULL,
//     city TEXT NOT NULL,
//     province TEXT NOT NULL,
//     postal TEXT NOT NULL,
//     phone1 TEXT NOT NULL,
//     phone2 TEXT NOT NULL,
//     email TEXT NOT NULL,
//     web TEXT NOT NULL
// );


package main
import (
  "fmt"
  "time"
  "os"
  "encoding/csv"
  "encoding/json"
  "database/sql"
  _ "github.com/lib/pq"
)

type Configuration struct {
  Users       []string
  Password    []string
  Name        []string
}

func main() {
  fmt.Printf("StartTime: %v\n", time.Now())
  settings := load_config("config.json")
  fmt.Println("### - Read file CSV and import data to PostgreSQL")
  // connect postgresql
  dbinfo := fmt.Sprintf("user=%s host=localhost password=%s dbname=%s sslmode=disable", settings.Users[0], settings.Password[0], settings.Name[0])
  db, err := sql.Open("postgres", dbinfo)
  show_error(err)

  drop_table, err := db.Exec("DROP TABLE userinfo")
  show_error(err)
  fmt.Println(drop_table)

  create_table, err := db.Exec("CREATE TABLE userinfo ( ID SERIAL PRIMARY KEY NOT NULL, first_name TEXT NOT NULL, last_name TEXT NOT NULL, company_name TEXT NOT NULL, address TEXT NOT NULL, city TEXT NOT NULL, province TEXT NOT NULL, postal TEXT NOT NULL, phone1 TEXT NOT NULL, phone2 TEXT NOT NULL, email TEXT NOT NULL, web TEXT NOT NULL )" )
  show_error(err)
  fmt.Println(create_table)

  csvFile, err := os.Open("data.csv")
  show_error(err)
  defer csvFile.Close()

  reader := csv.NewReader(csvFile)
  reader.FieldsPerRecord = -1
  csvData, err := reader.ReadAll()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for _, row := range csvData {
    // Insert data
    _, err := db.Exec("INSERT INTO userinfo (first_name, last_name, company_name, address, city, province, postal, phone1, phone2, email, web) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10])
    show_error(err)
  }

  db.Close()
  fmt.Printf("StopTime: %v\n", time.Now())
}

func show_error(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func load_config(file_name string) Configuration {
  file, _ := os.Open(file_name)
  decoder := json.NewDecoder(file)
  configuration := Configuration{}
  err := decoder.Decode(&configuration)
  if err != nil {
    fmt.Println("error:", err)
  }
  return configuration
}