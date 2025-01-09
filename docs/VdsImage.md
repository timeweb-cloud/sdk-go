# VdsImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID образа сервера. | 
**Name** | **string** | Название образа сервера. | 
**IsCustom** | **bool** | Является ли образ кастомным. | 

## Methods

### NewVdsImage

`func NewVdsImage(id string, name string, isCustom bool, ) *VdsImage`

NewVdsImage instantiates a new VdsImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdsImageWithDefaults

`func NewVdsImageWithDefaults() *VdsImage`

NewVdsImageWithDefaults instantiates a new VdsImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VdsImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VdsImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VdsImage) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VdsImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdsImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdsImage) SetName(v string)`

SetName sets Name field to given value.


### GetIsCustom

`func (o *VdsImage) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *VdsImage) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *VdsImage) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


