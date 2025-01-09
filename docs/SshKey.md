# SshKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID SSH-ключа. | 
**Name** | **string** | Название SSH-ключа. | 
**Body** | **string** | Тело SSH-ключа. | 
**CreatedAt** | **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан SSH-ключ. | 
**UsedBy** | [**[]SshKeyUsedByInner**](SshKeyUsedByInner.md) | Список серверов, которые используют SSH-ключ. | 
**IsDefault** | Pointer to **bool** | Это логическое значение, которое показывает, будет ли выбираться SSH-ключ по умолчанию при создании сервера. | [optional] 

## Methods

### NewSshKey

`func NewSshKey(id float32, name string, body string, createdAt time.Time, usedBy []SshKeyUsedByInner, ) *SshKey`

NewSshKey instantiates a new SshKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyWithDefaults

`func NewSshKeyWithDefaults() *SshKey`

NewSshKeyWithDefaults instantiates a new SshKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SshKey) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshKey) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshKey) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *SshKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKey) SetName(v string)`

SetName sets Name field to given value.


### GetBody

`func (o *SshKey) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SshKey) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SshKey) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *SshKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SshKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SshKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUsedBy

`func (o *SshKey) GetUsedBy() []SshKeyUsedByInner`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *SshKey) GetUsedByOk() (*[]SshKeyUsedByInner, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *SshKey) SetUsedBy(v []SshKeyUsedByInner)`

SetUsedBy sets UsedBy field to given value.


### GetIsDefault

`func (o *SshKey) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SshKey) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SshKey) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SshKey) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


