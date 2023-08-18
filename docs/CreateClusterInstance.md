# CreateClusterInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название базы данных | [optional] 
**Description** | Pointer to **string** | Описание базы данных | [optional] 

## Methods

### NewCreateClusterInstance

`func NewCreateClusterInstance() *CreateClusterInstance`

NewCreateClusterInstance instantiates a new CreateClusterInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterInstanceWithDefaults

`func NewCreateClusterInstanceWithDefaults() *CreateClusterInstance`

NewCreateClusterInstanceWithDefaults instantiates a new CreateClusterInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateClusterInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateClusterInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateClusterInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateClusterInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateClusterInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


