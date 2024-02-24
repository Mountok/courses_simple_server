package middleware

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func RequestLog(next http.Handler) http.Handler {
	var wg sync.WaitGroup

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s Method: %s", r.RequestURI, r.Method)
		wg.Add(1)
		logInFile(fmt.Sprintf("%s %s | Request: %s Method: %s",formatDate(),formatTime(), r.RequestURI, r.Method),&wg)
		wg.Wait()
		next.ServeHTTP(w, r) //главная функция для продолжения работы, без нее ваш Middleware поломает ответы на запросы, поскольку не передаст управление функциям из Router
	}) // задача Handler - обработка запросов, поэтому Middleware должен вернуть handler, мы используем HandlerFunc для простоты
}

func logInFile(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Unable to open log file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "%s\n", str)
		if err != nil {
			log.Println("Unable to write to log file", http.StatusInternalServerError)
			return
		}

		log.Println("Data logged successfully")
}

func formatDate() string {
	date := time.Now()
	return fmt.Sprintf("%02d/%02d/%d", date.Day(), date.Month(), date.Year())
}
func formatTime() string {
	time := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d", time.Hour(), time.Minute(), time.Second())
}
