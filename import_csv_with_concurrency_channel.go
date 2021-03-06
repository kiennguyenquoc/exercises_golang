package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"sync"
	"time"
)

type Configuration struct {
	Users    []string
	Password []string
	Name     []string
}

func main() {
	settings := load_config("config.json")
	fmt.Println("### - Read file CSV and import data to PostgreSQL")
	fmt.Printf("StartTime: %v\n", time.Now())
	// connect postgresql
	dbinfo := fmt.Sprintf("user=%s host=localhost password=%s dbname=%s sslmode=disable", settings.Users[0], settings.Password[0], settings.Name[0])
	db, err := sql.Open("postgres", dbinfo)
	show_error(err)
	defer db.Close()

	drop_and_create_table(db)
	c := make(chan [][]string)
	var wg sync.WaitGroup
	go get_data_from_csv("data.csv", c)
	go get_data_from_csv("ca-500.csv", c)
	go get_data_from_csv("uk-500.csv", c)
	csvData1, csvData2, csvData3 := <-c, <-c, <-c
	wg.Add(3)
	go import_data_csv(csvData1, db, &wg)
	go import_data_csv(csvData2, db, &wg)
	go import_data_csv(csvData3, db, &wg)
	wg.Wait()
	fmt.Printf("StopTime: %v\n", time.Now())
}

func show_error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func get_data_from_csv(file_name string, c chan [][]string) {
	csvFile, err := os.Open(file_name)
	show_error(err)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c <- csvData
}

func import_data_csv(data [][]string, db *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	insertData, err := db.Prepare("INSERT INTO userinfo_channel (first_name, last_name, company_name, address, city, province, postal, phone1, phone2, email, web) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	show_error(err)
	defer insertData.Close()

	tx, err := db.Begin()
	for _, row := range data {
		_, err := tx.Stmt(insertData).Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10])
		show_error(err)
	}
	tx.Commit()
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

func drop_and_create_table(db *sql.DB) {
	db.Exec("DROP IF EXISTS TABLE userinfo_channel")
	db.Exec("CREATE IF NOT EXISTS TABLE userinfo_channel ( ID SERIAL PRIMARY KEY NOT NULL, first_name TEXT NOT NULL, last_name TEXT NOT NULL, company_name TEXT NOT NULL, address TEXT NOT NULL, city TEXT NOT NULL, province TEXT NOT NULL, postal TEXT NOT NULL, phone1 TEXT NOT NULL, phone2 TEXT NOT NULL, email TEXT NOT NULL, web TEXT NOT NULL )")
}
