# Simple TCP  LoadBalancer
## Example Usage:
Open sample three different port in your local laptop with using nodejs command
```bazaar
 $ npx http-server -p 5001
 $ npx http-server -p 5002
 $ npx http-server -p 5003
```
Run your app as below
```bazaar
 $ go run main.go
    counter=1, backedn=localhost:5001
    counter=2, backedn=localhost:5002
    counter=3, backedn=localhost:5003
    counter=4, backedn=localhost:5001
    counter=5, backedn=localhost:5002
    counter=6, backedn=localhost:5003
    counter=7, backedn=localhost:5001
```
Send simple curl request load balancer endpoint
```bazaar
$ curl localhost:8080
```
App logs ...
```bazaar
counter=1, backedn=localhost:5001
counter=2, backedn=localhost:5002
counter=3, backedn=localhost:5003
counter=4, backedn=localhost:5001
counter=5, backedn=localhost:5002
counter=6, backedn=localhost:5003
counter=7, backedn=localhost:5001
```


## Coming Features
* Config Properties From File
* Backend health checks
* Timeouts
* Retries
* Different LB Algoritms(Source IP Hash, Resource Based , Least Connection, Weighted Round Robin , St)
* Active Passive LB