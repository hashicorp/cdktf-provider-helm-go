// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahelmtemplate


type DataHelmTemplateSetString struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.17.0/docs/data-sources/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.17.0/docs/data-sources/template#value DataHelmTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

