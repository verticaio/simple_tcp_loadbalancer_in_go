# Simple TCP  LoadBalancer

## Coming Features
* Config Properties From File
* Backends health check
* Timeouts(Client,Server)
* Retries
* Different LB Algoritms(Source IP Hash, Resource Based , Least Connection, Weighted Round Robin)
* Active Passive LB
* Dynamic Backend Size
* Path,Port, Header based routing

## Example Usage:
Open sample three different port in your local laptop with using nodejs command in which using as  backend servers
```bazaar
 $ npx http-server -p 5001
 $ npx http-server -p 5002
 $ npx http-server -p 5003
```
Run your app as below(set lb endpoint as argument)
```bazaar
 $ go run main.go 8080
```
Send simple curl request load balancer endpoint one more time
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