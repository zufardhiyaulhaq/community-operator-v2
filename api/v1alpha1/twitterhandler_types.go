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

// TwitterHandlerSpec defines the desired state of TwitterHandler
type TwitterHandlerSpec struct {
	Authentication *TwitterHandlerSpec_Authentication `json:"authentication"`
}

type TwitterHandlerSpec_Authentication struct {
	ApiKey            *TwitterHandlerSpec_Authentication_ApiKey            `json:"apiKey"`
	ApiKeySecret      *TwitterHandlerSpec_Authentication_ApiKeySecret      `json:"apiKeySecret"`
	AccessToken       *TwitterHandlerSpec_Authentication_AccessToken       `json:"accessToken"`
	AccessTokenSecret *TwitterHandlerSpec_Authentication_AccessTokenSecret `json:"accessTokenSecret"`
}

type TwitterHandlerSpec_Authentication_ApiKey struct {
	Secret Secret `json:"secret"`
}

type TwitterHandlerSpec_Authentication_ApiKeySecret struct {
	Secret Secret `json:"secret"`
}

type TwitterHandlerSpec_Authentication_AccessToken struct {
	Secret Secret `json:"secret"`
}

type TwitterHandlerSpec_Authentication_AccessTokenSecret struct {
	Secret Secret `json:"secret"`
}

// TwitterHandlerStatus defines the observed state of TwitterHandler
type TwitterHandlerStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TwitterHandler is the Schema for the twitterhandlers API
type TwitterHandler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TwitterHandlerSpec   `json:"spec,omitempty"`
	Status TwitterHandlerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TwitterHandlerList contains a list of TwitterHandler
type TwitterHandlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TwitterHandler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TwitterHandler{}, &TwitterHandlerList{})
}
