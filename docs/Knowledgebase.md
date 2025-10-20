# Knowledgebase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Уникальный идентификатор базы знаний | 
**Name** | **string** | Название базы знаний | 
**Description** | Pointer to **NullableString** | Описание базы знаний | [optional] 
**DbaasId** | **float32** | ID базы данных opensearch | 
**Status** | **string** | Статус базы знаний | 
**LastSync** | Pointer to **NullableTime** | Дата последней синхронизации | [optional] 
**TotalTokens** | **float32** | Всего токенов выделено | 
**UsedTokens** | **float32** | Использовано токенов | 
**RemainingTokens** | **float32** | Осталось токенов | 
**TokenPackageId** | **float32** | ID пакета токенов | 
**SubscriptionRenewalDate** | **time.Time** | Дата обновления подписки | 
**Documents** | [**[]Document**](Document.md) | Документы в базе знаний | 
**AgentsIds** | **[]float32** | ID агентов, связанных с базой знаний | 
**CreatedAt** | **time.Time** | Дата создания базы знаний | 

## Methods

### NewKnowledgebase

`func NewKnowledgebase(id float32, name string, dbaasId float32, status string, totalTokens float32, usedTokens float32, remainingTokens float32, tokenPackageId float32, subscriptionRenewalDate time.Time, documents []Document, agentsIds []float32, createdAt time.Time, ) *Knowledgebase`

NewKnowledgebase instantiates a new Knowledgebase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgebaseWithDefaults

`func NewKnowledgebaseWithDefaults() *Knowledgebase`

NewKnowledgebaseWithDefaults instantiates a new Knowledgebase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Knowledgebase) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Knowledgebase) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Knowledgebase) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Knowledgebase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Knowledgebase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Knowledgebase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Knowledgebase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Knowledgebase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Knowledgebase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Knowledgebase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Knowledgebase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Knowledgebase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDbaasId

`func (o *Knowledgebase) GetDbaasId() float32`

GetDbaasId returns the DbaasId field if non-nil, zero value otherwise.

### GetDbaasIdOk

`func (o *Knowledgebase) GetDbaasIdOk() (*float32, bool)`

GetDbaasIdOk returns a tuple with the DbaasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasId

`func (o *Knowledgebase) SetDbaasId(v float32)`

SetDbaasId sets DbaasId field to given value.


### GetStatus

`func (o *Knowledgebase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Knowledgebase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Knowledgebase) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSync

`func (o *Knowledgebase) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *Knowledgebase) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *Knowledgebase) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *Knowledgebase) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *Knowledgebase) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *Knowledgebase) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetTotalTokens

`func (o *Knowledgebase) GetTotalTokens() float32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *Knowledgebase) GetTotalTokensOk() (*float32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *Knowledgebase) SetTotalTokens(v float32)`

SetTotalTokens sets TotalTokens field to given value.


### GetUsedTokens

`func (o *Knowledgebase) GetUsedTokens() float32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *Knowledgebase) GetUsedTokensOk() (*float32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *Knowledgebase) SetUsedTokens(v float32)`

SetUsedTokens sets UsedTokens field to given value.


### GetRemainingTokens

`func (o *Knowledgebase) GetRemainingTokens() float32`

GetRemainingTokens returns the RemainingTokens field if non-nil, zero value otherwise.

### GetRemainingTokensOk

`func (o *Knowledgebase) GetRemainingTokensOk() (*float32, bool)`

GetRemainingTokensOk returns a tuple with the RemainingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTokens

`func (o *Knowledgebase) SetRemainingTokens(v float32)`

SetRemainingTokens sets RemainingTokens field to given value.


### GetTokenPackageId

`func (o *Knowledgebase) GetTokenPackageId() float32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *Knowledgebase) GetTokenPackageIdOk() (*float32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *Knowledgebase) SetTokenPackageId(v float32)`

SetTokenPackageId sets TokenPackageId field to given value.


### GetSubscriptionRenewalDate

`func (o *Knowledgebase) GetSubscriptionRenewalDate() time.Time`

GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field if non-nil, zero value otherwise.

### GetSubscriptionRenewalDateOk

`func (o *Knowledgebase) GetSubscriptionRenewalDateOk() (*time.Time, bool)`

GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRenewalDate

`func (o *Knowledgebase) SetSubscriptionRenewalDate(v time.Time)`

SetSubscriptionRenewalDate sets SubscriptionRenewalDate field to given value.


### GetDocuments

`func (o *Knowledgebase) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Knowledgebase) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Knowledgebase) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.


### GetAgentsIds

`func (o *Knowledgebase) GetAgentsIds() []float32`

GetAgentsIds returns the AgentsIds field if non-nil, zero value otherwise.

### GetAgentsIdsOk

`func (o *Knowledgebase) GetAgentsIdsOk() (*[]float32, bool)`

GetAgentsIdsOk returns a tuple with the AgentsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsIds

`func (o *Knowledgebase) SetAgentsIds(v []float32)`

SetAgentsIds sets AgentsIds field to given value.


### GetCreatedAt

`func (o *Knowledgebase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Knowledgebase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Knowledgebase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


