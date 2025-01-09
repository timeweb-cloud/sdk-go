# UpdateKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Тело SSH-ключа. | [optional] 
**IsDefault** | Pointer to **bool** | Будет ли выбираться SSH-ключ по умолчанию при создании сервера. | [optional] 
**Name** | Pointer to **string** | Название SSH-ключа. | [optional] 

## Methods

### NewUpdateKeyRequest

`func NewUpdateKeyRequest() *UpdateKeyRequest`

NewUpdateKeyRequest instantiates a new UpdateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKeyRequestWithDefaults

`func NewUpdateKeyRequestWithDefaults() *UpdateKeyRequest`

NewUpdateKeyRequestWithDefaults instantiates a new UpdateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *UpdateKeyRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateKeyRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateKeyRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateKeyRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetIsDefault

`func (o *UpdateKeyRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdateKeyRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdateKeyRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UpdateKeyRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *UpdateKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


