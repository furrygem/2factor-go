package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/furrygem/authentication_server/internal/app/model"
)

func (s *apiserver) WriteResponse(msg string, code int, rw http.ResponseWriter) {
	rw.WriteHeader(code)
	rw.Write([]byte(msg))
}

func (s *apiserver) WriteError(err error, code int, rw http.ResponseWriter) {
	s.logger.Error(err.Error())
	s.WriteResponse(http.StatusText(code), code, rw)
}

func (s *apiserver) HandleUsers() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			//getting information about users
			model := model.NewUser()
			//Decoding json request to user model and cathing error
			if err := json.NewDecoder(r.Body).Decode(model); err != nil {
				s.WriteError(err, http.StatusUnprocessableEntity, rw) // Writing error to response and logging it
			}
			s.store.GetUserByModel(model)
		}
	}
}
