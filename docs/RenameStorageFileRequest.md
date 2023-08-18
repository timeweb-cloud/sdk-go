# RenameStorageFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewFilename** | **string** | Новое название файла или папки. Названия папок должны быть указаны с \&quot;/\&quot; в конце, например: \&quot;dirname/\&quot;. | 
**OldFilename** | **string** | Старое название файла или папки. Названия папок должны быть указаны с \&quot;/\&quot; в конце, например: \&quot;dirname/\&quot;. | 

## Methods

### NewRenameStorageFileRequest

`func NewRenameStorageFileRequest(newFilename string, oldFilename string, ) *RenameStorageFileRequest`

NewRenameStorageFileRequest instantiates a new RenameStorageFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameStorageFileRequestWithDefaults

`func NewRenameStorageFileRequestWithDefaults() *RenameStorageFileRequest`

NewRenameStorageFileRequestWithDefaults instantiates a new RenameStorageFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewFilename

`func (o *RenameStorageFileRequest) GetNewFilename() string`

GetNewFilename returns the NewFilename field if non-nil, zero value otherwise.

### GetNewFilenameOk

`func (o *RenameStorageFileRequest) GetNewFilenameOk() (*string, bool)`

GetNewFilenameOk returns a tuple with the NewFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFilename

`func (o *RenameStorageFileRequest) SetNewFilename(v string)`

SetNewFilename sets NewFilename field to given value.


### GetOldFilename

`func (o *RenameStorageFileRequest) GetOldFilename() string`

GetOldFilename returns the OldFilename field if non-nil, zero value otherwise.

### GetOldFilenameOk

`func (o *RenameStorageFileRequest) GetOldFilenameOk() (*string, bool)`

GetOldFilenameOk returns a tuple with the OldFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldFilename

`func (o *RenameStorageFileRequest) SetOldFilename(v string)`

SetOldFilename sets OldFilename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


