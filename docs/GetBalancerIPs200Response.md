# GetBalancerIPs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Ips** | **[]string** |  | 

## Methods

### NewGetBalancerIPs200Response

`func NewGetBalancerIPs200Response(meta Meta, ips []string, ) *GetBalancerIPs200Response`

NewGetBalancerIPs200Response instantiates a new GetBalancerIPs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalancerIPs200ResponseWithDefaults

`func NewGetBalancerIPs200ResponseWithDefaults() *GetBalancerIPs200Response`

NewGetBalancerIPs200ResponseWithDefaults instantiates a new GetBalancerIPs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBalancerIPs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBalancerIPs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBalancerIPs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetIps

`func (o *GetBalancerIPs200Response) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *GetBalancerIPs200Response) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *GetBalancerIPs200Response) SetIps(v []string)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


