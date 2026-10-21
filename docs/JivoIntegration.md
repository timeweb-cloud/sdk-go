# JivoIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID интеграции | 
**AccountId** | **string** | Логин аккаунта, которому принадлежит интеграция | 
**AgentId** | **int32** | ID агента | 
**AccessId** | **string** | Access ID агента. Это токен бота для Jivo и последний сегмент адреса, на который Jivo отправляет события | 
**Status** | **string** | Статус интеграции: &#x60;pending&#x60; — ожидает подключения, &#x60;active&#x60; — работает, &#x60;error&#x60; — ошибка подключения | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**UpdatedAt** | **time.Time** | Дата и время последнего обновления | 

## Methods

### NewJivoIntegration

`func NewJivoIntegration(id string, accountId string, agentId int32, accessId string, status string, createdAt time.Time, updatedAt time.Time, ) *JivoIntegration`

NewJivoIntegration instantiates a new JivoIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJivoIntegrationWithDefaults

`func NewJivoIntegrationWithDefaults() *JivoIntegration`

NewJivoIntegrationWithDefaults instantiates a new JivoIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JivoIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JivoIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JivoIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *JivoIntegration) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *JivoIntegration) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *JivoIntegration) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAgentId

`func (o *JivoIntegration) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *JivoIntegration) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *JivoIntegration) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetAccessId

`func (o *JivoIntegration) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *JivoIntegration) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *JivoIntegration) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetStatus

`func (o *JivoIntegration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JivoIntegration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JivoIntegration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *JivoIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JivoIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JivoIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *JivoIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *JivoIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *JivoIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


