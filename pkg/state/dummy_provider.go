/*
2023 NVIDIA CORPORATION & AFFILIATES

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package state

import (
	"github.com/Mellanox/network-operator/pkg/clustertype"
	"github.com/Mellanox/network-operator/pkg/nodeinfo"
	"github.com/Mellanox/network-operator/pkg/staticconfig"
)

type dummyProvider struct {
}

func (d *dummyProvider) GetClusterType() clustertype.Type {
	return clustertype.Kubernetes
}

func (d *dummyProvider) IsKubernetes() bool {
	return true
}

func (d *dummyProvider) IsOpenshift() bool {
	return false
}

func (d *dummyProvider) GetStaticConfig() staticconfig.StaticConfig {
	return staticconfig.StaticConfig{CniBinDirectory: ""}
}

func (d *dummyProvider) GetNodesAttributes(...nodeinfo.Filter) []nodeinfo.NodeAttributes {
	nodeAttr := make(map[nodeinfo.AttributeType]string)
	nodeAttr[nodeinfo.AttrTypeCPUArch] = "amd64"
	nodeAttr[nodeinfo.AttrTypeOSName] = "ubuntu"
	nodeAttr[nodeinfo.AttrTypeOSVer] = "20.04"

	return []nodeinfo.NodeAttributes{{Attributes: nodeAttr}}
}

func getDummyCatalog() InfoCatalog {
	catalog := NewInfoCatalog()
	catalog.Add(InfoTypeNodeInfo, &dummyProvider{})
	catalog.Add(InfoTypeStaticConfig, &dummyProvider{})
	catalog.Add(InfoTypeClusterType, &dummyProvider{})

	return catalog
}
