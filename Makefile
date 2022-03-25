start-env:
	docker compose up

start-minikube:
	minikube start --driver docker

enable-ingress:
	minikube addons enable ingress

list-namespaces:
	kubectl get namespace

create-namespace:
	kubectl create -f ./deployment/namespace.yaml

list-deployments:
	kubectl get deployments --namespace monitoring

deploy-elasticsearch:
	kubectl apply -f ./deployment/elasticsearch/deployment.yaml && kubectl apply -f ./deployment/elasticsearch/service.yaml

deploy-kibana:
	kubectl apply -f ./deployment/kibana/deployment.yaml && kubectl apply -f ./deployment/kibana/service.yaml && kubectl apply -f ./deployment/kibana/ingress.yaml

list-pods:
	kubectl get pods --namespace monitoring

list-services:
	kubectl get services --namespace monitoring

list-all:
	kubectl get all --namespace monitoring