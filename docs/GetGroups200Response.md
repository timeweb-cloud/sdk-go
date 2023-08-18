# GetGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Groups** | [**[]FirewallGroupOutAPI**](FirewallGroupOutAPI.md) | Массив объектов Групп правил | 

## Methods

### NewGetGroups200Response

`func NewGetGroups200Response(responseId string, meta Meta, groups []FirewallGroupOutAPI, ) *GetGroups200Response`

NewGetGroups200Response instantiates a new GetGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroups200ResponseWithDefaults

`func NewGetGroups200ResponseWithDefaults() *GetGroups200Response`

NewGetGroups200ResponseWithDefaults instantiates a new GetGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetGroups200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetGroups200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetGroups200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetGroups200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetGroups200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetGroups200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetGroups

`func (o *GetGroups200Response) GetGroups() []FirewallGroupOutAPI`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetGroups200Response) GetGroupsOk() (*[]FirewallGroupOutAPI, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetGroups200Response) SetGroups(v []FirewallGroupOutAPI)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


