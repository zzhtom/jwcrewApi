package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	_ "strconv"
	"strings"

	"github.com/gorilla/mux"
)

func AdminInfo(w http.ResponseWriter, r *http.Request) {

	var admin Admin

	if strings.Contains(r.Method, "POST") {

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &admin); err != nil || strings.EqualFold("", admin.UserName) {
			w.WriteHeader(422) // unprocessable entity
			if err == nil {
				err = errors.New("username isn't null")
			}
			if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 422, Data: nil, Message: err.Error()}); err != nil {
				panic(err)
			}
			return
		}

		if rst, _ := admin.isExist(admin.UserName); rst {
			if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 205, Data: nil, Message: "sql: user have been exist."}); err != nil {
				panic(err)
			}

			return
		}
		if rst, err := admin.add(); rst {
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 200, Data: nil, Message: ""}); err != nil {
				panic(err)
			}
		} else {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 204, Data: nil, Message: err.Error()}); err != nil {
				panic(err)
			}
		}
		return
	}

	id := mux.Vars(r)["id"]

	w.WriteHeader(http.StatusOK)
	if rst, err := admin.isExist(id); !rst {
		if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 404, Data: nil, Message: err.Error()}); err != nil {
			panic(err)
		}

		return
	}
	if strings.Contains(r.Method, "DELETE") {
		if rst, err := admin.deleteById(id); err != nil && !rst {
			if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 204, Data: nil, Message: err.Error()}); err != nil {
				panic(err)
			}
		} else {
			if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 200, Data: nil, Message: ""}); err != nil {
				panic(err)
			}
		}

		return
	}

	if data, err := admin.querySingleRowById(id); err != nil {
		if err := json.NewEncoder(w).Encode(jwcrewApi{Code: http.StatusNoContent, Data: nil, Message: err.Error()}); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(jwcrewApi{Code: http.StatusOK, Data: data, Message: ""}); err != nil {
			panic(err)
		}
	}

}

func AdminsList(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	var admin Admin
	if data, err := admin.queryRows(); err != nil {
		if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 204, Data: nil, Message: err.Error()}); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(jwcrewApi{Code: 200, Data: data, Message: ""}); err != nil {
			panic(err)
		}
	}

}
