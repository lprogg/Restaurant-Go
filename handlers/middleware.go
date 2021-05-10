package handlers

import (
	"Ex5Validation/data"
	"context"
	"fmt"
	"net/http"
)

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			p.l.Println("Unable to validate product", err)
			http.Error(
				w,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), &KeyProduct{}, product)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
