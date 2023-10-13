#!/bin/bash
kind create cluster --name kong-poc --config clusterconfig.yaml
kubectl cluster-info --context kind-kong-poc
