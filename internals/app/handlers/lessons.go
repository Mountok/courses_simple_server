package handlers

import (
	"courses_serve/example/internals/app/processors"
	"net/http"
	"strconv"
	"strings"
)

type LessonsHandler struct {
	processor *processors.LesssonsProgessor
}

func NewLessonsHandler(processor *processors.LesssonsProgessor) *LessonsHandler {
	handler := new(LessonsHandler)
	handler.processor = processor
	return handler
}

func (handler *LessonsHandler) List(w http.ResponseWriter, r *http.Request) {
	idstr := strings.Trim(r.URL.Query().Get("id"),"\"")

	id, err := strconv.Atoi(idstr)
	if err != nil {WrapError(w,err)}

	list, err := handler.processor.ListLessons(id)
	if err != nil {WrapError(w,err)}

	var m = map[string]interface{} {
		"result": "OK",
		"data": list,
	}

	WrapOK(w,m)
}

func (handler *LessonsHandler) Lesson(w http.ResponseWriter, r *http.Request) {
	idstr := strings.Trim(r.URL.Query().Get("subject_id"),"\"")
	subject_id, err := strconv.Atoi(idstr)
	if err != nil { WrapError(w,err) }
	idstr = strings.Trim(r.URL.Query().Get("id"),"\"")
	lesson_id, err := strconv.Atoi(idstr)
	if err != nil { WrapError(w,err) }

	list, err := handler.processor.LessonContent(subject_id,lesson_id)
	if err != nil {
		WrapError(w,err)
	}
	var m = map[string]interface{} {
		"result": "OK",
		"data": list,
	}

	WrapOK(w,m)
}