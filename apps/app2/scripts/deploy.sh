#!/usr/bin/env bash

set -e

COMMIT_VERSION=$(git rev-parse HEAD | cut -c1-7)

SCRIPTPATH=$(dirname $(readlink -f ${0}))
APPPATH=$(eval "cd ${SCRIPTPATH}/..; pwd")
APPNAME=$(basename $(eval "cd ${APPPATH}; pwd"))
APPREPO=${APPPATH#*go/src/}
ROOTPATH=$(eval "cd ${APPPATH}/../../; pwd")
ROOTREPO=${ROOTPATH#*go/src/}

. ${ROOTPATH}/scripts/config.sh

# Env vars:
FRAMEWORK_NAME=${APPNAME}
MIN_REPLICAS=1
MAX_REPLICAS=10
# /Env vars

echo "================================================================================"
echo "  Framework:       ${FRAMEWORK_NAME}"
echo "  Docker register: ${DOCKER_REG}"
echo "  Repository:      ${ROOTREPO}"
echo "  App repository:  ${APPREPO}"
echo "  Commit version:  ${COMMIT_VERSION}"
echo "  Min replicas:    ${MIN_REPLICAS}"
echo "  Max replicas:    ${MAX_REPLICAS}"
echo "================================================================================"
echo "Beginning deploy to cluster(s)."

echo "Building image ${FRAMEWORK_NAME}:${COMMIT_VERSION}"
docker build \
	-f deployments/Dockerfile \
	--build-arg GOROOTREPO="${ROOTREPO}" \
	--build-arg GOAPPREPO="${APPREPO}" \
	-t ${DOCKER_REG}/${FRAMEWORK_NAME}:${COMMIT_VERSION} ${ROOTPATH}

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