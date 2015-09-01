package main

import "net/http"

/*
	This is the entry point for the API, the purpose of the API is to shorten url links via a REST HTTP request
	Once the code is built and running, to create a short url to a long url mapping, send a POST request to http://localhost:5100/Create, the POST request should include the shortUrl and the longUrl as follows:
	{'shorturl':'cosmosfading','longurl':'https://cosmosmagazine.com/space/universe-slowly-fading-away'}

	You could then consume a shorturl by issuing a GET request to http://localhost:5100/<the short url>
	The program makes use of the gorilla mux library for routing as well as the mgo library to interface with mongo database
*/
func main() {
	//Create a new API shortner API
	LinkShortener := NewUrlLinkShortenerAPI()
	//Create the needed routes for the API
	routes := CreateRoutes(LinkShortener)
	//Initiate the API routers
	router := NewLinkShortenerRouter(routes)
	//This will start the web server on local port 5100
	http.ListenAndServe(":5100", router)
}
