package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	flagFormatStr       bool
	flagFormatTimestamp bool
	flagFormatInt       bool
	flagFormatCustom    string
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "something about time",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		format := "2006-01-02 15:04:05"

		var time1 time.Time
		if flagFormatStr {
			time1 = now
			if len(args) > 0 {
				i, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					fmt.Println("invalid timestamp")
					return
				}
				time1 = time.Unix(i, 0)
			}
			fmt.Println(time1.Format(format))
			return
		}
		if flagFormatTimestamp || flagFormatInt {
			time1 = now
			if len(args) > 0 {
				t, err := time.Parse(format, args[0])
				if err != nil {
					fmt.Println("invalid time string")
					return
				}
				time1 = t
			}
			fmt.Println(time1.Unix())
			return
		}
		if flagFormatCustom != "" {
			time1 = now
			if len(args) > 0 {
				i, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					fmt.Println("invalid timestamp")
					return
				}
				time1 = time.Unix(i, 0)
			}
			fmt.Println(time1.Format(flagFormatCustom))
			return
		}

		fmt.Println(now)
		return
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().BoolVarP(&flagFormatStr, "string", "s", false, "show time with format string")
	timeCmd.Flags().BoolVarP(&flagFormatTimestamp, "timestamp", "t", false, "show time with format timestamp")
	timeCmd.Flags().BoolVarP(&flagFormatInt, "int", "i", false, "show time with format timestamp")
	timeCmd.Flags().StringVarP(&flagFormatCustom, "format", "f", "", "show time with custom format")
}
