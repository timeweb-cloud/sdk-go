# UpdateAgentEntityLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Лимит в токенах за период. Не действует, если задан &#x60;money_limit&#x60;. &#x60;0&#x60; — запрещены все запросы. | [optional] 
**MoneyLimit** | Pointer to **NullableFloat32** | Лимит расходов за период в валюте аккаунта, до 8 знаков после запятой. Приоритетнее &#x60;limit&#x60;. &#x60;0&#x60; — запрещены все запросы. &#x60;null&#x60; — снять денежный лимит (только вместе с &#x60;limit&#x60;). | [optional] 

## Methods

### NewUpdateAgentEntityLimit

`func NewUpdateAgentEntityLimit() *UpdateAgentEntityLimit`

NewUpdateAgentEntityLimit instantiates a new UpdateAgentEntityLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentEntityLimitWithDefaults

`func NewUpdateAgentEntityLimitWithDefaults() *UpdateAgentEntityLimit`

NewUpdateAgentEntityLimitWithDefaults instantiates a new UpdateAgentEntityLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *UpdateAgentEntityLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UpdateAgentEntityLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UpdateAgentEntityLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UpdateAgentEntityLimit) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMoneyLimit

`func (o *UpdateAgentEntityLimit) GetMoneyLimit() float32`

GetMoneyLimit returns the MoneyLimit field if non-nil, zero value otherwise.

### GetMoneyLimitOk

`func (o *UpdateAgentEntityLimit) GetMoneyLimitOk() (*float32, bool)`

GetMoneyLimitOk returns a tuple with the MoneyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyLimit

`func (o *UpdateAgentEntityLimit) SetMoneyLimit(v float32)`

SetMoneyLimit sets MoneyLimit field to given value.

### HasMoneyLimit

`func (o *UpdateAgentEntityLimit) HasMoneyLimit() bool`

HasMoneyLimit returns a boolean if a field has been set.

### SetMoneyLimitNil

`func (o *UpdateAgentEntityLimit) SetMoneyLimitNil(b bool)`

 SetMoneyLimitNil sets the value for MoneyLimit to be an explicit nil

### UnsetMoneyLimit
`func (o *UpdateAgentEntityLimit) UnsetMoneyLimit()`

UnsetMoneyLimit ensures that no value is present for MoneyLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


