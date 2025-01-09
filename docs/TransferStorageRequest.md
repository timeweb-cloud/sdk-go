# TransferStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | ID доступа стороннего S3-хранилища. | 
**SecretKey** | **string** | Пароль доступа стороннего S3-хранилища. | 
**Location** | **string** | Регион хранилища источника. | 
**IsForcePathStyle** | **bool** | Это логическое значение, которое показывает, следует ли принудительно указывать URL-адреса для объектов S3. | 
**Endpoint** | **string** | URL S3-хранилища источника. | 
**BucketName** | **string** | Имя хранилища источника. | 
**NewBucketName** | **string** | Имя хранилища получателя. | 

## Methods

### NewTransferStorageRequest

`func NewTransferStorageRequest(accessKey string, secretKey string, location string, isForcePathStyle bool, endpoint string, bucketName string, newBucketName string, ) *TransferStorageRequest`

NewTransferStorageRequest instantiates a new TransferStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferStorageRequestWithDefaults

`func NewTransferStorageRequestWithDefaults() *TransferStorageRequest`

NewTransferStorageRequestWithDefaults instantiates a new TransferStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *TransferStorageRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *TransferStorageRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *TransferStorageRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *TransferStorageRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *TransferStorageRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *TransferStorageRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetLocation

`func (o *TransferStorageRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TransferStorageRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TransferStorageRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetIsForcePathStyle

`func (o *TransferStorageRequest) GetIsForcePathStyle() bool`

GetIsForcePathStyle returns the IsForcePathStyle field if non-nil, zero value otherwise.

### GetIsForcePathStyleOk

`func (o *TransferStorageRequest) GetIsForcePathStyleOk() (*bool, bool)`

GetIsForcePathStyleOk returns a tuple with the IsForcePathStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForcePathStyle

`func (o *TransferStorageRequest) SetIsForcePathStyle(v bool)`

SetIsForcePathStyle sets IsForcePathStyle field to given value.


### GetEndpoint

`func (o *TransferStorageRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *TransferStorageRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *TransferStorageRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetBucketName

`func (o *TransferStorageRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *TransferStorageRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *TransferStorageRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetNewBucketName

`func (o *TransferStorageRequest) GetNewBucketName() string`

GetNewBucketName returns the NewBucketName field if non-nil, zero value otherwise.

### GetNewBucketNameOk

`func (o *TransferStorageRequest) GetNewBucketNameOk() (*string, bool)`

GetNewBucketNameOk returns a tuple with the NewBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewBucketName

`func (o *TransferStorageRequest) SetNewBucketName(v string)`

SetNewBucketName sets NewBucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


