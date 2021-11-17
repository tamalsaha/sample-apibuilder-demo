
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

// MongoDBInfoView
// +k8s:openapi-gen=true
type MongoDBInfoView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBInfoViewSpec   `json:"spec,omitempty"`
	Status MongoDBInfoViewStatus `json:"status,omitempty"`
}

// MongoDBInfoViewList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MongoDBInfoViewList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MongoDBInfoView `json:"items"`
}

// MongoDBInfoViewSpec defines the desired state of MongoDBInfoView
type MongoDBInfoViewSpec struct {
}

var _ resource.Object = &MongoDBInfoView{}
var _ resourcestrategy.Validater = &MongoDBInfoView{}

func (in *MongoDBInfoView) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *MongoDBInfoView) NamespaceScoped() bool {
	return false
}

func (in *MongoDBInfoView) New() runtime.Object {
	return &MongoDBInfoView{}
}

func (in *MongoDBInfoView) NewList() runtime.Object {
	return &MongoDBInfoViewList{}
}

func (in *MongoDBInfoView) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "kubedb.ui.bytebuilder.dev",
		Version:  "v1alpha1",
		Resource: "mongodbinfoviews",
	}
}

func (in *MongoDBInfoView) IsStorageVersion() bool {
	return true
}

func (in *MongoDBInfoView) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &MongoDBInfoViewList{}

func (in *MongoDBInfoViewList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}
// MongoDBInfoViewStatus defines the observed state of MongoDBInfoView
type MongoDBInfoViewStatus struct {
}

func (in MongoDBInfoViewStatus) SubResourceName() string {
	return "status"
}

// MongoDBInfoView implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &MongoDBInfoView{}

func (in *MongoDBInfoView) GetStatus() resource.StatusSubResource {
	return in.Status
}

// MongoDBInfoViewStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &MongoDBInfoViewStatus{}

func (in MongoDBInfoViewStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*MongoDBInfoView).Status = in
}
