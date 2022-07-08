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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	communityv1alpha1 "github.com/zufardhiyaulhaq/community-operator/api/v1alpha1"
	"github.com/zufardhiyaulhaq/community-operator/pkg/helper"
)

// WeeklyReconciler reconciles a Weekly object
type WeeklyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=community.zufardhiyaulhaq.com,resources=weeklies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=community.zufardhiyaulhaq.com,resources=weeklies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=community.zufardhiyaulhaq.com,resources=weeklies/finalizers,verbs=update

func (r *WeeklyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("Start Weekly Reconciler")

	log.Info("find weekly configuration")
	weekly := &communityv1alpha1.Weekly{}
	err := r.Client.Get(ctx, req.NamespacedName, weekly)
	if err != nil {
		return ctrl.Result{}, nil
	}

	log.Info("list community configuration")
	communities := &communityv1alpha1.CommunityList{}
	err = r.Client.List(ctx, communities)
	if err != nil {
		return ctrl.Result{}, err
	}

	var filteredCommunities []communityv1alpha1.Community
	for _, community := range communities.Items {
		if community.Namespace == req.Namespace && helper.StringInSlice(community.Name, weekly.Spec.Community) {
			filteredCommunities = append(filteredCommunities, community)
		}
	}

	if len(filteredCommunities) != len(weekly.Spec.Community) {
		return ctrl.Result{}, fmt.Errorf("cannot find community object in weekly.Spec.Community")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WeeklyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&communityv1alpha1.Weekly{}).
		Complete(r)
}
