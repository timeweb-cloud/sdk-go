# DatabaseClusterDomainsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | Доменное имя кластера базы данных. | 
**CertStatus** | **string** | Статус SSL-сертификата домена. | 

## Methods

### NewDatabaseClusterDomainsInner

`func NewDatabaseClusterDomainsInner(fqdn string, certStatus string, ) *DatabaseClusterDomainsInner`

NewDatabaseClusterDomainsInner instantiates a new DatabaseClusterDomainsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterDomainsInnerWithDefaults

`func NewDatabaseClusterDomainsInnerWithDefaults() *DatabaseClusterDomainsInner`

NewDatabaseClusterDomainsInnerWithDefaults instantiates a new DatabaseClusterDomainsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *DatabaseClusterDomainsInner) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DatabaseClusterDomainsInner) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DatabaseClusterDomainsInner) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetCertStatus

`func (o *DatabaseClusterDomainsInner) GetCertStatus() string`

GetCertStatus returns the CertStatus field if non-nil, zero value otherwise.

### GetCertStatusOk

`func (o *DatabaseClusterDomainsInner) GetCertStatusOk() (*string, bool)`

GetCertStatusOk returns a tuple with the CertStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStatus

`func (o *DatabaseClusterDomainsInner) SetCertStatus(v string)`

SetCertStatus sets CertStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


