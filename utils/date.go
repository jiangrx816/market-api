package utils

import (
	"log"
	"time"
)

const (
	DEFAULT_LAYOUT_DATE_TIME = "2006-01-02 15:04:05"
	DEFAULT_LAYOUT_DATE_TIME_1 = "2006-01-02 15:04"
	DEFAULT_LAYOUT_DATE      = "2006-01-02"
	DEFAULT_LAYOUT_DATE_YMD  = "20060102"
)

// GetCurrentDateTime 获取当前日期时间
func GetCurrentDateTime() (dateTime string) {
	dateTime = time.Now().Format(DEFAULT_LAYOUT_DATE_TIME)
	return
}

// GetCurrentDate 获取当前日期Y-M-D
func GetCurrentDate() (dateTime string) {
	dateTime = time.Now().Format(DEFAULT_LAYOUT_DATE)
	return
}

// GetCurrentDateYMD 获取当前日期YMD
func GetCurrentDateYMD() (dateTime string) {
	dateTime = time.Now().Format(DEFAULT_LAYOUT_DATE_YMD)
	return
}

// GetCurrentUnixTimestamp 获取当前的时间戳
func GetCurrentUnixTimestamp() (timestamp int64) {
	timestamp = time.Now().Unix()
	return
}

// GetDateToUnixTimestamp 日期时间格式转换成秒时间戳
func GetDateToUnixTimestamp(inputDateTime string) (timestamp int64) {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai") //指定时区
	dateTime, err := time.ParseInLocation(DEFAULT_LAYOUT_DATE_TIME, inputDateTime, TimeLocation)
	if err != nil {
		return
	}
	timestamp = dateTime.Unix()
	return
}

// GetDateToUnixNanoTimestamp 日期时间格式转换成秒时间戳
func GetDateToUnixNanoTimestamp(inputDateTime string) (timestamp int64) {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai") //指定时区
	dateTime, err := time.ParseInLocation(DEFAULT_LAYOUT_DATE_TIME, inputDateTime, TimeLocation)
	if err != nil {
		return
	}
	timestamp = dateTime.UnixNano()
	return
}

//GetUnixTimeToDate 时间戳转日期时间
func GetUnixTimeToDateTime(timestamp int64) (dateTime string) {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai") //指定时区
	dateTime = time.Unix(timestamp, 0).In(TimeLocation).Format(DEFAULT_LAYOUT_DATE_TIME)
	return
}

//GetUnixTimeToDate 时间戳转日期时间
func GetUnixTimeToDateTime1(timestamp int64) (dateTime string) {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai") //指定时区
	dateTime = time.Unix(timestamp, 0).In(TimeLocation).Format(DEFAULT_LAYOUT_DATE_TIME_1)
	return
}

//GetUnixTimeToDate 时间戳转日期Y-M-D
func GetUnixTimeToDate(timestamp int64) (dateTime string) {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai") //指定时区
	dateTime = time.Unix(timestamp, 0).In(TimeLocation).Format(DEFAULT_LAYOUT_DATE)
	return
}

//GetUnixTimeToDateYMD 时间戳转日期YMD
func GetUnixTimeToDateYMD(timestamp int64) (dateTime string) {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai") //指定时区
	dateTime = time.Unix(timestamp, 0).In(TimeLocation).Format(DEFAULT_LAYOUT_DATE_YMD)
	return
}

//defer GetTimeCost(time.Now(), "GetUnixTimeToDateYMD")
func GetTimeCost(start time.Time, tips string) {
	tc := time.Since(start)
	log.Printf("%s Time Const --> %#v \n", tips, tc)
}