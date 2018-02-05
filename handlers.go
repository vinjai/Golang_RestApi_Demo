package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome amigo!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := Todos{
		Todo{Name: "Write Presentation"},
		Todo{Name: "Host meetup"},
		Todo{Name: "Bridge Session"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo Show: ", todoId)
}
