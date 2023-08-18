# GetK8SNetworkDrivers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**NetworkDrivers** | **[]string** | Массив сетевых драйверов k8s | 

## Methods

### NewGetK8SNetworkDrivers200Response

`func NewGetK8SNetworkDrivers200Response(responseId string, meta Meta, networkDrivers []string, ) *GetK8SNetworkDrivers200Response`

NewGetK8SNetworkDrivers200Response instantiates a new GetK8SNetworkDrivers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8SNetworkDrivers200ResponseWithDefaults

`func NewGetK8SNetworkDrivers200ResponseWithDefaults() *GetK8SNetworkDrivers200Response`

NewGetK8SNetworkDrivers200ResponseWithDefaults instantiates a new GetK8SNetworkDrivers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetK8SNetworkDrivers200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetK8SNetworkDrivers200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetK8SNetworkDrivers200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetK8SNetworkDrivers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetK8SNetworkDrivers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetK8SNetworkDrivers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetNetworkDrivers

`func (o *GetK8SNetworkDrivers200Response) GetNetworkDrivers() []string`

GetNetworkDrivers returns the NetworkDrivers field if non-nil, zero value otherwise.

### GetNetworkDriversOk

`func (o *GetK8SNetworkDrivers200Response) GetNetworkDriversOk() (*[]string, bool)`

GetNetworkDriversOk returns a tuple with the NetworkDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDrivers

`func (o *GetK8SNetworkDrivers200Response) SetNetworkDrivers(v []string)`

SetNetworkDrivers sets NetworkDrivers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


