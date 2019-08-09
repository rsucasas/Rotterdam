//
// ROTTERDAM application
// CLASS Project: https://class-project.eu/
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     https://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Created on 11 June 2019
// @author: Roi Sucasas - ATOS
//

package openshift

import (
	"atos/rotterdam/caas/adapters/openshift/impl"
	structs "atos/rotterdam/caas/common/structs"
)

// Openshift Adapter
type OpenshiftAdapter struct{}

// DeployTask: Deploy a task (k8s: deployment & service & volumes ...)
func (a OpenshiftAdapter) DeployTask(cluster int, namespace string, task structs.CLASS_TASK) (string, error) {
	return impl.DeployTask(cluster, namespace, task)
}

// DeployTaskCompss: Deploy a COMPSs task (k8s: deployment & service & volumes ...)
func (a OpenshiftAdapter) DeployTaskCompss(cluster int, namespace string, task structs.CLASS_TASK) (string, error) {
	return impl.DeployTaskCompss(cluster, namespace, task)
}

// GetTask: Gets a deployment
func (a OpenshiftAdapter) GetTask(cluster int, namespace string, name string) (structs.DB_TASK, error) {
	return impl.GetTask(cluster, namespace, name)
}

// GetConfig: k8s configuration
func (a OpenshiftAdapter) GetConfig() (string, error) {
	return impl.GetConfig()
}

// ScaleUpDown: Scale up task
func (a OpenshiftAdapter) ScaleUpDown(cluster int, task structs.DB_TASK, replicas int) (string, error) {
	return impl.ScaleUpDown(cluster, task, replicas)
}

// RemoveTask: Deletes a task
func (a OpenshiftAdapter) RemoveTask(cluster int, namespace string, name string) (string, error) {
	return impl.RemoveTask(cluster, namespace, name)
}