# NetworkDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID сетевого диска. | 
**Name** | **string** | Название сетевого диска. | 
**Comment** | **NullableString** | Комментарий | 
**Size** | **float32** | Размер диска в Гб | 
**ServiceList** | [**[]NetworkDriveServiceListInner**](NetworkDriveServiceListInner.md) | Список сервисов к которым подключен диск. | 
**Location** | **string** | Локация сетевого диска. | 
**Status** | **string** | Статус сетевого диска. | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) |  | 
**Type** | **string** | Тип сетевого диска. | 
**PresetId** | **int32** | ID тарифа. | 

## Methods

### NewNetworkDrive

`func NewNetworkDrive(id string, name string, comment NullableString, size float32, serviceList []NetworkDriveServiceListInner, location string, status string, availabilityZone AvailabilityZone, type_ string, presetId int32, ) *NetworkDrive`

NewNetworkDrive instantiates a new NetworkDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDriveWithDefaults

`func NewNetworkDriveWithDefaults() *NetworkDrive`

NewNetworkDriveWithDefaults instantiates a new NetworkDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDrive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDrive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDrive) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NetworkDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDrive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDrive) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *NetworkDrive) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NetworkDrive) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NetworkDrive) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *NetworkDrive) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *NetworkDrive) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetSize

`func (o *NetworkDrive) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NetworkDrive) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NetworkDrive) SetSize(v float32)`

SetSize sets Size field to given value.


### GetServiceList

`func (o *NetworkDrive) GetServiceList() []NetworkDriveServiceListInner`

GetServiceList returns the ServiceList field if non-nil, zero value otherwise.

### GetServiceListOk

`func (o *NetworkDrive) GetServiceListOk() (*[]NetworkDriveServiceListInner, bool)`

GetServiceListOk returns a tuple with the ServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceList

`func (o *NetworkDrive) SetServiceList(v []NetworkDriveServiceListInner)`

SetServiceList sets ServiceList field to given value.


### GetLocation

`func (o *NetworkDrive) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetworkDrive) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetworkDrive) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetStatus

`func (o *NetworkDrive) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDrive) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDrive) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAvailabilityZone

`func (o *NetworkDrive) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *NetworkDrive) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *NetworkDrive) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetType

`func (o *NetworkDrive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkDrive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkDrive) SetType(v string)`

SetType sets Type field to given value.


### GetPresetId

`func (o *NetworkDrive) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *NetworkDrive) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *NetworkDrive) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


