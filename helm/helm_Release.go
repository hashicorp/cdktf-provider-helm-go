// Prebuilt helm Provider for Terraform CDK (cdktf)
package helm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-helm-go/helm/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-helm-go/helm/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/helm/r/release helm_release}.
type Release interface {
	cdktf.TerraformResource
	Atomic() interface{}
	SetAtomic(val interface{})
	AtomicInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Chart() *string
	SetChart(val *string)
	ChartInput() *string
	CleanupOnFail() interface{}
	SetCleanupOnFail(val interface{})
	CleanupOnFailInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateNamespace() interface{}
	SetCreateNamespace(val interface{})
	CreateNamespaceInput() interface{}
	DependencyUpdate() interface{}
	SetDependencyUpdate(val interface{})
	DependencyUpdateInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Devel() interface{}
	SetDevel(val interface{})
	DevelInput() interface{}
	DisableCrdHooks() interface{}
	SetDisableCrdHooks(val interface{})
	DisableCrdHooksInput() interface{}
	DisableOpenapiValidation() interface{}
	SetDisableOpenapiValidation(val interface{})
	DisableOpenapiValidationInput() interface{}
	DisableWebhooks() interface{}
	SetDisableWebhooks(val interface{})
	DisableWebhooksInput() interface{}
	ForceUpdate() interface{}
	SetForceUpdate(val interface{})
	ForceUpdateInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Keyring() *string
	SetKeyring(val *string)
	KeyringInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Lint() interface{}
	SetLint(val interface{})
	LintInput() interface{}
	Manifest() *string
	MaxHistory() *float64
	SetMaxHistory(val *float64)
	MaxHistoryInput() *float64
	Metadata() ReleaseMetadataList
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	PassCredentials() interface{}
	SetPassCredentials(val interface{})
	PassCredentialsInput() interface{}
	Postrender() ReleasePostrenderOutputReference
	PostrenderInput() *ReleasePostrender
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RecreatePods() interface{}
	SetRecreatePods(val interface{})
	RecreatePodsInput() interface{}
	RenderSubchartNotes() interface{}
	SetRenderSubchartNotes(val interface{})
	RenderSubchartNotesInput() interface{}
	Replace() interface{}
	SetReplace(val interface{})
	ReplaceInput() interface{}
	Repository() *string
	SetRepository(val *string)
	RepositoryCaFile() *string
	SetRepositoryCaFile(val *string)
	RepositoryCaFileInput() *string
	RepositoryCertFile() *string
	SetRepositoryCertFile(val *string)
	RepositoryCertFileInput() *string
	RepositoryInput() *string
	RepositoryKeyFile() *string
	SetRepositoryKeyFile(val *string)
	RepositoryKeyFileInput() *string
	RepositoryPassword() *string
	SetRepositoryPassword(val *string)
	RepositoryPasswordInput() *string
	RepositoryUsername() *string
	SetRepositoryUsername(val *string)
	RepositoryUsernameInput() *string
	ResetValues() interface{}
	SetResetValues(val interface{})
	ResetValuesInput() interface{}
	ReuseValues() interface{}
	SetReuseValues(val interface{})
	ReuseValuesInput() interface{}
	Set() ReleaseSetList
	SetInput() interface{}
	SetSensitive() ReleaseSetSensitiveList
	SetSensitiveInput() interface{}
	SkipCrds() interface{}
	SetSkipCrds(val interface{})
	SkipCrdsInput() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	Verify() interface{}
	SetVerify(val interface{})
	VerifyInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Wait() interface{}
	SetWait(val interface{})
	WaitForJobs() interface{}
	SetWaitForJobs(val interface{})
	WaitForJobsInput() interface{}
	WaitInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutPostrender(value *ReleasePostrender)
	PutSet(value interface{})
	PutSetSensitive(value interface{})
	ResetAtomic()
	ResetCleanupOnFail()
	ResetCreateNamespace()
	ResetDependencyUpdate()
	ResetDescription()
	ResetDevel()
	ResetDisableCrdHooks()
	ResetDisableOpenapiValidation()
	ResetDisableWebhooks()
	ResetForceUpdate()
	ResetId()
	ResetKeyring()
	ResetLint()
	ResetMaxHistory()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassCredentials()
	ResetPostrender()
	ResetRecreatePods()
	ResetRenderSubchartNotes()
	ResetReplace()
	ResetRepository()
	ResetRepositoryCaFile()
	ResetRepositoryCertFile()
	ResetRepositoryKeyFile()
	ResetRepositoryPassword()
	ResetRepositoryUsername()
	ResetResetValues()
	ResetReuseValues()
	ResetSet()
	ResetSetSensitive()
	ResetSkipCrds()
	ResetTfValues()
	ResetTimeout()
	ResetVerify()
	ResetVersion()
	ResetWait()
	ResetWaitForJobs()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Release
type jsiiProxy_Release struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Release) Atomic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atomic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) AtomicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atomicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Chart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ChartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CleanupOnFail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupOnFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CleanupOnFailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupOnFailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CreateNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CreateNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DependencyUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DependencyUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Devel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"develInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableCrdHooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCrdHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableCrdHooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCrdHooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableOpenapiValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOpenapiValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableOpenapiValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOpenapiValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableWebhooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableWebhooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableWebhooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableWebhooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ForceUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ForceUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Keyring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) KeyringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Lint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) LintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Manifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) MaxHistory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) MaxHistoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Metadata() ReleaseMetadataList {
	var returns ReleaseMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) PassCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) PassCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Postrender() ReleasePostrenderOutputReference {
	var returns ReleasePostrenderOutputReference
	_jsii_.Get(
		j,
		"postrender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) PostrenderInput() *ReleasePostrender {
	var returns *ReleasePostrender
	_jsii_.Get(
		j,
		"postrenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RecreatePods() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatePods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RecreatePodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatePodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RenderSubchartNotes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderSubchartNotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RenderSubchartNotesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderSubchartNotesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Replace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCaFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCaFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ResetValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ResetValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ReuseValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ReuseValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Set() ReleaseSetList {
	var returns ReleaseSetList
	_jsii_.Get(
		j,
		"set",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SetSensitive() ReleaseSetSensitiveList {
	var returns ReleaseSetSensitiveList
	_jsii_.Get(
		j,
		"setSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SetSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SkipCrds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCrds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SkipCrdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCrdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Verify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) VerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Wait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) WaitForJobs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForJobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) WaitForJobsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForJobsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) WaitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/helm/r/release helm_release} Resource.
func NewRelease(scope constructs.Construct, id *string, config *ReleaseConfig) Release {
	_init_.Initialize()

	j := jsiiProxy_Release{}

	_jsii_.Create(
		"@cdktf/provider-helm.Release",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/helm/r/release helm_release} Resource.
func NewRelease_Override(r Release, scope constructs.Construct, id *string, config *ReleaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.Release",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Release) SetAtomic(val interface{}) {
	_jsii_.Set(
		j,
		"atomic",
		val,
	)
}

func (j *jsiiProxy_Release) SetChart(val *string) {
	_jsii_.Set(
		j,
		"chart",
		val,
	)
}

func (j *jsiiProxy_Release) SetCleanupOnFail(val interface{}) {
	_jsii_.Set(
		j,
		"cleanupOnFail",
		val,
	)
}

func (j *jsiiProxy_Release) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Release) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Release) SetCreateNamespace(val interface{}) {
	_jsii_.Set(
		j,
		"createNamespace",
		val,
	)
}

func (j *jsiiProxy_Release) SetDependencyUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"dependencyUpdate",
		val,
	)
}

func (j *jsiiProxy_Release) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Release) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Release) SetDevel(val interface{}) {
	_jsii_.Set(
		j,
		"devel",
		val,
	)
}

func (j *jsiiProxy_Release) SetDisableCrdHooks(val interface{}) {
	_jsii_.Set(
		j,
		"disableCrdHooks",
		val,
	)
}

func (j *jsiiProxy_Release) SetDisableOpenapiValidation(val interface{}) {
	_jsii_.Set(
		j,
		"disableOpenapiValidation",
		val,
	)
}

func (j *jsiiProxy_Release) SetDisableWebhooks(val interface{}) {
	_jsii_.Set(
		j,
		"disableWebhooks",
		val,
	)
}

func (j *jsiiProxy_Release) SetForceUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"forceUpdate",
		val,
	)
}

func (j *jsiiProxy_Release) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Release) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Release) SetKeyring(val *string) {
	_jsii_.Set(
		j,
		"keyring",
		val,
	)
}

func (j *jsiiProxy_Release) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Release) SetLint(val interface{}) {
	_jsii_.Set(
		j,
		"lint",
		val,
	)
}

func (j *jsiiProxy_Release) SetMaxHistory(val *float64) {
	_jsii_.Set(
		j,
		"maxHistory",
		val,
	)
}

func (j *jsiiProxy_Release) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Release) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_Release) SetPassCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"passCredentials",
		val,
	)
}

func (j *jsiiProxy_Release) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Release) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Release) SetRecreatePods(val interface{}) {
	_jsii_.Set(
		j,
		"recreatePods",
		val,
	)
}

func (j *jsiiProxy_Release) SetRenderSubchartNotes(val interface{}) {
	_jsii_.Set(
		j,
		"renderSubchartNotes",
		val,
	)
}

func (j *jsiiProxy_Release) SetReplace(val interface{}) {
	_jsii_.Set(
		j,
		"replace",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepository(val *string) {
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryCaFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryCaFile",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryCertFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryCertFile",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryKeyFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryKeyFile",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryPassword(val *string) {
	_jsii_.Set(
		j,
		"repositoryPassword",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryUsername(val *string) {
	_jsii_.Set(
		j,
		"repositoryUsername",
		val,
	)
}

func (j *jsiiProxy_Release) SetResetValues(val interface{}) {
	_jsii_.Set(
		j,
		"resetValues",
		val,
	)
}

func (j *jsiiProxy_Release) SetReuseValues(val interface{}) {
	_jsii_.Set(
		j,
		"reuseValues",
		val,
	)
}

func (j *jsiiProxy_Release) SetSkipCrds(val interface{}) {
	_jsii_.Set(
		j,
		"skipCrds",
		val,
	)
}

func (j *jsiiProxy_Release) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_Release) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (j *jsiiProxy_Release) SetVerify(val interface{}) {
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (j *jsiiProxy_Release) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_Release) SetWait(val interface{}) {
	_jsii_.Set(
		j,
		"wait",
		val,
	)
}

func (j *jsiiProxy_Release) SetWaitForJobs(val interface{}) {
	_jsii_.Set(
		j,
		"waitForJobs",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Release_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.Release",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Release_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-helm.Release",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Release) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Release) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Release) PutPostrender(value *ReleasePostrender) {
	_jsii_.InvokeVoid(
		r,
		"putPostrender",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Release) PutSet(value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"putSet",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Release) PutSetSensitive(value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"putSetSensitive",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Release) ResetAtomic() {
	_jsii_.InvokeVoid(
		r,
		"resetAtomic",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetCleanupOnFail() {
	_jsii_.InvokeVoid(
		r,
		"resetCleanupOnFail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetCreateNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetCreateNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDependencyUpdate() {
	_jsii_.InvokeVoid(
		r,
		"resetDependencyUpdate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDescription() {
	_jsii_.InvokeVoid(
		r,
		"resetDescription",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDevel() {
	_jsii_.InvokeVoid(
		r,
		"resetDevel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDisableCrdHooks() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableCrdHooks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDisableOpenapiValidation() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableOpenapiValidation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDisableWebhooks() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableWebhooks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetForceUpdate() {
	_jsii_.InvokeVoid(
		r,
		"resetForceUpdate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetKeyring() {
	_jsii_.InvokeVoid(
		r,
		"resetKeyring",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetLint() {
	_jsii_.InvokeVoid(
		r,
		"resetLint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetMaxHistory() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxHistory",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetPassCredentials() {
	_jsii_.InvokeVoid(
		r,
		"resetPassCredentials",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetPostrender() {
	_jsii_.InvokeVoid(
		r,
		"resetPostrender",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRecreatePods() {
	_jsii_.InvokeVoid(
		r,
		"resetRecreatePods",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRenderSubchartNotes() {
	_jsii_.InvokeVoid(
		r,
		"resetRenderSubchartNotes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetReplace() {
	_jsii_.InvokeVoid(
		r,
		"resetReplace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepository() {
	_jsii_.InvokeVoid(
		r,
		"resetRepository",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryCaFile() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryCaFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryCertFile() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryCertFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryKeyFile() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryKeyFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetResetValues() {
	_jsii_.InvokeVoid(
		r,
		"resetResetValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetReuseValues() {
	_jsii_.InvokeVoid(
		r,
		"resetReuseValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetSet() {
	_jsii_.InvokeVoid(
		r,
		"resetSet",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetSetSensitive() {
	_jsii_.InvokeVoid(
		r,
		"resetSetSensitive",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetSkipCrds() {
	_jsii_.InvokeVoid(
		r,
		"resetSkipCrds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetTfValues() {
	_jsii_.InvokeVoid(
		r,
		"resetTfValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetVerify() {
	_jsii_.InvokeVoid(
		r,
		"resetVerify",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetWait() {
	_jsii_.InvokeVoid(
		r,
		"resetWait",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetWaitForJobs() {
	_jsii_.InvokeVoid(
		r,
		"resetWaitForJobs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

