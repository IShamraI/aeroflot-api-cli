/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	default_direction string = "from"
)


// exchangeCmd represents the exchange command
var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Search for exchange rates",
	Long: `Search all available exchanges and payments for specified days since fly`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exchange called")
	},
}

func init() {
	rootCmd.AddCommand(exchangeCmd)

	exchangeCmd.PersistentFlags().String("booking-code", "SFSHFR", "Booking code")
	exchangeCmd.PersistentFlags().String("last-name", "IVANOV", "Last name in Latin letters")
	exchangeCmd.PersistentFlags().String("direction", default_direction, "Direction of flight to look for exchange")
	exchangeCmd.PersistentFlags().Int("period", 10, "How many days to look for exchange since date of flight")

}
