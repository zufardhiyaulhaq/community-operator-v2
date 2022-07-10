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
	"github.com/zufardhiyaulhaq/community-operator/pkg/message"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MeetupSpec defines the desired state of Meetup
type MeetupSpec struct {
	Community []string        `json:"community"`
	Spec      MeetupSpec_Spec `json:"spec"`
}

type MeetupSpec_Spec struct {
	Name            string                    `json:"name"`
	Date            string                    `json:"date"`
	Time            string                    `json:"time"`
	Place           string                    `json:"place"`
	RegistrationUrl string                    `json:"registrationUrl"`
	ImageUrl        string                    `json:"imageUrl"`
	Tags            []string                  `json:"tags"`
	Sponsors        []MeetupSpec_Spec_Sponsor `json:"sponsors"`
	Speakers        []MeetupSpec_Spec_Speaker `json:"speakers"`
}

func (m MeetupSpec_Spec) ToMessageMeetup() message.Meetup {
	messageMeetup := message.Meetup{
		Name:            m.Name,
		Date:            m.Date,
		Time:            m.Time,
		Place:           m.Place,
		RegistrationUrl: m.RegistrationUrl,
		ImageUrl:        m.ImageUrl,
		Tags:            m.Tags,
	}

	var messageMeetupSponsors []message.Meetup_Sponsor
	for _, sponsor := range m.Sponsors {
		messageMeetupSponsors = append(messageMeetupSponsors, message.Meetup_Sponsor{
			Name:     sponsor.Name,
			ImageUrl: sponsor.ImageUrl,
			Website:  sponsor.Website,
		})
	}
	messageMeetup.Sponsors = messageMeetupSponsors

	var messageMeetupSpeakers []message.Meetup_Speaker
	for _, speaker := range m.Speakers {
		messageMeetupSpeakers = append(messageMeetupSpeakers, message.Meetup_Speaker{
			Name:     speaker.Name,
			Position: speaker.Position,
			Company:  speaker.Company,
			Title:    speaker.Title,
			ImageUrl: speaker.ImageUrl,
		})
	}
	messageMeetup.Speakers = messageMeetupSpeakers

	return messageMeetup
}

type MeetupSpec_Spec_Speaker struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Company  string `json:"company"`
	Title    string `json:"title"`
	// +optional
	ImageUrl *string `json:"imageUrl"`
}

type MeetupSpec_Spec_Sponsor struct {
	Name string `json:"name"`
	// +optional
	ImageUrl *string `json:"imageUrl"`
	// +optional
	Website *string `json:"website"`
}

// MeetupStatus defines the observed state of Meetup
type MeetupStatus struct {
	Status  string               `json:"status"`
	Details MeetupStatus_Details `json:"details"`
}

type MeetupStatus_Details struct {
	Community map[string]MeetupStatus_Details_Community `json:"community"`
}

type MeetupStatus_Details_Community struct {
	Handler MeetupStatus_Details_Community_Handler `json:"handler"`
}
type MeetupStatus_Details_Community_Handler struct {
	Telegram map[string]MeetupStatus_Details_Community_Handler_Telegram `json:"telegram"`
}

type MeetupStatus_Details_Community_Handler_Telegram struct {
	Status string `json:"status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Meetup is the Schema for the meetups API
type Meetup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MeetupSpec   `json:"spec,omitempty"`
	Status MeetupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MeetupList contains a list of Meetup
type MeetupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Meetup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Meetup{}, &MeetupList{})
}
