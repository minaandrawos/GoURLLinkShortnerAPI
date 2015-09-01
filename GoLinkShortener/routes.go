package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

/*
	Create the routes for the API. The API supports three URLs:
		1- GET "/" => Shows a description for the API
		2- GET "/{shorturl}" => If the shortUrl exists in the backend database, redirect to the long url that corresponds to func init() {
		3- Post "/Create" => Takes a post request with http body of {
																	shorturl: "short Link"
																	longurl:  "original long link"
																	}
		 Causes the API to create a mapping between the short url and the long url in the backend database
*/

func CreateRoutes(LS *LinkShortnerAPI) Routes {
	return Routes{
		Route{
			"UrlRoot",
			"GET",
			"/",
			LS.UrlRoot,
		},
		Route{
			"UrlShow",
			"GET",
			"/{shorturl}",
			LS.UrlShow,
		},
		Route{
			"UrlCreate",
			"POST",
			"/Create",
			LS.UrlCreate,
		},
	}
}
