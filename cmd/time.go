package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nosixtools/solarlunar"
	"github.com/nosixtools/solarlunar/festival"
	"github.com/spf13/cobra"
)

var (
	flagFormatStr       bool
	flagFormatTimestamp bool
	flagFormatCustom    string
	flagWeek            bool
	flagNongli          bool
	flagYangli          bool
	flagRun             bool
	flagJieri           bool
	flagOrigin          bool
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "something about time",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		format1 := "2006-01-02 15:04:05"
		format2 := "2006-01-02"
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
			fmt.Println(time1.Format(format1))
			return
		}
		if flagFormatTimestamp {
			time1 = now
			if len(args) > 0 {
				t, err := time.Parse(format1, args[0])
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

		if flagWeek {
			time1 = now
			if len(args) > 0 {
				t, err := time.Parse(format2, args[0])
				if err != nil {
					fmt.Println("invalid time string")
					return
				}
				time1 = t
			}
			fmt.Println(time1.Weekday())
			return
		}

		if flagJieri {
			time1 = now
			festival := festival.NewFestival("./festival.json")
			if len(args) > 0 {
				t, err := time.Parse(format2, args[0])
				if err != nil {
					fmt.Println("invalid time string")
					return
				}
				time1 = t
				if flagNongli {
					yl := solarlunar.LunarToSolar(time1.Format(format2), flagRun)
					fmt.Println(festival.GetFestivals(yl))
					return
				}
			}
			fmt.Println(festival.GetFestivals(time1.Format(format2)))
			return
		}

		if flagNongli {
			time1 = now
			if len(args) > 0 {
				t, err := time.Parse(format2, args[0])
				if err != nil {
					fmt.Println("invalid time string")
					return
				}
				time1 = t
			}
			fmt.Println(solarlunar.SolarToChineseLuanr(time1.Format(format2)))
			return
		}

		if flagYangli {
			time1 = now
			if len(args) > 0 {
				t, err := time.Parse(format2, args[0])
				if err != nil {
					fmt.Println("invalid time string")
					return
				}
				time1 = t
				fmt.Println(solarlunar.LunarToSolar(time1.Format(format2), flagRun))
				return
			}

			fmt.Println(time1.Format(format2))
			return
		}

		if flagOrigin {
			time1 = now
			if len(args) > 0 {
				t, err := time.Parse(format1, args[0])
				if err != nil {
					fmt.Println("invalid time string")
					return
				}
				time1 = t
			}
			fmt.Println(time1)
			return
		}

		week := now.Weekday()
		nl := solarlunar.SolarToChineseLuanr(now.Format(format2))
		fmt.Println(now.Format(format1), nl, week)
		return
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().BoolVarP(&flagFormatStr, "string", "s", false, "show time with format string")
	timeCmd.Flags().BoolVarP(&flagFormatTimestamp, "timestamp", "t", false, "show time with format timestamp")
	timeCmd.Flags().StringVarP(&flagFormatCustom, "format", "f", "", "show time with custom format")
	timeCmd.Flags().BoolVarP(&flagWeek, "week", "w", false, "show weekday")
	timeCmd.Flags().BoolVarP(&flagNongli, "nl", "n", false, "show lunar calendar")
	timeCmd.Flags().BoolVarP(&flagYangli, "yl", "y", false, "solar calendar")
	timeCmd.Flags().BoolVarP(&flagRun, "run", "r", false, "run lunar calendar")
	timeCmd.Flags().BoolVarP(&flagJieri, "jr", "j", false, "show jieri")
	timeCmd.Flags().BoolVarP(&flagOrigin, "origin", "o", false, "show original time")
}
