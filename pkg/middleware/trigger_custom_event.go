package middleware

import (
	"net/http"

	"github.com/dmji/gosudarevlist/pkg/expose_header_utils"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func HxTriggerMiddleware(eventName string) func(handler http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := expose_header_utils.WriterExposeHeader(w, "HX-Trigger-After-Swap", eventName)
			if err != nil {
				logger.Errorw(r.Context(), "HxTrigger Middleware | WriteHxTrigger failed", "error", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			handler(w, r)
		}
	}
}
