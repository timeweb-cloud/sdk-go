# AppDiskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Used** | Pointer to **float32** | Использованное пространство диска. | [optional] 
**Size** | Pointer to **float32** | Размер диска. | [optional] 
**DiskId** | Pointer to **float32** | ID диска. | [optional] 

## Methods

### NewAppDiskStatus

`func NewAppDiskStatus() *AppDiskStatus`

NewAppDiskStatus instantiates a new AppDiskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDiskStatusWithDefaults

`func NewAppDiskStatusWithDefaults() *AppDiskStatus`

NewAppDiskStatusWithDefaults instantiates a new AppDiskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsed

`func (o *AppDiskStatus) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *AppDiskStatus) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *AppDiskStatus) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *AppDiskStatus) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetSize

`func (o *AppDiskStatus) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AppDiskStatus) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AppDiskStatus) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AppDiskStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDiskId

`func (o *AppDiskStatus) GetDiskId() float32`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *AppDiskStatus) GetDiskIdOk() (*float32, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *AppDiskStatus) SetDiskId(v float32)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *AppDiskStatus) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


