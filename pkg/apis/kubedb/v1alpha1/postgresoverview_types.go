/*


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
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresOverview
// +k8s:openapi-gen=true
type PostgresOverview struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgresOverviewSpec   `json:"spec,omitempty"`
	Status PostgresOverviewStatus `json:"status,omitempty"`
}

// PostgresOverviewList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PostgresOverviewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []PostgresOverview `json:"items"`
}

// PostgresOverviewSpec defines the desired state of PostgresOverview
type PostgresOverviewSpec struct {
}

var _ resource.Object = &PostgresOverview{}
var _ resourcestrategy.Validater = &PostgresOverview{}

func (in *PostgresOverview) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *PostgresOverview) NamespaceScoped() bool {
	return false
}

func (in *PostgresOverview) New() runtime.Object {
	return &PostgresOverview{}
}

func (in *PostgresOverview) NewList() runtime.Object {
	return &PostgresOverviewList{}
}

func (in *PostgresOverview) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "kubedb.ui.bytebuilder.dev",
		Version:  "v1alpha1",
		Resource: "postgresoverviews",
	}
}

func (in *PostgresOverview) IsStorageVersion() bool {
	return true
}

func (in *PostgresOverview) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &PostgresOverviewList{}

func (in *PostgresOverviewList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// PostgresOverviewStatus defines the observed state of PostgresOverview
type PostgresOverviewStatus struct {
}

func (in PostgresOverviewStatus) SubResourceName() string {
	return "status"
}

// PostgresOverview implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &PostgresOverview{}

func (in *PostgresOverview) GetStatus() resource.StatusSubResource {
	return in.Status
}

// PostgresOverviewStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &PostgresOverviewStatus{}

func (in PostgresOverviewStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*PostgresOverview).Status = in
}
