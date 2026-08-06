# CertificateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID задачи на выпуск сертификата. | 
**Status** | **string** | Статус выпуска сертификата. - &#x60;in_progress&#x60; — сертификат выпускается; - &#x60;success&#x60; — сертификат выпущен и привязан к ресурсу; - &#x60;failed&#x60; — выпустить сертификат не удалось. | 
**Domains** | **[]string** | Доменные имена, для которых выпускается сертификат. | 
**ResourceId** | **int32** | ID CDN-ресурса, для которого выпускается сертификат. | 

## Methods

### NewCertificateTask

`func NewCertificateTask(id int32, status string, domains []string, resourceId int32, ) *CertificateTask`

NewCertificateTask instantiates a new CertificateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateTaskWithDefaults

`func NewCertificateTaskWithDefaults() *CertificateTask`

NewCertificateTaskWithDefaults instantiates a new CertificateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateTask) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *CertificateTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateTask) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDomains

`func (o *CertificateTask) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CertificateTask) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CertificateTask) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetResourceId

`func (o *CertificateTask) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CertificateTask) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CertificateTask) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


