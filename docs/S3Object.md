# S3Object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Название файла или папки. | [optional] 
**LastModified** | Pointer to **time.Time** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была сделана последняя модификация файла или папки. | [optional] 
**Etag** | Pointer to **string** | Тег. | [optional] 
**Size** | Pointer to **int32** | Размер (в байтах) файла или папки. | [optional] 
**StorageClass** | Pointer to **string** | Класс хранилища. | [optional] 
**ChecksumAlgorithm** | Pointer to **[]string** | Алгоритм | [optional] 
**Owner** | Pointer to [**S3ObjectOwner**](S3ObjectOwner.md) |  | [optional] 
**Type** | **string** | Тип (файл или папка). | 

## Methods

### NewS3Object

`func NewS3Object(type_ string, ) *S3Object`

NewS3Object instantiates a new S3Object object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ObjectWithDefaults

`func NewS3ObjectWithDefaults() *S3Object`

NewS3ObjectWithDefaults instantiates a new S3Object object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *S3Object) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3Object) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3Object) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3Object) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastModified

`func (o *S3Object) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *S3Object) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *S3Object) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *S3Object) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetEtag

`func (o *S3Object) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *S3Object) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *S3Object) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *S3Object) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetSize

`func (o *S3Object) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *S3Object) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *S3Object) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *S3Object) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageClass

`func (o *S3Object) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *S3Object) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *S3Object) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *S3Object) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetChecksumAlgorithm

`func (o *S3Object) GetChecksumAlgorithm() []string`

GetChecksumAlgorithm returns the ChecksumAlgorithm field if non-nil, zero value otherwise.

### GetChecksumAlgorithmOk

`func (o *S3Object) GetChecksumAlgorithmOk() (*[]string, bool)`

GetChecksumAlgorithmOk returns a tuple with the ChecksumAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumAlgorithm

`func (o *S3Object) SetChecksumAlgorithm(v []string)`

SetChecksumAlgorithm sets ChecksumAlgorithm field to given value.

### HasChecksumAlgorithm

`func (o *S3Object) HasChecksumAlgorithm() bool`

HasChecksumAlgorithm returns a boolean if a field has been set.

### GetOwner

`func (o *S3Object) GetOwner() S3ObjectOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *S3Object) GetOwnerOk() (*S3ObjectOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *S3Object) SetOwner(v S3ObjectOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *S3Object) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetType

`func (o *S3Object) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3Object) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3Object) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


