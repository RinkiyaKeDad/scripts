#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

declare -a repos=("test-infra" "autoscaler" "ingress-gce" "node-problem-detector" "kubeadm" "cloud-provider-vsphere" "dns" "utils" "kube-state-metrics" "perf-tests" "client-go" "klog" "kubectl" "cli-runtime" "ingress-nginx" "cloud-provider-alibaba-cloud" "legacy-cloud-providers" "cluster-bootstrap" "kube-controller-manager" "controller-manager" "kube-proxy" "sample-apiserver" "kube-aggregator" "apiserver" "component-base" "api" "code-generator" "git-sync" "system-validators" "frakti" "minikube" "dashboard" "kops" "kube-openapi" "enhancements" "cloud-provider-openstack" "release" "cloud-provider-aws" "sample-cli-plugin" "csi-translation-lib" "cloud-provider" "kube-scheduler" "kubelet" "metrics" "apiextensions-apiserver" "sample-controller" "component-helpers" "apimachinery" "mount-utils" "cri-api" "kompose" "publishing-bot" "gengo" "node-api" "csi-api")
declare  pass=0
rm -rf "/Users/arsh/depstat-temp"

for repo in "${repos[@]}"
do
echo ""
echo "Running for ${repo}"
echo ""
repoUrl="https://github.com/kubernetes/${repo}.git"

localFolder="/Users/arsh/depstat-temp"

echo "${repoUrl}"
git clone "$repoUrl" "$localFolder"

cd "$localFolder"
if 
    timeout 180s depstat stats;
then
    echo 'OK'; #if you want a positive response
    pass=$((pass+1))
else
    echo 'Not OK';
fi
cd ..

rm -rf "$localFolder"
done

echo "Repos passed: ${pass}"

