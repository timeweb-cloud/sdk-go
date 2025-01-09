# DatabaseInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра базы данных. Автоматически генерируется при создании. | 
**CreatedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных. | 
**Name** | **string** | Название базы данных. | 
**Description** | **string** | Описание базы данных | 

## Methods

### NewDatabaseInstance

`func NewDatabaseInstance(id float32, createdAt string, name string, description string, ) *DatabaseInstance`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


