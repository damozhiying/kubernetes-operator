#!/bin/bash
# 指定 apiserver 地址
KUBE_APISERVER="https://10.0.4.15:6443"

[ -e ../deploy/config.sh ] && . ../deploy/config.sh || exit

# 生成 Bootstrap Token
BOOTSTRAP_TOKEN_ID=$(head -c 6 /dev/urandom | md5sum | head -c 6)
BOOTSTRAP_TOKEN_SECRET=$(head -c 16 /dev/urandom | md5sum | head -c 16)
BOOTSTRAP_TOKEN="${BOOTSTRAP_TOKEN_ID}.${BOOTSTRAP_TOKEN_SECRET}"
CERTS_DIR="/etc/kubernetes/ssl"

[ -d output ] || mkdir output

# 生成 kubectl 配置文件
echo "Create kubectl kubeconfig..."
kubectl config set-cluster kubernetes \
  --certificate-authority=${CERTS_DIR}/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=output/kubectl.kubeconfig
kubectl config set-credentials "system:masters" \
  --client-certificate=${CERTS_DIR}/admin.pem \
  --client-key=${CERTS_DIR}/admin-key.pem \
  --embed-certs=true \
  --kubeconfig=output/kubectl.kubeconfig
kubectl config set-context default \
  --cluster=kubernetes \
  --user=system:masters \
  --kubeconfig=output/kubectl.kubeconfig
kubectl config use-context default --kubeconfig=output/kubectl.kubeconfig


# 生成 kube-controller-manager 配置文件
echo "Create kube-controller-manager kubeconfig..."
kubectl config set-cluster kubernetes \
  --certificate-authority=${CERTS_DIR}/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=output/kube-controller-manager.kubeconfig
kubectl config set-credentials "system:kube-controller-manager" \
  --client-certificate=${CERTS_DIR}/kube-controller-manager.pem \
  --client-key=${CERTS_DIR}/kube-controller-manager-key.pem \
  --embed-certs=true \
  --kubeconfig=output/kube-controller-manager.kubeconfig
kubectl config set-context default \
  --cluster=kubernetes \
  --user=system:kube-controller-manager \
  --kubeconfig=output/kube-controller-manager.kubeconfig
kubectl config use-context default --kubeconfig=output/kube-controller-manager.kubeconfig

# 生成 kube-scheduler 配置文件
echo "Create kube-scheduler kubeconfig..."
kubectl config set-cluster kubernetes \
  --certificate-authority=${CERTS_DIR}/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=output/kube-scheduler.kubeconfig
kubectl config set-credentials "system:kube-scheduler" \
  --client-certificate=${CERTS_DIR}/kube-scheduler.pem \
  --client-key=${CERTS_DIR}/kube-scheduler-key.pem \
  --embed-certs=true \
  --kubeconfig=output/kube-scheduler.kubeconfig
kubectl config set-context default \
  --cluster=kubernetes \
  --user=system:kube-scheduler \
  --kubeconfig=output/kube-scheduler.kubeconfig
kubectl config use-context default --kubeconfig=output/kube-scheduler.kubeconfig


# 生成 kubelet tls bootstrap 配置
echo "Create kubelet bootstrapping kubeconfig..."
kubectl config set-cluster kubernetes \
  --certificate-authority=${CERTS_DIR}/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=output/bootstrap.kubeconfig
kubectl config set-credentials "system:bootstrap:${BOOTSTRAP_TOKEN_ID}" \
  --token=${BOOTSTRAP_TOKEN} \
  --kubeconfig=output/bootstrap.kubeconfig
kubectl config set-context default \
  --cluster=kubernetes \
  --user="system:bootstrap:${BOOTSTRAP_TOKEN_ID}" \
  --kubeconfig=output/bootstrap.kubeconfig
kubectl config use-context default --kubeconfig=output/bootstrap.kubeconfig

# 生成 kubelet 配置文件
echo "Create kubelet kubeconfig..."
kubectl config set-cluster kubernetes \
  --certificate-authority=${CERTS_DIR}/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=output/kubelet.kubeconfig
kubectl config set-credentials "system:masters" \
  --client-certificate=${CERTS_DIR}/kubelet-client.pem \
  --client-key=${CERTS_DIR}/kubelet-client-key.pem \
  --embed-certs=true \
  --kubeconfig=output/kubelet.kubeconfig
kubectl config set-context default \
  --cluster=kubernetes \
  --user=system:masters \
  --kubeconfig=output/kubelet.kubeconfig
kubectl config use-context default --kubeconfig=output/kubelet.kubeconfig

# 生成 kube-proxy 配置文件
echo "Create kube-proxy kubeconfig..."
kubectl config set-cluster kubernetes \
  --certificate-authority=${CERTS_DIR}/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=output/kube-proxy.kubeconfig
kubectl config set-credentials "system:kube-proxy" \
  --client-certificate=${CERTS_DIR}/kube-proxy.pem \
  --client-key=${CERTS_DIR}/kube-proxy-key.pem \
  --embed-certs=true \
  --kubeconfig=output/kube-proxy.kubeconfig
kubectl config set-context default \
  --cluster=kubernetes \
  --user=system:kube-proxy \
  --kubeconfig=output/kube-proxy.kubeconfig
kubectl config use-context default --kubeconfig=output/kube-proxy.kubeconfig
