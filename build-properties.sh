#!/usr/bin/env bash

# exit when an error occurs
set -e

echo
echo "------------------------------------------------"
echo "setting the app variables..."
echo "------------------------------------------------"
echo

export APP_NAME="echo-go"
export GO_MODULE="roberto.local/echo"

#export REPO="robertorocha" # docker hub
#export REPO="nexus-devops.tigersoftware.local:5003" # local repo
#export REPO="sonatype-nexus-service.devops.svc.cluster.local:5003" # local repo

export APP_VERSION="1.0"
export PORT=3000
export ENDPOINT="/echo"
export APP_TYPE="service"
