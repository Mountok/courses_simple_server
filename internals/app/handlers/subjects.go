package handlers

import (
	"courses_serve/example/internals/app/processors"
	"net/http"
)

type SubjectHandler struct {
	processor *processors.SubjectsProcessor
}

func NewSubjectsHandler(processor *processors.SubjectsProcessor) *SubjectHandler {
	handler := new(SubjectHandler)
	handler.processor = processor
	return handler
}

func (handler *SubjectHandler) List(w http.ResponseWriter, r *http.Request) {

	list, err := handler.processor.ListSubjects()
	if err != nil {
		WrapError(w,err)
		return
	}
	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}
	WrapOK(w,m)
}
