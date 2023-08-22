# GetStorageSubdomains200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subdomains** | [**[]S3Subdomain**](S3Subdomain.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetStorageSubdomains200Response

`func NewGetStorageSubdomains200Response(subdomains []S3Subdomain, meta Meta, ) *GetStorageSubdomains200Response`

NewGetStorageSubdomains200Response instantiates a new GetStorageSubdomains200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageSubdomains200ResponseWithDefaults

`func NewGetStorageSubdomains200ResponseWithDefaults() *GetStorageSubdomains200Response`

NewGetStorageSubdomains200ResponseWithDefaults instantiates a new GetStorageSubdomains200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubdomains

`func (o *GetStorageSubdomains200Response) GetSubdomains() []S3Subdomain`

GetSubdomains returns the Subdomains field if non-nil, zero value otherwise.

### GetSubdomainsOk

`func (o *GetStorageSubdomains200Response) GetSubdomainsOk() (*[]S3Subdomain, bool)`

GetSubdomainsOk returns a tuple with the Subdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomains

`func (o *GetStorageSubdomains200Response) SetSubdomains(v []S3Subdomain)`

SetSubdomains sets Subdomains field to given value.


### GetMeta

`func (o *GetStorageSubdomains200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetStorageSubdomains200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetStorageSubdomains200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


