# CrispIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID интеграции | 
**AccountId** | **string** | Логин аккаунта, которому принадлежит интеграция | 
**AgentId** | **int32** | ID агента | 
**WebsiteId** | **string** | Website ID сайта в Crisp | 
**WebhookSecret** | **string** | Секрет вебхука — последний сегмент адреса вебхука, который нужно добавить в Crisp: &#x60;https://api.timeweb.cloud/api/v1/crisp/{website_id}/webhook/{webhook_secret}&#x60; | 
**Status** | **string** | Статус интеграции: &#x60;pending&#x60; — ожидает подключения, &#x60;active&#x60; — работает, &#x60;error&#x60; — ошибка подключения | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**UpdatedAt** | **time.Time** | Дата и время последнего обновления | 

## Methods

### NewCrispIntegration

`func NewCrispIntegration(id string, accountId string, agentId int32, websiteId string, webhookSecret string, status string, createdAt time.Time, updatedAt time.Time, ) *CrispIntegration`

NewCrispIntegration instantiates a new CrispIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrispIntegrationWithDefaults

`func NewCrispIntegrationWithDefaults() *CrispIntegration`

NewCrispIntegrationWithDefaults instantiates a new CrispIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CrispIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrispIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrispIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *CrispIntegration) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CrispIntegration) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CrispIntegration) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAgentId

`func (o *CrispIntegration) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CrispIntegration) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CrispIntegration) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetWebsiteId

`func (o *CrispIntegration) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *CrispIntegration) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *CrispIntegration) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetWebhookSecret

`func (o *CrispIntegration) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *CrispIntegration) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *CrispIntegration) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.


### GetStatus

`func (o *CrispIntegration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CrispIntegration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CrispIntegration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *CrispIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrispIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrispIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CrispIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrispIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrispIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


