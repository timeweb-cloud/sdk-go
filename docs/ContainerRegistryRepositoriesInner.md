# ContainerRegistryRepositoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название репозитория | 
**Tags** | [**Tags**](Tags.md) |  | 

## Methods

### NewContainerRegistryRepositoriesInner

`func NewContainerRegistryRepositoriesInner(name string, tags Tags, ) *ContainerRegistryRepositoriesInner`

NewContainerRegistryRepositoriesInner instantiates a new ContainerRegistryRepositoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryRepositoriesInnerWithDefaults

`func NewContainerRegistryRepositoriesInnerWithDefaults() *ContainerRegistryRepositoriesInner`

NewContainerRegistryRepositoriesInnerWithDefaults instantiates a new ContainerRegistryRepositoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerRegistryRepositoriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerRegistryRepositoriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerRegistryRepositoriesInner) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *ContainerRegistryRepositoriesInner) GetTags() Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContainerRegistryRepositoriesInner) GetTagsOk() (*Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContainerRegistryRepositoriesInner) SetTags(v Tags)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


