/*
Copyright 2020 The Skaffold Authors

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

package validator

import (
	"fmt"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/GoogleContainerTools/skaffold/proto"
)

type Resource struct {
	namespace string
	kind      string
	name      string
	status    Status
	err       error
	ErrCode   proto.ErrorCode
}

func (r Resource) Kind() string      { return r.kind }
func (r Resource) Name() string      { return r.name }
func (r Resource) Namespace() string { return r.namespace }
func (r Resource) Status() Status    { return r.status }
func (r Resource) Error() error      { return r.err }
func (r Resource) String() string {
	return fmt.Sprintf("{%s:%s/%s}", r.kind, r.namespace, r.name)
}

// NewResource creates new Resource of kind
func NewResource(namespace, kind, name string, status Status, err error, errCode proto.ErrorCode) Resource {
	return Resource{namespace: namespace, kind: kind, name: name, status: status, err: err, ErrCode: errCode}
}

// objectWithMetadata is any k8s object that has kind and object metadata.
type objectWithMetadata interface {
	runtime.Object
	meta_v1.Object
}

// NewResourceFromObject creates new Resource with fields populated from object metadata.
func NewResourceFromObject(object objectWithMetadata, status Status, err error, errCode proto.ErrorCode) Resource {
	return NewResource(object.GetNamespace(), object.GetObjectKind().GroupVersionKind().Kind, object.GetName(), status, err, errCode)
}
