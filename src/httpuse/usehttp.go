package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type CounrtyVO struct {
	Title     string
	Countries []CountryInfo
}
type JsonResponse struct {
	Resp RestResponse `json:"RestResponse"`
}

type CountryDeatils struct {
	Name    string `json:"name"`
	Alpha2  string `json:"alpha2Code"`
	Alpha3  string `json:"alpha3Code"`
	Capital string `json:"capital"`
}

type RestResponse struct {
	Messages []string `json:"messages"`
	Results  []Result `json:"result"`
}

type Result struct {
	Name        string `json:"name"`
	Alpha2_code string `json:"alpha2_code"`
	Alpha3_code string `json:"alpha3_code"`
}

func (r Result) String() string {
	return fmt.Sprintf(r.Name)
}

type CountryInfo struct {
	Idx     int
	Name    string
	Alpha2  string
	Alpha3  string
	Capital string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey this is index</h1>")
	fmt.Fprintf(w, "<a href='/countries'>Click here</a>")
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	var jresp JsonResponse
	var countryInfo []CountryInfo
	countryChennel := make(chan CountryDeatils, 249)
	resp, _ := http.Get("http://services.groupkt.com/country/get/all")
	bytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bytes, &jresp)
	resp.Body.Close()
	for _, results := range jresp.Resp.Results {
		wg.Add(1)
		go getCapital(countryChennel, results.Alpha3_code)
	}
	wg.Wait()
	close(countryChennel)
	idx := 1
	for elem := range countryChennel {
		countryInfo = append(countryInfo, CountryInfo{idx, elem.Name, elem.Alpha2, elem.Alpha3, elem.Capital})
		idx++
	}

	t, _ := template.ParseFiles("usehttp.html")
	viewData := CounrtyVO{Title: "Country Codes", Countries: countryInfo}
	fmt.Println(t.Execute(w, viewData))
}

func getCapital(c chan CountryDeatils, code string) {
	defer wg.Done()
	var capResp CountryDeatils
	resp, _ := http.Get("https://restcountries.eu/rest/v2/alpha/" + code + "?fields=capital;name;alpha2Code;alpha3Code")
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &capResp)
	resp.Body.Close()
	fmt.Println(capResp)
	c <- capResp
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/countries", countriesHandler)
	http.ListenAndServe("localhost:8000", nil)
}
