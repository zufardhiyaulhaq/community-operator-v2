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

// TelegramHandlerSpec defines the desired state of TelegramHandler
type TelegramHandlerSpec struct {
	//+kubebuilder:validation:Enum=group;channel
	Type           string                              `json:"type"`
	Credential     string                              `json:"credential"`
	Authentication *TelegramHandlerSpec_Authentication `json:"authentication"`
}

type TelegramHandlerSpec_Authentication struct {
	Token *TelegramHandlerSpec_Authentication_Token `json:"token"`
}

type TelegramHandlerSpec_Authentication_Token struct {
	Secret Secret `json:"secret"`
}

// TelegramHandlerStatus defines the observed state of TelegramHandler
type TelegramHandlerStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TelegramHandler is the Schema for the telegramhandlers API
type TelegramHandler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TelegramHandlerSpec   `json:"spec,omitempty"`
	Status TelegramHandlerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TelegramHandlerList contains a list of TelegramHandler
type TelegramHandlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TelegramHandler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TelegramHandler{}, &TelegramHandlerList{})
}
