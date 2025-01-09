# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID токена. | 
**CreatedAt** | **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан токен. | 
**Name** | **string** | Имя токена. | 
**ExpiredAt** | **NullableTime** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда истекает токен. | 
**IsAbleToDelete** | **bool** | Это логическое значение, которое показывает, можно ли удалять управляемые сервисы при помощи данного токена без подтверждения через Телеграм, когда это подтверждение включено. | 

## Methods

### NewApiKey

`func NewApiKey(id string, createdAt time.Time, name string, expiredAt NullableTime, isAbleToDelete bool, ) *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKey) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ApiKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *ApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpiredAt

`func (o *ApiKey) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ApiKey) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ApiKey) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *ApiKey) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *ApiKey) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetIsAbleToDelete

`func (o *ApiKey) GetIsAbleToDelete() bool`

GetIsAbleToDelete returns the IsAbleToDelete field if non-nil, zero value otherwise.

### GetIsAbleToDeleteOk

`func (o *ApiKey) GetIsAbleToDeleteOk() (*bool, bool)`

GetIsAbleToDeleteOk returns a tuple with the IsAbleToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbleToDelete

`func (o *ApiKey) SetIsAbleToDelete(v bool)`

SetIsAbleToDelete sets IsAbleToDelete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


