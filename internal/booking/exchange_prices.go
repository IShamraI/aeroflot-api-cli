package booking

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PricesReq struct {
	Lang                     string           `json:"lang"`
	PnrLocator               string           `json:"pnr_locator"`
	PnrKey                   string           `json:"pnr_key"`
	CancellingSegmentNumbers []int            `json:"cancelling_segment_numbers"`
	NewSegments              []NewSegments    `json:"new_segments"`
	Client                   Client           `json:"client"`
	PosCountry               string           `json:"pos_country"`
	Brand                    string           `json:"brand"`
	ExchangeOffers           []ExchangeOffers `json:"exchange_offers"`
}
type NewSegments struct {
	Origin            string `json:"origin"`
	Destination       string `json:"destination"`
	Departure         string `json:"departure"`
	AirlineCode       string `json:"airline_code"`
	FlightNumber      string `json:"flight_number"`
	BookingClass      string `json:"booking_class"`
	InboundConnection bool   `json:"inbound_connection"`
	Arrival           string `json:"arrival"`
}
type Client struct {
	GaClientID   string `json:"ga_client_id"`
	LoyaltyID    string `json:"loyalty_id"`
	LoyaltyLevel string `json:"loyalty_level"`
}
type ExchangeOffers struct {
	Brand                 string   `json:"brand"`
	SegmentBookingClasses []string `json:"segment_booking_classes"`
}

func (p *PricesReq) Json() (json_data string, err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	json_data = string(data)
	return
}

type PricesResp struct {
	Data    Data `json:"data"`
	Error   any  `json:"error"`
	Success bool `json:"success"`
}
type BaggageInfo struct {
	NotPermitted []any `json:"not_permitted"`
	Rules        []any `json:"rules"`
	SegmentInfo  []any `json:"segment_info"`
}

//	type ExchangeCost struct {
//		Base        string `json:"base"`
//		Currency    string `json:"currency"`
//		Fees        string `json:"fees"`
//		Tax         string `json:"tax"`
//		TotalAmount string `json:"total_amount"`
//	}
//
//	type FaresOnSegments struct {
//		BookingClass  string `json:"booking_class"`
//		FareBase      string `json:"fare_base"`
//		Hit           bool   `json:"hit"`
//		SegmentNumber int    `json:"segment_number"`
//		Shopped       bool   `json:"shopped"`
//	}
type PaxPrices struct {
	NewBase             string `json:"new_base"`
	NewFees             string `json:"new_fees"`
	NewTaxes            string `json:"new_taxes"`
	NewTotalAmount      string `json:"new_total_amount"`
	OriginalBase        string `json:"original_base"`
	OriginalTaxes       string `json:"original_taxes"`
	OriginalTotalAmount string `json:"original_total_amount"`
	PaxCount            int    `json:"pax_count"`
	PaxType             string `json:"pax_type"`
	PaxTypeName         string `json:"pax_type_name"`
	TicketNumbers       []any  `json:"ticket_numbers"`
}
type Total struct {
	NewCurrency         string `json:"new_currency"`
	NewTotalAmount      string `json:"new_total_amount"`
	OriginalCurrency    string `json:"original_currency"`
	OriginalTotalAmount string `json:"original_total_amount"`
	Surcharge           string `json:"surcharge"`
}

// type FaresVariants struct {
// 	BaggageInfo           BaggageInfo       `json:"baggage_info"`
// 	BaggageText           string            `json:"baggage_text"`
// 	BaggageWeightRuleText string            `json:"baggage_weight_rule_text"`
// 	Brand                 string            `json:"brand"`
// 	BrandName             string            `json:"brand_name"`
// 	Cabin                 string            `json:"cabin"`
// 	CabinName             string            `json:"cabin_name"`
// 	Exchange              bool              `json:"exchange"`
// 	ExchangeCost          ExchangeCost      `json:"exchange_cost"`
// 	ExchangeText          string            `json:"exchange_text"`
// 	FareGroup             string            `json:"fare_group"`
// 	FareGroupName         string            `json:"fare_group_name"`
// 	FaresOnSegments       []FaresOnSegments `json:"fares_on_segments"`
// 	MilesPercentText      string            `json:"miles_percent_text"`
// 	MilesUpgrade          bool              `json:"miles_upgrade"`
// 	MilesUpgradeText      string            `json:"miles_upgrade_text"`
// 	NoBaggage             bool              `json:"no_baggage"`
// 	OperationType         string            `json:"operation_type"`
// 	PaxPrices             []PaxPrices       `json:"pax_prices"`
// 	Refund                bool              `json:"refund"`
// 	RefundText            string            `json:"refund_text"`
// 	SeatPreselectText     string            `json:"seat_preselect_text"`
// 	Total                 Total             `json:"total"`
// }

// type Destination struct {
// 	Airport     string `json:"airport"`
// 	AirportName string `json:"airport_name"`
// 	City        string `json:"city"`
// 	CityName    string `json:"city_name"`
// 	Country     string `json:"country"`
// 	CountryName string `json:"country_name"`
// }
// type Origin struct {
// 	Airport     string `json:"airport"`
// 	AirportName string `json:"airport_name"`
// 	City        string `json:"city"`
// 	CityName    string `json:"city_name"`
// 	Country     string `json:"country"`
// 	CountryName string `json:"country_name"`
// }

//	type Segments struct {
//		Airline           string      `json:"airline"`
//		AirlineName       string      `json:"airline_name"`
//		Arrival           string      `json:"arrival"`
//		ArrivalDateName   string      `json:"arrival_date_name"`
//		ArrivalOffset     int         `json:"arrival_offset"`
//		ArrivalTime       string      `json:"arrival_time"`
//		Departure         string      `json:"departure"`
//		DepartureDateName string      `json:"departure_date_name"`
//		DepartureOffset   int         `json:"departure_offset"`
//		DepartureTime     string      `json:"departure_time"`
//		Destination       Destination `json:"destination"`
//		FlightNumber      string      `json:"flight_number"`
//		FlightTimeName    string      `json:"flight_time_name"`
//		Origin            Origin      `json:"origin"`
//		SegmentNumber     int         `json:"segment_number"`
//	}
type Legs struct {
	LegTimeName    string     `json:"leg_time_name"`
	LegWeekdayName string     `json:"leg_weekday_name"`
	Segments       []Segments `json:"segments"`
}

// type Data struct {
// 	FaresVariants []FaresVariants `json:"fares_variants"`
// 	Legs          []Legs          `json:"legs"`
// }

func (p *PricesResp) Sprintf() string {
	var price_view string = ""
	for _, price := range p.Data.FaresVariants {
		price_view = fmt.Sprintf("%s %s: %s,", price_view, price.BrandName, price.ExchangeCost.TotalAmount)
	}
	price_view = strings.TrimLeft(price_view, " ")
	price_view = strings.TrimRight(price_view, ",")
	return price_view
}
