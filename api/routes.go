package api

import (
	"courses_serve/example/internals/app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoute(
	lessonHandler *handlers.LessonsHandler, // хендлеры для роутинга
	subjectHandler *handlers.SubjectHandler,
) *mux.Router { // возвращаем указать на роутер
	router := mux.NewRouter() // создаем новый роутер для обработки путей (основной роутер для нашего сервера)

	// пример: http://localhost:8080/subjects/list
	router.HandleFunc("/subjects/list", subjectHandler.List).Methods(http.MethodGet) //каждая функция реализует один и тот же интерфейс
	
	// пример http://localhost:8080/lessons/list?id=35342
	router.HandleFunc("/lessons/list", lessonHandler.List).Methods(http.MethodGet) //Methods определяют какой глагол можно использовать, если будет другой - вернется 404
	
	// пример http://localhost:8080/lessons/content?subject_id=2&id=1
	router.HandleFunc("/lessons/content", lessonHandler.Lesson).Methods(http.MethodGet)

	router.NotFoundHandler = router.NewRoute().HandlerFunc(handlers.NotFound).GetHandler() //оборачиваем 404, для обработки NotFound
 
	return router
}
