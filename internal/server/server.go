package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"flights/internal/route"
)

type Leg struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type legRequest struct {
	Legs []Leg
}

type RouteServer struct {
}

func (r RouteServer) Find(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	var req legRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusOK)

	src, dst, err := r.findSrcDst(req.Legs)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(writer).Encode(Leg{
		Source:      src,
		Destination: dst,
	})
}

func (r RouteServer) findSrcDst(legs []Leg) (string, string, error) {
	if len(legs) == 0 {
		return "", "", errors.New("no legs provided")
	}
	if len(legs) == 1 {
		return legs[0].Source, legs[0].Destination, nil
	}

	var ro route.Route
	for _, l := range legs {
		ro.AddLeg(l.Source, l.Destination)
	}

	if !ro.IsValid() {
		return "", "", errors.New("inconsistent route")
	}
	return ro.Source(), ro.Destination(), nil
}
