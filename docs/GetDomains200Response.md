# GetDomains200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Domains** | [**[]Domain**](Domain.md) |  | 

## Methods

### NewGetDomains200Response

`func NewGetDomains200Response(meta Meta, domains []Domain, ) *GetDomains200Response`

NewGetDomains200Response instantiates a new GetDomains200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomains200ResponseWithDefaults

`func NewGetDomains200ResponseWithDefaults() *GetDomains200Response`

NewGetDomains200ResponseWithDefaults instantiates a new GetDomains200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDomains200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDomains200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDomains200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDomains

`func (o *GetDomains200Response) GetDomains() []Domain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *GetDomains200Response) GetDomainsOk() (*[]Domain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *GetDomains200Response) SetDomains(v []Domain)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


