# ConfigHttpHeadersCors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | **[]string** | Домены, которым разрешены кросс-доменные запросы. Укажите &#x60;*&#x60;, чтобы разрешить все домены. | 
**Always** | **bool** | Отдавать CORS-заголовки во всех ответах, а не только при наличии заголовка &#x60;Origin&#x60; в запросе. | 

## Methods

### NewConfigHttpHeadersCors

`func NewConfigHttpHeadersCors(domains []string, always bool, ) *ConfigHttpHeadersCors`

NewConfigHttpHeadersCors instantiates a new ConfigHttpHeadersCors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigHttpHeadersCorsWithDefaults

`func NewConfigHttpHeadersCorsWithDefaults() *ConfigHttpHeadersCors`

NewConfigHttpHeadersCorsWithDefaults instantiates a new ConfigHttpHeadersCors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *ConfigHttpHeadersCors) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ConfigHttpHeadersCors) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ConfigHttpHeadersCors) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetAlways

`func (o *ConfigHttpHeadersCors) GetAlways() bool`

GetAlways returns the Always field if non-nil, zero value otherwise.

### GetAlwaysOk

`func (o *ConfigHttpHeadersCors) GetAlwaysOk() (*bool, bool)`

GetAlwaysOk returns a tuple with the Always field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlways

`func (o *ConfigHttpHeadersCors) SetAlways(v bool)`

SetAlways sets Always field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


