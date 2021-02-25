package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hussain-alhassan/sensor-readings-go/models"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	sensorID := "sensor1"
	sensorType := "temperature"

	for {
		rand.Seed(time.Now().Unix())
		value := rand.Float64() * 50 - 25
		alert := false

		if value < -20 || value > 15 {
			alert = true
		}
		reading := models.Reading{
			ID:        sensorID,
			Type:      sensorType,
			Value:     value,
			Alert:     alert,
			Timestamp: time.Now().UTC(),
		}

		readingMarshal, _ := json.Marshal(reading)
		fmt.Println("Sending reading: ", string(readingMarshal))

		_, err := http.Post("http://127.0.0.1:5000/sensor-readings",
			"application/json", bytes.NewReader(readingMarshal))

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second)
	}
}