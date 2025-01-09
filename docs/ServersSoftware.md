# ServersSoftware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | ID ПО из маркетплейса. | [optional] 
**Name** | Pointer to **string** | Имя ПО из маркетплейса. | [optional] 
**OsIds** | Pointer to **[]float32** | Список id операционных систем, на которых доступна установка ПО. | [optional] 
**Description** | Pointer to **string** | Описание ПО из маркетплейса. | [optional] 
**Installations** | Pointer to **float32** | Количество установок ПО. | [optional] 
**Requirements** | Pointer to [**ServersSoftwareRequirements**](ServersSoftwareRequirements.md) |  | [optional] 

## Methods

### NewServersSoftware

`func NewServersSoftware() *ServersSoftware`

NewServersSoftware instantiates a new ServersSoftware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersSoftwareWithDefaults

`func NewServersSoftwareWithDefaults() *ServersSoftware`

NewServersSoftwareWithDefaults instantiates a new ServersSoftware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServersSoftware) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServersSoftware) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServersSoftware) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ServersSoftware) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServersSoftware) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServersSoftware) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServersSoftware) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServersSoftware) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsIds

`func (o *ServersSoftware) GetOsIds() []float32`

GetOsIds returns the OsIds field if non-nil, zero value otherwise.

### GetOsIdsOk

`func (o *ServersSoftware) GetOsIdsOk() (*[]float32, bool)`

GetOsIdsOk returns a tuple with the OsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsIds

`func (o *ServersSoftware) SetOsIds(v []float32)`

SetOsIds sets OsIds field to given value.

### HasOsIds

`func (o *ServersSoftware) HasOsIds() bool`

HasOsIds returns a boolean if a field has been set.

### GetDescription

`func (o *ServersSoftware) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServersSoftware) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServersSoftware) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServersSoftware) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstallations

`func (o *ServersSoftware) GetInstallations() float32`

GetInstallations returns the Installations field if non-nil, zero value otherwise.

### GetInstallationsOk

`func (o *ServersSoftware) GetInstallationsOk() (*float32, bool)`

GetInstallationsOk returns a tuple with the Installations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallations

`func (o *ServersSoftware) SetInstallations(v float32)`

SetInstallations sets Installations field to given value.

### HasInstallations

`func (o *ServersSoftware) HasInstallations() bool`

HasInstallations returns a boolean if a field has been set.

### GetRequirements

`func (o *ServersSoftware) GetRequirements() ServersSoftwareRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ServersSoftware) GetRequirementsOk() (*ServersSoftwareRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ServersSoftware) SetRequirements(v ServersSoftwareRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ServersSoftware) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


