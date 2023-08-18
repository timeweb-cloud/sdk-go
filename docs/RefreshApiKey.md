# RefreshApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expire** | Pointer to **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда истекает токен. | [optional] 

## Methods

### NewRefreshApiKey

`func NewRefreshApiKey() *RefreshApiKey`

NewRefreshApiKey instantiates a new RefreshApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshApiKeyWithDefaults

`func NewRefreshApiKeyWithDefaults() *RefreshApiKey`

NewRefreshApiKeyWithDefaults instantiates a new RefreshApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpire

`func (o *RefreshApiKey) GetExpire() time.Time`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *RefreshApiKey) GetExpireOk() (*time.Time, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *RefreshApiKey) SetExpire(v time.Time)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *RefreshApiKey) HasExpire() bool`

HasExpire returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


