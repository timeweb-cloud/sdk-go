# AddCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** | Приватный ключ сертификата в формате PEM. | 
**Certificate** | **string** | Сертификат вместе с цепочкой в формате PEM. | 
**Comment** | Pointer to **string** | Комментарий к сертификату. | [optional] [default to "SSL Certificate"]

## Methods

### NewAddCertificate

`func NewAddCertificate(privateKey string, certificate string, ) *AddCertificate`

NewAddCertificate instantiates a new AddCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCertificateWithDefaults

`func NewAddCertificateWithDefaults() *AddCertificate`

NewAddCertificateWithDefaults instantiates a new AddCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *AddCertificate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AddCertificate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AddCertificate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetCertificate

`func (o *AddCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AddCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AddCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetComment

`func (o *AddCertificate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AddCertificate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AddCertificate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AddCertificate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


