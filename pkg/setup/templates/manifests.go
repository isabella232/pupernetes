// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package templates

const (
	// ManifestStaticPod path for static pods
	ManifestStaticPod = "/manifest-static-pod"
	// ManifestAPI path for kube-apiserver manifests
	ManifestAPI = "/manifest-api"
	// ManifestConfig path for configuration related to Kubernetes
	// like kubeconfig, audit files, ...
	ManifestConfig = "/manifest-config"
	// ManifestSystemdUnit path where systemd units are stored
	ManifestSystemdUnit = "/manifest-systemd-unit"
)

// Manifest represent a file to be rendered in a destination
type Manifest struct {
	Name        string
	Destination string
	Content     []byte
}

// Manifests is the map catalog where all Kubernetes major.minor are stored
var Manifests map[string][]Manifest

// KubeTaggedVersions is a mapping between string tags and real kube versions to ease usage
var KubeTaggedVersions map[string]string

// TODO add a layer for flavor like, http, https
func init() {
	Manifests = map[string][]Manifest{
		"1.18": manifest1o18,
		"1.17": manifest1o17,
		"1.16": manifest1o16,
		"1.15": manifest1o15,
		"1.14": manifest1o14,
		"1.13": manifest1o13,
		"1.12": manifest1o12,
		"1.11": manifest1o11,
		"1.10": manifest1o10,
		"1.9":  manifest1o9,
		"1.8":  manifest1o8,
		"1.7":  manifest1o7,
		"1.6":  manifest1o6,
		"1.5":  manifest1o5,
	}

	KubeTaggedVersions = map[string]string{
		"latest": "1.16.3",
		"next":   "1.17.0",
	}
}
