# GetLocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Locations** | [**[]LocationDto**](LocationDto.md) |  | 

## Methods

### NewGetLocations200Response

`func NewGetLocations200Response(meta Meta, locations []LocationDto, ) *GetLocations200Response`

NewGetLocations200Response instantiates a new GetLocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLocations200ResponseWithDefaults

`func NewGetLocations200ResponseWithDefaults() *GetLocations200Response`

NewGetLocations200ResponseWithDefaults instantiates a new GetLocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetLocations200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetLocations200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetLocations200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLocations

`func (o *GetLocations200Response) GetLocations() []LocationDto`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GetLocations200Response) GetLocationsOk() (*[]LocationDto, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GetLocations200Response) SetLocations(v []LocationDto)`

SetLocations sets Locations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


