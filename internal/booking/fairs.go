package booking

import "encoding/json"

type FairsReq struct {
	Lang       string  `json:"lang"`
	PnrLocator string  `json:"pnr_locator"`
	PnrKey     string  `json:"pnr_key"`
	Routes     []Route `json:"routes"`
}
type Route struct {
	DepartureDate string `json:"departure_date"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
}

func (f *FairsReq) Json() (json_data string, err error) {
	data, err := json.Marshal(f)
	if err != nil {
		return
	}
	json_data = string(data)
	return
}

type FairsResp struct {
	Data    []Data `json:"data"`
	Error   any    `json:"error"`
	Success bool   `json:"success"`
}
