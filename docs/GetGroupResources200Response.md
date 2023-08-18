# GetGroupResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Resources** | [**[]FirewallGroupResourceOutAPI**](FirewallGroupResourceOutAPI.md) |  | 

## Methods

### NewGetGroupResources200Response

`func NewGetGroupResources200Response(responseId string, meta Meta, resources []FirewallGroupResourceOutAPI, ) *GetGroupResources200Response`

NewGetGroupResources200Response instantiates a new GetGroupResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupResources200ResponseWithDefaults

`func NewGetGroupResources200ResponseWithDefaults() *GetGroupResources200Response`

NewGetGroupResources200ResponseWithDefaults instantiates a new GetGroupResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetGroupResources200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetGroupResources200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetGroupResources200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetGroupResources200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetGroupResources200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetGroupResources200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetResources

`func (o *GetGroupResources200Response) GetResources() []FirewallGroupResourceOutAPI`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetGroupResources200Response) GetResourcesOk() (*[]FirewallGroupResourceOutAPI, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetGroupResources200Response) SetResources(v []FirewallGroupResourceOutAPI)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


