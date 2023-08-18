# EditApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Имя, установленное для токена. | [optional] 
**IsAbleToDelete** | Pointer to **bool** | Это логическое значение, которое показывает, можно ли удалять управляемые сервисы при помощи данного токена без подтверждения через Телеграм, когда это подтверждение включено. | [optional] 

## Methods

### NewEditApiKey

`func NewEditApiKey() *EditApiKey`

NewEditApiKey instantiates a new EditApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditApiKeyWithDefaults

`func NewEditApiKeyWithDefaults() *EditApiKey`

NewEditApiKeyWithDefaults instantiates a new EditApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsAbleToDelete

`func (o *EditApiKey) GetIsAbleToDelete() bool`

GetIsAbleToDelete returns the IsAbleToDelete field if non-nil, zero value otherwise.

### GetIsAbleToDeleteOk

`func (o *EditApiKey) GetIsAbleToDeleteOk() (*bool, bool)`

GetIsAbleToDeleteOk returns a tuple with the IsAbleToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbleToDelete

`func (o *EditApiKey) SetIsAbleToDelete(v bool)`

SetIsAbleToDelete sets IsAbleToDelete field to given value.

### HasIsAbleToDelete

`func (o *EditApiKey) HasIsAbleToDelete() bool`

HasIsAbleToDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


