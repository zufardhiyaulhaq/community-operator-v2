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

package controllers

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	communityv1alpha1 "github.com/zufardhiyaulhaq/community-operator-v2/api/v1alpha1"
	operatorClient "github.com/zufardhiyaulhaq/community-operator-v2/pkg/client"
	operatorHandler "github.com/zufardhiyaulhaq/community-operator-v2/pkg/handler"
	operatorHelper "github.com/zufardhiyaulhaq/community-operator-v2/pkg/helper"
)

// MeetupReconciler reconciles a Meetup object
type MeetupReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=community.zufardhiyaulhaq.com,resources=meetups,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=community.zufardhiyaulhaq.com,resources=meetups/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=community.zufardhiyaulhaq.com,resources=meetups/finalizers,verbs=update

func (r *MeetupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("Start Meetup Reconciler")

	meetup := &communityv1alpha1.Meetup{}
	err := r.Client.Get(ctx, req.NamespacedName, meetup)
	if err != nil {
		return ctrl.Result{}, nil
	}

	if meetup.Status.Status == "send" {
		log.Info("meetup has been send for all communities, skip reconcile")
		return ctrl.Result{}, nil
	}

	meetupMessage := meetup.Spec.Spec.ToMessageMeetup()

	communities := &communityv1alpha1.CommunityList{}
	err = r.Client.List(ctx, communities)
	if err != nil {
		return ctrl.Result{}, err
	}

	var filteredCommunities []communityv1alpha1.Community
	for _, community := range communities.Items {
		if community.Namespace == req.Namespace && operatorHelper.StringInSlice(community.Name, meetup.Spec.Community) {
			filteredCommunities = append(filteredCommunities, community)
		}
	}

	if len(filteredCommunities) != len(meetup.Spec.Community) {
		return ctrl.Result{}, fmt.Errorf("cannot find community object in meetup.Spec.Community")
	}

	for _, community := range communities.Items {
		telegramHandlers := &communityv1alpha1.TelegramHandlerList{}
		err = r.Client.List(ctx, telegramHandlers)
		if err != nil {
			return ctrl.Result{}, err
		}

		var filteredtelegramHandlers []communityv1alpha1.TelegramHandler
		for _, telegramHandler := range telegramHandlers.Items {
			if telegramHandler.Namespace == req.Namespace && operatorHelper.StringInSlice(telegramHandler.Name, community.Spec.SocialMedia.Telegram) {
				filteredtelegramHandlers = append(filteredtelegramHandlers, telegramHandler)
			}
		}

		if len(filteredtelegramHandlers) != len(community.Spec.SocialMedia.Telegram) {
			return ctrl.Result{}, fmt.Errorf("cannot find TelegramHandler object in community.Spec.SocialMedia.Telegram")
		}

		for _, telegramHandler := range filteredtelegramHandlers {
			if _, ok := meetup.Status.Details.Community[community.Name]; ok {
				if _, ok := meetup.Status.Details.Community[community.Name].Handler.Telegram[telegramHandler.Name]; ok {
					if meetup.Status.Details.Community[community.Name].Handler.Telegram[telegramHandler.Name].Status == "send" {
						log.Info(fmt.Sprintf("skip sending meetup on community %s on telegram handler %s", community.Name, telegramHandler.Name))
						continue
					}
				}
			}

			secret := &corev1.Secret{}
			err = r.Client.Get(context.TODO(), types.NamespacedName{Name: telegramHandler.Spec.Authentication.Token.Secret.Name, Namespace: req.Namespace}, secret)
			if err != nil {
				return ctrl.Result{}, err
			}

			tokenByte, ok := secret.Data[telegramHandler.Spec.Authentication.Token.Secret.Key]
			if !ok {
				return ctrl.Result{}, err
			}

			token := string(tokenByte)
			credential := telegramHandler.Spec.Credential
			handlerType, ok := operatorClient.ParseTelegramChatType(telegramHandler.Spec.Type)
			if !ok {
				return ctrl.Result{}, fmt.Errorf("cannot parse Telegram Handler Type")
			}

			client, err := operatorClient.NewTelegramClient(token, credential, handlerType)
			if err != nil {
				return ctrl.Result{}, err
			}

			log.Info(fmt.Sprintf("sending meetup on community %s on telegram handler %s", community.Name, telegramHandler.Name))

			handler := operatorHandler.NewTelegramHandler(client)
			err = handler.SendMessage(meetupMessage)
			if err != nil {
				return ctrl.Result{}, err
			}

			if meetup.Status.Details.Community == nil {
				meetup.Status.Details = communityv1alpha1.MeetupStatus_Details{
					Community: map[string]communityv1alpha1.MeetupStatus_Details_Community{
						community.Name: {
							Handler: communityv1alpha1.MeetupStatus_Details_Community_Handler{
								Telegram: map[string]communityv1alpha1.MeetupStatus_Details_Community_Handler_Telegram{
									telegramHandler.Name: {
										Status: "send",
									},
								},
							},
						},
					},
				}
			} else if _, ok := meetup.Status.Details.Community[community.Name]; ok {
				meetup.Status.Details.Community[community.Name].Handler.Telegram[telegramHandler.Name] = communityv1alpha1.MeetupStatus_Details_Community_Handler_Telegram{
					Status: "send",
				}
			} else {
				meetup.Status.Details.Community[community.Name] = communityv1alpha1.MeetupStatus_Details_Community{
					Handler: communityv1alpha1.MeetupStatus_Details_Community_Handler{
						Telegram: map[string]communityv1alpha1.MeetupStatus_Details_Community_Handler_Telegram{
							telegramHandler.Name: {
								Status: "send",
							},
						},
					},
				}
			}

			err = r.Client.Status().Update(context.TODO(), meetup)
			if err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	meetup.Status.Status = "send"
	err = r.Client.Status().Update(context.TODO(), meetup)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MeetupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&communityv1alpha1.Meetup{}).
		Complete(r)
}
