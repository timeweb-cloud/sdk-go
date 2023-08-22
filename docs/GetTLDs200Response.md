# GetTLDs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**TopLevelDomains** | [**[]TopLevelDomain**](TopLevelDomain.md) |  | 

## Methods

### NewGetTLDs200Response

`func NewGetTLDs200Response(meta Meta, topLevelDomains []TopLevelDomain, ) *GetTLDs200Response`

NewGetTLDs200Response instantiates a new GetTLDs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTLDs200ResponseWithDefaults

`func NewGetTLDs200ResponseWithDefaults() *GetTLDs200Response`

NewGetTLDs200ResponseWithDefaults instantiates a new GetTLDs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTLDs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTLDs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTLDs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetTopLevelDomains

`func (o *GetTLDs200Response) GetTopLevelDomains() []TopLevelDomain`

GetTopLevelDomains returns the TopLevelDomains field if non-nil, zero value otherwise.

### GetTopLevelDomainsOk

`func (o *GetTLDs200Response) GetTopLevelDomainsOk() (*[]TopLevelDomain, bool)`

GetTopLevelDomainsOk returns a tuple with the TopLevelDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelDomains

`func (o *GetTLDs200Response) SetTopLevelDomains(v []TopLevelDomain)`

SetTopLevelDomains sets TopLevelDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


