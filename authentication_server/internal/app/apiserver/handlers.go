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

func (s *apiserver) WriteRawResponse(msg []byte, code int, rw http.ResponseWriter) {
	rw.WriteHeader(code)
	rw.Write(msg)
}

func (s *apiserver) HandleUsers() http.HandlerFunc {
	type UpdateRequest struct {
		TargetUser    *model.User `json:"target_user"`
		UpdatedValues *model.User `json:"updated_values"`
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET": // READ
			//getting information about users
			// Creating new user model
			model := model.NewUser()
			//Decoding json request to user model and cathing error
			if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
				s.WriteError(err, http.StatusUnprocessableEntity, rw) // Writing error to response and logging it
				return
			}
			models, err := s.store.GetUserByModel(model)
			if err != nil {
				s.WriteError(err, http.StatusInternalServerError, rw)
				return
			}
			result, err := json.Marshal(models)
			if err != nil {
				s.WriteError(err, http.StatusInternalServerError, rw)
				return
			}
			s.WriteRawResponse(result, 200, rw)
		case "POST": // CREATE
			model := model.NewUser()

			if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
				s.WriteError(err, http.StatusUnprocessableEntity, rw) // Writing error to response and logging it
				return
			}
			model.Validate()
			model.Prepare()
			err := s.store.CreateUserWithModel(model)
			if err != nil {
				s.WriteError(err, http.StatusBadRequest, rw)
				return
			}
			response, err := json.Marshal(model)
			if err != nil {
				s.WriteError(err, http.StatusBadRequest, rw)
				return
			}
			s.WriteRawResponse(response, http.StatusCreated, rw)
			return
		case "PUT": // UPDATE
			ur := &UpdateRequest{}
			err := json.NewDecoder(r.Body).Decode(ur)
			if err != nil {
				s.WriteError(err, http.StatusUnprocessableEntity, rw)
				return
			}
			if err := s.store.UpdateUserWithModel(ur.TargetUser, ur.UpdatedValues); err != nil {
				s.WriteError(err, http.StatusBadRequest, rw)
				return
			}
			response, err := json.Marshal(ur.UpdatedValues)
			if err != nil {
				s.WriteError(err, http.StatusBadRequest, rw)
				return
			}
			s.WriteRawResponse(response, http.StatusCreated, rw)
			return
		case "DELETE": // DELETE
			model := model.NewUser()
			if err := json.NewDecoder(r.Body).Decode(model); err != nil {
				s.WriteError(err, http.StatusUnprocessableEntity, rw)
				return
			}
			err := s.store.DeleteUserWithModel(model)
			if err != nil {
				s.WriteError(err, http.StatusBadRequest, rw)
				return
			}
			response, err := json.Marshal(model)
			if err != nil {
				s.WriteError(err, http.StatusBadRequest, rw)
				return
			}
			s.WriteRawResponse(response, http.StatusCreated, rw)
			return
		}
	}

}
