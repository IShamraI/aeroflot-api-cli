/*
Copyright © 2023 Ivan Shamrai isshamray@gmail.com
*/
package cmd

import (
	"aeroflot-api-cli/internal/booking"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// exchangeCmd represents the exchange command
var exchangeCmd = &cobra.Command{
	Use:   "exchange [flags] booking_code last_name first_name",
	Args:  cobra.MinimumNArgs(3),
	Short: "Search for exchange rates",
	Long:  `Search all available exchanges and payments for specified days since fly`,
	Run: func(cmd *cobra.Command, args []string) {
		booking_code := args[0]
		log.Printf("Booking code %s", booking_code)

		last_name := args[1]
		log.Printf("Last name: %s", last_name)

		first_name := args[2]
		log.Printf("First name: %s", first_name)

		period, _ := cmd.Flags().GetInt("period")

		log.Printf("Looking avaliable flights for booking %s in %d days", booking_code, period)

		booking, _ := booking.New(booking_code, last_name, first_name)
		booking.GetInfo()
		departure, _ := time.Parse("2006-01-02", booking.Departure)
		end := departure.AddDate(0, 0, 1+period)
		for d := departure.AddDate(0, 0, 1); d.After(end) == false; d = d.AddDate(0, 0, 1) {
			booking.GetFairPrices(d.Format("2006-01-02"))
		}
	},
}

func init() {
	rootCmd.AddCommand(exchangeCmd)

	exchangeCmd.PersistentFlags().IntP("period", "p", 10, "how many days to look for exchange since date of flight")
}
