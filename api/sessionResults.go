package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SessionResults struct {
	DriverNumber int     `json:"driver_number"`
	Position     int     `json:"position"`
	Points       float32 `json:"points"`
	Dnf          bool    `json:"dnf"`
	Dns          bool    `json:"dns"`
	Dsq          bool    `json:"dsq"`
	duration     float64 `json:"duration"`
}

type Drivers struct {
	DriverNumber int    `json:"driver_number"`
	FullName     string `json:"full_name"`
	Acronym      string `json:"name_acronym"`
	TeamName     string `json:"team_name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	HeadshotUrl  string `json:"headshot_url"`
}

type CombinedData struct {
	Position     int
	DriverNumber int
	FullName     string
	Acronym      string
	TeamName     string
	FirstName    string
	LastName     string
	HeadshotUrl  string
	Points       float32
	Dnf          bool
	Dns          bool
	Dsq          bool
	duration     float64
}

func GetDriverData() {
	//drivers data ------------------------------------------------------------------------------
	driversUrl := "https://api.openf1.org/v1/drivers?session_key=11388"

	resp, err := http.Get(driversUrl)
	if err != nil {
		fmt.Println("Ошибка загрузки данных", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var d []Drivers

	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("Ошибка парсинга данных", err)
	}

	//session results data ----------------------------------------------------------------------
	sessionUrl := "https://api.openf1.org/v1/session_result?session_key=11388"

	sessionResp, err := http.Get(sessionUrl)
	if err != nil {
		fmt.Println("Ошибка загрузки данных", err)
	}
	defer sessionResp.Body.Close()

	sessionBody, err := io.ReadAll(sessionResp.Body)

	var s []SessionResults

	err = json.Unmarshal(sessionBody, &s)
	if err != nil {
		fmt.Println("Ошибка парсинга данных", err)
	}

	for _, v := range CombineData(s, d) {
		fmt.Println(v.FullName, v.Position, v.Points, v.TeamName)
	}
}

func CombineData(results []SessionResults, drivers []Drivers) []CombinedData {
	driversMap := make(map[int]Drivers)
	for _, d := range drivers {
		driversMap[d.DriverNumber] = d
	}

	var combined []CombinedData

	for _, r := range results {
		driver, exists := driversMap[r.DriverNumber]
		if !exists {
			continue
		}

		data := CombinedData{
			Position:     r.Position,
			DriverNumber: r.DriverNumber,
			FullName:     driver.FullName,
			Acronym:      driver.Acronym,
			TeamName:     driver.TeamName,
			FirstName:    driver.FirstName,
			LastName:     driver.LastName,
			HeadshotUrl:  driver.HeadshotUrl,
			Points:       r.Points,
			Dnf:          r.Dnf,
			Dns:          r.Dns,
			Dsq:          r.Dsq,
			duration:     r.duration,
		}
		combined = append(combined, data)
	}
	return combined
}
