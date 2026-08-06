# IssueCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **int32** | ID CDN-ресурса, для доменов которого нужно выпустить сертификат. | 

## Methods

### NewIssueCertificate

`func NewIssueCertificate(resourceId int32, ) *IssueCertificate`

NewIssueCertificate instantiates a new IssueCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCertificateWithDefaults

`func NewIssueCertificateWithDefaults() *IssueCertificate`

NewIssueCertificateWithDefaults instantiates a new IssueCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *IssueCertificate) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *IssueCertificate) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *IssueCertificate) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


