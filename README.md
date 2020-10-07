# go-jaeger
A simple implementation of [Golang opentracing](https://opentracing.io/guides/golang/quick-start) with [Jaeger](https://www.jaegertracing.io).

## How to Run
### Using Docker Compose
Go to the root directory of the repository and run `docker-compose up`
<br />
There will be two applications running on multiple ports, however, ports 5000 and 16686 are important to us.

Port <b>16686</b> is the Jaegers' UI and <b>5000</b> is a sample API.

Once you redirect to `localhost:5000/publish`, you can see your record on Jaeger UI.
