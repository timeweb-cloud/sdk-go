# RestorePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID снапшота. | 
**CreatedAt** | **time.Time** | Дата и время создания снапшота в формате ISO 8601. | 
**ExpiredAt** | **NullableTime** | Дата и время истечения снапшота в формате ISO 8601. | 
**Status** | **string** | Статус снапшота.  - &#x60;creating&#x60; — создаётся; - &#x60;created&#x60; — создан; - &#x60;committed&#x60; — зафиксирован; - &#x60;rolled_back&#x60; — откачен; - &#x60;error&#x60; — ошибка; - &#x60;deleted&#x60; — удалён. | 
**VdsId** | **float32** | ID облачного сервера (VDS), к которому относится снапшот. | 
**AccountId** | **string** | ID аккаунта-владельца снапшота. | 

## Methods

### NewRestorePoint

`func NewRestorePoint(id float32, createdAt time.Time, expiredAt NullableTime, status string, vdsId float32, accountId string, ) *RestorePoint`

NewRestorePoint instantiates a new RestorePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorePointWithDefaults

`func NewRestorePointWithDefaults() *RestorePoint`

NewRestorePointWithDefaults instantiates a new RestorePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestorePoint) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestorePoint) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestorePoint) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *RestorePoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RestorePoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RestorePoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiredAt

`func (o *RestorePoint) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *RestorePoint) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *RestorePoint) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *RestorePoint) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *RestorePoint) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetStatus

`func (o *RestorePoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestorePoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestorePoint) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVdsId

`func (o *RestorePoint) GetVdsId() float32`

GetVdsId returns the VdsId field if non-nil, zero value otherwise.

### GetVdsIdOk

`func (o *RestorePoint) GetVdsIdOk() (*float32, bool)`

GetVdsIdOk returns a tuple with the VdsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdsId

`func (o *RestorePoint) SetVdsId(v float32)`

SetVdsId sets VdsId field to given value.


### GetAccountId

`func (o *RestorePoint) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RestorePoint) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RestorePoint) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


