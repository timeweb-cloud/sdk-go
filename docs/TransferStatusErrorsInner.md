# TransferStatusErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Текст ошибки. | 
**Try** | **float32** | Количество попыток. | 

## Methods

### NewTransferStatusErrorsInner

`func NewTransferStatusErrorsInner(value string, try float32, ) *TransferStatusErrorsInner`

NewTransferStatusErrorsInner instantiates a new TransferStatusErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferStatusErrorsInnerWithDefaults

`func NewTransferStatusErrorsInnerWithDefaults() *TransferStatusErrorsInner`

NewTransferStatusErrorsInnerWithDefaults instantiates a new TransferStatusErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TransferStatusErrorsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransferStatusErrorsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransferStatusErrorsInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetTry

`func (o *TransferStatusErrorsInner) GetTry() float32`

GetTry returns the Try field if non-nil, zero value otherwise.

### GetTryOk

`func (o *TransferStatusErrorsInner) GetTryOk() (*float32, bool)`

GetTryOk returns a tuple with the Try field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTry

`func (o *TransferStatusErrorsInner) SetTry(v float32)`

SetTry sets Try field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


