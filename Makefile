# Define variables
IMAGE_NAME=my-app-image
IMAGE_TAG=latest
REGISTRY_URL=shn27# Replace with your Docker registry (Docker Hub or private registry)
DOCKERFILE_PATH=./Dockerfile
K8S_DEPLOYMENT_FILE=k8s/deployment.yaml
K8S_CLUSTER_ROLE_FILE=k8s/cluster_role.yaml
KIND_CLUSTER_NAME=kind
NAMESPACE=default

# Build the Docker image
.PHONY: build
build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f $(DOCKERFILE_PATH) .

# Tag the Docker image for the registry
.PHONY: tag
tag:
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(REGISTRY_URL)/$(IMAGE_NAME):$(IMAGE_TAG)

# Push the Docker image to the registry
.PHONY: push
push: tag
	docker push $(REGISTRY_URL)/$(IMAGE_NAME):$(IMAGE_TAG)


# Push the Docker image to kind
.PHONY: push-to-kind
push-to-kind: build
	kind load docker-image $(IMAGE_NAME):$(IMAGE_TAG) --name $(KIND_CLUSTER_NAME)


# Deploy Docker image to Kind cluster
.PHONY: deploy
deploy:push-to-kind
	# Apply the Kubernetes deployment and service YAML
	kubectl apply -f $(K8S_DEPLOYMENT_FILE) -n $(NAMESPACE)
	kubectl apply -f $(K8S_CLUSTER_ROLE_FILE)

# Clean up the environment (optional)
.PHONY: clean
clean:
	kubectl delete -f $(K8S_DEPLOYMENT_FILE) -n $(NAMESPACE)

# Help command to display available commands
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build   - Build the Docker image"
	@echo "  make push    - Tag and push the Docker image to the registry"
	@echo "  make deploy  - Build and deploy the Docker image to the Kind cluster"
	@echo "  make clean   - Remove the deployed resources from the Kubernetes cluster"
