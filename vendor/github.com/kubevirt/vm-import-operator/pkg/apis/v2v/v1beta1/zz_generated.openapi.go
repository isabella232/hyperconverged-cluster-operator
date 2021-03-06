// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1.VirtualMachineImport":       schema_pkg_apis_v2v_v1beta1_VirtualMachineImport(ref),
		"github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1.VirtualMachineImportSpec":   schema_pkg_apis_v2v_v1beta1_VirtualMachineImportSpec(ref),
		"github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1.VirtualMachineImportStatus": schema_pkg_apis_v2v_v1beta1_VirtualMachineImportStatus(ref),
	}
}

func schema_pkg_apis_v2v_v1beta1_VirtualMachineImport(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtualMachineImport is the Schema for the virtualmachineimports API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1.VirtualMachineImportSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1.VirtualMachineImportStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1.VirtualMachineImportSpec", "github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1beta1.VirtualMachineImportStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v2v_v1beta1_VirtualMachineImportSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtualMachineImportSpec defines the desired state of VirtualMachineImport",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_v2v_v1beta1_VirtualMachineImportStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtualMachineImportStatus defines the observed state of VirtualMachineImport",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
