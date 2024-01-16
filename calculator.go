package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var history []string

func main() {
	http.HandleFunc("/calc/soma/", handleCalculation(addition))
	http.HandleFunc("/calc/sub/", handleCalculation(subtraction))
	http.HandleFunc("/calc/mul/", handleCalculation(multiplication))
	http.HandleFunc("/calc/div/", handleCalculation(division))
	http.HandleFunc("/calc/historico", handleHistory)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleCalculation(operation func(float64, float64) float64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.Path, "/")
		a, _ := strconv.ParseFloat(urlParts[3], 64)
		b, _ := strconv.ParseFloat(urlParts[4], 64)

		result := operation(a, b)

		operationName := strings.TrimSuffix(strings.TrimPrefix(urlParts[2], "calc/"), "/")
		historyEntry := fmt.Sprintf("%s: %.2f %s %.2f = %.2f", operationName, a, getOperatorSymbol(operationName), b, result)
		history = append(history, historyEntry)

		fmt.Fprintf(w, "%.2f", result)
	}
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	for _, entry := range history {
		fmt.Fprintln(w, entry)
	}
}

func addition(a, b float64) float64 {
	return a + b
}

func subtraction(a, b float64) float64 {
	return a - b
}

func multiplication(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	return 0
}

func getOperatorSymbol(operationName string) string {
	switch operationName {
	case "soma":
		return "+"
	case "sub":
		return "-"
	case "mul":
		return "*"
	case "div":
		return "/"
	default:
		return ""
	}
}
