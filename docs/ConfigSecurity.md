# ConfigSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redirect** | Pointer to **bool** | Перенаправлять запросы с HTTP на HTTPS. | [optional] 
**CertificateId** | Pointer to **NullableInt32** | ID SSL-сертификата, который используется для доменов ресурса. &#x60;null&#x60; — отвязать сертификат. | [optional] 
**SecureToken** | Pointer to [**NullableConfigSecuritySecureToken**](ConfigSecuritySecureToken.md) |  | [optional] 

## Methods

### NewConfigSecurity

`func NewConfigSecurity() *ConfigSecurity`

NewConfigSecurity instantiates a new ConfigSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSecurityWithDefaults

`func NewConfigSecurityWithDefaults() *ConfigSecurity`

NewConfigSecurityWithDefaults instantiates a new ConfigSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirect

`func (o *ConfigSecurity) GetRedirect() bool`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *ConfigSecurity) GetRedirectOk() (*bool, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *ConfigSecurity) SetRedirect(v bool)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *ConfigSecurity) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetCertificateId

`func (o *ConfigSecurity) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *ConfigSecurity) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *ConfigSecurity) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *ConfigSecurity) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *ConfigSecurity) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *ConfigSecurity) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetSecureToken

`func (o *ConfigSecurity) GetSecureToken() ConfigSecuritySecureToken`

GetSecureToken returns the SecureToken field if non-nil, zero value otherwise.

### GetSecureTokenOk

`func (o *ConfigSecurity) GetSecureTokenOk() (*ConfigSecuritySecureToken, bool)`

GetSecureTokenOk returns a tuple with the SecureToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureToken

`func (o *ConfigSecurity) SetSecureToken(v ConfigSecuritySecureToken)`

SetSecureToken sets SecureToken field to given value.

### HasSecureToken

`func (o *ConfigSecurity) HasSecureToken() bool`

HasSecureToken returns a boolean if a field has been set.

### SetSecureTokenNil

`func (o *ConfigSecurity) SetSecureTokenNil(b bool)`

 SetSecureTokenNil sets the value for SecureToken to be an explicit nil

### UnsetSecureToken
`func (o *ConfigSecurity) UnsetSecureToken()`

UnsetSecureToken ensures that no value is present for SecureToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


