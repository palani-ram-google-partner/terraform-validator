// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "fmt"

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const PrivatecaCaPoolIAMAssetType string = "privateca.googleapis.com/CaPool"

func resourceConverterPrivatecaCaPoolIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         PrivatecaCaPoolIAMAssetType,
		Convert:           GetPrivatecaCaPoolIamPolicyCaiObject,
		MergeCreateUpdate: MergePrivatecaCaPoolIamPolicy,
	}
}

func resourceConverterPrivatecaCaPoolIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         PrivatecaCaPoolIAMAssetType,
		Convert:           GetPrivatecaCaPoolIamBindingCaiObject,
		FetchFullResource: FetchPrivatecaCaPoolIamPolicy,
		MergeCreateUpdate: MergePrivatecaCaPoolIamBinding,
		MergeDelete:       MergePrivatecaCaPoolIamBindingDelete,
	}
}

func resourceConverterPrivatecaCaPoolIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         PrivatecaCaPoolIAMAssetType,
		Convert:           GetPrivatecaCaPoolIamMemberCaiObject,
		FetchFullResource: FetchPrivatecaCaPoolIamPolicy,
		MergeCreateUpdate: MergePrivatecaCaPoolIamMember,
		MergeDelete:       MergePrivatecaCaPoolIamMemberDelete,
	}
}

func GetPrivatecaCaPoolIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newPrivatecaCaPoolIamAsset(d, config, expandIamPolicyBindings)
}

func GetPrivatecaCaPoolIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newPrivatecaCaPoolIamAsset(d, config, expandIamRoleBindings)
}

func GetPrivatecaCaPoolIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newPrivatecaCaPoolIamAsset(d, config, expandIamMemberBindings)
}

func MergePrivatecaCaPoolIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergePrivatecaCaPoolIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergePrivatecaCaPoolIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergePrivatecaCaPoolIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergePrivatecaCaPoolIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newPrivatecaCaPoolIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//privateca.googleapis.com/{{capool}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: PrivatecaCaPoolIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchPrivatecaCaPoolIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{capool}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		PrivatecaCaPoolIamUpdaterProducer,
		d,
		config,
		"//privateca.googleapis.com/{{capool}}",
		PrivatecaCaPoolIAMAssetType,
	)
}
