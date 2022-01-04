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
	"github.com/aquasecurity/aqua-operator/pkg/apis/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigAuditReportsSpec defines the desired state of ConfigAuditReports
type ConfigAuditReportsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Infrastructure                *v1alpha1.AquaInfrastructure `json:"infra,omitempty"`
	AllowAnyVersion               bool                         `json:"allowAnyVersion,omitempty"`
	StarboardService              *v1alpha1.AquaService        `json:"deploy,required"`
	Config                        v1alpha1.AquaStarboardConfig `json:"config"`
	RegistryData                  *v1alpha1.AquaDockerRegistry `json:"registry,omitempty"`
	ImageData                     *v1alpha1.AquaImage          `json:"image,omitempty"`
	Envs                          []corev1.EnvVar              `json:"env,omitempty"`
	LogDevMode                    bool                         `json:"logDevMode,omitempty"`
	ConcurrentScanJobsLimit       string                       `json:"concurrentScanJobsLimit,omitempty"`
	ScanJobRetryAfter             string                       `json:"scanJobRetryAfter,omitempty"`
	MetricsBindAddress            string                       `json:"metricsBindAddress,omitempty"`
	HealthProbeBindAddress        string                       `json:"healthProbeBindAddress,omitempty"`
	CisKubernetesBenchmarkEnabled string                       `json:"cisKubernetesBenchmarkEnabled,omitempty"`
	VulnerabilityScannerEnabled   string                       `json:"vulnerabilityScannerEnabled,omitempty"`
	BatchDeleteLimit              string                       `json:"batchDeleteLimit,omitempty"`
	BatchDeleteDelay              string                       `json:"batchDeleteDelay,omitempty"`
}

// ConfigAuditReportsStatus defines the observed state of ConfigAuditReports
type ConfigAuditReportsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string                     `json:"nodes"`
	State v1alpha1.AquaDeploymentState `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigAuditReports is the Schema for the configauditreports API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=configauditreports,scope=Namespaced
type ConfigAuditReports struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigAuditReportsSpec   `json:"spec,omitempty"`
	Status ConfigAuditReportsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigAuditReportsList contains a list of ConfigAuditReports
type ConfigAuditReportsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigAuditReports `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigAuditReports{}, &ConfigAuditReportsList{})
}
