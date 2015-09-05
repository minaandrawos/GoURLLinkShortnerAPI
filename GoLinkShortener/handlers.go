package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type LinkShortnerAPI struct {
	myconnection *MongoConnection
}

type UrlMapping struct {
	ShortUrl string `json:shorturl`
	LongUrl  string `json:longurl`
}

type APIResponse struct {
	StatusMessage string `json:statusmessage`
}

func NewUrlLinkShortenerAPI() *LinkShortnerAPI {
	LS := &LinkShortnerAPI{
		myconnection: NewDBConnection(),
	}
	return LS
}

func (Ls *LinkShortnerAPI) UrlRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello and welcome to the Go link shortner API \n"+
		"Do a Get request with the short Link to get the long Link \n"+
		"Do a POST request with long Link to get a short Link \n")
}

func (Ls *LinkShortnerAPI) UrlCreate(w http.ResponseWriter, r *http.Request) {
	reqBodyStruct := new(UrlMapping)
	responseEncoder := json.NewEncoder(w)
	if err := json.NewDecoder(r.Body).Decode(&reqBodyStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := responseEncoder.Encode(&APIResponse{StatusMessage: err.Error()}); err != nil {
			fmt.Fprintf(w, "Error occured while processing post request %v \n", err.Error())
		}
		return
	}
	err := Ls.myconnection.AddUrls(reqBodyStruct.LongUrl, reqBodyStruct.ShortUrl)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		if err := responseEncoder.Encode(&APIResponse{StatusMessage: err.Error()}); err != nil {
			fmt.Fprintf(w, "Error %s occured while trying to add the url \n", err.Error())
		}
		return
	}
	responseEncoder.Encode(&APIResponse{StatusMessage: "Ok"})
}

func (Ls *LinkShortnerAPI) UrlShow(w http.ResponseWriter, r *http.Request) {
	//retrieve the variable from the request
	vars := mux.Vars(r)
	sUrl := vars["shorturl"]
	if len(sUrl) > 0 {
		//find long url that corresponds to the short url
		lUrl, err := Ls.myconnection.FindlongUrl(sUrl)
		if err != nil {
			fmt.Fprintf(w, "Could not find saved long url that corresponds to the short url %s \n", sUrl)
			return
		}
		//Ensure we are dealing with an absolute path
		http.Redirect(w, r, lUrl, http.StatusFound)
	}
}
