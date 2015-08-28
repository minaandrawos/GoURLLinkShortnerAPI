# Go URL Link Shortener API
A URL Link shortener API code written in Go. The purpose of the API is to shorten url links via a REST HTTP request. 
Once the code is built and running, to create a short url to a long url mapping, send a POST request to http://localhost:5100/Create, the POST request should include the shortUrl and the longUrl as follows:
{'shorturl':'cosmosfading','longurl':'https://cosmosmagazine.com/space/universe-slowly-fading-away'}

You could then consume a shorturl by issuing a GET request to http://localhost:5100/shorturl

The program makes use of the gorilla mux library (http://www.gorillatoolkit.org/pkg/mux) for routing as well as the mgo library (https://labix.org/mgo) to interface with mongo database

Like this? check my website for more Go articles: www.minaandrawos.com
