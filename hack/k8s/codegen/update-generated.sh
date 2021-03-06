#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
  "all" \
  "github.com/GLYASAI/rainbond-operator/pkg/generated" \
  "github.com/GLYASAI/rainbond-operator/pkg/apis" \
  "rainbond:v1alpha1" \
  --go-header-file "./hack/k8s/codegen/boilerplate.go.txt" \
  $@
