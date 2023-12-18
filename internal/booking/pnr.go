package booking

import "encoding/json"

type pnr struct {
	PnrLocator string `json:"pnr_locator"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Lang       string `json:"lang"`
}

func (p *pnr) Json() (json_data string, err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	json_data = string(data)
	return
}
