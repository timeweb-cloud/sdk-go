# GetCdnCertificateTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateTasks** | [**[]CertificateTask**](CertificateTask.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetCdnCertificateTasks200Response

`func NewGetCdnCertificateTasks200Response(certificateTasks []CertificateTask, meta Meta, ) *GetCdnCertificateTasks200Response`

NewGetCdnCertificateTasks200Response instantiates a new GetCdnCertificateTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCdnCertificateTasks200ResponseWithDefaults

`func NewGetCdnCertificateTasks200ResponseWithDefaults() *GetCdnCertificateTasks200Response`

NewGetCdnCertificateTasks200ResponseWithDefaults instantiates a new GetCdnCertificateTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateTasks

`func (o *GetCdnCertificateTasks200Response) GetCertificateTasks() []CertificateTask`

GetCertificateTasks returns the CertificateTasks field if non-nil, zero value otherwise.

### GetCertificateTasksOk

`func (o *GetCdnCertificateTasks200Response) GetCertificateTasksOk() (*[]CertificateTask, bool)`

GetCertificateTasksOk returns a tuple with the CertificateTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTasks

`func (o *GetCdnCertificateTasks200Response) SetCertificateTasks(v []CertificateTask)`

SetCertificateTasks sets CertificateTasks field to given value.


### GetMeta

`func (o *GetCdnCertificateTasks200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCdnCertificateTasks200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCdnCertificateTasks200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


