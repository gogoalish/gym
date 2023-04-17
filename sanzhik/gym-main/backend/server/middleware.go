package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user *User
		var err error
		c, err := r.Cookie("access_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, nil)))
				return
			}
			h.badRequestResponse(w, r, err)
			return
		}

		user, err = h.Service.ParseToken(c.Value)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, nil)))
			return
		}
		if user.Expires.Before(time.Now()) {
			if err := h.Service.DeleteToken(c.Value); err != nil {
				h.errorResponse(w, r, 500, err)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, nil)))
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, user)))
	}
}

func (h *Handler) recoverPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")

				h.serverErrorResponse(w, r, nil)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func (h *Handler) authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// user := r.Context().Value(ctxKeyUser).(*User)
		u := r.Context().Value(ctxKeyUser)

		if u == nil {
			fmt.Println("middleware:authorized: user is not authorized")
			h.notPermittedResponse(w, r)
			return
		}

		// user := u.(*User)

		next.ServeHTTP(w, r)
	}
}
