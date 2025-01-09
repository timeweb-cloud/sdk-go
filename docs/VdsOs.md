# VdsOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID операционной системы. | 
**Name** | **string** | Тип операционной системы. | 
**Version** | **NullableString** | Версия операционной системы. | 

## Methods

### NewVdsOs

`func NewVdsOs(id float32, name string, version NullableString, ) *VdsOs`

NewVdsOs instantiates a new VdsOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdsOsWithDefaults

`func NewVdsOsWithDefaults() *VdsOs`

NewVdsOsWithDefaults instantiates a new VdsOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VdsOs) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VdsOs) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VdsOs) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *VdsOs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdsOs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdsOs) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *VdsOs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VdsOs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VdsOs) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *VdsOs) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *VdsOs) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


