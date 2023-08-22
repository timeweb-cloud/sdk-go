# GetDedicatedServersPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**DedicatedServersPresets** | [**[]DedicatedServerPreset**](DedicatedServerPreset.md) |  | 

## Methods

### NewGetDedicatedServersPresets200Response

`func NewGetDedicatedServersPresets200Response(meta Meta, dedicatedServersPresets []DedicatedServerPreset, ) *GetDedicatedServersPresets200Response`

NewGetDedicatedServersPresets200Response instantiates a new GetDedicatedServersPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDedicatedServersPresets200ResponseWithDefaults

`func NewGetDedicatedServersPresets200ResponseWithDefaults() *GetDedicatedServersPresets200Response`

NewGetDedicatedServersPresets200ResponseWithDefaults instantiates a new GetDedicatedServersPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDedicatedServersPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDedicatedServersPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDedicatedServersPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDedicatedServersPresets

`func (o *GetDedicatedServersPresets200Response) GetDedicatedServersPresets() []DedicatedServerPreset`

GetDedicatedServersPresets returns the DedicatedServersPresets field if non-nil, zero value otherwise.

### GetDedicatedServersPresetsOk

`func (o *GetDedicatedServersPresets200Response) GetDedicatedServersPresetsOk() (*[]DedicatedServerPreset, bool)`

GetDedicatedServersPresetsOk returns a tuple with the DedicatedServersPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedServersPresets

`func (o *GetDedicatedServersPresets200Response) SetDedicatedServersPresets(v []DedicatedServerPreset)`

SetDedicatedServersPresets sets DedicatedServersPresets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


