# ServersOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | ID операционной системы. | [optional] 
**Family** | Pointer to **string** | Семейство операционной системы. | [optional] 
**Name** | Pointer to **string** | Название операционной системы. | [optional] 
**Version** | Pointer to **string** | Версия операционной системы. | [optional] 
**VersionCodename** | Pointer to **string** | Кодовое имя версии операционной системы. | [optional] 
**Description** | Pointer to **string** | Описание операционной системы. | [optional] 
**Requirements** | Pointer to [**ServersOsRequirements**](ServersOsRequirements.md) |  | [optional] 

## Methods

### NewServersOs

`func NewServersOs() *ServersOs`

NewServersOs instantiates a new ServersOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersOsWithDefaults

`func NewServersOsWithDefaults() *ServersOs`

NewServersOsWithDefaults instantiates a new ServersOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServersOs) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServersOs) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServersOs) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ServersOs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFamily

`func (o *ServersOs) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ServersOs) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ServersOs) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *ServersOs) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetName

`func (o *ServersOs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServersOs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServersOs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServersOs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *ServersOs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServersOs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServersOs) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServersOs) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionCodename

`func (o *ServersOs) GetVersionCodename() string`

GetVersionCodename returns the VersionCodename field if non-nil, zero value otherwise.

### GetVersionCodenameOk

`func (o *ServersOs) GetVersionCodenameOk() (*string, bool)`

GetVersionCodenameOk returns a tuple with the VersionCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCodename

`func (o *ServersOs) SetVersionCodename(v string)`

SetVersionCodename sets VersionCodename field to given value.

### HasVersionCodename

`func (o *ServersOs) HasVersionCodename() bool`

HasVersionCodename returns a boolean if a field has been set.

### GetDescription

`func (o *ServersOs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServersOs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServersOs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServersOs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequirements

`func (o *ServersOs) GetRequirements() ServersOsRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ServersOs) GetRequirementsOk() (*ServersOsRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ServersOs) SetRequirements(v ServersOsRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ServersOs) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


