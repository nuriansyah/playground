package main

import (
	"encoding/json"
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Movie struct {
	ID      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

//Baca README untuk tau jumlah request yang perlu dilakukan dan targetnya
//untuk durasi cukup gunakan satu detik

//menambahkan movie baru
//untuk data yang dikirim adalah JSON
//gunakan struct Movie diatas, cukup gunakan field episode dan name
//ID sudah auto increment
func addMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	result := make([]Movie, 0)
	for i := 1; i <= 25; i++ {
		movie := Movie{
			ID:      i,
			Episode: i,
			Name:    fmt.Sprintf("Movie %d", i),
		}
		result = append(result, movie)
	}
	body, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body:   body,
	})
	metrics = vegetaAttack(targeter, 10, time.Second)

	// fmt.Sprintf(`{"episode": %d, "name": "%s"}`, Movie.Episode, Movie.Name)
	// json.Marshal(Movie)

	return metrics
}

//mendapatkan informasi movie dengan ID 1-25
//vegeta.NewStaticTargeter() adalah variadic function
//kita bisa menggunakannya untuk menentukan multiple target vegeta attack
func getMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	for i := 1; i <= 25; i++ {
		json.Marshal(Movie{
			ID: i,
		})
	}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	})
	metrics = vegetaAttack(targeter, 25, time.Second)
	return metrics
}

//mendapatkan semua informasi movie
func getMoviesTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	for i := 1; i <= 25; i++ {
		json.Marshal(Movie{
			ID:      i,
			Episode: i,
			Name:    fmt.Sprintf("Movie %d", i),
		})
	}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
	})
	metrics = vegetaAttack(targeter, 20, time.Second)

	return metrics
}

func vegetaAttack(targeter vegeta.Targeter, frequency int, duration time.Duration) *vegeta.Metrics {
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res)
	}
	metrics.Close()
	return &metrics
}
