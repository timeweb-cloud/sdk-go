# CreateKnowledgebase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название базы знаний | 
**Description** | Pointer to **string** | Описание базы знаний | [optional] 
**DbaasPresetId** | **float32** | ID пресета базы данных | 
**NetworkId** | **string** | ID сети | 
**TokenPackageId** | **float32** | ID пакета токенов | 
**ProjectId** | Pointer to **float32** | ID проекта | [optional] 

## Methods

### NewCreateKnowledgebase

`func NewCreateKnowledgebase(name string, dbaasPresetId float32, networkId string, tokenPackageId float32, ) *CreateKnowledgebase`

NewCreateKnowledgebase instantiates a new CreateKnowledgebase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKnowledgebaseWithDefaults

`func NewCreateKnowledgebaseWithDefaults() *CreateKnowledgebase`

NewCreateKnowledgebaseWithDefaults instantiates a new CreateKnowledgebase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKnowledgebase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKnowledgebase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKnowledgebase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateKnowledgebase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateKnowledgebase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateKnowledgebase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateKnowledgebase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDbaasPresetId

`func (o *CreateKnowledgebase) GetDbaasPresetId() float32`

GetDbaasPresetId returns the DbaasPresetId field if non-nil, zero value otherwise.

### GetDbaasPresetIdOk

`func (o *CreateKnowledgebase) GetDbaasPresetIdOk() (*float32, bool)`

GetDbaasPresetIdOk returns a tuple with the DbaasPresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasPresetId

`func (o *CreateKnowledgebase) SetDbaasPresetId(v float32)`

SetDbaasPresetId sets DbaasPresetId field to given value.


### GetNetworkId

`func (o *CreateKnowledgebase) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreateKnowledgebase) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreateKnowledgebase) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetTokenPackageId

`func (o *CreateKnowledgebase) GetTokenPackageId() float32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *CreateKnowledgebase) GetTokenPackageIdOk() (*float32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *CreateKnowledgebase) SetTokenPackageId(v float32)`

SetTokenPackageId sets TokenPackageId field to given value.


### GetProjectId

`func (o *CreateKnowledgebase) GetProjectId() float32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateKnowledgebase) GetProjectIdOk() (*float32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateKnowledgebase) SetProjectId(v float32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateKnowledgebase) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


