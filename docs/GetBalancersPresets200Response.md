# GetBalancersPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**BalancersPresets** | [**[]PresetsBalancer**](PresetsBalancer.md) |  | 

## Methods

### NewGetBalancersPresets200Response

`func NewGetBalancersPresets200Response(meta Meta, balancersPresets []PresetsBalancer, ) *GetBalancersPresets200Response`

NewGetBalancersPresets200Response instantiates a new GetBalancersPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalancersPresets200ResponseWithDefaults

`func NewGetBalancersPresets200ResponseWithDefaults() *GetBalancersPresets200Response`

NewGetBalancersPresets200ResponseWithDefaults instantiates a new GetBalancersPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBalancersPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBalancersPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBalancersPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetBalancersPresets

`func (o *GetBalancersPresets200Response) GetBalancersPresets() []PresetsBalancer`

GetBalancersPresets returns the BalancersPresets field if non-nil, zero value otherwise.

### GetBalancersPresetsOk

`func (o *GetBalancersPresets200Response) GetBalancersPresetsOk() (*[]PresetsBalancer, bool)`

GetBalancersPresetsOk returns a tuple with the BalancersPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancersPresets

`func (o *GetBalancersPresets200Response) SetBalancersPresets(v []PresetsBalancer)`

SetBalancersPresets sets BalancersPresets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


