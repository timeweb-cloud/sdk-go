# DatabaseType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название кластера базы данных. | 
**Version** | **string** | Версия кластера базы данных. | 
**Type** | **string** | Тип кластера базы данных. Передается при создании кластера в поле &#x60;type&#x60; | 

## Methods

### NewDatabaseType

`func NewDatabaseType(name string, version string, type_ string, ) *DatabaseType`

NewDatabaseType instantiates a new DatabaseType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseTypeWithDefaults

`func NewDatabaseTypeWithDefaults() *DatabaseType`

NewDatabaseTypeWithDefaults instantiates a new DatabaseType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseType) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *DatabaseType) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseType) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseType) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *DatabaseType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseType) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


