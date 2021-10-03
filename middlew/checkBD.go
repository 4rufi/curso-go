package middlew

import (
	"net/http"

	"github.com/4rufi/curso-go/db"
)

/* CheckDB - middleware for validate state mongoDB */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if	!db.ConnectionCheck() {
			http.Error(w, "Conection DB loss", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}