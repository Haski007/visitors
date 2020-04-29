package handlers

import (
	"net/http"
)

// DoNothing - prevents twice call of handle func.
func DoNothing(w http.ResponseWriter, r *http.Request) {}