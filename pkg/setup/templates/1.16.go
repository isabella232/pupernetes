package templates

var (
	manifest1o16 = []Manifest{
		{
			Name:        "kubelet.service",
			Destination: ManifestSystemdUnit,
			Content: []byte(`[Unit]
Description=Hyperkube kubelet for pupernetes
After=network.target

[Service]
ExecStart={{.RootABSPath}}/bin/hyperkube kubelet \
  --v=4 \
  --hairpin-mode=none \
  --config={{.RootABSPath}}/manifest-config/kubelet-config.yaml \
	--pod-manifest-path={{.RootABSPath}}/manifest-static-pod \
	--hostname-override={{ .Hostname }} \
	--root-dir=/var/lib/p8s-kubelet \
	--healthz-port=10248 \
	--kubeconfig={{.RootABSPath}}/manifest-config/kubeconfig-insecure.yaml \
	--resolv-conf={{.RootABSPath}}/net.d/resolv-conf \
	--cluster-dns={{ .DNSClusterIP }} \
	--cluster-domain=cluster.local \
	--cert-dir={{.RootABSPath}}/secrets \
	--client-ca-file={{.RootABSPath}}/secrets/kubernetes.issuing_ca \
	--tls-cert-file={{.RootABSPath}}/secrets/kubernetes.certificate \
	--tls-private-key-file={{.RootABSPath}}/secrets/kubernetes.private_key \
	--read-only-port=0 \
	--anonymous-auth=false \
	--authentication-token-webhook \
	--authentication-token-webhook-cache-ttl=5s \
	--authorization-mode=Webhook  \
	--cgroups-per-qos=true \
	--cgroup-driver={{ .CgroupDriver }} \
	--max-pods=60 \
	--node-ip={{ .NodeIP }} \
	--node-labels=p8s=mononode \
	--application-metrics-count-limit=50 \
	--network-plugin=cni \
	--cni-conf-dir={{.RootABSPath}}/net.d \
	--cni-bin-dir={{.RootABSPath}}/bin \
	--container-runtime={{.ContainerRuntime}} \
	--runtime-request-timeout=15m \
	--container-runtime-endpoint=unix://{{.ContainerRuntimeEndpoint}}

Restart=no
`),
		},
		{
			Name:        "containerd.service",
			Destination: ManifestSystemdUnit,
			Content: []byte(`[Unit]
Description=containerd
After=network.target

[Service]
KillMode=process
Environment=PATH=/bin:/sbin:/usr/bin:/usr/sbin/:/usr/local/bin:/usr/local/sbin:{{.RootABSPath}}/bin
ExecStart={{.RootABSPath}}/bin/containerd \
	--config {{.RootABSPath}}/manifest-config/containerd-config.toml

Restart=no
`),
		},
		{
			Name:        "kube-apiserver.service",
			Destination: ManifestSystemdUnit,
			Content: []byte(`[Unit]
Description=Hyperkube apiserver for pupernetes
After=network.target

[Service]
ExecStart={{.RootABSPath}}/bin/hyperkube kube-apiserver \
	--apiserver-count=1 \
	--insecure-bind-address=127.0.0.1 \
	--insecure-port=8080 \
	--allow-privileged=true \
	--service-cluster-ip-range={{ .ServiceClusterIPRange }} \
	--admission-control=NamespaceLifecycle,PodPreset,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota,EventRateLimit \
	--kubelet-preferred-address-types=InternalIP,LegacyHostIP,ExternalDNS,InternalDNS,Hostname \
	--authorization-mode=RBAC \
	--etcd-servers=http://127.0.0.1:2379 \
	--anonymous-auth=false \
	--service-account-lookup=true \
	--runtime-config=api/all=true \
	--client-ca-file={{.RootABSPath}}/secrets/kubernetes.issuing_ca \
	--tls-cert-file={{.RootABSPath}}/secrets/kubernetes.certificate \
	--tls-private-key-file={{.RootABSPath}}/secrets/kubernetes.private_key \
	--service-account-key-file={{.RootABSPath}}/secrets/service-accounts.rsa \
	--kubelet-client-certificate={{.RootABSPath}}/secrets/kubernetes.certificate \
	--kubelet-client-key={{.RootABSPath}}/secrets/kubernetes.private_key \
	--kubelet-https \
	--requestheader-client-ca-file={{.RootABSPath}}/secrets/kubernetes.issuing_ca \
	--requestheader-allowed-names=aggregator,p8s \
	--requestheader-extra-headers-prefix=X-Remote-Extra- \
	--requestheader-group-headers=X-Remote-Group \
	--requestheader-username-headers=X-Remote-User \
	--proxy-client-cert-file={{.RootABSPath}}/secrets/kubernetes.certificate \
	--proxy-client-key-file={{.RootABSPath}}/secrets/kubernetes.private_key \
	--kubelet-https \
	--kubelet-certificate-authority={{.RootABSPath}}/secrets/kubernetes.issuing_ca \
	--target-ram-mb=0 \
	--watch-cache=false \
	--default-watch-cache-size=0 \
	--watch-cache-sizes="" \
	--deserialization-cache-size=0 \
	--audit-log-path={{.RootABSPath}}/logs/audit.log \
	--audit-policy-file={{.RootABSPath}}/manifest-config/audit.yaml \
	--etcd-compaction-interval=0 \
	--event-ttl=10m \
	--admission-control-config-file={{.RootABSPath}}/manifest-config/admission.yaml

Restart=no
`),
		},
		{
			Name:        "etcd.service",
			Destination: ManifestSystemdUnit,
			Content: []byte(`[Unit]
Description=etcd for pupernetes
After=network.target

[Service]
ExecStart={{.RootABSPath}}/bin/etcd \
	--name=etcdv3 \
	--data-dir={{.RootABSPath}}/etcd-data \
	--auto-compaction-retention=0 \
	--quota-backend-bytes=0 \
	--metrics=basic \
	--cert-file={{.RootABSPath}}/secrets/etcd.certificate \
	--key-file={{.RootABSPath}}/secrets/etcd.private_key \
	--client-cert-auth=true \
	--trusted-ca-file={{.RootABSPath}}/secrets/etcd.issuing_ca \
	--listen-client-urls=http://127.0.0.1:2379,https://{{ .NodeIP }}:2379 \
	--advertise-client-urls=http://127.0.0.1:2379,https://{{ .NodeIP }}:2379

Restart=no
`),
		},
		{
			Name:        "kubeconfig-auth.yaml",
			Destination: ManifestConfig,
			Content: []byte(`---
apiVersion: v1
kind: Config
clusters:
  - cluster:
      server: https://127.0.0.1:6443
      certificate-authority: "{{.RootABSPath}}/secrets/kubernetes.issuing_ca"
    name: p8s
contexts:
  - context:
      cluster: p8s
      user: p8s
    name: p8s
current-context: p8s
users:
  - name: p8s
    username: p8s
    client-certificate: "{{.RootABSPath}}/secrets/kubernetes.certificate"
    client-key: "{{.RootABSPath}}/secrets/kubernetes.private_key"
`),
		},
		{
			Name:        "kubelet-config.yaml",
			Destination: ManifestConfig,
			Content: []byte(`---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
failSwapOn: false
`),
		},
		{
			Name:        "containerd-config.toml",
			Destination: ManifestConfig,
			Content: []byte(`
root = "/var/lib/containerd"
state = "/run/containerd"
oom_score = 0

[grpc]
  address = "{{.ContainerRuntimeEndpoint}}"
  uid = 0
  gid = 0
  max_recv_message_size = 16777216
  max_send_message_size = 16777216

[debug]
  address = ""
  uid = 0
  gid = 0
  level = ""

[metrics]
  address = "127.0.0.1:1338"
  grpc_histogram = false

[cgroup]
  path = ""

[plugins]
  [plugins.cgroups]
    no_prometheus = false
  [plugins.cri]
    stream_server_address = ""
    stream_server_port = "10010"
    enable_selinux = false
    sandbox_image = "k8s.gcr.io/pause:3.1"
    stats_collect_period = 10
    systemd_cgroup = false
    enable_tls_streaming = false
    [plugins.cri.containerd]
      snapshotter = "overlayfs"
      [plugins.cri.containerd.default_runtime]
        runtime_type = "io.containerd.runtime.v1.linux"
        runtime_engine = ""
        runtime_root = ""
      [plugins.cri.containerd.untrusted_workload_runtime]
        runtime_type = ""
        runtime_engine = ""
        runtime_root = ""
    [plugins.cri.cni]
      bin_dir = "{{.RootABSPath}}/bin"
      conf_dir = "{{.RootABSPath}}/net.d"
      conf_template = ""
    [plugins.cri.registry]
      [plugins.cri.registry.mirrors]
        [plugins.cri.registry.mirrors."docker.io"]
          endpoint = ["https://registry-1.docker.io"]
  [plugins.diff-service]
    default = ["walking"]
  [plugins.linux]
    shim = "containerd-shim"
    runtime = "runc"
    runtime_root = ""
    no_shim = false
    shim_debug = false
  [plugins.scheduler]
    pause_threshold = 0.02
    deletion_threshold = 0
    mutation_threshold = 100
    schedule_delay = "0s"
    startup_delay = "100ms"
`),
		},
		{
			Name:        "kubeconfig-insecure.yaml",
			Destination: ManifestConfig,
			Content: []byte(`---
apiVersion: v1
kind: Config
clusters:
  - cluster:
      server: http://127.0.0.1:8080
    name: p8s
contexts:
  - context:
      cluster: p8s
      user: p8s
    name: p8s
current-context: p8s
users:
  - name: p8s
    username: p8s
`),
		},
		{
			Name:        "audit.yaml",
			Destination: ManifestConfig,
			Content: []byte(`---
apiVersion: audit.k8s.io/v1
kind: Policy
rules:
  - level: Request
    verbs:
      - create
    omitStages:
      - RequestReceived
    resources:
    - group: ""
      resources:
        - events
  - level: Metadata
    omitStages:
      - RequestReceived
`),
		},
		{
			Name:        "admission.yaml",
			Destination: ManifestConfig,
			Content: []byte(`---
kind: AdmissionConfiguration
apiVersion: apiserver.k8s.io/v1alpha1
plugins:
- name: EventRateLimit
  path: eventconfig.yaml
`),
		},
		{
			Name:        "eventconfig.yaml",
			Destination: ManifestConfig,
			Content: []byte(`---
kind: Configuration
apiVersion: eventratelimit.admission.k8s.io/v1alpha1
limits:
- type: Namespace
  qps: 50
  burst: 100
  cacheSize: 2000
- type: User
  qps: 10
  burst: 50
`),
		},
		{
			Name:        "kube-controller-manager.yaml",
			Destination: ManifestAPI,
			Content: []byte(`---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-controller-manager
  namespace: kube-system
automountServiceAccountToken: false
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: kube-controller-manager
  name: kube-controller-manager
  namespace: kube-system
spec:
  serviceAccountName: kube-controller-manager
  automountServiceAccountToken: false
  nodeName: "{{ .Hostname }}"
  hostNetwork: true
  volumes:
  - name: secrets
    hostPath:
      path: "{{.RootABSPath}}/secrets"
  containers:
  - name: kube-controller-manager
    image: "{{ .HyperkubeImageURL }}"
    imagePullPolicy: IfNotPresent
    command:
    - /hyperkube
    - kube-controller-manager
    - --master=http://127.0.0.1:8080
    - --leader-elect=true
    - --leader-elect-lease-duration=150s
    - --leader-elect-renew-deadline=100s
    - --leader-elect-retry-period=20s
    - --cluster-signing-cert-file=/etc/secrets/pupernetes.certificate
    - --cluster-signing-key-file=/etc/secrets/pupernetes.private_key
    - --root-ca-file=/etc/secrets/pupernetes.issuing_ca
    - --service-account-private-key-file=/etc/secrets/service-accounts.rsa
    - --concurrent-deployment-syncs=2
    - --concurrent-endpoint-syncs=2
    - --concurrent-gc-syncs=5
    - --concurrent-namespace-syncs=3
    - --concurrent-replicaset-syncs=2
    - --concurrent-resource-quota-syncs=2
    - --concurrent-service-syncs=1
    - --concurrent-serviceaccount-token-syncs=2
    volumeMounts:
      - name: secrets
        mountPath: /etc/secrets
    livenessProbe:
      httpGet:
        path: /healthz
        port: 10252
      initialDelaySeconds: 15
    readinessProbe:
      httpGet:
        path: /healthz
        port: 10252
      initialDelaySeconds: 5
    resources:
      requests:
        cpu: "100m"
      limits:
        cpu: "250m"
`),
		},
		{
			Name:        "kube-scheduler.yaml",
			Destination: ManifestAPI,
			Content: []byte(`---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-scheduler
  namespace: kube-system
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: system:kube-scheduler
subjects:
  - kind: ServiceAccount
    name: kube-scheduler
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: system:kube-scheduler
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: kube-scheduler
  name: kube-scheduler
  namespace: kube-system
spec:
  serviceAccountName: kube-scheduler
  nodeName: "{{ .Hostname }}"
  hostNetwork: true
  volumes:
  - name: secrets
    hostPath:
      path: "{{.RootABSPath}}/secrets"
  containers:
      - name: kube-scheduler
        image: "{{ .HyperkubeImageURL }}"
        imagePullPolicy: IfNotPresent
        command:
        - /hyperkube
        - kube-scheduler
        - --master=http://127.0.0.1:8080
        - --leader-elect=true
        livenessProbe:
          httpGet:
            path: /healthz
            port: 10251
          initialDelaySeconds: 15
        readinessProbe:
          httpGet:
            path: /healthz
            port: 10251
          initialDelaySeconds: 5
        resources:
          requests:
            cpu: "100m"
          limits:
            cpu: "200m"
`),
		},
		{
			Name:        "kube-proxy.yaml",
			Destination: ManifestAPI,
			Content: []byte(`---
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-proxy
  namespace: kube-system
data:
  config.yaml: |
    apiVersion: kubeproxy.config.k8s.io/v1alpha1
    kind: KubeProxyConfiguration
    bindAddress: 0.0.0.0
    clientConnection:
      kubeconfig: /var/lib/kubernetes/kubeconfig.yaml
    clusterCIDR: "{{ .ServiceClusterIPRange }}"
    healthzBindAddress: 0.0.0.0:10256
    hostnameOverride: "{{ .Hostname }}"
    iptables:
      masqueradeAll: true
    metricsBindAddress: 127.0.0.1:10249
    mode: iptables

  kubeconfig.yaml: |
    apiVersion: v1
    kind: Config
    clusters:
      - name: kube
        cluster:
          server: https://127.0.0.1:6443
          certificate-authority: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    users:
      - name: service-account
        user:
          tokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token
    contexts:
      - name: kube
        context:
          cluster: kube
          user: service-account
    current-context: kube
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-proxy
  namespace: kube-system
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: system:kube-proxy
subjects:
  - kind: ServiceAccount
    name: kube-proxy
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: system:node-proxier
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: kube-proxy
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: kube-proxy
  template:
    metadata:
      labels:
        app: kube-proxy
    spec:
      hostNetwork: true
      serviceAccountName: kube-proxy
      containers:
      - name: kube-proxy
        image: "{{ .HyperkubeImageURL }}"
        imagePullPolicy: IfNotPresent
        command:
        - /hyperkube
        - kube-proxy
        - --config=/var/lib/kubernetes/config.yaml
        securityContext:
          privileged: true
        volumeMounts:
        - name: config
          mountPath: /var/lib/kubernetes/
        livenessProbe:
          httpGet:
            path: /healthz
            port: 10256
        readinessProbe:
          httpGet:
            path: /healthz
            port: 10256
        resources:
          requests:
            cpu: "50m"
          limits:
            cpu: "100m"
      volumes:
      - name: config
        configMap:
          name: kube-proxy
`),
		},
		{
			Name:        "p8s-user-admin.yaml",
			Destination: ManifestAPI,
			Content: []byte(`---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: p8s-admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: User
  name: p8s
`),
		},
		{
			Name:        "coredns.yaml",
			Destination: ManifestAPI,
			Content: []byte(`---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: coredns
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: system:coredns
rules:
- apiGroups:
  - ""
  resources:
  - endpoints
  - services
  - pods
  - namespaces
  verbs:
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: system:coredns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:coredns
subjects:
- kind: ServiceAccount
  name: coredns
  namespace: kube-system
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns
  namespace: kube-system
data:
  Corefile: |
    .:53 {
        errors
        log
        health
        kubernetes cluster.local {{ .ServiceClusterIPRange }} {
          pods insecure
        }
        prometheus :9153
        forward . /etc/resolv.conf 8.8.8.8 8.8.4.4
        cache 30
    }
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: coredns
  namespace: kube-system
  labels:
    dns: coredns
    kubernetes.io/name: "CoreDNS"
spec:
  replicas: 1
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 0
  selector:
    matchLabels:
      dns: coredns
  template:
    metadata:
      labels:
        dns: coredns
    spec:
      serviceAccountName: coredns
      tolerations:
        - key: "CriticalAddonsOnly"
          operator: "Exists"
      containers:
      - name: coredns
        image: coredns/coredns:1.6.2
        imagePullPolicy: IfNotPresent
        args: [ "-conf", "/etc/coredns/Corefile" ]
        volumeMounts:
        - name: config-volume
          mountPath: /etc/coredns
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        - containerPort: 9153
          name: metrics
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
        resources:
          requests:
            cpu: "50m"
          limits:
            cpu: "100m"
      dnsPolicy: Default
      volumes:
      - name: config-volume
        configMap:
          name: coredns
          items:
          - key: Corefile
            path: Corefile
---
apiVersion: v1
kind: Service
metadata:
  name: coredns
  namespace: kube-system
  annotations:
  labels:
    dns: coredns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "CoreDNS"
spec:
  selector:
    dns: coredns
  clusterIP: {{ .DNSClusterIP }}
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP

`),
		},
	}
)
