# UpdateDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Пароль для подключения к базе данных. | [optional] 
**Name** | Pointer to **string** | Название базы данных. | [optional] 
**PresetId** | Pointer to **int32** | ID тарифа. | [optional] 
**ConfigParameters** | Pointer to [**ConfigParameters**](ConfigParameters.md) |  | [optional] 
**IsExternalIp** | Pointer to **bool** | Использовать или нет внешний IP. | [optional] 

## Methods

### NewUpdateDb

`func NewUpdateDb() *UpdateDb`

NewUpdateDb instantiates a new UpdateDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDbWithDefaults

`func NewUpdateDbWithDefaults() *UpdateDb`

NewUpdateDbWithDefaults instantiates a new UpdateDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateDb) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateDb) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateDb) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateDb) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetName

`func (o *UpdateDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDb) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDb) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdateDb) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateDb) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateDb) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateDb) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetConfigParameters

`func (o *UpdateDb) GetConfigParameters() ConfigParameters`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *UpdateDb) GetConfigParametersOk() (*ConfigParameters, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *UpdateDb) SetConfigParameters(v ConfigParameters)`

SetConfigParameters sets ConfigParameters field to given value.

### HasConfigParameters

`func (o *UpdateDb) HasConfigParameters() bool`

HasConfigParameters returns a boolean if a field has been set.

### GetIsExternalIp

`func (o *UpdateDb) GetIsExternalIp() bool`

GetIsExternalIp returns the IsExternalIp field if non-nil, zero value otherwise.

### GetIsExternalIpOk

`func (o *UpdateDb) GetIsExternalIpOk() (*bool, bool)`

GetIsExternalIpOk returns a tuple with the IsExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalIp

`func (o *UpdateDb) SetIsExternalIp(v bool)`

SetIsExternalIp sets IsExternalIp field to given value.

### HasIsExternalIp

`func (o *UpdateDb) HasIsExternalIp() bool`

HasIsExternalIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


