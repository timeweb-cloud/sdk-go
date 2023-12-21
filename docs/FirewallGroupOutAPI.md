# FirewallGroupOutAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор группы правил | 
**CreatedAt** | **time.Time** | Дата и время создания | 
**UpdatedAt** | **time.Time** | Дата и время последнего обновления | 
**Name** | **string** | Имя группы правил | 
**Description** | **string** | Описание группы правил | 
**Policy** | [**Policy**](Policy.md) |  | 

## Methods

### NewFirewallGroupOutAPI

`func NewFirewallGroupOutAPI(id string, createdAt time.Time, updatedAt time.Time, name string, description string, policy Policy, ) *FirewallGroupOutAPI`

NewFirewallGroupOutAPI instantiates a new FirewallGroupOutAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallGroupOutAPIWithDefaults

`func NewFirewallGroupOutAPIWithDefaults() *FirewallGroupOutAPI`

NewFirewallGroupOutAPIWithDefaults instantiates a new FirewallGroupOutAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallGroupOutAPI) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallGroupOutAPI) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallGroupOutAPI) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FirewallGroupOutAPI) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallGroupOutAPI) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallGroupOutAPI) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FirewallGroupOutAPI) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FirewallGroupOutAPI) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FirewallGroupOutAPI) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *FirewallGroupOutAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallGroupOutAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallGroupOutAPI) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FirewallGroupOutAPI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallGroupOutAPI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallGroupOutAPI) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPolicy

`func (o *FirewallGroupOutAPI) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *FirewallGroupOutAPI) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *FirewallGroupOutAPI) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


