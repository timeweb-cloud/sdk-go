# AddStorageSubdomains200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subdomains** | [**[]AddedSubdomain**](AddedSubdomain.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewAddStorageSubdomains200Response

`func NewAddStorageSubdomains200Response(subdomains []AddedSubdomain, meta Meta, ) *AddStorageSubdomains200Response`

NewAddStorageSubdomains200Response instantiates a new AddStorageSubdomains200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageSubdomains200ResponseWithDefaults

`func NewAddStorageSubdomains200ResponseWithDefaults() *AddStorageSubdomains200Response`

NewAddStorageSubdomains200ResponseWithDefaults instantiates a new AddStorageSubdomains200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubdomains

`func (o *AddStorageSubdomains200Response) GetSubdomains() []AddedSubdomain`

GetSubdomains returns the Subdomains field if non-nil, zero value otherwise.

### GetSubdomainsOk

`func (o *AddStorageSubdomains200Response) GetSubdomainsOk() (*[]AddedSubdomain, bool)`

GetSubdomainsOk returns a tuple with the Subdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomains

`func (o *AddStorageSubdomains200Response) SetSubdomains(v []AddedSubdomain)`

SetSubdomains sets Subdomains field to given value.


### GetMeta

`func (o *AddStorageSubdomains200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddStorageSubdomains200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddStorageSubdomains200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


