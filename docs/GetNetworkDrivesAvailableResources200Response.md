# GetNetworkDrivesAvailableResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**AvailableResources** | Pointer to [**[]NetworkDriveAvailableResource**](NetworkDriveAvailableResource.md) |  | [optional] 

## Methods

### NewGetNetworkDrivesAvailableResources200Response

`func NewGetNetworkDrivesAvailableResources200Response(meta Meta, ) *GetNetworkDrivesAvailableResources200Response`

NewGetNetworkDrivesAvailableResources200Response instantiates a new GetNetworkDrivesAvailableResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDrivesAvailableResources200ResponseWithDefaults

`func NewGetNetworkDrivesAvailableResources200ResponseWithDefaults() *GetNetworkDrivesAvailableResources200Response`

NewGetNetworkDrivesAvailableResources200ResponseWithDefaults instantiates a new GetNetworkDrivesAvailableResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetNetworkDrivesAvailableResources200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetNetworkDrivesAvailableResources200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetNetworkDrivesAvailableResources200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetAvailableResources

`func (o *GetNetworkDrivesAvailableResources200Response) GetAvailableResources() []NetworkDriveAvailableResource`

GetAvailableResources returns the AvailableResources field if non-nil, zero value otherwise.

### GetAvailableResourcesOk

`func (o *GetNetworkDrivesAvailableResources200Response) GetAvailableResourcesOk() (*[]NetworkDriveAvailableResource, bool)`

GetAvailableResourcesOk returns a tuple with the AvailableResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableResources

`func (o *GetNetworkDrivesAvailableResources200Response) SetAvailableResources(v []NetworkDriveAvailableResource)`

SetAvailableResources sets AvailableResources field to given value.

### HasAvailableResources

`func (o *GetNetworkDrivesAvailableResources200Response) HasAvailableResources() bool`

HasAvailableResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


