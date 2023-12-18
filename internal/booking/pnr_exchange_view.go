package booking

type PnrExchangeView struct {
	Data    []Data `json:"data"`
	Error   any    `json:"error"`
	Success bool   `json:"success"`
}
type ExchangeCost struct {
	Base        string `json:"base"`
	Currency    string `json:"currency"`
	Fees        string `json:"fees"`
	Tax         string `json:"tax"`
	TotalAmount string `json:"total_amount"`
}
type FaresOnSegments struct {
	BookingClass     string `json:"booking_class"`
	BookingClassName string `json:"booking_class_name"`
	FareBase         string `json:"fare_base"`
	Hit              bool   `json:"hit"`
	MealCode         string `json:"meal_code"`
	MealName         string `json:"meal_name"`
	SegmentNumber    int    `json:"segment_number"`
	Shopped          bool   `json:"shopped"`
}
type FaresVariants struct {
	Brand           string            `json:"brand"`
	BrandName       string            `json:"brand_name"`
	Cabin           string            `json:"cabin"`
	ExchangeCost    ExchangeCost      `json:"exchange_cost"`
	FaresOnSegments []FaresOnSegments `json:"fares_on_segments"`
	OperationType   string            `json:"operation_type"`
}
type Destination struct {
	AirportCode string `json:"airport_code"`
	AirportName string `json:"airport_name"`
	CityCode    string `json:"city_code"`
	CityName    string `json:"city_name"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}
type Origin struct {
	AirportCode string `json:"airport_code"`
	AirportName string `json:"airport_name"`
	CityCode    string `json:"city_code"`
	CityName    string `json:"city_name"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}
type Segments struct {
	AirlineCode          string      `json:"airline_code"`
	AirlineName          string      `json:"airline_name"`
	Arrival              string      `json:"arrival"`
	ArrivalDateName      string      `json:"arrival_date_name"`
	ArrivalOffset        int         `json:"arrival_offset"`
	ArrivalTime          string      `json:"arrival_time"`
	Departure            string      `json:"departure"`
	DepartureDateName    string      `json:"departure_date_name"`
	DepartureOffset      int         `json:"departure_offset"`
	DepartureTime        string      `json:"departure_time"`
	Destination          Destination `json:"destination"`
	FlightNumber         string      `json:"flight_number"`
	FlightTime           int         `json:"flight_time"`
	FlightTimeName       string      `json:"flight_time_name"`
	InboundConnection    bool        `json:"inbound_connection"`
	OperatingAirlineCode string      `json:"operating_airline_code"`
	OperatingAirlineName string      `json:"operating_airline_name"`
	Origin               Origin      `json:"origin"`
	SegmentNumber        int         `json:"segment_number"`
}
type Leg struct {
	LegTime     int        `json:"leg_time"`
	LegTimeName string     `json:"leg_time_name"`
	Segments    []Segments `json:"segments"`
}
type Data struct {
	FaresVariants []FaresVariants `json:"fares_variants"`
	Legs          []Leg           `json:"legs"`
}
