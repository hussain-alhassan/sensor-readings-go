package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/hussain-alhassan/sensor-readings-go/models"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	database, _ := sql.Open("sqlite3", "./readings.db")
	statement, _ := database.Prepare(
		"CREATE TABLE IF NOT EXISTS readings" +
			"(id INTEGER PRIMARY KEY," +
			"sensor_id VARCHAR(255)," +
			"type VARCHAR(255)," +
			"value VARCHAR(255)," +
			"alert BOOLEAN," +
			"timestamp TIMESTAMP)",
	)

	http.HandleFunc("/sensor-readings", func(writer http.ResponseWriter, request *http.Request) {
		reqBody, err := ioutil.ReadAll(request.Body)

		if err != nil {
			log.Println(err)
		}

		var reading models.Reading
		err = json.Unmarshal(reqBody, &reading)

		if err != nil {
			log.Println(err)
		}

		_, err = statement.Exec()
		if err != nil {
			log.Println(err)
		}
		statement, _ = database.Prepare(
			"INSERT INTO readings (sensor_id, type, value, alert, timestamp) VALUES (?, ?, ?, ?, ?)",
		)
		_, err = statement.Exec(reading.ID, reading.Type, reading.Value, reading.Alert, reading.Timestamp)
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/get-sensor-readings", func(writer http.ResponseWriter, request *http.Request) {
		rows, _ := database.Query("SELECT id, sensor_id, type, value, timestamp FROM readings ORDER BY id desc LIMIT 5")

		readings := make([]*models.Reading, 0)
		defer rows.Close()
		for rows.Next() {
			oneReading := new(models.Reading)
			if err := rows.Scan(&oneReading.ID, &oneReading.Type, &oneReading.Value, &oneReading.Timestamp); err != nil {
				panic(err)
			}
			readings = append(readings, oneReading)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

		result, _ := json.Marshal(readings)

		fmt.Fprintf(writer, string(result))
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Start Server")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}