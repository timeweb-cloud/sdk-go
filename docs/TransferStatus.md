# TransferStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Общий статус трансфера. | 
**Tries** | **float32** | Количество попыток. | 
**TotalCount** | **float32** | Общее количество затронутых объектов. | 
**TotalSize** | **float32** | Общий размер затронутых объектов. | 
**UploadedCount** | **float32** | Количество перемещенных объектов. | 
**UploadedSize** | **float32** | Размер перемещенных объектов. | 
**Errors** | [**[]TransferStatusErrorsInner**](TransferStatusErrorsInner.md) |  | 

## Methods

### NewTransferStatus

`func NewTransferStatus(status string, tries float32, totalCount float32, totalSize float32, uploadedCount float32, uploadedSize float32, errors []TransferStatusErrorsInner, ) *TransferStatus`

NewTransferStatus instantiates a new TransferStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferStatusWithDefaults

`func NewTransferStatusWithDefaults() *TransferStatus`

NewTransferStatusWithDefaults instantiates a new TransferStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TransferStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTries

`func (o *TransferStatus) GetTries() float32`

GetTries returns the Tries field if non-nil, zero value otherwise.

### GetTriesOk

`func (o *TransferStatus) GetTriesOk() (*float32, bool)`

GetTriesOk returns a tuple with the Tries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTries

`func (o *TransferStatus) SetTries(v float32)`

SetTries sets Tries field to given value.


### GetTotalCount

`func (o *TransferStatus) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TransferStatus) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TransferStatus) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetTotalSize

`func (o *TransferStatus) GetTotalSize() float32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *TransferStatus) GetTotalSizeOk() (*float32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *TransferStatus) SetTotalSize(v float32)`

SetTotalSize sets TotalSize field to given value.


### GetUploadedCount

`func (o *TransferStatus) GetUploadedCount() float32`

GetUploadedCount returns the UploadedCount field if non-nil, zero value otherwise.

### GetUploadedCountOk

`func (o *TransferStatus) GetUploadedCountOk() (*float32, bool)`

GetUploadedCountOk returns a tuple with the UploadedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedCount

`func (o *TransferStatus) SetUploadedCount(v float32)`

SetUploadedCount sets UploadedCount field to given value.


### GetUploadedSize

`func (o *TransferStatus) GetUploadedSize() float32`

GetUploadedSize returns the UploadedSize field if non-nil, zero value otherwise.

### GetUploadedSizeOk

`func (o *TransferStatus) GetUploadedSizeOk() (*float32, bool)`

GetUploadedSizeOk returns a tuple with the UploadedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedSize

`func (o *TransferStatus) SetUploadedSize(v float32)`

SetUploadedSize sets UploadedSize field to given value.


### GetErrors

`func (o *TransferStatus) GetErrors() []TransferStatusErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TransferStatus) GetErrorsOk() (*[]TransferStatusErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TransferStatus) SetErrors(v []TransferStatusErrorsInner)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


