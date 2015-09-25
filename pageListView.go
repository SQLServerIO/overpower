package main

import (
	"fmt"
	"mule/planetattack/attack"
	"net/http"
)

var TPLIST = MixTemp("frame", "listview", "listcoord")

func (g *Game) FactionView(w http.ResponseWriter, r *http.Request, f *attack.Faction) {
	if r.Method == "POST" {
		action := r.FormValue("action")
		var err error
		switch action {
		case "recenter":
			err = UserRecenter(r, f)
		case "setlorder":
			err = UserSetLaunchOrder(r, f)
		default:
			http.Error(w, "Action "+action+" not found", http.StatusInternalServerError)
			return
		}
		if err == nil {
			//g.Save()
		} else {
			fmt.Println(err)
		}
		http.Redirect(w, r, r.URL.Path, http.StatusFound)
		return
	}
	Apply(TPLIST, w, f)
}

func UserRecenter(r *http.Request, f *attack.Faction) error {
	m, err := GetInts(r, "adjX", "adjY")
	if err != nil {
		return err
	}
	f.CenterTV([2]int{m["adjX"], m["adjY"]})
	return nil
}

func UserSetLaunchOrder(r *http.Request, f *attack.Faction) error {
	m, err := GetInts(r, "amount", "tarID", "sourceID")
	if err != nil {
		return err
	}
	amount := m["amount"]
	source, ok := f.GetPlanetView(m["sourceID"])
	if !ok || !source.Yours {
		return makeE("Bad launch order:", m["sourceID"], "not found/owned by", f.Name)
	}
	if f.NumAvail(source.Location) < amount {
		return makeE("Bad launch order:", amount, "not avail at", source.Location)
	}
	target, ok := f.GetPlanetView(m["tarID"])
	if !ok {
		return makeE("Bad launch order:", m["tarID"], "not found by", f.Name)
	}
	f.SetOrder(amount, source, target)
	return nil
}
