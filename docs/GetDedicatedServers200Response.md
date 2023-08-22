# GetDedicatedServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**DedicatedServers** | [**[]DedicatedServer**](DedicatedServer.md) |  | 

## Methods

### NewGetDedicatedServers200Response

`func NewGetDedicatedServers200Response(meta Meta, dedicatedServers []DedicatedServer, ) *GetDedicatedServers200Response`

NewGetDedicatedServers200Response instantiates a new GetDedicatedServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDedicatedServers200ResponseWithDefaults

`func NewGetDedicatedServers200ResponseWithDefaults() *GetDedicatedServers200Response`

NewGetDedicatedServers200ResponseWithDefaults instantiates a new GetDedicatedServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDedicatedServers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDedicatedServers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDedicatedServers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDedicatedServers

`func (o *GetDedicatedServers200Response) GetDedicatedServers() []DedicatedServer`

GetDedicatedServers returns the DedicatedServers field if non-nil, zero value otherwise.

### GetDedicatedServersOk

`func (o *GetDedicatedServers200Response) GetDedicatedServersOk() (*[]DedicatedServer, bool)`

GetDedicatedServersOk returns a tuple with the DedicatedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedServers

`func (o *GetDedicatedServers200Response) SetDedicatedServers(v []DedicatedServer)`

SetDedicatedServers sets DedicatedServers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


