package main

import (
	"Assignment3/dataStruct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

func main() {
	DataStatus := dataStruct.DataStatus{}
	var filepath = "public/index.html"
	var jsonpath = "public/weather.json"
	go WriteJson(DataStatus)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        
		//read json file
		file, _ := ioutil.ReadFile(jsonpath)
		json.Unmarshal(file, &DataStatus)
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}


		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"StatusWater": DataStatus.StatusWater,
			"StatusWind": DataStatus.StatusWind,
			"Water":  DataStatus.DataStatus.Water,
			"Wind": DataStatus.DataStatus.Wind,	
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
    })

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}

func WriteJson(dataStatus dataStruct.DataStatus){
	for{
		dataStatus.StatusWater = "aman"
		dataStatus.StatusWind = "bahaya"
		dataStatus.DataStatus.Water = rand.Intn(100)
		dataStatus.DataStatus.Wind = rand.Intn(100)

		if dataStatus.DataStatus.Water < 5{
			dataStatus.StatusWater = "aman"
		} 
		if (dataStatus.DataStatus.Water >=6 && dataStatus.DataStatus.Water <=8){
			dataStatus.StatusWater = "siaga"
		}
		if dataStatus.DataStatus.Water > 8{
			dataStatus.StatusWater = "bahaya"
		} 

		if dataStatus.DataStatus.Wind < 6{
			dataStatus.StatusWind = "aman"
		} 
		if (dataStatus.DataStatus.Wind >=7 && dataStatus.DataStatus.Wind <=15){
			dataStatus.StatusWind = "siaga"
		}
		if dataStatus.DataStatus.Wind > 15{
			dataStatus.StatusWind = "bahaya"
		} 

		//write json file 
		jsonString, _ := json.Marshal(dataStatus) 
		ioutil.WriteFile("public/weather.json", jsonString, os.ModePerm)

		time.Sleep(15 * time.Second)
	}
}

