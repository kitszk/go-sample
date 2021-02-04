package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("hello world.")

	now := time.Now()
	fmt.Println("Original:", now, ", Type:", reflect.TypeOf(now))
	// Original: 2021-02-04 23:59:04.5925843 +0900 JST m=+0.001999401 , Type: time.Time

	layout := "2006-01-02T15:04:05.999999Z07:00"
	formatedNow := now.Format(layout)
	fmt.Println("Formated:", formatedNow, ", Type:", reflect.TypeOf(formatedNow))
	// Formated: 2021-02-04T23:59:04.592584+09:00 , Type: string

	parsedNow, err := time.Parse(layout, formatedNow)
	fmt.Println("Parsed:", parsedNow, ", Type:", reflect.TypeOf(parsedNow))
	// Parsed: 2021-02-04 23:59:04.592584 +0900 JST , Type: time.Time
	if err != nil {
		fmt.Println("Err:", err)
	}

	now, err = time.Parse(layout, time.Now().Format(layout))
	fmt.Println("Merged:", now, ", Type:", reflect.TypeOf(now))
	// Merged: 2021-02-04 23:59:04.631546 +0900 JST , Type: time.Time
	if err != nil {
		fmt.Println("Err:", err)
	}

	roundedNow := now.Round(time.Microsecond)
	fmt.Println("Rounded:", roundedNow, ", Type:", reflect.TypeOf(roundedNow))
	// Rounded: 2021-02-04 23:59:04.631546 +0900 JST , Type: time.Time
}
