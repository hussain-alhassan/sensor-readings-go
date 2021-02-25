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
			fmt.Println(err)
		}

		var reading models.Reading
		err = json.Unmarshal(reqBody, &reading)

		if err != nil {
			fmt.Println(err)
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
			fmt.Println(err)
		}
	})

	fmt.Println("Start Server")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}