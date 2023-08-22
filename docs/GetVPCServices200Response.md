# GetVPCServices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Services** | [**[]VpcService**](VpcService.md) |  | 

## Methods

### NewGetVPCServices200Response

`func NewGetVPCServices200Response(meta Meta, services []VpcService, ) *GetVPCServices200Response`

NewGetVPCServices200Response instantiates a new GetVPCServices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVPCServices200ResponseWithDefaults

`func NewGetVPCServices200ResponseWithDefaults() *GetVPCServices200Response`

NewGetVPCServices200ResponseWithDefaults instantiates a new GetVPCServices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetVPCServices200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetVPCServices200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetVPCServices200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetServices

`func (o *GetVPCServices200Response) GetServices() []VpcService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetVPCServices200Response) GetServicesOk() (*[]VpcService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetVPCServices200Response) SetServices(v []VpcService)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


