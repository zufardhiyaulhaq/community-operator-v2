//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Announcement) DeepCopyInto(out *Announcement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Announcement.
func (in *Announcement) DeepCopy() *Announcement {
	if in == nil {
		return nil
	}
	out := new(Announcement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Announcement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementList) DeepCopyInto(out *AnnouncementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Announcement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementList.
func (in *AnnouncementList) DeepCopy() *AnnouncementList {
	if in == nil {
		return nil
	}
	out := new(AnnouncementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnnouncementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementSpec) DeepCopyInto(out *AnnouncementSpec) {
	*out = *in
	if in.Community != nil {
		in, out := &in.Community, &out.Community
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementSpec.
func (in *AnnouncementSpec) DeepCopy() *AnnouncementSpec {
	if in == nil {
		return nil
	}
	out := new(AnnouncementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementSpec_Spec) DeepCopyInto(out *AnnouncementSpec_Spec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementSpec_Spec.
func (in *AnnouncementSpec_Spec) DeepCopy() *AnnouncementSpec_Spec {
	if in == nil {
		return nil
	}
	out := new(AnnouncementSpec_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementStatus) DeepCopyInto(out *AnnouncementStatus) {
	*out = *in
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementStatus.
func (in *AnnouncementStatus) DeepCopy() *AnnouncementStatus {
	if in == nil {
		return nil
	}
	out := new(AnnouncementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementStatus_Details) DeepCopyInto(out *AnnouncementStatus_Details) {
	*out = *in
	if in.Community != nil {
		in, out := &in.Community, &out.Community
		*out = make(map[string]AnnouncementStatus_Details_Community, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementStatus_Details.
func (in *AnnouncementStatus_Details) DeepCopy() *AnnouncementStatus_Details {
	if in == nil {
		return nil
	}
	out := new(AnnouncementStatus_Details)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementStatus_Details_Community) DeepCopyInto(out *AnnouncementStatus_Details_Community) {
	*out = *in
	in.Handler.DeepCopyInto(&out.Handler)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementStatus_Details_Community.
func (in *AnnouncementStatus_Details_Community) DeepCopy() *AnnouncementStatus_Details_Community {
	if in == nil {
		return nil
	}
	out := new(AnnouncementStatus_Details_Community)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementStatus_Details_Community_Handler) DeepCopyInto(out *AnnouncementStatus_Details_Community_Handler) {
	*out = *in
	if in.Telegram != nil {
		in, out := &in.Telegram, &out.Telegram
		*out = make(map[string]AnnouncementStatus_Details_Community_Handler_Telegram, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Twitter != nil {
		in, out := &in.Twitter, &out.Twitter
		*out = make(map[string]AnnouncementStatus_Details_Community_Handler_Twitter, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementStatus_Details_Community_Handler.
func (in *AnnouncementStatus_Details_Community_Handler) DeepCopy() *AnnouncementStatus_Details_Community_Handler {
	if in == nil {
		return nil
	}
	out := new(AnnouncementStatus_Details_Community_Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementStatus_Details_Community_Handler_Telegram) DeepCopyInto(out *AnnouncementStatus_Details_Community_Handler_Telegram) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementStatus_Details_Community_Handler_Telegram.
func (in *AnnouncementStatus_Details_Community_Handler_Telegram) DeepCopy() *AnnouncementStatus_Details_Community_Handler_Telegram {
	if in == nil {
		return nil
	}
	out := new(AnnouncementStatus_Details_Community_Handler_Telegram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnnouncementStatus_Details_Community_Handler_Twitter) DeepCopyInto(out *AnnouncementStatus_Details_Community_Handler_Twitter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnnouncementStatus_Details_Community_Handler_Twitter.
func (in *AnnouncementStatus_Details_Community_Handler_Twitter) DeepCopy() *AnnouncementStatus_Details_Community_Handler_Twitter {
	if in == nil {
		return nil
	}
	out := new(AnnouncementStatus_Details_Community_Handler_Twitter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Community) DeepCopyInto(out *Community) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Community.
func (in *Community) DeepCopy() *Community {
	if in == nil {
		return nil
	}
	out := new(Community)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Community) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommunityList) DeepCopyInto(out *CommunityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Community, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommunityList.
func (in *CommunityList) DeepCopy() *CommunityList {
	if in == nil {
		return nil
	}
	out := new(CommunityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommunityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommunitySpec) DeepCopyInto(out *CommunitySpec) {
	*out = *in
	in.SocialMedia.DeepCopyInto(&out.SocialMedia)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommunitySpec.
func (in *CommunitySpec) DeepCopy() *CommunitySpec {
	if in == nil {
		return nil
	}
	out := new(CommunitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommunitySpec_SocialMedia) DeepCopyInto(out *CommunitySpec_SocialMedia) {
	*out = *in
	if in.Telegram != nil {
		in, out := &in.Telegram, &out.Telegram
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Twitter != nil {
		in, out := &in.Twitter, &out.Twitter
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommunitySpec_SocialMedia.
func (in *CommunitySpec_SocialMedia) DeepCopy() *CommunitySpec_SocialMedia {
	if in == nil {
		return nil
	}
	out := new(CommunitySpec_SocialMedia)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommunityStatus) DeepCopyInto(out *CommunityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommunityStatus.
func (in *CommunityStatus) DeepCopy() *CommunityStatus {
	if in == nil {
		return nil
	}
	out := new(CommunityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Meetup) DeepCopyInto(out *Meetup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Meetup.
func (in *Meetup) DeepCopy() *Meetup {
	if in == nil {
		return nil
	}
	out := new(Meetup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Meetup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupList) DeepCopyInto(out *MeetupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Meetup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupList.
func (in *MeetupList) DeepCopy() *MeetupList {
	if in == nil {
		return nil
	}
	out := new(MeetupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeetupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupSpec) DeepCopyInto(out *MeetupSpec) {
	*out = *in
	if in.Community != nil {
		in, out := &in.Community, &out.Community
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupSpec.
func (in *MeetupSpec) DeepCopy() *MeetupSpec {
	if in == nil {
		return nil
	}
	out := new(MeetupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupSpec_Spec) DeepCopyInto(out *MeetupSpec_Spec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Sponsors != nil {
		in, out := &in.Sponsors, &out.Sponsors
		*out = make([]MeetupSpec_Spec_Sponsor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Speakers != nil {
		in, out := &in.Speakers, &out.Speakers
		*out = make([]MeetupSpec_Spec_Speaker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupSpec_Spec.
func (in *MeetupSpec_Spec) DeepCopy() *MeetupSpec_Spec {
	if in == nil {
		return nil
	}
	out := new(MeetupSpec_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupSpec_Spec_Speaker) DeepCopyInto(out *MeetupSpec_Spec_Speaker) {
	*out = *in
	if in.ImageUrl != nil {
		in, out := &in.ImageUrl, &out.ImageUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupSpec_Spec_Speaker.
func (in *MeetupSpec_Spec_Speaker) DeepCopy() *MeetupSpec_Spec_Speaker {
	if in == nil {
		return nil
	}
	out := new(MeetupSpec_Spec_Speaker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupSpec_Spec_Sponsor) DeepCopyInto(out *MeetupSpec_Spec_Sponsor) {
	*out = *in
	if in.ImageUrl != nil {
		in, out := &in.ImageUrl, &out.ImageUrl
		*out = new(string)
		**out = **in
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupSpec_Spec_Sponsor.
func (in *MeetupSpec_Spec_Sponsor) DeepCopy() *MeetupSpec_Spec_Sponsor {
	if in == nil {
		return nil
	}
	out := new(MeetupSpec_Spec_Sponsor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupStatus) DeepCopyInto(out *MeetupStatus) {
	*out = *in
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupStatus.
func (in *MeetupStatus) DeepCopy() *MeetupStatus {
	if in == nil {
		return nil
	}
	out := new(MeetupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupStatus_Details) DeepCopyInto(out *MeetupStatus_Details) {
	*out = *in
	if in.Community != nil {
		in, out := &in.Community, &out.Community
		*out = make(map[string]MeetupStatus_Details_Community, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupStatus_Details.
func (in *MeetupStatus_Details) DeepCopy() *MeetupStatus_Details {
	if in == nil {
		return nil
	}
	out := new(MeetupStatus_Details)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupStatus_Details_Community) DeepCopyInto(out *MeetupStatus_Details_Community) {
	*out = *in
	in.Handler.DeepCopyInto(&out.Handler)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupStatus_Details_Community.
func (in *MeetupStatus_Details_Community) DeepCopy() *MeetupStatus_Details_Community {
	if in == nil {
		return nil
	}
	out := new(MeetupStatus_Details_Community)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupStatus_Details_Community_Handler) DeepCopyInto(out *MeetupStatus_Details_Community_Handler) {
	*out = *in
	if in.Telegram != nil {
		in, out := &in.Telegram, &out.Telegram
		*out = make(map[string]MeetupStatus_Details_Community_Handler_Telegram, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Twitter != nil {
		in, out := &in.Twitter, &out.Twitter
		*out = make(map[string]MeetupStatus_Details_Community_Handler_Twitter, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupStatus_Details_Community_Handler.
func (in *MeetupStatus_Details_Community_Handler) DeepCopy() *MeetupStatus_Details_Community_Handler {
	if in == nil {
		return nil
	}
	out := new(MeetupStatus_Details_Community_Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupStatus_Details_Community_Handler_Telegram) DeepCopyInto(out *MeetupStatus_Details_Community_Handler_Telegram) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupStatus_Details_Community_Handler_Telegram.
func (in *MeetupStatus_Details_Community_Handler_Telegram) DeepCopy() *MeetupStatus_Details_Community_Handler_Telegram {
	if in == nil {
		return nil
	}
	out := new(MeetupStatus_Details_Community_Handler_Telegram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeetupStatus_Details_Community_Handler_Twitter) DeepCopyInto(out *MeetupStatus_Details_Community_Handler_Twitter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeetupStatus_Details_Community_Handler_Twitter.
func (in *MeetupStatus_Details_Community_Handler_Twitter) DeepCopy() *MeetupStatus_Details_Community_Handler_Twitter {
	if in == nil {
		return nil
	}
	out := new(MeetupStatus_Details_Community_Handler_Twitter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelegramHandler) DeepCopyInto(out *TelegramHandler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelegramHandler.
func (in *TelegramHandler) DeepCopy() *TelegramHandler {
	if in == nil {
		return nil
	}
	out := new(TelegramHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TelegramHandler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelegramHandlerList) DeepCopyInto(out *TelegramHandlerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TelegramHandler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelegramHandlerList.
func (in *TelegramHandlerList) DeepCopy() *TelegramHandlerList {
	if in == nil {
		return nil
	}
	out := new(TelegramHandlerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TelegramHandlerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelegramHandlerSpec) DeepCopyInto(out *TelegramHandlerSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(TelegramHandlerSpec_Authentication)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelegramHandlerSpec.
func (in *TelegramHandlerSpec) DeepCopy() *TelegramHandlerSpec {
	if in == nil {
		return nil
	}
	out := new(TelegramHandlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelegramHandlerSpec_Authentication) DeepCopyInto(out *TelegramHandlerSpec_Authentication) {
	*out = *in
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(TelegramHandlerSpec_Authentication_Token)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelegramHandlerSpec_Authentication.
func (in *TelegramHandlerSpec_Authentication) DeepCopy() *TelegramHandlerSpec_Authentication {
	if in == nil {
		return nil
	}
	out := new(TelegramHandlerSpec_Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelegramHandlerSpec_Authentication_Token) DeepCopyInto(out *TelegramHandlerSpec_Authentication_Token) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelegramHandlerSpec_Authentication_Token.
func (in *TelegramHandlerSpec_Authentication_Token) DeepCopy() *TelegramHandlerSpec_Authentication_Token {
	if in == nil {
		return nil
	}
	out := new(TelegramHandlerSpec_Authentication_Token)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelegramHandlerStatus) DeepCopyInto(out *TelegramHandlerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelegramHandlerStatus.
func (in *TelegramHandlerStatus) DeepCopy() *TelegramHandlerStatus {
	if in == nil {
		return nil
	}
	out := new(TelegramHandlerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandler) DeepCopyInto(out *TwitterHandler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandler.
func (in *TwitterHandler) DeepCopy() *TwitterHandler {
	if in == nil {
		return nil
	}
	out := new(TwitterHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwitterHandler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerList) DeepCopyInto(out *TwitterHandlerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TwitterHandler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerList.
func (in *TwitterHandlerList) DeepCopy() *TwitterHandlerList {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwitterHandlerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerSpec) DeepCopyInto(out *TwitterHandlerSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(TwitterHandlerSpec_Authentication)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerSpec.
func (in *TwitterHandlerSpec) DeepCopy() *TwitterHandlerSpec {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerSpec_Authentication) DeepCopyInto(out *TwitterHandlerSpec_Authentication) {
	*out = *in
	if in.ApiKey != nil {
		in, out := &in.ApiKey, &out.ApiKey
		*out = new(TwitterHandlerSpec_Authentication_ApiKey)
		**out = **in
	}
	if in.ApiKeySecret != nil {
		in, out := &in.ApiKeySecret, &out.ApiKeySecret
		*out = new(TwitterHandlerSpec_Authentication_ApiKeySecret)
		**out = **in
	}
	if in.AccessToken != nil {
		in, out := &in.AccessToken, &out.AccessToken
		*out = new(TwitterHandlerSpec_Authentication_AccessToken)
		**out = **in
	}
	if in.AccessTokenSecret != nil {
		in, out := &in.AccessTokenSecret, &out.AccessTokenSecret
		*out = new(TwitterHandlerSpec_Authentication_AccessTokenSecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerSpec_Authentication.
func (in *TwitterHandlerSpec_Authentication) DeepCopy() *TwitterHandlerSpec_Authentication {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerSpec_Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerSpec_Authentication_AccessToken) DeepCopyInto(out *TwitterHandlerSpec_Authentication_AccessToken) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerSpec_Authentication_AccessToken.
func (in *TwitterHandlerSpec_Authentication_AccessToken) DeepCopy() *TwitterHandlerSpec_Authentication_AccessToken {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerSpec_Authentication_AccessToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerSpec_Authentication_AccessTokenSecret) DeepCopyInto(out *TwitterHandlerSpec_Authentication_AccessTokenSecret) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerSpec_Authentication_AccessTokenSecret.
func (in *TwitterHandlerSpec_Authentication_AccessTokenSecret) DeepCopy() *TwitterHandlerSpec_Authentication_AccessTokenSecret {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerSpec_Authentication_AccessTokenSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerSpec_Authentication_ApiKey) DeepCopyInto(out *TwitterHandlerSpec_Authentication_ApiKey) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerSpec_Authentication_ApiKey.
func (in *TwitterHandlerSpec_Authentication_ApiKey) DeepCopy() *TwitterHandlerSpec_Authentication_ApiKey {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerSpec_Authentication_ApiKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerSpec_Authentication_ApiKeySecret) DeepCopyInto(out *TwitterHandlerSpec_Authentication_ApiKeySecret) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerSpec_Authentication_ApiKeySecret.
func (in *TwitterHandlerSpec_Authentication_ApiKeySecret) DeepCopy() *TwitterHandlerSpec_Authentication_ApiKeySecret {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerSpec_Authentication_ApiKeySecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwitterHandlerStatus) DeepCopyInto(out *TwitterHandlerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwitterHandlerStatus.
func (in *TwitterHandlerStatus) DeepCopy() *TwitterHandlerStatus {
	if in == nil {
		return nil
	}
	out := new(TwitterHandlerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Weekly) DeepCopyInto(out *Weekly) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Weekly.
func (in *Weekly) DeepCopy() *Weekly {
	if in == nil {
		return nil
	}
	out := new(Weekly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Weekly) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyList) DeepCopyInto(out *WeeklyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Weekly, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyList.
func (in *WeeklyList) DeepCopy() *WeeklyList {
	if in == nil {
		return nil
	}
	out := new(WeeklyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WeeklyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklySpec) DeepCopyInto(out *WeeklySpec) {
	*out = *in
	if in.Community != nil {
		in, out := &in.Community, &out.Community
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklySpec.
func (in *WeeklySpec) DeepCopy() *WeeklySpec {
	if in == nil {
		return nil
	}
	out := new(WeeklySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklySpec_Article) DeepCopyInto(out *WeeklySpec_Article) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklySpec_Article.
func (in *WeeklySpec_Article) DeepCopy() *WeeklySpec_Article {
	if in == nil {
		return nil
	}
	out := new(WeeklySpec_Article)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklySpec_Spec) DeepCopyInto(out *WeeklySpec_Spec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Articles != nil {
		in, out := &in.Articles, &out.Articles
		*out = make([]WeeklySpec_Article, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklySpec_Spec.
func (in *WeeklySpec_Spec) DeepCopy() *WeeklySpec_Spec {
	if in == nil {
		return nil
	}
	out := new(WeeklySpec_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStatus) DeepCopyInto(out *WeeklyStatus) {
	*out = *in
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStatus.
func (in *WeeklyStatus) DeepCopy() *WeeklyStatus {
	if in == nil {
		return nil
	}
	out := new(WeeklyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStatus_Details) DeepCopyInto(out *WeeklyStatus_Details) {
	*out = *in
	if in.Community != nil {
		in, out := &in.Community, &out.Community
		*out = make(map[string]WeeklyStatus_Details_Community, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStatus_Details.
func (in *WeeklyStatus_Details) DeepCopy() *WeeklyStatus_Details {
	if in == nil {
		return nil
	}
	out := new(WeeklyStatus_Details)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStatus_Details_Community) DeepCopyInto(out *WeeklyStatus_Details_Community) {
	*out = *in
	in.Handler.DeepCopyInto(&out.Handler)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStatus_Details_Community.
func (in *WeeklyStatus_Details_Community) DeepCopy() *WeeklyStatus_Details_Community {
	if in == nil {
		return nil
	}
	out := new(WeeklyStatus_Details_Community)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStatus_Details_Community_Handler) DeepCopyInto(out *WeeklyStatus_Details_Community_Handler) {
	*out = *in
	if in.Telegram != nil {
		in, out := &in.Telegram, &out.Telegram
		*out = make(map[string]WeeklyStatus_Details_Community_Handler_Telegram, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Twitter != nil {
		in, out := &in.Twitter, &out.Twitter
		*out = make(map[string]WeeklyStatus_Details_Community_Handler_Twitter, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStatus_Details_Community_Handler.
func (in *WeeklyStatus_Details_Community_Handler) DeepCopy() *WeeklyStatus_Details_Community_Handler {
	if in == nil {
		return nil
	}
	out := new(WeeklyStatus_Details_Community_Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStatus_Details_Community_Handler_Telegram) DeepCopyInto(out *WeeklyStatus_Details_Community_Handler_Telegram) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStatus_Details_Community_Handler_Telegram.
func (in *WeeklyStatus_Details_Community_Handler_Telegram) DeepCopy() *WeeklyStatus_Details_Community_Handler_Telegram {
	if in == nil {
		return nil
	}
	out := new(WeeklyStatus_Details_Community_Handler_Telegram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStatus_Details_Community_Handler_Twitter) DeepCopyInto(out *WeeklyStatus_Details_Community_Handler_Twitter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStatus_Details_Community_Handler_Twitter.
func (in *WeeklyStatus_Details_Community_Handler_Twitter) DeepCopy() *WeeklyStatus_Details_Community_Handler_Twitter {
	if in == nil {
		return nil
	}
	out := new(WeeklyStatus_Details_Community_Handler_Twitter)
	in.DeepCopyInto(out)
	return out
}
