# GetCdnCertificates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]Certificate**](Certificate.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetCdnCertificates200Response

`func NewGetCdnCertificates200Response(certificates []Certificate, meta Meta, ) *GetCdnCertificates200Response`

NewGetCdnCertificates200Response instantiates a new GetCdnCertificates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCdnCertificates200ResponseWithDefaults

`func NewGetCdnCertificates200ResponseWithDefaults() *GetCdnCertificates200Response`

NewGetCdnCertificates200ResponseWithDefaults instantiates a new GetCdnCertificates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *GetCdnCertificates200Response) GetCertificates() []Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *GetCdnCertificates200Response) GetCertificatesOk() (*[]Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *GetCdnCertificates200Response) SetCertificates(v []Certificate)`

SetCertificates sets Certificates field to given value.


### GetMeta

`func (o *GetCdnCertificates200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCdnCertificates200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCdnCertificates200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


