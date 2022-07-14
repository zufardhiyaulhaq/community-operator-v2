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
	"github.com/zufardhiyaulhaq/community-operator-v2/pkg/message"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AnnouncementSpec defines the desired state of Announcement
type AnnouncementSpec struct {
	Community []string              `json:"community"`
	Spec      AnnouncementSpec_Spec `json:"spec"`
}

type AnnouncementSpec_Spec struct {
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
	ImageUrl string   `json:"imageUrl"`
	Tags     []string `json:"tags"`
}

func (a AnnouncementSpec_Spec) ToMessageAnnouncement() message.Announcement {
	return message.Announcement{
		Subject:  a.Subject,
		Body:     a.Body,
		ImageUrl: a.ImageUrl,
		Tags:     a.Tags,
	}
}

// AnnouncementStatus defines the observed state of Announcement
type AnnouncementStatus struct {
	Status  string                     `json:"status"`
	Details AnnouncementStatus_Details `json:"details"`
}

type AnnouncementStatus_Details struct {
	Community map[string]AnnouncementStatus_Details_Community `json:"community"`
}

type AnnouncementStatus_Details_Community struct {
	Handler AnnouncementStatus_Details_Community_Handler `json:"handler"`
}
type AnnouncementStatus_Details_Community_Handler struct {
	Telegram map[string]AnnouncementStatus_Details_Community_Handler_Telegram `json:"telegram"`
	Twitter  map[string]AnnouncementStatus_Details_Community_Handler_Twitter  `json:"twitter"`
}

type AnnouncementStatus_Details_Community_Handler_Telegram struct {
	Status string `json:"status"`
}

type AnnouncementStatus_Details_Community_Handler_Twitter struct {
	Status string `json:"status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Announcement is the Schema for the announcements API
type Announcement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnnouncementSpec   `json:"spec,omitempty"`
	Status AnnouncementStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AnnouncementList contains a list of Announcement
type AnnouncementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Announcement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Announcement{}, &AnnouncementList{})
}
