# AppV2DiskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Used** | Pointer to **float32** | Использованное пространство диска. | [optional] 
**Size** | Pointer to **float32** | Размер диска. | [optional] 
**DiskId** | Pointer to **float32** | ID диска. | [optional] 

## Methods

### NewAppV2DiskStatus

`func NewAppV2DiskStatus() *AppV2DiskStatus`

NewAppV2DiskStatus instantiates a new AppV2DiskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppV2DiskStatusWithDefaults

`func NewAppV2DiskStatusWithDefaults() *AppV2DiskStatus`

NewAppV2DiskStatusWithDefaults instantiates a new AppV2DiskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsed

`func (o *AppV2DiskStatus) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *AppV2DiskStatus) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *AppV2DiskStatus) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *AppV2DiskStatus) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetSize

`func (o *AppV2DiskStatus) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AppV2DiskStatus) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AppV2DiskStatus) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AppV2DiskStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDiskId

`func (o *AppV2DiskStatus) GetDiskId() float32`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *AppV2DiskStatus) GetDiskIdOk() (*float32, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *AppV2DiskStatus) SetDiskId(v float32)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *AppV2DiskStatus) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


