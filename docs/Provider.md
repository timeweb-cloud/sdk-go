# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **string** | ID провайдера. | 
**Login** | **string** | Логин пользователя. | 
**ProviderType** | [**Providers**](Providers.md) |  | 

## Methods

### NewProvider

`func NewProvider(providerId string, login string, providerType Providers, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *Provider) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Provider) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Provider) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetLogin

`func (o *Provider) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Provider) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Provider) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetProviderType

`func (o *Provider) GetProviderType() Providers`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *Provider) GetProviderTypeOk() (*Providers, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *Provider) SetProviderType(v Providers)`

SetProviderType sets ProviderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


