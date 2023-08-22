# GetServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Servers** | [**[]Vds**](Vds.md) |  | 

## Methods

### NewGetServers200Response

`func NewGetServers200Response(meta Meta, servers []Vds, ) *GetServers200Response`

NewGetServers200Response instantiates a new GetServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServers200ResponseWithDefaults

`func NewGetServers200ResponseWithDefaults() *GetServers200Response`

NewGetServers200ResponseWithDefaults instantiates a new GetServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetServers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetServers

`func (o *GetServers200Response) GetServers() []Vds`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetServers200Response) GetServersOk() (*[]Vds, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetServers200Response) SetServers(v []Vds)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


