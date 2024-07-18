package main

import (
	"fmt"
	"time"
)

func main() {
	//timeStr := "2023-11-30 11:00:00"
	//parseTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	//fmt.Println("after time.Parse, the location is : ", parseTime.Location(),
	//	"the actual time is : ", parseTime.String())

	now := time.Now()
	fmt.Println("now is: ", now.String())
	asiaLocation, _ := time.LoadLocation("Asia/Shanghai")
	asiaTime := now.In(asiaLocation)
	fmt.Println("asia time: ", asiaTime.String(), ", location: ", asiaTime.Location(), ", timestamp: ", asiaTime.Unix())

	americaLocation, _ := time.LoadLocation("America/Los_Angeles")
	americaTime := now.In(americaLocation)
	fmt.Println("america time: ", americaTime.String(), ", location: ", americaTime.Location(), ", timestamp: ", americaTime.Unix())

	utcTime := now.In(time.UTC)
	fmt.Println("utc time: ", utcTime.String(), ", location: ", utcTime.Location(), ", timestamp: ", utcTime.Unix())

	Review()

	formatDuration(2)
}

func Review() {
	// 北京时间
	sendTime := "2023-11-30 11:12:00"
	// 北京时间
	pushEndTime := "2023-11-30 19:10:00"
	parseTime, _ := time.Parse("2006-01-02 15:04:05", pushEndTime)
	fmt.Println("pushEndTime after time.Parse is : ", parseTime.String(), ", timestamp is : ", parseTime.Unix())
	parseInLocalTime, _ := time.ParseInLocation("2006-01-02 15:04:05", pushEndTime, time.Local)
	fmt.Println("pushEndTime after time.ParseInLocation is : ", parseInLocalTime.String(), ", timestamp is : ", parseInLocalTime.Unix())
	// get pushEndTime timestamp
	utcTimeStamp := parseTime.Unix()
	cstTimeStamp := parseInLocalTime.Unix()
	parseSendTime, _ := time.Parse("2006-01-02 15:04:05", sendTime)
	fmt.Println("sendTime after time.Parse Location : ", parseSendTime.Location())
	fmt.Println("cstTimeStamp: ", cstTimeStamp, ", utcTimeStamp: ", utcTimeStamp, " sendTimeStamp: ", parseSendTime.Unix())

	d, _ := time.ParseDuration("7h8m30s")
	fmt.Println(d)

	start := time.Now()
	time.Sleep(3 * time.Second)
	duration := time.Since(start)
	fmt.Println("since start, duration : ", duration)
}

func formatDuration(d time.Duration) {
	if d < time.Second {
		fmt.Println(d)
	}
}

func GetStartTimeOfDayWithInterval(interval int64, location *time.Location) (t1, t2 time.Time) {
	if location == nil || interval < 0 {
		return
	}

	now := time.Now().In(location)
	yearOfNow, monthOfNow, dayOfNow := now.Date()

	before := now.Add(-1 * time.Second * time.Duration(interval))
	yearOfBefore, monthOfBefore, dayOfBefore := before.Date()

	t1 = time.Date(yearOfBefore, monthOfBefore, dayOfBefore, 0, 0, 0, 0, location)
	t2 = time.Date(yearOfNow, monthOfNow, dayOfNow-1, 23, 59, 59, 1e9-1, location)

	return
}
