// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package add_kubernetes_metadata

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes/metadata"
)

type kubeAnnotatorConfig struct {
	KubeConfig        string                       `config:"kube_config"`
	KubeClientOptions kubernetes.KubeClientOptions `config:"kube_client_options"`
	Host              string                       `config:"host"`
	Scope             string                       `config:"scope"`
	Namespace         string                       `config:"namespace"`
	SyncPeriod        time.Duration                `config:"sync_period"`
	// Annotations are kept after pod is removed, until they haven't been accessed
	// for a full `cleanup_timeout`:
	CleanupTimeout  time.Duration `config:"cleanup_timeout"`
	Indexers        PluginConfig  `config:"indexers"`
	Matchers        PluginConfig  `config:"matchers"`
	DefaultMatchers Enabled       `config:"default_matchers"`
	DefaultIndexers Enabled       `config:"default_indexers"`

	AddResourceMetadata *metadata.AddResourceMetadataConfig `config:"add_resource_metadata"`
}

type Enabled struct {
	Enabled bool `config:"enabled"`
}

type PluginConfig []map[string]common.Config

func defaultKubernetesAnnotatorConfig() kubeAnnotatorConfig {
	return kubeAnnotatorConfig{
		SyncPeriod:      10 * time.Minute,
		CleanupTimeout:  60 * time.Second,
		DefaultMatchers: Enabled{true},
		DefaultIndexers: Enabled{true},
		Scope:           "node",
	}
}

func (k *kubeAnnotatorConfig) Validate() error {
	if k.Scope != "node" && k.Scope != "cluster" {
		return fmt.Errorf("invalid scope %s, valid values include `cluster`, `node`", k.Scope)
	}

	if k.Scope == "cluster" {
		k.Host = ""
	}

	return nil
}
