import requests
import json
data = {'shorturl':'cosmosfading','longurl':'https://cosmosmagazine.com/space/universe-slowly-fading-away'}
r = requests.post('http://localhost:5100/Create',data=json.dumps(data))
print "Status code received: " + str(r.status_code) + " response body: " + r.text
