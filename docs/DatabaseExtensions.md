# DatabaseExtensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionList** | [**[]DbExtension**](DbExtension.md) | Список расширений базы данных. | 
**IsGlobal** | **bool** | Установить расширения во всех базах данных кластера (&#x60;true&#x60;) или только в этой базе данных (&#x60;false&#x60;). | 

## Methods

### NewDatabaseExtensions

`func NewDatabaseExtensions(extensionList []DbExtension, isGlobal bool, ) *DatabaseExtensions`

NewDatabaseExtensions instantiates a new DatabaseExtensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseExtensionsWithDefaults

`func NewDatabaseExtensionsWithDefaults() *DatabaseExtensions`

NewDatabaseExtensionsWithDefaults instantiates a new DatabaseExtensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionList

`func (o *DatabaseExtensions) GetExtensionList() []DbExtension`

GetExtensionList returns the ExtensionList field if non-nil, zero value otherwise.

### GetExtensionListOk

`func (o *DatabaseExtensions) GetExtensionListOk() (*[]DbExtension, bool)`

GetExtensionListOk returns a tuple with the ExtensionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionList

`func (o *DatabaseExtensions) SetExtensionList(v []DbExtension)`

SetExtensionList sets ExtensionList field to given value.


### GetIsGlobal

`func (o *DatabaseExtensions) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *DatabaseExtensions) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *DatabaseExtensions) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


