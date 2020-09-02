package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/milQA/int-slice-crud-bst/internal/api/store"
	"go.uber.org/zap"
)

func (s *Server) insert() http.HandlerFunc {

	type InputData struct {
		Value int `json:"val"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.With(
			zap.String(
				requestIDKey, r.Context().Value(RequestIDKey{}).(string),
			),
		).Sugar()

		var inputData InputData

		if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Errorf("cannot unmarshal body: %s", err.Error())
			return
		}

		if err := s.service.IntSlice().Insert(inputData.Value); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Errorf("cannot save val = %v: %s", inputData.Value, err.Error())
			return
		}

		logger.Infof("save val = %v", inputData.Value)
		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) delete() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.With(
			zap.String(
				requestIDKey, r.Context().Value(RequestIDKey{}).(string),
			),
		).Sugar()

		val := r.FormValue("val")
		value, err := strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Errorf(
				"cannot parse val = %s: %s", val, err.Error(),
			)
			return
		}

		if err := s.service.IntSlice().Delete(value); err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			logger.Errorf(
				"cannot delete val = %v: %s", value, err.Error(),
			)
			return
		}

		logger.Infof("delete val = %v", value)
		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) search() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.With(
			zap.String(
				requestIDKey, r.Context().Value(RequestIDKey{}).(string),
			),
		).Sugar()

		val := r.FormValue("val")
		value, err := strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Errorf(
				"cannot parse val = %s: %s", val, err.Error(),
			)
			return
		}

		answer, err := s.service.IntSlice().Search(value)
		if err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			logger.Errorf(
				"cannot find val = %v: %s", value, err.Error(),
			)
			return
		}

		logger.Infof("find val = %v", answer)
		w.WriteHeader(http.StatusOK)
	})
}
