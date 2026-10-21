# UploadResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadedFiles** | **[]int32** | ID созданных документов | 
**Documents** | [**[]Document**](Document.md) | Созданные документы | 

## Methods

### NewUploadResult

`func NewUploadResult(uploadedFiles []int32, documents []Document, ) *UploadResult`

NewUploadResult instantiates a new UploadResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadResultWithDefaults

`func NewUploadResultWithDefaults() *UploadResult`

NewUploadResultWithDefaults instantiates a new UploadResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadedFiles

`func (o *UploadResult) GetUploadedFiles() []int32`

GetUploadedFiles returns the UploadedFiles field if non-nil, zero value otherwise.

### GetUploadedFilesOk

`func (o *UploadResult) GetUploadedFilesOk() (*[]int32, bool)`

GetUploadedFilesOk returns a tuple with the UploadedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedFiles

`func (o *UploadResult) SetUploadedFiles(v []int32)`

SetUploadedFiles sets UploadedFiles field to given value.


### GetDocuments

`func (o *UploadResult) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *UploadResult) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *UploadResult) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


