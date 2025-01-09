# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра хранилища. Автоматически генерируется при создании. | 
**Name** | **string** | Удобочитаемое имя, установленное для хранилища. | 
**DiskStats** | [**BucketDiskStats**](BucketDiskStats.md) |  | 
**Type** | **string** | Тип хранилища. | 
**PresetId** | **NullableFloat32** | ID тарифа хранилища. | 
**Status** | **string** | Статус хранилища. | 
**ObjectAmount** | **float32** | Количество файлов в хранилище. | 
**Location** | **string** | Регион хранилища. | 
**Hostname** | **string** | Адрес хранилища для подключения. | 
**AccessKey** | **string** | Ключ доступа от хранилища. | 
**SecretKey** | **string** | Секретный ключ доступа от хранилища. | 

## Methods

### NewBucket

`func NewBucket(id float32, name string, diskStats BucketDiskStats, type_ string, presetId NullableFloat32, status string, objectAmount float32, location string, hostname string, accessKey string, secretKey string, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bucket) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bucket) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bucket) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.


### GetDiskStats

`func (o *Bucket) GetDiskStats() BucketDiskStats`

GetDiskStats returns the DiskStats field if non-nil, zero value otherwise.

### GetDiskStatsOk

`func (o *Bucket) GetDiskStatsOk() (*BucketDiskStats, bool)`

GetDiskStatsOk returns a tuple with the DiskStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStats

`func (o *Bucket) SetDiskStats(v BucketDiskStats)`

SetDiskStats sets DiskStats field to given value.


### GetType

`func (o *Bucket) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Bucket) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Bucket) SetType(v string)`

SetType sets Type field to given value.


### GetPresetId

`func (o *Bucket) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *Bucket) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *Bucket) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### SetPresetIdNil

`func (o *Bucket) SetPresetIdNil(b bool)`

 SetPresetIdNil sets the value for PresetId to be an explicit nil

### UnsetPresetId
`func (o *Bucket) UnsetPresetId()`

UnsetPresetId ensures that no value is present for PresetId, not even an explicit nil
### GetStatus

`func (o *Bucket) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bucket) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bucket) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectAmount

`func (o *Bucket) GetObjectAmount() float32`

GetObjectAmount returns the ObjectAmount field if non-nil, zero value otherwise.

### GetObjectAmountOk

`func (o *Bucket) GetObjectAmountOk() (*float32, bool)`

GetObjectAmountOk returns a tuple with the ObjectAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectAmount

`func (o *Bucket) SetObjectAmount(v float32)`

SetObjectAmount sets ObjectAmount field to given value.


### GetLocation

`func (o *Bucket) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Bucket) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Bucket) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetHostname

`func (o *Bucket) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Bucket) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Bucket) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetAccessKey

`func (o *Bucket) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Bucket) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Bucket) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *Bucket) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *Bucket) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *Bucket) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


