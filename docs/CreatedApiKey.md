# CreatedApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | Созданный токен, будет показан только один раз, его необходимо сохранить. | 
**Id** | **string** | ID токена. | 
**CreatedAt** | **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан токен. | 
**Name** | **string** | Имя токена. | 
**ExpiredAt** | **NullableTime** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда истекает токен. | 

## Methods

### NewCreatedApiKey

`func NewCreatedApiKey(token string, id string, createdAt time.Time, name string, expiredAt NullableTime, ) *CreatedApiKey`

NewCreatedApiKey instantiates a new CreatedApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedApiKeyWithDefaults

`func NewCreatedApiKeyWithDefaults() *CreatedApiKey`

NewCreatedApiKeyWithDefaults instantiates a new CreatedApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CreatedApiKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreatedApiKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreatedApiKey) SetToken(v string)`

SetToken sets Token field to given value.


### GetId

`func (o *CreatedApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedApiKey) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreatedApiKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreatedApiKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreatedApiKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *CreatedApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpiredAt

`func (o *CreatedApiKey) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CreatedApiKey) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CreatedApiKey) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *CreatedApiKey) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *CreatedApiKey) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


