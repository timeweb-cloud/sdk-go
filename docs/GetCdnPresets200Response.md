# GetCdnPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpResourcePresets** | [**[]Preset**](Preset.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetCdnPresets200Response

`func NewGetCdnPresets200Response(httpResourcePresets []Preset, meta Meta, ) *GetCdnPresets200Response`

NewGetCdnPresets200Response instantiates a new GetCdnPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCdnPresets200ResponseWithDefaults

`func NewGetCdnPresets200ResponseWithDefaults() *GetCdnPresets200Response`

NewGetCdnPresets200ResponseWithDefaults instantiates a new GetCdnPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpResourcePresets

`func (o *GetCdnPresets200Response) GetHttpResourcePresets() []Preset`

GetHttpResourcePresets returns the HttpResourcePresets field if non-nil, zero value otherwise.

### GetHttpResourcePresetsOk

`func (o *GetCdnPresets200Response) GetHttpResourcePresetsOk() (*[]Preset, bool)`

GetHttpResourcePresetsOk returns a tuple with the HttpResourcePresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResourcePresets

`func (o *GetCdnPresets200Response) SetHttpResourcePresets(v []Preset)`

SetHttpResourcePresets sets HttpResourcePresets field to given value.


### GetMeta

`func (o *GetCdnPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCdnPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCdnPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


