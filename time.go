package main

import (
	"fmt"
	"time"
)

func Maintance_3600() {
	Timer := 0
	start_time := time.Date(2022, 12, 13, 12, 57, 30, 30, time.UTC)
	finish_time := start_time.Add(3600 * time.Hour)
	fmt.Println("Tahmini 3600 saat  bakım tarihi:", finish_time)

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(finish_time)

		Timer++
		if start_time == finish_time {
			fmt.Println("Bakım zamanı geldi.")
			break
		} else {

			//fmt.Printf("Bakıma kalan süre: %02d\n", timeRemaining)
			fmt.Printf("Days: %02d Hours: %02d Minutes: %02d Seconds: %02d\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)
			if Timer == 10 {
				break
			}

		}

	}
}
func Maintance_7200() {
	Timer := 0
	start_time := time.Date(2022, 12, 13, 12, 57, 30, 30, time.UTC)
	finish_time := start_time.Add(7200 * time.Hour)
	fmt.Println("Tahmini 7200 saat  bakım tarihi:", finish_time)

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(finish_time)

		Timer++
		if start_time == finish_time {
			fmt.Println("Bakım zamanı geldi.")
			break
		} else {

			//fmt.Printf("Bakıma kalan süre: %02d\n", timeRemaining)
			fmt.Printf("Days: %02d Hours: %02d Minutes: %02d Seconds: %02d\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)
			if Timer == 10 {
				break
			}

		}

	}
}

func main() {
	var bakim string
	fmt.Println("giriş yapılan zaman:", time.Now())

	fmt.Println("a : 3600 saat bakımı \nb : 7200 saat bakımı")
	fmt.Printf("bakım türünü seçiniz: ")
	fmt.Scan(&bakim)

	switch bakim {
	case "a", "A":
		Maintance_3600()
		break
	case "b", "B":
		Maintance_7200()
		break

	}

}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
