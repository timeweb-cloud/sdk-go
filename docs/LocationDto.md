# LocationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**Location**](Location.md) |  | 
**LocationCode** | **string** | Код локации в формате &#x60;ISO 3166&#x60; | 
**AvailabilityZones** | **[]string** | Список зон, доступных в данной локации | 

## Methods

### NewLocationDto

`func NewLocationDto(location Location, locationCode string, availabilityZones []string, ) *LocationDto`

NewLocationDto instantiates a new LocationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationDtoWithDefaults

`func NewLocationDtoWithDefaults() *LocationDto`

NewLocationDtoWithDefaults instantiates a new LocationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *LocationDto) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationDto) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationDto) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetLocationCode

`func (o *LocationDto) GetLocationCode() string`

GetLocationCode returns the LocationCode field if non-nil, zero value otherwise.

### GetLocationCodeOk

`func (o *LocationDto) GetLocationCodeOk() (*string, bool)`

GetLocationCodeOk returns a tuple with the LocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCode

`func (o *LocationDto) SetLocationCode(v string)`

SetLocationCode sets LocationCode field to given value.


### GetAvailabilityZones

`func (o *LocationDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *LocationDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *LocationDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


