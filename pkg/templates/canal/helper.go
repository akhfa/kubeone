/*
Copyright 2019 The KubeOne Authors.

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

package canal

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	dynclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func simpleCreateOrUpdate(ctx context.Context, client dynclient.Client, obj runtime.Object) error {
	okFunc := func(runtime.Object) error { return nil }
	_, err := controllerutil.CreateOrUpdate(ctx, client, obj, okFunc)
	return err
}
