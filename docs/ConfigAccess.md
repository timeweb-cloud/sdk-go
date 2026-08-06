# ConfigAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedMethods** | Pointer to **[]string** | Дополнительные HTTP-методы, разрешенные при обращении к ресурсу. &#x60;null&#x60; — используется набор методов по умолчанию. | [optional] [readonly] 

## Methods

### NewConfigAccess

`func NewConfigAccess() *ConfigAccess`

NewConfigAccess instantiates a new ConfigAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigAccessWithDefaults

`func NewConfigAccessWithDefaults() *ConfigAccess`

NewConfigAccessWithDefaults instantiates a new ConfigAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedMethods

`func (o *ConfigAccess) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *ConfigAccess) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *ConfigAccess) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *ConfigAccess) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### SetAllowedMethodsNil

`func (o *ConfigAccess) SetAllowedMethodsNil(b bool)`

 SetAllowedMethodsNil sets the value for AllowedMethods to be an explicit nil

### UnsetAllowedMethods
`func (o *ConfigAccess) UnsetAllowedMethods()`

UnsetAllowedMethods ensures that no value is present for AllowedMethods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


