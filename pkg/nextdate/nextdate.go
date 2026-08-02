package nextdate

import (
	"strconv"
	"strings"
	"time"

	"github.com/nickchervov/go-diplom-project/internal/domain"
)

var weekdayNumber = map[string]int{
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
	"Sunday":    7,
}

var monthDays = map[string]int{
	"January":   31,
	"February":  28,
	"March":     31,
	"April":     30,
	"May":       31,
	"June":      30,
	"July":      31,
	"August":    31,
	"September": 30,
	"October":   31,
	"November":  30,
	"December":  31,
}

var monthNumber = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func NextDate(now time.Time, dstart, repeat string) (string, error) {
	var acceptedDaysOfWeek = make(map[int]bool)
	var acceptedDaysOfMonth = make(map[int]bool)
	var acceptedMonth = make(map[int]bool)

	if now.Year()%4 == 0 {
		monthDays["February"] = 29
	} else {
		monthDays["February"] = 28
	}

	repeatRules := strings.Split(repeat, " ")

	switch repeatRules[0] {
	case "d":
		if len(repeatRules) != 2 {
			return "", domain.ErrIncorrectRepeatRule
		}

		days, err := strconv.Atoi(repeatRules[1])
		if err != nil {
			return "", domain.ErrIncorrectRepeatRule
		}

		if days > 400 {
			return "", domain.ErrIncorrectRepeatRule
		}

		date, err := time.Parse("20060102", dstart)
		if err != nil {
			return "", domain.ErrIncorrectDate
		}

		for {
			date = date.AddDate(0, 0, days)
			if date.After(now) {
				break
			}
		}

		return date.Format("20060102"), nil
	case "y":
		if len(repeatRules) > 1 {
			return "", domain.ErrIncorrectRepeatRule
		}

		date, err := time.Parse("20060102", dstart)
		if err != nil {
			return "", domain.ErrIncorrectDate
		}

		for {
			date = date.AddDate(1, 0, 0)
			if date.After(now) {
				break
			}
		}
		return date.Format("20060102"), nil
	case "w":
		if len(repeatRules) != 2 {
			return "", domain.ErrIncorrectRepeatRule
		}

		date, err := time.Parse("20060102", dstart)
		if err != nil {
			return "", domain.ErrIncorrectDate
		}

		days := strings.Split(repeatRules[1], ",")
		for _, d := range days {
			numDay, err := strconv.Atoi(d)
			if err != nil {
				return "", domain.ErrIncorrectRepeatRule
			}
			if numDay > 7 || numDay < 1 {
				return "", domain.ErrIncorrectRepeatRule
			}
			acceptedDaysOfWeek[numDay] = true
		}

		for {
			date = date.AddDate(0, 0, 1)
			dayWeek := weekdayNumber[date.Weekday().String()]
			if acceptedDaysOfWeek[dayWeek] {
				if date.After(now) {
					break
				}
			}
		}
		return date.Format("20060102"), nil
	case "m":
		if len(repeatRules) > 3 {
			return "", domain.ErrIncorrectRepeatRule
		}

		if len(repeatRules) == 3 {
			monthDay := strings.Split(repeatRules[1], ",")
			month := strings.Split(repeatRules[2], ",")
			isLastDay := false
			isPreLastDay := false
			for _, md := range monthDay {
				numMd, err := strconv.Atoi(md)
				if err != nil {
					return "", domain.ErrIncorrectRepeatRule
				}
				if numMd < -2 || numMd > 31 || numMd == 0 {
					return "", domain.ErrIncorrectRepeatRule
				}
				if numMd == -1 {
					isLastDay = true
					continue
				}
				if numMd == -2 {
					isPreLastDay = true
					continue
				}
				acceptedDaysOfMonth[numMd] = true
			}
			for _, m := range month {
				numM, err := strconv.Atoi(m)
				if err != nil {
					return "", domain.ErrIncorrectRepeatRule
				}

				if numM > 12 || numM < 1 {
					return "", domain.ErrIncorrectRepeatRule
				}

				acceptedMonth[numM] = true
			}

			date, err := time.Parse("20060102", dstart)
			if err != nil {
				return "", domain.ErrIncorrectDate
			}

			if isPreLastDay {
				for {
					date = date.AddDate(0, 0, 1)
					if acceptedMonth[monthNumber[date.Month().String()]] {
						if acceptedDaysOfMonth[date.Day()] {
							if date.After(now) {
								break
							}
						}
						if date.Day() == monthDays[date.Month().String()]-1 {
							if date.After(now) {
								break
							}
						}
					}
				}
				return date.Format("20060102"), nil
			}

			if isLastDay {
				for {
					date = date.AddDate(0, 0, 1)
					if acceptedMonth[monthNumber[date.Month().String()]] {
						if acceptedDaysOfMonth[date.Day()] {
							if date.After(now) {
								break
							}
						}
						if date.Day() == monthDays[date.Month().String()] {
							if date.After(now) {
								break
							}
						}
					}
				}
				return date.Format("20060102"), nil
			}

			for {
				date = date.AddDate(0, 0, 1)
				if acceptedMonth[monthNumber[date.Month().String()]] {
					if acceptedDaysOfMonth[date.Day()] {
						if date.After(now) {
							break
						}
					}
				}
			}
			return date.Format("20060102"), nil

		} else {
			monthDay := strings.Split(repeatRules[1], ",")
			isLastDay := false
			isPreLastDay := false
			for _, md := range monthDay {
				numMd, err := strconv.Atoi(md)
				if err != nil {
					return "", domain.ErrIncorrectRepeatRule
				}
				if numMd < -2 || numMd > 31 || numMd == 0 {
					return "", domain.ErrIncorrectRepeatRule
				}
				if numMd == -1 {
					isLastDay = true
					continue
				}
				if numMd == -2 {
					isPreLastDay = true
					continue
				}
				acceptedDaysOfMonth[numMd] = true
			}

			date, err := time.Parse("20060102", dstart)
			if err != nil {
				return "", domain.ErrIncorrectDate
			}

			if isPreLastDay {
				for {
					date = date.AddDate(0, 0, 1)
					if acceptedDaysOfMonth[date.Day()] {
						if date.After(now) {
							break
						}
					}
					if date.Day() == monthDays[date.Month().String()]-1 {
						if date.After(now) {
							break
						}
					}
				}
				return date.Format("20060102"), nil
			}

			if isLastDay {
				for {
					date = date.AddDate(0, 0, 1)
					if acceptedDaysOfMonth[date.Day()] {
						if date.After(now) {
							break
						}
					}
					if date.Day() == monthDays[date.Month().String()] {
						if date.After(now) {
							break
						}
					}
				}
				return date.Format("20060102"), nil
			}

			for {
				date = date.AddDate(0, 0, 1)
				if acceptedDaysOfMonth[date.Day()] {
					if date.After(now) {
						break
					}
				}
			}
			return date.Format("20060102"), nil
		}

	default:
		return "", domain.ErrIncorrectRepeatRule
	}
}
