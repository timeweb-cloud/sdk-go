# S3Subdomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID поддомена. | 
**Subdomain** | **string** | Поддомен. | 
**CertReleased** | **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был выдан SSL сертификат. | 
**Tries** | **float32** | Количество попыток перевыпустить SSL сертификат. | 
**Status** | **string** | Поддомен. | 

## Methods

### NewS3Subdomain

`func NewS3Subdomain(id float32, subdomain string, certReleased time.Time, tries float32, status string, ) *S3Subdomain`

NewS3Subdomain instantiates a new S3Subdomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3SubdomainWithDefaults

`func NewS3SubdomainWithDefaults() *S3Subdomain`

NewS3SubdomainWithDefaults instantiates a new S3Subdomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3Subdomain) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3Subdomain) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3Subdomain) SetId(v float32)`

SetId sets Id field to given value.


### GetSubdomain

`func (o *S3Subdomain) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *S3Subdomain) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *S3Subdomain) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.


### GetCertReleased

`func (o *S3Subdomain) GetCertReleased() time.Time`

GetCertReleased returns the CertReleased field if non-nil, zero value otherwise.

### GetCertReleasedOk

`func (o *S3Subdomain) GetCertReleasedOk() (*time.Time, bool)`

GetCertReleasedOk returns a tuple with the CertReleased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertReleased

`func (o *S3Subdomain) SetCertReleased(v time.Time)`

SetCertReleased sets CertReleased field to given value.


### GetTries

`func (o *S3Subdomain) GetTries() float32`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *S3Subdomain) GetTriesOk() (*float32, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *S3Subdomain) SetTries(v float32)`

SetTries sets Tries field to given value.


### GetStatus

`func (o *S3Subdomain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *S3Subdomain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *S3Subdomain) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


