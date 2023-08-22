# GetProjectDedicatedServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedServers** | [**[]DedicatedServer**](DedicatedServer.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetProjectDedicatedServers200Response

`func NewGetProjectDedicatedServers200Response(dedicatedServers []DedicatedServer, meta Meta, ) *GetProjectDedicatedServers200Response`

NewGetProjectDedicatedServers200Response instantiates a new GetProjectDedicatedServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectDedicatedServers200ResponseWithDefaults

`func NewGetProjectDedicatedServers200ResponseWithDefaults() *GetProjectDedicatedServers200Response`

NewGetProjectDedicatedServers200ResponseWithDefaults instantiates a new GetProjectDedicatedServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedServers

`func (o *GetProjectDedicatedServers200Response) GetDedicatedServers() []DedicatedServer`

GetDedicatedServers returns the DedicatedServers field if non-nil, zero value otherwise.

### GetDedicatedServersOk

`func (o *GetProjectDedicatedServers200Response) GetDedicatedServersOk() (*[]DedicatedServer, bool)`

GetDedicatedServersOk returns a tuple with the DedicatedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedServers

`func (o *GetProjectDedicatedServers200Response) SetDedicatedServers(v []DedicatedServer)`

SetDedicatedServers sets DedicatedServers field to given value.


### GetMeta

`func (o *GetProjectDedicatedServers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProjectDedicatedServers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProjectDedicatedServers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


