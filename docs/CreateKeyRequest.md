# CreateKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | Тело SSH-ключа. | 
**IsDefault** | **bool** | Будет ли выбираться SSH-ключ по умолчанию при создании сервера.   | 
**Name** | **string** | Название SSH-ключа. | 

## Methods

### NewCreateKeyRequest

`func NewCreateKeyRequest(body string, isDefault bool, name string, ) *CreateKeyRequest`

NewCreateKeyRequest instantiates a new CreateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyRequestWithDefaults

`func NewCreateKeyRequestWithDefaults() *CreateKeyRequest`

NewCreateKeyRequestWithDefaults instantiates a new CreateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CreateKeyRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateKeyRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateKeyRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetIsDefault

`func (o *CreateKeyRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateKeyRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateKeyRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetName

`func (o *CreateKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKeyRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


