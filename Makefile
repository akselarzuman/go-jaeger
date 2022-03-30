start-env:
	docker compose up

start-minikube:
	minikube start --driver=docker

stop-minikube:
	minikube stop

enable-ingress:
	minikube addons enable ingress

list-namespaces:
	kubectl get namespace

create-namespace:
	kubectl create -f ./deployment/namespace.yaml

deploy-elasticsearch:
	kubectl apply -f ./deployment/elasticsearch/deployment.yaml && kubectl apply -f ./deployment/elasticsearch/service.yaml

deploy-kibana:
	kubectl apply -f ./deployment/kibana/deployment.yaml && kubectl apply -f ./deployment/kibana/service.yaml

deploy-jaeger-collector:
	kubectl apply -f ./deployment/jaeger-collector/deployment.yaml && kubectl apply -f ./deployment/jaeger-collector/service.yaml

deploy-jaeger-query:
	kubectl apply -f ./deployment/jaeger-query/deployment.yaml && kubectl apply -f ./deployment/jaeger-query/service.yaml

deploy-jaeger: deploy-jaeger-collector deploy-jaeger-query

deploy: deploy-elasticsearch deploy-kibana deploy-jaeger

forward-kibana:
	kubectl port-forward deployment/kibana-deployment 5601:5601

forward-collector:
	kubectl port-forward deployment/jaeger-collector-deployment 14268:14268

forward-query:
	kubectl port-forward deployment/jaeger-query-deployment 16686:16686

list-all:
	kubectl get all