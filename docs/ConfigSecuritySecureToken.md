# ConfigSecuritySecureToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretKey** | **string** | Секретный ключ, которым подписывается ссылка. | 
**RestrictByIp** | **bool** | Привязывать подписанную ссылку к IP-адресу пользователя. | 

## Methods

### NewConfigSecuritySecureToken

`func NewConfigSecuritySecureToken(secretKey string, restrictByIp bool, ) *ConfigSecuritySecureToken`

NewConfigSecuritySecureToken instantiates a new ConfigSecuritySecureToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSecuritySecureTokenWithDefaults

`func NewConfigSecuritySecureTokenWithDefaults() *ConfigSecuritySecureToken`

NewConfigSecuritySecureTokenWithDefaults instantiates a new ConfigSecuritySecureToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretKey

`func (o *ConfigSecuritySecureToken) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ConfigSecuritySecureToken) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ConfigSecuritySecureToken) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetRestrictByIp

`func (o *ConfigSecuritySecureToken) GetRestrictByIp() bool`

GetRestrictByIp returns the RestrictByIp field if non-nil, zero value otherwise.

### GetRestrictByIpOk

`func (o *ConfigSecuritySecureToken) GetRestrictByIpOk() (*bool, bool)`

GetRestrictByIpOk returns a tuple with the RestrictByIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictByIp

`func (o *ConfigSecuritySecureToken) SetRestrictByIp(v bool)`

SetRestrictByIp sets RestrictByIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


