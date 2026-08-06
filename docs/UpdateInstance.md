# UpdateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название инстанса базы данных | [optional] 
**Description** | Pointer to **string** | Описание инстанса базы данных | [optional] 
**OwnerId** | Pointer to **float32** | ID пользователя базы данных, который станет владельцем инстанса. Доступно только в кластерах PostgreSQL | [optional] 
**Extensions** | Pointer to [**DatabaseExtensions**](DatabaseExtensions.md) |  | [optional] 
**ConfigParameters** | Pointer to [**UpdateKafkaConfigParameters**](UpdateKafkaConfigParameters.md) |  | [optional] 

## Methods

### NewUpdateInstance

`func NewUpdateInstance() *UpdateInstance`

NewUpdateInstance instantiates a new UpdateInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceWithDefaults

`func NewUpdateInstanceWithDefaults() *UpdateInstance`

NewUpdateInstanceWithDefaults instantiates a new UpdateInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerId

`func (o *UpdateInstance) GetOwnerId() float32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpdateInstance) GetOwnerIdOk() (*float32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpdateInstance) SetOwnerId(v float32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UpdateInstance) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetExtensions

`func (o *UpdateInstance) GetExtensions() DatabaseExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *UpdateInstance) GetExtensionsOk() (*DatabaseExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *UpdateInstance) SetExtensions(v DatabaseExtensions)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *UpdateInstance) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetConfigParameters

`func (o *UpdateInstance) GetConfigParameters() UpdateKafkaConfigParameters`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *UpdateInstance) GetConfigParametersOk() (*UpdateKafkaConfigParameters, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *UpdateInstance) SetConfigParameters(v UpdateKafkaConfigParameters)`

SetConfigParameters sets ConfigParameters field to given value.

### HasConfigParameters

`func (o *UpdateInstance) HasConfigParameters() bool`

HasConfigParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


