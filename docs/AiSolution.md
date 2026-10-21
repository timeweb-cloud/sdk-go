# AiSolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID решения | 
**AccountId** | **string** | ID аккаунта, которому принадлежит решение | 
**Name** | **string** | Название решения | 
**Description** | **NullableString** | Описание решения | 
**Type** | **NullableString** | Тип решения. Сейчас всегда &#x60;openclaw&#x60; | 
**VdsId** | **int32** | ID облачного сервера, на котором развёрнуто решение. Сведения о сервере — в разделе «Облачные серверы» | 
**AiKeyId** | **string** | ID ключа AI Gateway, выпущенного для решения. Сведения о ключе — в разделе AI Gateway; перевыпустить ключ можно только через решение | 
**ModelId** | **int32** | ID модели, с которой работает решение | 
**Status** | **string** | Статус решения, выводится из состояния его сервера: &#x60;creating&#x60; — сервер разворачивается или выполняет длительную операцию; &#x60;active&#x60; — сервер работает; &#x60;stopped&#x60; — сервер выключен или перезагружается; &#x60;blocked&#x60; — сервер заблокирован или не оплачен; &#x60;error&#x60; — сервер в состоянии ошибки или не найден; &#x60;deleted&#x60; — сервер удаляется. Готовность агента к общению показывает метод получения статуса чата | 
**CreatedAt** | **time.Time** | Дата создания | 
**UpdatedAt** | **time.Time** | Дата последнего изменения | 

## Methods

### NewAiSolution

`func NewAiSolution(id int32, accountId string, name string, description NullableString, type_ NullableString, vdsId int32, aiKeyId string, modelId int32, status string, createdAt time.Time, updatedAt time.Time, ) *AiSolution`

NewAiSolution instantiates a new AiSolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionWithDefaults

`func NewAiSolutionWithDefaults() *AiSolution`

NewAiSolutionWithDefaults instantiates a new AiSolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiSolution) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiSolution) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiSolution) SetId(v int32)`

SetId sets Id field to given value.


### GetAccountId

`func (o *AiSolution) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AiSolution) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AiSolution) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *AiSolution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiSolution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiSolution) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AiSolution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiSolution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiSolution) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AiSolution) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AiSolution) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *AiSolution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiSolution) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiSolution) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AiSolution) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AiSolution) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVdsId

`func (o *AiSolution) GetVdsId() int32`

GetVdsId returns the VdsId field if non-nil, zero value otherwise.

### GetVdsIdOk

`func (o *AiSolution) GetVdsIdOk() (*int32, bool)`

GetVdsIdOk returns a tuple with the VdsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdsId

`func (o *AiSolution) SetVdsId(v int32)`

SetVdsId sets VdsId field to given value.


### GetAiKeyId

`func (o *AiSolution) GetAiKeyId() string`

GetAiKeyId returns the AiKeyId field if non-nil, zero value otherwise.

### GetAiKeyIdOk

`func (o *AiSolution) GetAiKeyIdOk() (*string, bool)`

GetAiKeyIdOk returns a tuple with the AiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeyId

`func (o *AiSolution) SetAiKeyId(v string)`

SetAiKeyId sets AiKeyId field to given value.


### GetModelId

`func (o *AiSolution) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AiSolution) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AiSolution) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetStatus

`func (o *AiSolution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiSolution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiSolution) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *AiSolution) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiSolution) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiSolution) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AiSolution) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AiSolution) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AiSolution) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


