package main

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//kita bisa menggunakan attack yang meningkat seperti ini untuk stress/capacity test

//melakukan attack selama 4 detik
//setiap detik frequency akan naik sesuai increaseValue
//nilai yang digunakan
/*
duration := 4 * time.Second
frequency := 1
increaseValue := 2
*/
func increaseTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 4 * time.Second
	frequency := 1
	increaseValue := 2
	rate := vegeta.ConstantPacer{Freq: frequency, Per: time.Second}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	})
	metrics = vegetaAttackIncreaseBySecond(targeter, rate, duration, increaseValue)
	return metrics
}

//frequency request akan naik sebesar increaseValue
//diawal tidak perlu menambahkan increaseValue
//contoh
//pertama 10 request
//kedua 15 request
//ketiga 20 request
//total 45 request

func vegetaAttackIncreaseBySecond(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration, increaseValue int) *vegeta.Metrics {
	var metrics vegeta.Metrics
	// TODO: answer here
	for i := 0; i < 3; i++ {
		rate.Freq = rate.Freq + i
	}
	/*for i := 0; i < int(duration.Seconds()); i++ {
		if i != 0 {
			rate.Freq = rate.Freq + increaseValue
		}
	}*/
	for res := range vegeta.NewAttacker().Attack(targeter, rate, duration, "Example") {
		metrics.Add(res)
		rate.Freq += increaseValue

	}
	metrics.Close()
	return &metrics
}
