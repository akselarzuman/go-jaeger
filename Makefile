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

deploy: deploy-elasticsearch deploy-kibana

forward-kibana:
	kubectl port-forward deployment/kibana-deployment 5601:5601

list-all:
	kubectl get all