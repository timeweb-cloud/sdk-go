# BucketDiskStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **float32** | Размер (в Кб) диска хранилища включенного в тариф. | 
**Used** | **float32** | Размер (в Кб) использованного пространства диска хранилища. | 
**IsUnlimited** | **bool** | Признак безлимитного размера хранилища. | 

## Methods

### NewBucketDiskStats

`func NewBucketDiskStats(size float32, used float32, isUnlimited bool, ) *BucketDiskStats`

NewBucketDiskStats instantiates a new BucketDiskStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketDiskStatsWithDefaults

`func NewBucketDiskStatsWithDefaults() *BucketDiskStats`

NewBucketDiskStatsWithDefaults instantiates a new BucketDiskStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *BucketDiskStats) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BucketDiskStats) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BucketDiskStats) SetSize(v float32)`

SetSize sets Size field to given value.


### GetUsed

`func (o *BucketDiskStats) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *BucketDiskStats) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *BucketDiskStats) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetIsUnlimited

`func (o *BucketDiskStats) GetIsUnlimited() bool`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *BucketDiskStats) GetIsUnlimitedOk() (*bool, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *BucketDiskStats) SetIsUnlimited(v bool)`

SetIsUnlimited sets IsUnlimited field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


