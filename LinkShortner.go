package main
import "net/http"

/*
	This is the entry point for the API, the purpose of the API is to shorten url links via a REST HTTP request
	The request should include the shortUrl and the longUrl as follows {'shorturl':'cosmosfading','longurl':'https://cosmosmagazine.com/space/universe-slowly-fading-away'}
    The program makes use of the gorilla mux library for routing
 */
func main(){
	//Create a new API shortner API
	LinkShortner := NewUrlLinkShortnerAPI()
	//Create the needed routes for the API
	routes := CreateRoutes(LinkShortner)
	//Initiate the API routers
	router := NewLinkShortnerRouter(routes)
	//This will start the web server on local port 5100
	http.ListenAndServe(":5100", router)
}
