package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getSession, err := r.Cookie("session_token")
		
		if err != nil {
			if err == http.ErrNoCookie {
				res := model.ErrorResponse{Error: "http: named cookie not present"}
				w.WriteHeader(http.StatusUnauthorized)

				if err := json.NewEncoder(w).Encode(res); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			res := model.ErrorResponse{Error: "Bad Request!"}
			if err := json.NewEncoder(w).Encode(res); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		
		
		sessionToken := getSession.Value // TODO: replace this

		sessionFound, err := api.sessionsRepo.CheckExpireToken(sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.Method
		if err != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Error: "Method is not allowed!",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.Method
		if err != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Error: "Method is not allowed!",
			})
			return
		}


		next.ServeHTTP(w, r)
	})
}
