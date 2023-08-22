# GetServerIPs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**ServerIps** | [**[]ServerIp**](ServerIp.md) |  | 

## Methods

### NewGetServerIPs200Response

`func NewGetServerIPs200Response(meta Meta, serverIps []ServerIp, ) *GetServerIPs200Response`

NewGetServerIPs200Response instantiates a new GetServerIPs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerIPs200ResponseWithDefaults

`func NewGetServerIPs200ResponseWithDefaults() *GetServerIPs200Response`

NewGetServerIPs200ResponseWithDefaults instantiates a new GetServerIPs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetServerIPs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServerIPs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServerIPs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetServerIps

`func (o *GetServerIPs200Response) GetServerIps() []ServerIp`

GetServerIps returns the ServerIps field if non-nil, zero value otherwise.

### GetServerIpsOk

`func (o *GetServerIPs200Response) GetServerIpsOk() (*[]ServerIp, bool)`

GetServerIpsOk returns a tuple with the ServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIps

`func (o *GetServerIPs200Response) SetServerIps(v []ServerIp)`

SetServerIps sets ServerIps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


