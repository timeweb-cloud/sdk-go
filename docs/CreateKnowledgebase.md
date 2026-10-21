# CreateKnowledgebase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название базы знаний | 
**Description** | Pointer to **string** | Описание базы знаний | [optional] 
**DbaasPresetId** | **int32** | ID тарифа кластера OpenSearch, который будет создан для базы знаний. Список тарифов баз данных — &#x60;GET /api/v2/presets/dbs&#x60; | 
**NetworkId** | **string** | ID приватной сети (VPC), в которой будет создан кластер | 
**TokenPackageId** | **int32** | ID пакета токенов подписки. Подходят только базовые пакеты (&#x60;package_type: base&#x60;) из &#x60;GET /api/v1/cloud-ai/token-packages/knowledge-bases&#x60; | 
**ProjectId** | Pointer to **int32** | ID проекта, в который нужно добавить базу знаний | [optional] 

## Methods

### NewCreateKnowledgebase

`func NewCreateKnowledgebase(name string, dbaasPresetId int32, networkId string, tokenPackageId int32, ) *CreateKnowledgebase`

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

`func (o *CreateKnowledgebase) GetDbaasPresetId() int32`

GetDbaasPresetId returns the DbaasPresetId field if non-nil, zero value otherwise.

### GetDbaasPresetIdOk

`func (o *CreateKnowledgebase) GetDbaasPresetIdOk() (*int32, bool)`

GetDbaasPresetIdOk returns a tuple with the DbaasPresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasPresetId

`func (o *CreateKnowledgebase) SetDbaasPresetId(v int32)`

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

`func (o *CreateKnowledgebase) GetTokenPackageId() int32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *CreateKnowledgebase) GetTokenPackageIdOk() (*int32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *CreateKnowledgebase) SetTokenPackageId(v int32)`

SetTokenPackageId sets TokenPackageId field to given value.


### GetProjectId

`func (o *CreateKnowledgebase) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateKnowledgebase) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateKnowledgebase) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateKnowledgebase) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


