#!/usr/bin/env bash

set -e

MIN_REPLICAS=1
MAX_REPLICAS=10

COMMIT_VERSION=$(git rev-parse HEAD | cut -c1-7)
DOCKER_REG="hscasn"

# Env vars:
FRAMEWORK_NAME="go-microservice"
# /Env vars

echo "Beginning deploy to cluster(s)."

echo "Building image ${FRAMEWORK_NAME}:${COMMIT_VERSION}"
docker build -f deployments/Dockerfile \
  -t ${DOCKER_REG}/${FRAMEWORK_NAME}:${COMMIT_VERSION} .
rm -rf node_modules

echo "Pushing image ${FRAMEWORK_NAME}:${COMMIT_VERSION} to Docker Hub"
docker push ${DOCKER_REG}/${FRAMEWORK_NAME}:${COMMIT_VERSION}

# Build k8s objects YAML file from template
echo "Building Kubernetes objects YAML file for development from template"
cat ./deployments/deployment.yml \
        | sed s/{{dockerReg}}/${DOCKER_REG}/g \
        | sed s/{{imageVersion}}/${COMMIT_VERSION}/g \
        | sed s/{{frameworkName}}/${FRAMEWORK_NAME}/g \
        | sed s/{{minReplicas}}/${MIN_REPLICAS}/g \
        | sed s/{{maxReplicas}}/${MAX_REPLICAS}/g \
        > ./deployments/_k8s_objects.yml


echo "Deploying to Kubernetes cluster(s)"
kubectl apply -f ./deployments/_k8s_objects.yml