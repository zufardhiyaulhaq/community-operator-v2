/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CommunitySpec defines the desired state of Community
type CommunitySpec struct {
	SocialMedia CommunitySpec_SocialMedia `json:"socialMedia" yaml:"socialMedia"`
}

type CommunitySpec_SocialMedia struct {
	Telegram []string `json:"telegram"`
	Twitter  []string `json:"twitter"`
}

// CommunityStatus defines the observed state of Community
type CommunityStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Community is the Schema for the communities API
type Community struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CommunitySpec   `json:"spec,omitempty"`
	Status CommunityStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CommunityList contains a list of Community
type CommunityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Community `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Community{}, &CommunityList{})
}
