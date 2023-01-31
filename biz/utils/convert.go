package utils

import (
	"github.com/lutasam/doctors/biz/common"
	"strconv"
	"time"
)

func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func StringToUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, common.USERINPUTERROR
	}
	return i, nil
}

func StringToFloat32(s string) (float32, error) {
	temp, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, common.USERINPUTERROR
	}
	return float32(temp), nil
}

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, common.USERINPUTERROR
	}
	return i, nil
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StringToTime(s string) (time.Time, error) {
	parseTime, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return time.Time{}, common.USERINPUTERROR
	}
	return parseTime, nil
}

func TimeToDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func DateStringToTime(date string) (time.Time, error) {
	parseTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, common.USERINPUTERROR
	}
	return parseTime, nil
}
