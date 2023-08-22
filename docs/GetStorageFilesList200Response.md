# GetStorageFilesList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]S3Object**](S3Object.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetStorageFilesList200Response

`func NewGetStorageFilesList200Response(files []S3Object, meta Meta, ) *GetStorageFilesList200Response`

NewGetStorageFilesList200Response instantiates a new GetStorageFilesList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageFilesList200ResponseWithDefaults

`func NewGetStorageFilesList200ResponseWithDefaults() *GetStorageFilesList200Response`

NewGetStorageFilesList200ResponseWithDefaults instantiates a new GetStorageFilesList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *GetStorageFilesList200Response) GetFiles() []S3Object`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetStorageFilesList200Response) GetFilesOk() (*[]S3Object, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetStorageFilesList200Response) SetFiles(v []S3Object)`

SetFiles sets Files field to given value.


### GetMeta

`func (o *GetStorageFilesList200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetStorageFilesList200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetStorageFilesList200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


