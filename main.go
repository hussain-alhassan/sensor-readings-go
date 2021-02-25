package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/hussain-alhassan/sensor-readings-go/models"
)

func main() {
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
		fmt.Println(reading.ID)
	})

	fmt.Println("Start Server")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}