# AiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID ключа (UUID) | 
**AccountId** | **string** | Логин аккаунта | 
**Name** | **string** | Название ключа | 
**Description** | **NullableString** | Описание ключа | 
**EntityType** | **string** | Тип ключа: &#x60;aikey&#x60; — ключ, созданный пользователем (в том числе ключ AI-решения), &#x60;k8s&#x60; — ключ, выпущенный платформой для кластера Kubernetes | 
**EntityId** | **NullableString** | ID ресурса, для которого платформа выпустила ключ (например, кластера Kubernetes). &#x60;null&#x60; — у ключей, созданных пользователем | 
**ProjectId** | **NullableInt32** | ID проекта, к которому привязан ключ | 
**ExpiresAt** | **NullableTime** | Дата и время окончания действия ключа. &#x60;null&#x60; — ключ бессрочный | 
**IsReadonly** | **bool** | Ключ управляется платформой (выпущен для AI-решения или кластера Kubernetes): изменить, перевыпустить и удалить его нельзя — такие запросы вернут 403 | 
**FallbackModels** | **[]string** | Запасные модели в порядке приоритета. &#x60;[]&#x60; — не настроены, &#x60;null&#x60; — настройки временно не удалось прочитать (повторите запрос позже). Запасные модели применяются к любой модели, запрошенной этим ключом, и только в &#x60;/v1/chat/completions&#x60; и &#x60;/v1/responses&#x60;: если запрошенная модель не ответила, запрос по порядку повторяется на запасных. Запасная модель, совпадающая с запрошенной, пропускается. Резерв средств и проверка лимитов делаются по цене запрошенной модели, списание и учёт расхода — по цене ответившей: если запасная модель дороже, итоговое списание может превысить резерв и денежный лимит ключа. | 
**CreatedAt** | **time.Time** | Дата и время создания ключа | 
**UpdatedAt** | **time.Time** | Дата и время изменения ключа | 

## Methods

### NewAiKey

`func NewAiKey(id string, accountId string, name string, description NullableString, entityType string, entityId NullableString, projectId NullableInt32, expiresAt NullableTime, isReadonly bool, fallbackModels []string, createdAt time.Time, updatedAt time.Time, ) *AiKey`

NewAiKey instantiates a new AiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiKeyWithDefaults

`func NewAiKeyWithDefaults() *AiKey`

NewAiKeyWithDefaults instantiates a new AiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiKey) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *AiKey) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AiKey) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AiKey) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *AiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiKey) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AiKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiKey) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AiKey) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AiKey) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEntityType

`func (o *AiKey) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AiKey) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AiKey) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *AiKey) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiKey) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiKey) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### SetEntityIdNil

`func (o *AiKey) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *AiKey) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetProjectId

`func (o *AiKey) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AiKey) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AiKey) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *AiKey) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *AiKey) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetExpiresAt

`func (o *AiKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AiKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AiKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *AiKey) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AiKey) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetIsReadonly

`func (o *AiKey) GetIsReadonly() bool`

GetIsReadonly returns the IsReadonly field if non-nil, zero value otherwise.

### GetIsReadonlyOk

`func (o *AiKey) GetIsReadonlyOk() (*bool, bool)`

GetIsReadonlyOk returns a tuple with the IsReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadonly

`func (o *AiKey) SetIsReadonly(v bool)`

SetIsReadonly sets IsReadonly field to given value.


### GetFallbackModels

`func (o *AiKey) GetFallbackModels() []string`

GetFallbackModels returns the FallbackModels field if non-nil, zero value otherwise.

### GetFallbackModelsOk

`func (o *AiKey) GetFallbackModelsOk() (*[]string, bool)`

GetFallbackModelsOk returns a tuple with the FallbackModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackModels

`func (o *AiKey) SetFallbackModels(v []string)`

SetFallbackModels sets FallbackModels field to given value.


### SetFallbackModelsNil

`func (o *AiKey) SetFallbackModelsNil(b bool)`

 SetFallbackModelsNil sets the value for FallbackModels to be an explicit nil

### UnsetFallbackModels
`func (o *AiKey) UnsetFallbackModels()`

UnsetFallbackModels ensures that no value is present for FallbackModels, not even an explicit nil
### GetCreatedAt

`func (o *AiKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AiKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AiKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AiKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


