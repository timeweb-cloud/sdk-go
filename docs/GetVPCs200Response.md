# GetVPCs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Vpcs** | [**[]Vpc**](Vpc.md) |  | 

## Methods

### NewGetVPCs200Response

`func NewGetVPCs200Response(meta Meta, vpcs []Vpc, ) *GetVPCs200Response`

NewGetVPCs200Response instantiates a new GetVPCs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVPCs200ResponseWithDefaults

`func NewGetVPCs200ResponseWithDefaults() *GetVPCs200Response`

NewGetVPCs200ResponseWithDefaults instantiates a new GetVPCs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetVPCs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetVPCs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetVPCs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetVpcs

`func (o *GetVPCs200Response) GetVpcs() []Vpc`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *GetVPCs200Response) GetVpcsOk() (*[]Vpc, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *GetVPCs200Response) SetVpcs(v []Vpc)`

SetVpcs sets Vpcs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


