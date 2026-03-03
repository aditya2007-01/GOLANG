package handler

import (
	"html/template"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		x, _ := strconv.ParseFloat(r.FormValue("x"), 64)
		y, _ := strconv.ParseFloat(r.FormValue("y"), 64)
		op := r.FormValue("op")

		var result float64
		var errMsg string

		switch op {
		case "+":
			result = x + y
		case "-":
			result = x - y
		case "*":
			result = x * y
		case "/":
			if y == 0 {
				errMsg = "Cannot divide by zero"
			} else {
				result = x / y
			}
		default:
			errMsg = "Invalid operation"
		}

		tpl.Execute(w, map[string]interface{}{
			"Result": result,
			"Error":  errMsg,
		})
		return
	}

	tpl.Execute(w, nil)
}
