# GetNetworkDrives200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**NetworkDrives** | [**[]NetworkDrive**](NetworkDrive.md) |  | 

## Methods

### NewGetNetworkDrives200Response

`func NewGetNetworkDrives200Response(meta Meta, networkDrives []NetworkDrive, ) *GetNetworkDrives200Response`

NewGetNetworkDrives200Response instantiates a new GetNetworkDrives200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDrives200ResponseWithDefaults

`func NewGetNetworkDrives200ResponseWithDefaults() *GetNetworkDrives200Response`

NewGetNetworkDrives200ResponseWithDefaults instantiates a new GetNetworkDrives200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetNetworkDrives200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetNetworkDrives200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetNetworkDrives200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetNetworkDrives

`func (o *GetNetworkDrives200Response) GetNetworkDrives() []NetworkDrive`

GetNetworkDrives returns the NetworkDrives field if non-nil, zero value otherwise.

### GetNetworkDrivesOk

`func (o *GetNetworkDrives200Response) GetNetworkDrivesOk() (*[]NetworkDrive, bool)`

GetNetworkDrivesOk returns a tuple with the NetworkDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDrives

`func (o *GetNetworkDrives200Response) SetNetworkDrives(v []NetworkDrive)`

SetNetworkDrives sets NetworkDrives field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


