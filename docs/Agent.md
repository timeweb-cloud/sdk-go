# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Уникальный идентификатор агента | 
**Name** | **string** | Название агента | 
**Description** | **string** | Описание агента | 
**ModelId** | **float32** | ID модели | 
**ProviderId** | **float32** | ID провайдера | 
**Settings** | [**AgentSettings**](AgentSettings.md) |  | 
**Status** | **string** | Статус агента | 
**AccessType** | **string** | Тип доступа к агенту | 
**TotalTokens** | **float32** | Всего токенов выделено агенту | 
**UsedTokens** | **float32** | Использовано токенов | 
**RemainingTokens** | **float32** | Осталось токенов | 
**TokenPackageId** | **float32** | ID пакета токенов | 
**SubscriptionRenewalDate** | **time.Time** | Дата обновления подписки | 
**KnowledgeBasesIds** | **[]float32** | ID баз знаний, связанных с агентом | 
**AccessId** | **float32** | ID доступа | 
**CreatedAt** | **time.Time** | Дата создания агента | 

## Methods

### NewAgent

`func NewAgent(id float32, name string, description string, modelId float32, providerId float32, settings AgentSettings, status string, accessType string, totalTokens float32, usedTokens float32, remainingTokens float32, tokenPackageId float32, subscriptionRenewalDate time.Time, knowledgeBasesIds []float32, accessId float32, createdAt time.Time, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agent) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModelId

`func (o *Agent) GetModelId() float32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *Agent) GetModelIdOk() (*float32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *Agent) SetModelId(v float32)`

SetModelId sets ModelId field to given value.


### GetProviderId

`func (o *Agent) GetProviderId() float32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Agent) GetProviderIdOk() (*float32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Agent) SetProviderId(v float32)`

SetProviderId sets ProviderId field to given value.


### GetSettings

`func (o *Agent) GetSettings() AgentSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Agent) GetSettingsOk() (*AgentSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Agent) SetSettings(v AgentSettings)`

SetSettings sets Settings field to given value.


### GetStatus

`func (o *Agent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Agent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Agent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccessType

`func (o *Agent) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Agent) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Agent) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetTotalTokens

`func (o *Agent) GetTotalTokens() float32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *Agent) GetTotalTokensOk() (*float32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *Agent) SetTotalTokens(v float32)`

SetTotalTokens sets TotalTokens field to given value.


### GetUsedTokens

`func (o *Agent) GetUsedTokens() float32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *Agent) GetUsedTokensOk() (*float32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *Agent) SetUsedTokens(v float32)`

SetUsedTokens sets UsedTokens field to given value.


### GetRemainingTokens

`func (o *Agent) GetRemainingTokens() float32`

GetRemainingTokens returns the RemainingTokens field if non-nil, zero value otherwise.

### GetRemainingTokensOk

`func (o *Agent) GetRemainingTokensOk() (*float32, bool)`

GetRemainingTokensOk returns a tuple with the RemainingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTokens

`func (o *Agent) SetRemainingTokens(v float32)`

SetRemainingTokens sets RemainingTokens field to given value.


### GetTokenPackageId

`func (o *Agent) GetTokenPackageId() float32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *Agent) GetTokenPackageIdOk() (*float32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *Agent) SetTokenPackageId(v float32)`

SetTokenPackageId sets TokenPackageId field to given value.


### GetSubscriptionRenewalDate

`func (o *Agent) GetSubscriptionRenewalDate() time.Time`

GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field if non-nil, zero value otherwise.

### GetSubscriptionRenewalDateOk

`func (o *Agent) GetSubscriptionRenewalDateOk() (*time.Time, bool)`

GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRenewalDate

`func (o *Agent) SetSubscriptionRenewalDate(v time.Time)`

SetSubscriptionRenewalDate sets SubscriptionRenewalDate field to given value.


### GetKnowledgeBasesIds

`func (o *Agent) GetKnowledgeBasesIds() []float32`

GetKnowledgeBasesIds returns the KnowledgeBasesIds field if non-nil, zero value otherwise.

### GetKnowledgeBasesIdsOk

`func (o *Agent) GetKnowledgeBasesIdsOk() (*[]float32, bool)`

GetKnowledgeBasesIdsOk returns a tuple with the KnowledgeBasesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBasesIds

`func (o *Agent) SetKnowledgeBasesIds(v []float32)`

SetKnowledgeBasesIds sets KnowledgeBasesIds field to given value.


### GetAccessId

`func (o *Agent) GetAccessId() float32`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *Agent) GetAccessIdOk() (*float32, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *Agent) SetAccessId(v float32)`

SetAccessId sets AccessId field to given value.


### GetCreatedAt

`func (o *Agent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Agent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Agent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


