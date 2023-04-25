/*
Copyright 2023 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
	networkv1 "k8s.io/cloud-provider-gcp/crd/apis/network/v1"
)

// FakeGKENetworkParamSetLists implements GKENetworkParamSetListInterface
type FakeGKENetworkParamSetLists struct {
	Fake *FakeNetworkingV1
}

var gkenetworkparamsetlistsResource = schema.GroupVersionResource{Group: "networking.gke.io", Version: "v1", Resource: "gkenetworkparamsetlists"}

var gkenetworkparamsetlistsKind = schema.GroupVersionKind{Group: "networking.gke.io", Version: "v1", Kind: "GKENetworkParamSetList"}

// Get takes name of the gKENetworkParamSetList, and returns the corresponding gKENetworkParamSetList object, and an error if there is any.
func (c *FakeGKENetworkParamSetLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkv1.GKENetworkParamSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(gkenetworkparamsetlistsResource, name), &networkv1.GKENetworkParamSetList{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkv1.GKENetworkParamSetList), err
}