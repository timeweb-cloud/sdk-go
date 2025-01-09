# UpdateStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PresetId** | Pointer to **float32** | ID тарифа. | [optional] 
**BucketType** | Pointer to **string** | Тип хранилища. | [optional] 

## Methods

### NewUpdateStorageRequest

`func NewUpdateStorageRequest() *UpdateStorageRequest`

NewUpdateStorageRequest instantiates a new UpdateStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageRequestWithDefaults

`func NewUpdateStorageRequestWithDefaults() *UpdateStorageRequest`

NewUpdateStorageRequestWithDefaults instantiates a new UpdateStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresetId

`func (o *UpdateStorageRequest) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateStorageRequest) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateStorageRequest) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateStorageRequest) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetBucketType

`func (o *UpdateStorageRequest) GetBucketType() string`

GetBucketType returns the BucketType field if non-nil, zero value otherwise.

### GetBucketTypeOk

`func (o *UpdateStorageRequest) GetBucketTypeOk() (*string, bool)`

GetBucketTypeOk returns a tuple with the BucketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketType

`func (o *UpdateStorageRequest) SetBucketType(v string)`

SetBucketType sets BucketType field to given value.

### HasBucketType

`func (o *UpdateStorageRequest) HasBucketType() bool`

HasBucketType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


