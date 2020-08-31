package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/milQA/int-slice-crud-bst/internal/api/service"
	"github.com/milQA/int-slice-crud-bst/internal/api/store"
	"go.uber.org/zap"
)

func insert(repo store.Store, log *zap.Logger) http.HandlerFunc {

	intSliceService := service.NewIntSliceService(repo)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.With(
			zap.String(
				requestIDKey, r.Context().Value(RequestIDKey{}).(string),
			),
		).Sugar()

		vals := make([]int, 0, 30)

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Errorf("cannot read bytes from body: %s", err.Error())
			return
		}

		if err := json.Unmarshal(buf, &vals); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Errorf("cannot unmarshal body to vals = %s: %s", string(buf), err.Error())
			return
		}

		if len(vals) != 30 {
			w.WriteHeader(http.StatusBadRequest)
			logger.Errorf("invalid vals len = %v", vals)
			return
		}

		if err := intSliceService.Insert(vals...); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Errorf("cannot save vals = %v: %s", vals, err.Error())
			return
		}

		logger.Infof("save vals = %v", vals)
		w.WriteHeader(http.StatusOK)
	})
}

func delete(repo store.Store, log *zap.Logger) http.HandlerFunc {

	intSliceService := service.NewIntSliceService(repo)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.With(
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

		if err := intSliceService.Delete(value); err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
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

func search(repo store.Store, log *zap.Logger) http.HandlerFunc {

	intSliceService := service.NewIntSliceService(repo)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.With(
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

		answer, err := intSliceService.Search(value)
		if err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
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
