# CreateApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя, установленное для токена. | 
**Expire** | Pointer to **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда истекает токен. | [optional] 
**IsAbleToDelete** | Pointer to **bool** | Это логическое значение, которое показывает, можно ли удалять управляемые сервисы при помощи данного токена без подтверждения через Телеграм, когда это подтверждение включено. | [optional] 
**Roles** | Pointer to **[]string** | Роли, которые могут быть назначены токену. | [optional] 
**Projects** | Pointer to **[]float32** | Список идентификаторов проектов, к которым привязан токен. Если передан null - доступ к проектам не ограничен. | [optional] 

## Methods

### NewCreateApiKey

`func NewCreateApiKey(name string, ) *CreateApiKey`

NewCreateApiKey instantiates a new CreateApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyWithDefaults

`func NewCreateApiKeyWithDefaults() *CreateApiKey`

NewCreateApiKeyWithDefaults instantiates a new CreateApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpire

`func (o *CreateApiKey) GetExpire() time.Time`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *CreateApiKey) GetExpireOk() (*time.Time, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *CreateApiKey) SetExpire(v time.Time)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *CreateApiKey) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetIsAbleToDelete

`func (o *CreateApiKey) GetIsAbleToDelete() bool`

GetIsAbleToDelete returns the IsAbleToDelete field if non-nil, zero value otherwise.

### GetIsAbleToDeleteOk

`func (o *CreateApiKey) GetIsAbleToDeleteOk() (*bool, bool)`

GetIsAbleToDeleteOk returns a tuple with the IsAbleToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbleToDelete

`func (o *CreateApiKey) SetIsAbleToDelete(v bool)`

SetIsAbleToDelete sets IsAbleToDelete field to given value.

### HasIsAbleToDelete

`func (o *CreateApiKey) HasIsAbleToDelete() bool`

HasIsAbleToDelete returns a boolean if a field has been set.

### GetRoles

`func (o *CreateApiKey) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateApiKey) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateApiKey) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateApiKey) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetProjects

`func (o *CreateApiKey) GetProjects() []float32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateApiKey) GetProjectsOk() (*[]float32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateApiKey) SetProjects(v []float32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CreateApiKey) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *CreateApiKey) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *CreateApiKey) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


