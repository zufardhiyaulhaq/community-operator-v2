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

// WeeklySpec defines the desired state of Weekly
type WeeklySpec struct {
	Community []string        `json:"community"`
	Spec      WeeklySpec_Spec `json:"spec"`
}

type WeeklySpec_Spec struct {
	Name     string               `json:"name"`
	Date     string               `json:"date"`
	ImageUrl string               `json:"image_url"`
	Tags     []string             `json:"tags"`
	Articles []WeeklySpec_Article `json:"articles"`
}

type WeeklySpec_Article struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Type  string `json:"type"`
}

// WeeklyStatus defines the observed state of Weekly
type WeeklyStatus struct {
	Status string `json:"status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Weekly is the Schema for the weeklies API
type Weekly struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WeeklySpec   `json:"spec,omitempty"`
	Status WeeklyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WeeklyList contains a list of Weekly
type WeeklyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Weekly `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Weekly{}, &WeeklyList{})
}
