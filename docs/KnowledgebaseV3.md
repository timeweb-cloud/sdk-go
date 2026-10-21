# KnowledgebaseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID базы знаний | 
**Name** | **string** | Название базы знаний | 
**Description** | **NullableString** | Описание базы знаний | 
**ProjectId** | Pointer to **int32** | ID проекта, в котором находится база знаний. Поля нет, если проект не определён | [optional] 
**DbaasId** | **int32** | ID кластера базы данных OpenSearch, который создан для базы знаний в вашем аккаунте | 
**Status** | **string** | Статус базы знаний: - &#x60;creating&#x60; — создаётся кластер базы данных, документы будут проиндексированы после перехода в &#x60;active&#x60; - &#x60;active&#x60; — база знаний работает - &#x60;blocked&#x60; — заблокирована: закончились токены или не удалось продлить подписку - &#x60;deleted&#x60; — удалена | 
**LastSync** | **NullableTime** | Дата и время последней синхронизации. &#x60;null&#x60;, если синхронизации не было | 
**TotalTokens** | **int32** | Количество токенов, доступных в текущем периоде подписки (пакет подписки и докупленные пакеты) | 
**UsedTokens** | **int32** | Количество израсходованных токенов | 
**RemainingTokens** | **int32** | Количество оставшихся токенов | 
**TokenPackageId** | **int32** | ID пакета токенов подписки базы знаний | 
**SubscriptionRenewalDate** | **NullableTime** | Дата следующего продления подписки (списания стоимости пакета токенов) | 
**DocumentsCount** | **int32** | Количество документов в базе знаний | 
**Agents** | [**[]RetrievalModeRef**](RetrievalModeRef.md) | Агенты, к которым привязана база знаний, и режим поиска для каждого | 
**CreatedAt** | **time.Time** | Дата и время создания базы знаний | 

## Methods

### NewKnowledgebaseV3

`func NewKnowledgebaseV3(id int32, name string, description NullableString, dbaasId int32, status string, lastSync NullableTime, totalTokens int32, usedTokens int32, remainingTokens int32, tokenPackageId int32, subscriptionRenewalDate NullableTime, documentsCount int32, agents []RetrievalModeRef, createdAt time.Time, ) *KnowledgebaseV3`

NewKnowledgebaseV3 instantiates a new KnowledgebaseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgebaseV3WithDefaults

`func NewKnowledgebaseV3WithDefaults() *KnowledgebaseV3`

NewKnowledgebaseV3WithDefaults instantiates a new KnowledgebaseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgebaseV3) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgebaseV3) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgebaseV3) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *KnowledgebaseV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgebaseV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgebaseV3) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KnowledgebaseV3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KnowledgebaseV3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KnowledgebaseV3) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *KnowledgebaseV3) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KnowledgebaseV3) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *KnowledgebaseV3) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KnowledgebaseV3) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KnowledgebaseV3) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *KnowledgebaseV3) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetDbaasId

`func (o *KnowledgebaseV3) GetDbaasId() int32`

GetDbaasId returns the DbaasId field if non-nil, zero value otherwise.

### GetDbaasIdOk

`func (o *KnowledgebaseV3) GetDbaasIdOk() (*int32, bool)`

GetDbaasIdOk returns a tuple with the DbaasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasId

`func (o *KnowledgebaseV3) SetDbaasId(v int32)`

SetDbaasId sets DbaasId field to given value.


### GetStatus

`func (o *KnowledgebaseV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KnowledgebaseV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KnowledgebaseV3) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSync

`func (o *KnowledgebaseV3) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *KnowledgebaseV3) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *KnowledgebaseV3) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.


### SetLastSyncNil

`func (o *KnowledgebaseV3) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *KnowledgebaseV3) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetTotalTokens

`func (o *KnowledgebaseV3) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *KnowledgebaseV3) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *KnowledgebaseV3) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### GetUsedTokens

`func (o *KnowledgebaseV3) GetUsedTokens() int32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *KnowledgebaseV3) GetUsedTokensOk() (*int32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *KnowledgebaseV3) SetUsedTokens(v int32)`

SetUsedTokens sets UsedTokens field to given value.


### GetRemainingTokens

`func (o *KnowledgebaseV3) GetRemainingTokens() int32`

GetRemainingTokens returns the RemainingTokens field if non-nil, zero value otherwise.

### GetRemainingTokensOk

`func (o *KnowledgebaseV3) GetRemainingTokensOk() (*int32, bool)`

GetRemainingTokensOk returns a tuple with the RemainingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTokens

`func (o *KnowledgebaseV3) SetRemainingTokens(v int32)`

SetRemainingTokens sets RemainingTokens field to given value.


### GetTokenPackageId

`func (o *KnowledgebaseV3) GetTokenPackageId() int32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *KnowledgebaseV3) GetTokenPackageIdOk() (*int32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *KnowledgebaseV3) SetTokenPackageId(v int32)`

SetTokenPackageId sets TokenPackageId field to given value.


### GetSubscriptionRenewalDate

`func (o *KnowledgebaseV3) GetSubscriptionRenewalDate() time.Time`

GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field if non-nil, zero value otherwise.

### GetSubscriptionRenewalDateOk

`func (o *KnowledgebaseV3) GetSubscriptionRenewalDateOk() (*time.Time, bool)`

GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRenewalDate

`func (o *KnowledgebaseV3) SetSubscriptionRenewalDate(v time.Time)`

SetSubscriptionRenewalDate sets SubscriptionRenewalDate field to given value.


### SetSubscriptionRenewalDateNil

`func (o *KnowledgebaseV3) SetSubscriptionRenewalDateNil(b bool)`

 SetSubscriptionRenewalDateNil sets the value for SubscriptionRenewalDate to be an explicit nil

### UnsetSubscriptionRenewalDate
`func (o *KnowledgebaseV3) UnsetSubscriptionRenewalDate()`

UnsetSubscriptionRenewalDate ensures that no value is present for SubscriptionRenewalDate, not even an explicit nil
### GetDocumentsCount

`func (o *KnowledgebaseV3) GetDocumentsCount() int32`

GetDocumentsCount returns the DocumentsCount field if non-nil, zero value otherwise.

### GetDocumentsCountOk

`func (o *KnowledgebaseV3) GetDocumentsCountOk() (*int32, bool)`

GetDocumentsCountOk returns a tuple with the DocumentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsCount

`func (o *KnowledgebaseV3) SetDocumentsCount(v int32)`

SetDocumentsCount sets DocumentsCount field to given value.


### GetAgents

`func (o *KnowledgebaseV3) GetAgents() []RetrievalModeRef`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *KnowledgebaseV3) GetAgentsOk() (*[]RetrievalModeRef, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *KnowledgebaseV3) SetAgents(v []RetrievalModeRef)`

SetAgents sets Agents field to given value.


### GetCreatedAt

`func (o *KnowledgebaseV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KnowledgebaseV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KnowledgebaseV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


