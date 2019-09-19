package kubeadm

import (
	"bytes"
	"text/template"

	"github.com/gosoon/glog"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/version"
)

// ConfigData is supplied to the kubeadm config template, with values populated
// by the cluster package
type ConfigData struct {
	ClusterName       string
	KubernetesVersion string
	// The ControlPlaneEndpoint, that is the address of the external loadbalancer
	// if defined or the bootstrap node
	ControlPlaneEndpoint string
	// The Local API Server port
	APIBindPort int
	// The API server external listen IP (which we will port forward)
	APIServerAddress string
	// ControlPlane flag specifies the node belongs to the control plane
	ControlPlane bool
	// The main IP address of the node
	NodeAddress string
	// The Token for TLS bootstrap
	Token string
	// The subnet used for pods
	PodSubnet string
	// The subnet used for services
	ServiceSubnet string
	// IPv4 values take precedence over IPv6 by default, if true set IPv6 default values
	IPv6 bool
	// DerivedConfigData is populated by Derive()
	// These auto-generated fields are available to Config templates,
	// but not meant to be set by hand
	//DerivedConfigData
}

// DerivedConfigData fields are automatically derived by
// ConfigData.Derive if they are not specified / zero valued
//type DerivedConfigData struct {
//// DockerStableTag is automatically derived from KubernetesVersion
//DockerStableTag string
//}

// Derive automatically derives DockerStableTag if not specified
//func (c *ConfigData) Derive() {
//if c.DockerStableTag == "" {
//c.DockerStableTag = strings.Replace(c.KubernetesVersion, "+", "_", -1)
//}
//}

// See docs for these APIs at:
// https://godoc.org/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm#pkg-subdirectories
// EG:
// https://godoc.org/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1

// ConfigTemplateAlphaV2 is the kubadm config template for v1alpha2
//
// NOTE: this is the v1.11 version of this API, breaking changes occurred
// in v1.12 vs v1.11, but v1.12 also has v1alpha3 which we use instead
//
// see: https://github.com/kubernetes/kubernetes/tree/release-1.11/cmd/kubeadm/app/apis/kubeadm/v1alpha2
const ConfigTemplateAlphaV2 = `# config generated by kind
{{if .ControlPlane}}# config for the control plane node
apiVersion: kubeadm.k8s.io/v1alpha2
kind: MasterConfiguration
metadata:
  name: config
kubernetesVersion: {{.KubernetesVersion}}
clusterName: "{{.ClusterName}}"
# we use a well know token for TLS bootstrap
bootstrapTokens:
- token: "{{ .Token }}"
controlPlaneEndpoint: "{{ .ControlPlaneEndpoint }}"
# we use a well know port for making the API server discoverable inside docker network. 
# from the host machine such port will be accessible via a random local port instead.
api:
  advertiseAddress: "{{ .NodeAddress }}"
  bindPort: {{.APIBindPort}}
# we need nsswitch.conf so we use /etc/hosts
# https://github.com/kubernetes/kubernetes/issues/69195
apiServerExtraVolumes:
- name: nsswitch
  mountPath: /etc/nsswitch.conf
  hostPath: /etc/nsswitch.conf
  writeable: false
  pathType: FileOrCreate
# on docker for mac we have to expose the api server via port forward,
# so we need to ensure the cert is valid for localhost so we can talk
# to the cluster after rewriting the kubeconfig to point to localhost
apiServerCertSANs: [localhost, "{{.APIServerAddress}}"]
kubeletConfiguration:
  baseConfig:
    # configure ipv6 addresses in IPv6 mode
    {{ if .IPv6 -}}
    address: "::"
    healthzBindAddress: "::"
    {{- end }}
    # disable disk resource management by default
    # kubelet will see the host disk that the inner container runtime
    # is ultimately backed by and attempt to recover disk space.
    # we don't want that.
    imageGCHighThresholdPercent: 100
    evictionHard:
      nodefs.available: "0%"
      nodefs.inodesFree: "0%"
      imagefs.available: "0%"
controllerManagerExtraArgs:
  enable-hostpath-provisioner: "true"
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
networking:
  podSubnet: "{{ .PodSubnet }}"
{{else}}# config for this worker node
apiVersion: kubeadm.k8s.io/v1alpha2
kind: NodeConfiguration
metadata:
  name: config
clusterName: "{{.ClusterName}}"
discoveryTokenAPIServers: ["{{ .ControlPlaneEndpoint }}"]
token: "{{ .Token }}"
discoveryTokenUnsafeSkipCAVerification: true
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
{{end}}
`

// ConfigTemplateAlphaV3 is the kubadm config template for API version v1alpha3
const ConfigTemplateAlphaV3 = `# config generated by kind
apiVersion: kubeadm.k8s.io/v1alpha3
kind: ClusterConfiguration
metadata:
  name: config
kubernetesVersion: {{.KubernetesVersion}}
clusterName: "{{.ClusterName}}"
controlPlaneEndpoint: "{{ .ControlPlaneEndpoint }}"
networking:
  podSubnet: "{{ .PodSubnet }}"
  serviceSubnet: "{{ .ServiceSubnet }}"
# we need nsswitch.conf so we use /etc/hosts
# https://github.com/kubernetes/kubernetes/issues/69195
apiServerExtraVolumes:
- name: nsswitch
  mountPath: /etc/nsswitch.conf
  hostPath: /etc/nsswitch.conf
  writeable: false
  pathType: FileOrCreate
# on docker for mac we have to expose the api server via port forward,
# so we need to ensure the cert is valid for localhost so we can talk
# to the cluster after rewriting the kubeconfig to point to localhost
apiServerCertSANs: [localhost, "{{.APIServerAddress}}"]
controllerManagerExtraArgs:
  enable-hostpath-provisioner: "true"
networking:
  podSubnet: "{{ .PodSubnet }}"
---
apiVersion: kubeadm.k8s.io/v1alpha3
kind: InitConfiguration
metadata:
  name: config
# we use a well know token for TLS bootstrap
bootstrapTokens:
- token: "{{ .Token }}"
# we use a well know port for making the API server discoverable inside docker network. 
# from the host machine such port will be accessible via a random local port instead.
apiEndpoint:
  advertiseAddress: "{{ .NodeAddress }}"
  bindPort: {{.APIBindPort}}
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
---
# no-op entry that exists solely so it can be patched
apiVersion: kubeadm.k8s.io/v1alpha3
kind: JoinConfiguration
metadata:
  name: config
discoveryTokenAPIServers: ["{{ .ControlPlaneEndpoint }}"]
token: "{{ .Token }}"
discoveryTokenUnsafeSkipCAVerification: true
controlPlane: {{ .ControlPlane }}
{{ if .ControlPlane -}}
apiEndpoint:
  advertiseAddress: "{{ .NodeAddress }}"
  bindPort: {{.APIBindPort}}
{{- end }}
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
metadata:
  name: config
# configure ipv6 addresses in IPv6 mode
{{ if .IPv6 -}}
address: "::"
healthzBindAddress: "::"
{{- end }}
# disable disk resource management by default
# kubelet will see the host disk that the inner container runtime
# is ultimately backed by and attempt to recover disk space. we don't want that.
imageGCHighThresholdPercent: 100
evictionHard:
  nodefs.available: "0%"
  nodefs.inodesFree: "0%"
  imagefs.available: "0%"
---
# no-op entry that exists solely so it can be patched
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
metadata:
  name: config
`

// ConfigTemplateBetaV1 is the kubadm config template for API version v1beta1
const ConfigTemplateBetaV1 = `# config generated by kind
apiVersion: kubeadm.k8s.io/v1beta1
kind: ClusterConfiguration
metadata:
  name: config
kubernetesVersion: {{.KubernetesVersion}}
clusterName: "{{.ClusterName}}"
controlPlaneEndpoint: "{{ .ControlPlaneEndpoint }}"
# on docker for mac we have to expose the api server via port forward,
# so we need to ensure the cert is valid for localhost so we can talk
# to the cluster after rewriting the kubeconfig to point to localhost
apiServer:
  certSANs: [localhost, "{{.APIServerAddress}}"]
controllerManager:
  extraArgs:
    enable-hostpath-provisioner: "true"
    # configure ipv6 default addresses for IPv6 clusters
    {{ if .IPv6 -}}
    bind-address: "::"
    {{- end }}
scheduler:
  extraArgs:
    # configure ipv6 default addresses for IPv6 clusters
    {{ if .IPv6 -}}
    address: "::"
    bind-address: "::1"
    {{- end }}
networking:
  podSubnet: "{{ .PodSubnet }}"
  serviceSubnet: "{{ .ServiceSubnet }}"
---
apiVersion: kubeadm.k8s.io/v1beta1
kind: InitConfiguration
metadata:
  name: config
# we use a well know token for TLS bootstrap
bootstrapTokens:
- token: "{{ .Token }}"
# we use a well know port for making the API server discoverable inside docker network. 
# from the host machine such port will be accessible via a random local port instead.
localAPIEndpoint:
  advertiseAddress: "{{ .NodeAddress }}"
  bindPort: {{.APIBindPort}}
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
---
# no-op entry that exists solely so it can be patched
apiVersion: kubeadm.k8s.io/v1beta1
kind: JoinConfiguration
metadata:
  name: config
{{ if .ControlPlane -}}	
controlPlane:
  localAPIEndpoint:
    advertiseAddress: "{{ .NodeAddress }}"
    bindPort: {{.APIBindPort}}
{{- end }}
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
discovery:
  bootstrapToken:
    apiServerEndpoint: "{{ .ControlPlaneEndpoint }}"
    token: "{{ .Token }}"
    unsafeSkipCAVerification: true
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
metadata:
  name: config
# configure ipv6 addresses in IPv6 mode
{{ if .IPv6 -}}
address: "::"
healthzBindAddress: "::"
{{- end }}
# disable disk resource management by default
# kubelet will see the host disk that the inner container runtime
# is ultimately backed by and attempt to recover disk space. we don't want that.
imageGCHighThresholdPercent: 100
evictionHard:
  nodefs.available: "0%"
  nodefs.inodesFree: "0%"
  imagefs.available: "0%"
---
# no-op entry that exists solely so it can be patched
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
metadata:
  name: config
`

// ConfigTemplateBetaV2 is the kubadm config template for API version v1beta2
const ConfigTemplateBetaV2 = `# config generated by kind
apiVersion: kubeadm.k8s.io/v1beta2
kind: ClusterConfiguration
metadata:
  name: config
kubernetesVersion: {{.KubernetesVersion}}
clusterName: "{{.ClusterName}}"
controlPlaneEndpoint: "{{ .ControlPlaneEndpoint }}"
# on docker for mac we have to expose the api server via port forward,
# so we need to ensure the cert is valid for localhost so we can talk
# to the cluster after rewriting the kubeconfig to point to localhost
apiServer:
  certSANs: [localhost, "{{.APIServerAddress}}"]
controllerManager:
  extraArgs:
    enable-hostpath-provisioner: "true"
    # configure ipv6 default addresses for IPv6 clusters
    {{ if .IPv6 -}}
    bind-address: "::"
    {{- end }}
scheduler:
  extraArgs:
    # configure ipv6 default addresses for IPv6 clusters
    {{ if .IPv6 -}}
    address: "::"
    bind-address: "::1"
    {{- end }}
networking:
  podSubnet: "{{ .PodSubnet }}"
  serviceSubnet: "{{ .ServiceSubnet }}"
---
apiVersion: kubeadm.k8s.io/v1beta2
kind: InitConfiguration
metadata:
  name: config
# we use a well know token for TLS bootstrap
bootstrapTokens:
- token: "{{ .Token }}"
# we use a well know port for making the API server discoverable inside docker network. 
# from the host machine such port will be accessible via a random local port instead.
localAPIEndpoint:
  advertiseAddress: "{{ .NodeAddress }}"
  bindPort: {{.APIBindPort}}
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
---
# no-op entry that exists solely so it can be patched
apiVersion: kubeadm.k8s.io/v1beta2
kind: JoinConfiguration
metadata:
  name: config
{{ if .ControlPlane -}}
controlPlane:
  localAPIEndpoint:
    advertiseAddress: "{{ .NodeAddress }}"
    bindPort: {{.APIBindPort}}
{{- end }}
nodeRegistration:
  criSocket: "/run/containerd/containerd.sock"
  kubeletExtraArgs:
    fail-swap-on: "false"
    node-ip: "{{ .NodeAddress }}"
discovery:
  bootstrapToken:
    apiServerEndpoint: "{{ .ControlPlaneEndpoint }}"
    token: "{{ .Token }}"
    unsafeSkipCAVerification: true
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
metadata:
  name: config
# configure ipv6 addresses in IPv6 mode
{{ if .IPv6 -}}
address: "::"
healthzBindAddress: "::"
{{- end }}
# disable disk resource management by default
# kubelet will see the host disk that the inner container runtime
# is ultimately backed by and attempt to recover disk space. we don't want that.
imageGCHighThresholdPercent: 100
evictionHard:
  nodefs.available: "0%"
  nodefs.inodesFree: "0%"
  imagefs.available: "0%"
---
# no-op entry that exists solely so it can be patched
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
metadata:
  name: config
`

// Config returns a kubeadm config generated from config data, in particular
// the kubernetes version
func Config(data ConfigData) (config string, err error) {
	glog.Infof("Configuration Input data: %v", data)
	ver, err := version.ParseGeneric(data.KubernetesVersion)
	if err != nil {
		return "", err
	}

	// assume the latest API version, then fallback if the k8s version is too low
	templateSource := ConfigTemplateBetaV2
	if ver.LessThan(version.MustParseSemantic("v1.12.0")) {
		templateSource = ConfigTemplateAlphaV2
	} else if ver.LessThan(version.MustParseSemantic("v1.13.0")) {
		templateSource = ConfigTemplateAlphaV3
	} else if ver.LessThan(version.MustParseSemantic("v1.15.0")) {
		templateSource = ConfigTemplateBetaV1
	}

	t, err := template.New("kubeadm-config").Parse(templateSource)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse config template")
	}

	// derive any automatic fields if not supplied
	//data.Derive()

	// execute the template
	var buff bytes.Buffer
	err = t.Execute(&buff, data)
	if err != nil {
		return "", errors.Wrap(err, "error executing config template")
	}
	glog.Infof("Configuration generated:\n %v", buff.String())
	return buff.String(), nil
}
