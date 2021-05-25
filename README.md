# Simple TCP  LoadBalancer

## Coming Features
* Config Properties From File
* Backends health check
* Timeouts(Client,Server)
* Retries
* Different LB Algoritms(Source IP Hash, Resource Based , Least Connection, Weighted Round Robin)
* Dynamic Backend Count from Service Discovery
* Path,Port, Header based routing
* Forward Client IPs to backend endpoint  
* Monitoring metrics

## Example Usage:
Prepare three backend services with  different port in your local laptop  using nodejs package
```bazaar
 $ npx http-server -p 5001
 $ npx http-server -p 5002
 $ npx http-server -p 5003
```
Run your app as below(set lb port as argument)
```bazaar
 $ go run main.go 8080
```
Send simple curl request load balancer endpoint one more time
```bazaar
$ curl localhost:8080
```
App logs ...
```bazaar
request counter=1, backedn=localhost:5001
request counter=2, backedn=localhost:5002
request counter=3, backedn=localhost:5003
request counter=4, backedn=localhost:5001
request counter=5, backedn=localhost:5002
request counter=6, backedn=localhost:5003
request counter=7, backedn=localhost:5001
```