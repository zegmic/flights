package server

import "testing"

func TestSingleLeg(t *testing.T) {
	l := []Leg{
		{
			"SFO",
			"EWR",
		},
	}

	var s RouteServer
	src, dst, err := s.findSrcDst(l)
	if err != nil {
		t.Error(err)
	}
	if src != "SFO" {
		t.Fail()
	}
	if dst != "EWR" {
		t.Fail()
	}
}

func TestTwoLegs(t *testing.T) {
	l := []Leg{
		{
			"ATL",
			"EWR",
		},
		{
			"SFO",
			"ATL",
		},
	}

	var s RouteServer
	src, dst, err := s.findSrcDst(l)
	if err != nil {
		t.Error(err)
	}

	if src != "SFO" {
		t.Fail()
	}
	if dst != "EWR" {
		t.Fail()
	}
}

func TestLongerPath(t *testing.T) {
	l := []Leg{
		{
			"IND",
			"EWR",
		},
		{
			"SFO",
			"ATL",
		},
		{
			"GSO",
			"IND",
		},
		{
			"ATL",
			"GSO",
		},
	}

	var s RouteServer
	src, dst, err := s.findSrcDst(l)
	if err != nil {
		t.Error(err)
	}

	if src != "SFO" {
		t.Fail()
	}
	if dst != "EWR" {
		t.Fail()
	}
}

func TestBrokenPath(t *testing.T) {
	l := []Leg{
		{
			"SFO",
			"ATL",
		},
		{
			"ATL",
			"JFK",
		},
		{
			"GSO",
			"EWR",
		},
	}

	var s RouteServer
	_, _, err := s.findSrcDst(l)
	if err == nil {
		t.Fail()
	}
}
