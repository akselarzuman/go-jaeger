# go-jaeger
A simple implementation of [Golang OpenTelemetry](https://opentelemetry.io/docs/instrumentation/go) with [Jaeger](https://www.jaegertracing.io).

## How to Run
### Using Docker Compose
Go to the root directory of the repository and run `make start-env` or `docker compose up`
<br />
There will be two applications running on multiple ports, however, ports 8080 and 16686 are important to us.

Port <b>16686</b> is the Jaegers' UI and <b>8080</b> is a sample API.

Once you redirect to [localhost:8080/health_check](localhost:8080/health_check), you can see your record on Jaeger UI.
