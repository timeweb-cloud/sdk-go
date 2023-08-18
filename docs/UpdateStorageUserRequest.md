# UpdateStorageUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretKey** | **string** | Новый пароль пользователя-администратора хранилища. | 

## Methods

### NewUpdateStorageUserRequest

`func NewUpdateStorageUserRequest(secretKey string, ) *UpdateStorageUserRequest`

NewUpdateStorageUserRequest instantiates a new UpdateStorageUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageUserRequestWithDefaults

`func NewUpdateStorageUserRequestWithDefaults() *UpdateStorageUserRequest`

NewUpdateStorageUserRequestWithDefaults instantiates a new UpdateStorageUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretKey

`func (o *UpdateStorageUserRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *UpdateStorageUserRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *UpdateStorageUserRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


