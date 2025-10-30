# KnowledgebaseV2

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
**DocumentsCount** | **float32** | Общее количество документов в базе знаний | 
**AgentsIds** | **[]float32** | ID агентов, связанных с базой знаний | 
**CreatedAt** | **time.Time** | Дата создания базы знаний | 

## Methods

### NewKnowledgebaseV2

`func NewKnowledgebaseV2(id float32, name string, dbaasId float32, status string, totalTokens float32, usedTokens float32, remainingTokens float32, tokenPackageId float32, subscriptionRenewalDate time.Time, documentsCount float32, agentsIds []float32, createdAt time.Time, ) *KnowledgebaseV2`

NewKnowledgebaseV2 instantiates a new KnowledgebaseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgebaseV2WithDefaults

`func NewKnowledgebaseV2WithDefaults() *KnowledgebaseV2`

NewKnowledgebaseV2WithDefaults instantiates a new KnowledgebaseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgebaseV2) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgebaseV2) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgebaseV2) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *KnowledgebaseV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgebaseV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgebaseV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KnowledgebaseV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KnowledgebaseV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KnowledgebaseV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KnowledgebaseV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KnowledgebaseV2) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KnowledgebaseV2) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDbaasId

`func (o *KnowledgebaseV2) GetDbaasId() float32`

GetDbaasId returns the DbaasId field if non-nil, zero value otherwise.

### GetDbaasIdOk

`func (o *KnowledgebaseV2) GetDbaasIdOk() (*float32, bool)`

GetDbaasIdOk returns a tuple with the DbaasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasId

`func (o *KnowledgebaseV2) SetDbaasId(v float32)`

SetDbaasId sets DbaasId field to given value.


### GetStatus

`func (o *KnowledgebaseV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KnowledgebaseV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KnowledgebaseV2) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSync

`func (o *KnowledgebaseV2) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *KnowledgebaseV2) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *KnowledgebaseV2) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *KnowledgebaseV2) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *KnowledgebaseV2) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *KnowledgebaseV2) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetTotalTokens

`func (o *KnowledgebaseV2) GetTotalTokens() float32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *KnowledgebaseV2) GetTotalTokensOk() (*float32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *KnowledgebaseV2) SetTotalTokens(v float32)`

SetTotalTokens sets TotalTokens field to given value.


### GetUsedTokens

`func (o *KnowledgebaseV2) GetUsedTokens() float32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *KnowledgebaseV2) GetUsedTokensOk() (*float32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *KnowledgebaseV2) SetUsedTokens(v float32)`

SetUsedTokens sets UsedTokens field to given value.


### GetRemainingTokens

`func (o *KnowledgebaseV2) GetRemainingTokens() float32`

GetRemainingTokens returns the RemainingTokens field if non-nil, zero value otherwise.

### GetRemainingTokensOk

`func (o *KnowledgebaseV2) GetRemainingTokensOk() (*float32, bool)`

GetRemainingTokensOk returns a tuple with the RemainingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTokens

`func (o *KnowledgebaseV2) SetRemainingTokens(v float32)`

SetRemainingTokens sets RemainingTokens field to given value.


### GetTokenPackageId

`func (o *KnowledgebaseV2) GetTokenPackageId() float32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *KnowledgebaseV2) GetTokenPackageIdOk() (*float32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *KnowledgebaseV2) SetTokenPackageId(v float32)`

SetTokenPackageId sets TokenPackageId field to given value.


### GetSubscriptionRenewalDate

`func (o *KnowledgebaseV2) GetSubscriptionRenewalDate() time.Time`

GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field if non-nil, zero value otherwise.

### GetSubscriptionRenewalDateOk

`func (o *KnowledgebaseV2) GetSubscriptionRenewalDateOk() (*time.Time, bool)`

GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRenewalDate

`func (o *KnowledgebaseV2) SetSubscriptionRenewalDate(v time.Time)`

SetSubscriptionRenewalDate sets SubscriptionRenewalDate field to given value.


### GetDocumentsCount

`func (o *KnowledgebaseV2) GetDocumentsCount() float32`

GetDocumentsCount returns the DocumentsCount field if non-nil, zero value otherwise.

### GetDocumentsCountOk

`func (o *KnowledgebaseV2) GetDocumentsCountOk() (*float32, bool)`

GetDocumentsCountOk returns a tuple with the DocumentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsCount

`func (o *KnowledgebaseV2) SetDocumentsCount(v float32)`

SetDocumentsCount sets DocumentsCount field to given value.


### GetAgentsIds

`func (o *KnowledgebaseV2) GetAgentsIds() []float32`

GetAgentsIds returns the AgentsIds field if non-nil, zero value otherwise.

### GetAgentsIdsOk

`func (o *KnowledgebaseV2) GetAgentsIdsOk() (*[]float32, bool)`

GetAgentsIdsOk returns a tuple with the AgentsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsIds

`func (o *KnowledgebaseV2) SetAgentsIds(v []float32)`

SetAgentsIds sets AgentsIds field to given value.


### GetCreatedAt

`func (o *KnowledgebaseV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KnowledgebaseV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KnowledgebaseV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


