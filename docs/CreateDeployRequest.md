# CreateDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | Pointer to **string** | Хэш коммита. | [optional] 

## Methods

### NewCreateDeployRequest

`func NewCreateDeployRequest() *CreateDeployRequest`

NewCreateDeployRequest instantiates a new CreateDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeployRequestWithDefaults

`func NewCreateDeployRequestWithDefaults() *CreateDeployRequest`

NewCreateDeployRequestWithDefaults instantiates a new CreateDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *CreateDeployRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CreateDeployRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CreateDeployRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *CreateDeployRequest) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


