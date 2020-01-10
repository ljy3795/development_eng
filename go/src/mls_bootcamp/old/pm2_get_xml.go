package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DailyAverageAirQuality struct {
	XMLName        xml.Name `xml:"DailyAverageAirQuality"`
	Text           string   `xml:",chardata"`
	ListTotalCount string   `xml:"list_total_count"`
	RESULT         struct {
		// Text    string `xml:",chardata"`
		CODE    string `xml:"CODE"`
		MESSAGE string `xml:"MESSAGE"`
	} `xml:"RESULT"`
	Row []struct {
		// Text     string `xml:",chardata"`
		MSRDTDE  string `xml:"MSRDT_DE"`
		MSRSTENM string `xml:"MSRSTE_NM"`
		NO2      string `xml:"NO2"`
		O3       string `xml:"O3"`
		CO       string `xml:"CO"`
		SO2      string `xml:"SO2"`
		PM10     string `xml:"PM10"`
		PM25     string `xml:"PM25"`
	} `xml:"row"`
}

func main() {
	// wd, _ := os.Getwd()
	// fmt.Println("Current dir is ", wd)

	// // Read XML file
	// xmlFile, err := os.Open(wd + "/" + "test.xml")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer xmlFile.Close()

	// byteValue, _ := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// GET data from Open AP
	res, err := http.Get("http://openAPI.seoul.go.kr:8088/6f536d76626c6a793534526f6f7a57/xml/DailyAverageAirQuality/1/185/20121225")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var pm2 DailyAverageAirQuality
	err = xml.Unmarshal(data, &pm2)

	if err != nil {
		panic(err.Error())
	}

	// check API data is okay
	if pm2.RESULT.CODE != "INFO-000" {
		panic("The API data is ERROR, " + pm2.RESULT.CODE + " " + pm2.RESULT.MESSAGE)
	}

	fmt.Println(len(pm2.Row))

	for i := 0; i < len(pm2.Row); i++ {
		tmp_v := pm2.Row[i]
		fmt.Printf("%+v\n", tmp_v)
	}

	// // XML 디코딩
	// var pm2 PM2
	// err = xml.Unmarshal(data, &pm2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(pm2)
}
