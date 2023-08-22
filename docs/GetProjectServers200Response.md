# GetProjectServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]Vds**](Vds.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetProjectServers200Response

`func NewGetProjectServers200Response(servers []Vds, meta Meta, ) *GetProjectServers200Response`

NewGetProjectServers200Response instantiates a new GetProjectServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectServers200ResponseWithDefaults

`func NewGetProjectServers200ResponseWithDefaults() *GetProjectServers200Response`

NewGetProjectServers200ResponseWithDefaults instantiates a new GetProjectServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *GetProjectServers200Response) GetServers() []Vds`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetProjectServers200Response) GetServersOk() (*[]Vds, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetProjectServers200Response) SetServers(v []Vds)`

SetServers sets Servers field to given value.


### GetMeta

`func (o *GetProjectServers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProjectServers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProjectServers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


