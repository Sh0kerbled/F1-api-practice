package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Driver struct {
	Driver          string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	FirstName       string `json:"givenName"`
	LastName        string `json:"familyName"`
	Nationality     string `json:"nationality"`
}

type JolpicaResponse struct {
	MRData struct {
		Total       string `json:"total"`
		DriverTable struct {
			Season  string   `json:"season"`
			Drivers []Driver `json:"Drivers"`
		} `json:"DriverTable"`
	} `json:"MRData"`
}

func GetDriver() ([]Driver, error) {
	url := "https://api.jolpi.ca/ergast/f1/2026/drivers.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сервер вернул статус: %d", resp.StatusCode)
	}

	var data JolpicaResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	return data.MRData.DriverTable.Drivers, nil

}
