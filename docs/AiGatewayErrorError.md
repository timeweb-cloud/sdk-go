# AiGatewayErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Текст ошибки | 
**Type** | Pointer to **NullableString** | Тип ошибки, например &#x60;auth_error&#x60; или &#x60;invalid_request_error&#x60;. Может быть &#x60;null&#x60; или строкой &#x60;\&quot;None\&quot;&#x60; | [optional] 
**Param** | Pointer to **NullableString** | Параметр запроса, к которому относится ошибка. Может быть &#x60;null&#x60; или строкой &#x60;\&quot;None\&quot;&#x60; | [optional] 
**Code** | Pointer to **NullableString** | HTTP-код ответа строкой | [optional] 

## Methods

### NewAiGatewayErrorError

`func NewAiGatewayErrorError(message string, ) *AiGatewayErrorError`

NewAiGatewayErrorError instantiates a new AiGatewayErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiGatewayErrorErrorWithDefaults

`func NewAiGatewayErrorErrorWithDefaults() *AiGatewayErrorError`

NewAiGatewayErrorErrorWithDefaults instantiates a new AiGatewayErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiGatewayErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiGatewayErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiGatewayErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *AiGatewayErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiGatewayErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiGatewayErrorError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AiGatewayErrorError) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AiGatewayErrorError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AiGatewayErrorError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetParam

`func (o *AiGatewayErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *AiGatewayErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *AiGatewayErrorError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *AiGatewayErrorError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *AiGatewayErrorError) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *AiGatewayErrorError) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil
### GetCode

`func (o *AiGatewayErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiGatewayErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiGatewayErrorError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AiGatewayErrorError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AiGatewayErrorError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AiGatewayErrorError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


