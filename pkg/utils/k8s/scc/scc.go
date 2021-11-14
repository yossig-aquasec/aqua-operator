package scc

import (
	"fmt"
	v1 "github.com/openshift/api/security/v1"
	security1 "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// CheckIfAquaSecurityContextConstraintsExists Check if aqua-scc exists
func CheckIfAquaSecurityContextConstraintsExists(v1Client security1.SecurityV1Client) bool {
	exist := true
	_, err := v1Client.SecurityContextConstraints().Get("aqua-scc", metav1.GetOptions{})
	if err != nil {
		exist = false
	}
	return exist
}

// CreateAquaSecurityContextConstraints create aqua-scc
func CreateAquaSecurityContextConstraints() *v1.SecurityContextConstraints {

	bool := true
	capabilities := []corev1.Capability{
		"SYS_ADMIN",
		"NET_ADMIN",
		"NET_RAW",
		"SYS_PTRACE",
		"KILL",
		"MKNOD",
		"SETGID",
		"SETUID",
		"SYS_MODULE",
		"AUDIT_CONTROL",
		"SYSLOG",
		"SYS_CHROOT",
	}

	annotations := map[string]string{
		"kubernetes.io/description":        "aqua scc provides all features of the restricted SCC\n      but allows users to run with any non-root UID and access hostPath. The user must\n      specify the UID or it must be specified on the by the manifest of the container runtime.",
		"release.openshift.io/create-only": "true",
	}

	constraints := v1.SecurityContextConstraints{
		TypeMeta: metav1.TypeMeta{Kind: "SecurityContextConstraints"},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: annotations,
			Name:        "aqua-scc",
		},
		Priority:                        nil,
		AllowPrivilegedContainer:        false,
		DefaultAddCapabilities:          nil,
		RequiredDropCapabilities:        nil,
		AllowedCapabilities:             capabilities,
		AllowHostDirVolumePlugin:        true,
		Volumes:                         nil,
		AllowedFlexVolumes:              nil,
		AllowHostNetwork:                false,
		AllowHostPorts:                  false,
		AllowHostPID:                    true,
		AllowHostIPC:                    false,
		DefaultAllowPrivilegeEscalation: nil,
		AllowPrivilegeEscalation:        &bool,
		SELinuxContext:                  v1.SELinuxContextStrategyOptions{},
		RunAsUser:                       v1.RunAsUserStrategyOptions{},
		SupplementalGroups:              v1.SupplementalGroupsStrategyOptions{},
		FSGroup: v1.FSGroupStrategyOptions{
			Type: v1.FSGroupStrategyRunAsAny},
		ReadOnlyRootFilesystem: false,
		Users:                  nil,
		Groups:                 []string{},
		SeccompProfiles:        nil,
		AllowedUnsafeSysctls:   nil,
		ForbiddenSysctls:       nil,
	}

	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}

	config, err := security1.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}
	contextConstraints, err := config.SecurityContextConstraints().Create(&constraints)
	//restClient := config.RESTClient()

	//security1.SecurityV1Client{}
	//security1.SecurityContextConstraintsInterface().Create(&constraints)
	//security1.NewForConfig()
	//client := security1.SecurityV1Client{}.RESTClient()
	//contextConstraints, err := client.SecurityContextConstraints().Create(&constraints)

	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}
	return contextConstraints
}
