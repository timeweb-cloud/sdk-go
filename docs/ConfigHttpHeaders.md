# ConfigHttpHeaders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to **map[string]string** | Статические заголовки, которые CDN добавляет к запросам на источник контента. Имена заголовков могут содержать символы &#x60;A-Za-z0-9-_.*()&#x3D;/;&#x60;, значения — печатные ASCII-символы. &#x60;null&#x60; — удалить все заголовки. | [optional] 
**Cors** | Pointer to [**NullableConfigHttpHeadersCors**](ConfigHttpHeadersCors.md) |  | [optional] 

## Methods

### NewConfigHttpHeaders

`func NewConfigHttpHeaders() *ConfigHttpHeaders`

NewConfigHttpHeaders instantiates a new ConfigHttpHeaders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigHttpHeadersWithDefaults

`func NewConfigHttpHeadersWithDefaults() *ConfigHttpHeaders`

NewConfigHttpHeadersWithDefaults instantiates a new ConfigHttpHeaders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *ConfigHttpHeaders) GetRequest() map[string]string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ConfigHttpHeaders) GetRequestOk() (*map[string]string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ConfigHttpHeaders) SetRequest(v map[string]string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ConfigHttpHeaders) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *ConfigHttpHeaders) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *ConfigHttpHeaders) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetCors

`func (o *ConfigHttpHeaders) GetCors() ConfigHttpHeadersCors`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *ConfigHttpHeaders) GetCorsOk() (*ConfigHttpHeadersCors, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *ConfigHttpHeaders) SetCors(v ConfigHttpHeadersCors)`

SetCors sets Cors field to given value.

### HasCors

`func (o *ConfigHttpHeaders) HasCors() bool`

HasCors returns a boolean if a field has been set.

### SetCorsNil

`func (o *ConfigHttpHeaders) SetCorsNil(b bool)`

 SetCorsNil sets the value for Cors to be an explicit nil

### UnsetCors
`func (o *ConfigHttpHeaders) UnsetCors()`

UnsetCors ensures that no value is present for Cors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


