# GetDomainRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Requests** | [**[]DomainRequest**](DomainRequest.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetDomainRequests200Response

`func NewGetDomainRequests200Response(meta Meta, requests []DomainRequest, responseId string, ) *GetDomainRequests200Response`

NewGetDomainRequests200Response instantiates a new GetDomainRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainRequests200ResponseWithDefaults

`func NewGetDomainRequests200ResponseWithDefaults() *GetDomainRequests200Response`

NewGetDomainRequests200ResponseWithDefaults instantiates a new GetDomainRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDomainRequests200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDomainRequests200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDomainRequests200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetRequests

`func (o *GetDomainRequests200Response) GetRequests() []DomainRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *GetDomainRequests200Response) GetRequestsOk() (*[]DomainRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *GetDomainRequests200Response) SetRequests(v []DomainRequest)`

SetRequests sets Requests field to given value.


### GetResponseId

`func (o *GetDomainRequests200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetDomainRequests200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetDomainRequests200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


