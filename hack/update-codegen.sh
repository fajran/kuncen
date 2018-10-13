#!/bin/bash

$GOPATH/src/k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" \
    github.com/fajran/kuncen/pkg/client \
    github.com/fajran/kuncen/pkg/apis \
    kuncen:v1alpha1 \
    --output-base $GOPATH/src \
    --go-header-file hack/boilerplate.go.txt

