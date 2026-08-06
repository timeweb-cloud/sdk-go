# DatabaseInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра базы данных. Автоматически генерируется при создании. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных. | 
**Name** | **string** | Название базы данных. | 
**Description** | **string** | Описание базы данных | 
**Extensions** | [**DatabaseExtensions**](DatabaseExtensions.md) |  | 
**OwnerId** | **NullableFloat32** | ID администратора базы данных, который является владельцем этой базы данных. &#x60;null&#x60;, если владелец не задан. | 
**ConfigParameters** | [**KafkaConfigParameters**](KafkaConfigParameters.md) |  | 

## Methods

### NewDatabaseInstance

`func NewDatabaseInstance(id float32, createdAt string, name string, description string, extensions DatabaseExtensions, ownerId NullableFloat32, configParameters KafkaConfigParameters, ) *DatabaseInstance`

NewDatabaseInstance instantiates a new DatabaseInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseInstanceWithDefaults

`func NewDatabaseInstanceWithDefaults() *DatabaseInstance`

NewDatabaseInstanceWithDefaults instantiates a new DatabaseInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseInstance) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseInstance) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseInstance) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DatabaseInstance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabaseInstance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabaseInstance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *DatabaseInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseInstance) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseInstance) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExtensions

`func (o *DatabaseInstance) GetExtensions() DatabaseExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *DatabaseInstance) GetExtensionsOk() (*DatabaseExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *DatabaseInstance) SetExtensions(v DatabaseExtensions)`

SetExtensions sets Extensions field to given value.


### GetOwnerId

`func (o *DatabaseInstance) GetOwnerId() float32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DatabaseInstance) GetOwnerIdOk() (*float32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DatabaseInstance) SetOwnerId(v float32)`

SetOwnerId sets OwnerId field to given value.


### SetOwnerIdNil

`func (o *DatabaseInstance) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *DatabaseInstance) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetConfigParameters

`func (o *DatabaseInstance) GetConfigParameters() KafkaConfigParameters`

GetConfigParameters returns the ConfigParameters field if non-nil, zero value otherwise.

### GetConfigParametersOk

`func (o *DatabaseInstance) GetConfigParametersOk() (*KafkaConfigParameters, bool)`

GetConfigParametersOk returns a tuple with the ConfigParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParameters

`func (o *DatabaseInstance) SetConfigParameters(v KafkaConfigParameters)`

SetConfigParameters sets ConfigParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


