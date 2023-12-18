package booking

import (
	"aeroflot-api-cli/internal/api"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Booking struct {
	Code         string
	LastName     string
	FirstName    string
	Departure    string
	api          api.AeroflotApi
	pnr_req_data pnr
	pnr_view     PnrView
	selected_leg PnrLegView
	fairs        FairsResp
	prices       PricesResp
}

func (b *Booking) SprintfLeg(leg PnrLegView) string {
	// var str_view string
	str_view := fmt.Sprintf("Рейс %s", leg.Segments[0].Departure)
	str_view = fmt.Sprintf("%s %s", str_view, leg.Segments[0].CabinName)
	str_view = fmt.Sprintf("%s: %s", str_view, leg.Segments[0].Origin.AirportName)
	str_view = fmt.Sprintf("%s(%s)", str_view, leg.Segments[0].Origin.CountryName)
	str_view = fmt.Sprintf("%s -> %s", str_view, leg.Segments[0].Destination.AirportName)
	str_view = fmt.Sprintf("%s(%s)", str_view, leg.Segments[0].Destination.CountryName)
	return str_view
}

func (b *Booking) GetInfo() {
	pnr_locator, _ := b.pnr_req_data.Json()
	data, err := b.api.PostJson("https://www.aeroflot.ru/se/api/app/pnr/view/v3", []byte(pnr_locator))
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &b.pnr_view)
	log.Printf("PNR Key: %+v", b.pnr_view.Data.PnrKey)
	b.selectLeg()
	b.Departure = strings.Split(b.selected_leg.Segments[0].Departure, " ")[0]

	// log.Printf("%#v", b.selected_leg)
}

func (b *Booking) selectLeg() {
	for j, leg := range b.pnr_view.Data.Legs {
		log.Printf("Номер %d: %s", j, b.SprintfLeg(leg))
	}
	var i int
	fmt.Printf("Введите номер рейса: ")
	fmt.Scanf("%d", &i)
	if i >= len(b.pnr_view.Data.Legs) {
		log.Fatalf("Некорректный номер рейса для поиска. Должен быть 0 <= x < %d.", len(b.pnr_view.Data.Legs))
	}
	b.selected_leg = b.pnr_view.Data.Legs[i]
	// leg, _ := json.Marshal(b.selected_leg)
	// log.Print(string(leg))
}

func (b *Booking) GetFaresVariants(date string) {
	fair_req := FairsReq{
		Lang:       b.pnr_req_data.Lang,
		PnrLocator: b.pnr_req_data.PnrLocator,
		PnrKey:     b.pnr_view.Data.PnrKey,
	}
	fair_req.Routes = append(fair_req.Routes,
		Route{
			DepartureDate: date,
			Origin:        b.selected_leg.Segments[0].Origin.AirportCode,
			Destination:   b.selected_leg.Segments[0].Destination.CityCode,
		})
	fair_req_json, _ := fair_req.Json()
	fair_resp, _ := b.api.PostJson("https://www.aeroflot.ru/se/api/app/pnr/exchange/search/v3", []byte(fair_req_json))
	json.Unmarshal(fair_resp, &b.fairs)
	// log.Printf("%#v", b.fairs)
	// log.Printf("%s", string(fair_resp))
}

func (b *Booking) GetFairPrices(date string) {
	b.GetFaresVariants(date)
	for _, fair_variant := range b.fairs.Data {
		// log.Printf("%+v", fair_variant)
		price_req := PricesReq{
			Lang:                     b.pnr_req_data.Lang,
			PnrLocator:               b.pnr_req_data.PnrLocator,
			PnrKey:                   b.pnr_view.Data.PnrKey,
			CancellingSegmentNumbers: []int{fair_variant.Legs[0].Segments[0].SegmentNumber},
			NewSegments: []NewSegments{
				NewSegments{
					Origin:            fair_variant.Legs[0].Segments[0].Origin.AirportCode,
					Destination:       fair_variant.Legs[0].Segments[0].Destination.AirportCode,
					Departure:         strings.TrimSuffix(fair_variant.Legs[0].Segments[0].Departure, ":00"),
					AirlineCode:       fair_variant.Legs[0].Segments[0].AirlineCode,
					FlightNumber:      fair_variant.Legs[0].Segments[0].FlightNumber,
					BookingClass:      "E",
					InboundConnection: false,
					Arrival:           strings.TrimSuffix(fair_variant.Legs[0].Segments[0].Arrival, ":00"),
				},
			},
			Client:     Client{},
			PosCountry: fair_variant.Legs[0].Segments[0].Destination.CountryCode,
			Brand:      "EC",
			// ExchangeOffers: fair_variant.FaresVariants,
		}

		for i, exc_offer := range fair_variant.FaresVariants {
			price_req.ExchangeOffers = append(price_req.ExchangeOffers,
				ExchangeOffers{
					Brand: exc_offer.Brand,
				})
			for _, fares := range exc_offer.FaresOnSegments {
				price_req.ExchangeOffers[i].SegmentBookingClasses = append(price_req.ExchangeOffers[i].SegmentBookingClasses, fares.BookingClass)
			}
		}

		price_req_json, _ := price_req.Json()
		// log.Printf("Req %+v", string(price_req_json))
		prices, _ := b.api.PostJson("https://www.aeroflot.ru/se/api/app/pnr/exchange/price/v3", []byte(price_req_json))
		// log.Printf("Prices: 	%+v", string(prices))
		json.Unmarshal(prices, &b.prices)
		fmt.Printf("%s: %s %s\n", date, fair_variant.Legs[0].Segments[0].DepartureTime, b.prices.Sprintf())

		// log.Println(err)
		// return // TODO: del
	}
}

func New(code string, last_name string, first_name string) (booking Booking, err error) {
	booking = Booking{
		Code:      code,
		LastName:  last_name,
		FirstName: first_name,
		api:       api.New(),
		pnr_req_data: pnr{
			PnrLocator: code,
			LastName:   last_name,
			FirstName:  first_name,
			Lang:       "ru",
		},
	}
	return
}
