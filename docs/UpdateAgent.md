# UpdateAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название агента | [optional] 
**Description** | Pointer to **string** | Описание агента | [optional] 
**AccessType** | Pointer to **string** | Тип доступа к агенту | [optional] 
**Status** | Pointer to **string** | Статус агента | [optional] 
**TokenPackageId** | Pointer to **float32** | ID пакета токенов | [optional] 
**Settings** | Pointer to [**UpdateAgentSettings**](UpdateAgentSettings.md) |  | [optional] 
**ProjectId** | Pointer to **float32** | ID проекта | [optional] 

## Methods

### NewUpdateAgent

`func NewUpdateAgent() *UpdateAgent`

NewUpdateAgent instantiates a new UpdateAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentWithDefaults

`func NewUpdateAgentWithDefaults() *UpdateAgent`

NewUpdateAgentWithDefaults instantiates a new UpdateAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAgent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAgent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAgent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAgent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAgent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAgent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAgent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAgent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessType

`func (o *UpdateAgent) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UpdateAgent) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UpdateAgent) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *UpdateAgent) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAgent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAgent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAgent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAgent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenPackageId

`func (o *UpdateAgent) GetTokenPackageId() float32`

GetTokenPackageId returns the TokenPackageId field if non-nil, zero value otherwise.

### GetTokenPackageIdOk

`func (o *UpdateAgent) GetTokenPackageIdOk() (*float32, bool)`

GetTokenPackageIdOk returns a tuple with the TokenPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackageId

`func (o *UpdateAgent) SetTokenPackageId(v float32)`

SetTokenPackageId sets TokenPackageId field to given value.

### HasTokenPackageId

`func (o *UpdateAgent) HasTokenPackageId() bool`

HasTokenPackageId returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateAgent) GetSettings() UpdateAgentSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateAgent) GetSettingsOk() (*UpdateAgentSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateAgent) SetSettings(v UpdateAgentSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateAgent) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetProjectId

`func (o *UpdateAgent) GetProjectId() float32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateAgent) GetProjectIdOk() (*float32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateAgent) SetProjectId(v float32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateAgent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


