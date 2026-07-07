NS ?= bible

.PHONY: up down logs build watch watch-down test

dev:
	@echo "============= Spinning Everything Up ============="
	docker compose up -d

down:
	@echo "============= Taking Everything Down ============="
	docker compose down

logs:
	@echo "============= View Logs ============="
	docker compose logs -f

build:
	@echo "============= Rebuilding Docker Images ============="
	docker compose build

watch:
	@echo "============= Starting Dev with Hot Reload ============="
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up

watch-down:
	@echo "============= Taking Dev Down ============="
	docker compose -f docker-compose.yml -f docker-compose.dev.yml down

test:
	@echo "============= Running Tests ============="
	docker compose -f docker-compose.yml -f docker-compose.dev.yml exec go-api go test ./...

push:
	bin/push.sh

.PHONY: k8s-namespace
k8s-namespace:
	kubectl apply -f infra/k8s/namespace.yaml

.PHONY: k8s-secret
k8s-secret:
	kubectl create secret generic bible-env --from-env-file=.env -n $(NS) --dry-run=client -o yaml | kubectl apply -f -

.PHONY: k8s-deploy
k8s-deploy: k8s-namespace k8s-secret
	kubectl apply -f infra/k8s/mariadb.yaml
	kubectl apply -f infra/k8s/elasticsearch.yaml
	kubectl apply -f infra/k8s/deployment.yaml
	kubectl apply -f infra/k8s/ingress.yaml

.PHONY: k8s-delete
k8s-delete:
	-kubectl delete -f infra/k8s/ingress.yaml
	-kubectl delete -f infra/k8s/deployment.yaml
	-kubectl delete -f infra/k8s/elasticsearch.yaml
	-kubectl delete -f infra/k8s/mariadb.yaml
	-kubectl delete secret bible-env -n $(NS) || true

.PHONY: k8s-status
k8s-status:
	kubectl get all -n $(NS)
