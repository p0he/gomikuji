package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Omikuji struct {
	 Fortune string
}

var FortuneMap = map[int]string{
	1: "Excellent luck",
	2: "Very good luck",
	3: "Good luck",
	4: "Slightly good luck",
	5: "Uncertain luck",
	6: "Bad luck",
}

func main() {
	http.HandleFunc("/omikuji", omikujiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	var omikuji Omikuji
	fortune := getFortune()

	omikuji.Fortune = fortune

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	if err := t.Execute(w, omikuji); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func getFortune() string {
	rand.Seed(time.Now().Unix())
	fortune, _ := FortuneMap[rand.Intn(10)]

	return fortune
}