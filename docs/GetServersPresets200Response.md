# GetServersPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**ServerPresets** | [**[]ServersPreset**](ServersPreset.md) |  | 

## Methods

### NewGetServersPresets200Response

`func NewGetServersPresets200Response(meta Meta, serverPresets []ServersPreset, ) *GetServersPresets200Response`

NewGetServersPresets200Response instantiates a new GetServersPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServersPresets200ResponseWithDefaults

`func NewGetServersPresets200ResponseWithDefaults() *GetServersPresets200Response`

NewGetServersPresets200ResponseWithDefaults instantiates a new GetServersPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetServersPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServersPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServersPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetServerPresets

`func (o *GetServersPresets200Response) GetServerPresets() []ServersPreset`

GetServerPresets returns the ServerPresets field if non-nil, zero value otherwise.

### GetServerPresetsOk

`func (o *GetServersPresets200Response) GetServerPresetsOk() (*[]ServersPreset, bool)`

GetServerPresetsOk returns a tuple with the ServerPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPresets

`func (o *GetServersPresets200Response) SetServerPresets(v []ServersPreset)`

SetServerPresets sets ServerPresets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


