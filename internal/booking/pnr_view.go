package booking

type PnrLegView struct {
	LegTimeName    string `json:"leg_time_name"`
	LegWeekdayName string `json:"leg_weekday_name"`
	Segments       []struct {
		AircraftType         string `json:"aircraft_type"`
		AircraftTypeName     string `json:"aircraft_type_name"`
		AirlineCode          string `json:"airline_code"`
		AirlineName          string `json:"airline_name"`
		Arrival              string `json:"arrival"`
		ArrivalDateName      string `json:"arrival_date_name"`
		ArrivalOffset        int    `json:"arrival_offset"`
		ArrivalTime          string `json:"arrival_time"`
		AwardExchangeAllowed bool   `json:"award_exchange_allowed"`
		BookingClass         string `json:"booking_class"`
		Cabin                string `json:"cabin"`
		CabinName            string `json:"cabin_name"`
		DayDifference        int    `json:"day_difference"`
		Departure            string `json:"departure"`
		DepartureDateName    string `json:"departure_date_name"`
		DepartureOffset      int    `json:"departure_offset"`
		DepartureTime        string `json:"departure_time"`
		Destination          struct {
			AirportCode  string `json:"airport_code"`
			AirportName  string `json:"airport_name"`
			CityCode     string `json:"city_code"`
			CityName     string `json:"city_name"`
			CountryCode  string `json:"country_code"`
			CountryName  string `json:"country_name"`
			TerminalCode string `json:"terminal_code"`
			TerminalName string `json:"terminal_name"`
		} `json:"destination"`
		ExcessBaggageAllowed    bool     `json:"excess_baggage_allowed"`
		ExchangeAllowed         bool     `json:"exchange_allowed"`
		FareBasisCode           string   `json:"fare_basis_code"`
		FareClassOfService      string   `json:"fare_class_of_service"`
		FareGroupName           string   `json:"fare_group_name"`
		FlightNumber            string   `json:"flight_number"`
		FlightTime              int      `json:"flight_time"`
		FlightTimeName          string   `json:"flight_time_name"`
		IcsDownloadURL          string   `json:"ics_download_url"`
		IsChanged               bool     `json:"is_changed"`
		MealCodes               []string `json:"meal_codes"`
		MealNames               string   `json:"meal_names"`
		OperatingAirlineCode    string   `json:"operating_airline_code"`
		OperatingAirlineMessage string   `json:"operating_airline_message"`
		OperatingAirlineName    string   `json:"operating_airline_name"`
		OperatingFlightNumber   string   `json:"operating_flight_number"`
		Origin                  struct {
			AirportCode  string `json:"airport_code"`
			AirportName  string `json:"airport_name"`
			CityCode     string `json:"city_code"`
			CityName     string `json:"city_name"`
			CountryCode  string `json:"country_code"`
			CountryName  string `json:"country_name"`
			TerminalCode string `json:"terminal_code"`
			TerminalName string `json:"terminal_name"`
		} `json:"origin"`
		OriginalArrivalDateName   string `json:"original_arrival_date_name"`
		OriginalArrivalTime       string `json:"original_arrival_time"`
		OriginalDepartureDateName string `json:"original_departure_date_name"`
		OriginalDepartureTime     string `json:"original_departure_time"`
		OtherLocator              string `json:"other_locator"`
		PaidMealAllowed           bool   `json:"paid_meal_allowed"`
		PartnerAirlineMessage     string `json:"partner_airline_message"`
		SeatPaxDisallowed         []any  `json:"seat_pax_disallowed"`
		SeatPreselect             bool   `json:"seat_preselect"`
		SegmentNumber             int    `json:"segment_number"`
		StatusCode                string `json:"status_code"`
		StatusName                string `json:"status_name"`
		TrainConfirmations        []any  `json:"train_confirmations"`
		TransferSameTerminal      bool   `json:"transfer_same_terminal"`
		TransferTime              string `json:"transfer_time"`
		Upcl4CrtAllowed           bool   `json:"upcl4crt_allowed"`
	} `json:"segments"`
}

type PnrView struct {
	Data struct {
		AddServicesLossWarning   string `json:"add_services_loss_warning"`
		AdditionalServicesStatus []struct {
			ChangeAllowed bool   `json:"change_allowed"`
			OrderAllowed  bool   `json:"order_allowed"`
			Ordered       bool   `json:"ordered"`
			ServiceType   string `json:"service_type"`
		} `json:"additional_services_status"`
		AncillaryServicesData string       `json:"ancillary_services_data"`
		AnimalBaggageAllowed  bool         `json:"animal_baggage_allowed"`
		AwardExchangeAllowed  bool         `json:"award_exchange_allowed"`
		BaggageInfo           any          `json:"baggage_info"`
		BassinetAllowed       bool         `json:"bassinet_allowed"`
		BookingMail           string       `json:"booking_mail"`
		BookingPhone          string       `json:"booking_phone"`
		BusinessPassInfo      any          `json:"business_pass_info"`
		CanBePaid             bool         `json:"can_be_paid"`
		CheckinURL            string       `json:"checkin_url"`
		EssURL                string       `json:"ess_url"`
		ExcessBaggageAllowed  bool         `json:"excess_baggage_allowed"`
		ExchangeAllowed       bool         `json:"exchange_allowed"`
		GetAppleWalletPass    string       `json:"get_apple_wallet_pass"`
		GetGooglePayPass      string       `json:"get_google_pay_pass"`
		HasChangedSegments    bool         `json:"has_changed_segments"`
		IsAward               bool         `json:"is_award"`
		IsBlockCharter        bool         `json:"is_block_charter"`
		IsBusinessPass        bool         `json:"is_business_pass"`
		IsCorpPnr             bool         `json:"is_corp_pnr"`
		IsResidenceRequired   bool         `json:"is_residence_required"`
		IsSubsidized          bool         `json:"is_subsidized"`
		IsVisaRequired        bool         `json:"is_visa_required"`
		Legs                  []PnrLegView `json:"legs"`
		NotificationRefused   bool         `json:"notification_refused"`
		Organizer             string       `json:"organizer"`
		Passengers            []struct {
			Birthdate                  string `json:"birthdate"`
			DocCountry                 string `json:"doc_country"`
			DocCountryName             string `json:"doc_country_name"`
			DocExpiration              string `json:"doc_expiration"`
			DocNumber                  string `json:"doc_number"`
			DocWithoutExpiration       bool   `json:"doc_without_expiration"`
			Email                      string `json:"email"`
			EmergencyContactCountry    string `json:"emergency_contact_country"`
			EmergencyContactName       string `json:"emergency_contact_name"`
			EmergencyContactPhone      string `json:"emergency_contact_phone"`
			EscortRef                  string `json:"escort_ref"`
			FirstName                  string `json:"first_name"`
			IsEmailRequired            bool   `json:"is_email_required"`
			IsEmergencyContactRequired bool   `json:"is_emergency_contact_required"`
			IsPhoneRequired            bool   `json:"is_phone_required"`
			LastName                   string `json:"last_name"`
			LoyaltyID                  string `json:"loyalty_id"`
			LoyaltyLevel               string `json:"loyalty_level"`
			LoyaltyLevelName           string `json:"loyalty_level_name"`
			LoyaltyProgram             string `json:"loyalty_program"`
			LoyaltyProgramName         string `json:"loyalty_program_name"`
			Nationality                string `json:"nationality"`
			NationalityName            string `json:"nationality_name"`
			PassengerAge               string `json:"passenger_age"`
			PaxKey                     string `json:"pax_key"`
			PaxType                    string `json:"pax_type"`
			Phone                      string `json:"phone"`
			RefNumber                  string `json:"ref_number"`
			ResidenceAddress           string `json:"residence_address"`
			ResidenceCity              string `json:"residence_city"`
			ResidenceCountry           string `json:"residence_country"`
			ResidenceLiveCountry       string `json:"residence_live_country"`
			ResidencePostalCode        string `json:"residence_postal_code"`
			ResidenceState             string `json:"residence_state"`
			SeatsCount                 int    `json:"seats_count"`
			SeatsRequirements          []any  `json:"seats_requirements"`
			Sex                        string `json:"sex"`
			SpecialServices            []any  `json:"special_services"`
			TicketingDocuments         struct {
				Additional []struct {
					Inactive bool   `json:"inactive"`
					Number   string `json:"number"`
					URL      string `json:"url"`
					URLPdf   string `json:"url_pdf"`
				} `json:"additional"`
				Tickets []struct {
					AirportCodes []string `json:"airport_codes"`
					Inactive     bool     `json:"inactive"`
					Number       string   `json:"number"`
					URL          string   `json:"url"`
					URLPdf       string   `json:"url_pdf"`
				} `json:"tickets"`
			} `json:"ticketing_documents"`
			VisaBirthPlace              string `json:"visa_birth_place"`
			VisaCountry                 string `json:"visa_country"`
			VisaIssueDate               string `json:"visa_issue_date"`
			VisaIssuePlace              string `json:"visa_issue_place"`
			VisaIssuePlaceDefault       string `json:"visa_issue_place_default"`
			VisaNumber                  string `json:"visa_number"`
			WaivervisaBirthdate         string `json:"waivervisa_birthdate"`
			WaivervisaExpiration        string `json:"waivervisa_expiration"`
			WaivervisaIssueCountry      string `json:"waivervisa_issue_country"`
			WaivervisaNationality       string `json:"waivervisa_nationality"`
			WaivervisaNumber            string `json:"waivervisa_number"`
			WaivervisaType              string `json:"waivervisa_type"`
			WaivervisaTypeName          string `json:"waivervisa_type_name"`
			WaivervisaWithoutExpiration bool   `json:"waivervisa_without_expiration"`
		} `json:"passengers"`
		PaymentURL                      string `json:"payment_url"`
		PnrChangeAuthenticationRequired bool   `json:"pnr_change_authentication_required"`
		PnrContactChangeAllowed         bool   `json:"pnr_contact_change_allowed"`
		PnrDate                         string `json:"pnr_date"`
		PnrKey                          string `json:"pnr_key"`
		PnrLocator                      string `json:"pnr_locator"`
		PosCountry                      string `json:"pos_country"`
		PosCountryBlocked               bool   `json:"pos_country_blocked"`
		RefundAllowed                   bool   `json:"refund_allowed"`
		Seats                           []any  `json:"seats"`
		Services                        struct {
			AdditionalInfo any `json:"additional_info"`
			Co2            struct {
				Method string `json:"method"`
				Urls   []struct {
					SegmentNumber int    `json:"segment_number"`
					URL           string `json:"url"`
				} `json:"urls"`
			} `json:"co2"`
			Miles struct {
				Method string `json:"method"`
				Urls   []any  `json:"urls"`
			} `json:"miles"`
		} `json:"services"`
		ShoulderbeltAllowed           bool   `json:"shoulderbelt_allowed"`
		ShoulderbeltShowControl       bool   `json:"shoulderbelt_show_control"`
		SpecialMealAllowed            bool   `json:"special_meal_allowed"`
		SpecialMeals                  []any  `json:"special_meals"`
		SportBaggageAllowed           bool   `json:"sport_baggage_allowed"`
		TicketsDocDownloadURL         string `json:"tickets_doc_download_url"`
		TicketsDocPrintURL            string `json:"tickets_doc_print_url"`
		TicketsDocSendURL             string `json:"tickets_doc_send_url"`
		TimeLimitExpired              bool   `json:"time_limit_expired"`
		Upcl4CrtAllowed               bool   `json:"upcl4crt_allowed"`
		Upcl4MlsAllowed               bool   `json:"upcl4mls_allowed"`
		UpclluckyMilesAllowed         bool   `json:"upcllucky_miles_allowed"`
		UpclluckyMoneyAllowed         bool   `json:"upcllucky_money_allowed"`
		VouchersRequestShowControl    bool   `json:"vouchers_request_show_control"`
		WalletDownloadURL             string `json:"wallet_download_url"`
		WalletSendURL                 string `json:"wallet_send_url"`
		WarningOfPobedaSelectSeats    string `json:"warning_of_pobeda_select_seats"`
		WarningOfSpecialCarryonRules  string `json:"warning_of_special_carryon_rules"`
		WarningOfUpdatingPersonalData string `json:"warning_of_updating_personal_data"`
		WarningOnItineraryChange      string `json:"warning_on_itinerary_change"`
		Warnings                      []struct {
			Description string `json:"description"`
			Title       string `json:"title"`
			URL         string `json:"url"`
		} `json:"warnings"`
	} `json:"data"`
	Error   any  `json:"error"`
	Success bool `json:"success"`
}
