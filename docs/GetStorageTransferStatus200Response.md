# GetStorageTransferStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferStatus** | [**TransferStatus**](TransferStatus.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetStorageTransferStatus200Response

`func NewGetStorageTransferStatus200Response(transferStatus TransferStatus, responseId string, ) *GetStorageTransferStatus200Response`

NewGetStorageTransferStatus200Response instantiates a new GetStorageTransferStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageTransferStatus200ResponseWithDefaults

`func NewGetStorageTransferStatus200ResponseWithDefaults() *GetStorageTransferStatus200Response`

NewGetStorageTransferStatus200ResponseWithDefaults instantiates a new GetStorageTransferStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferStatus

`func (o *GetStorageTransferStatus200Response) GetTransferStatus() TransferStatus`

GetTransferStatus returns the TransferStatus field if non-nil, zero value otherwise.

### GetTransferStatusOk

`func (o *GetStorageTransferStatus200Response) GetTransferStatusOk() (*TransferStatus, bool)`

GetTransferStatusOk returns a tuple with the TransferStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferStatus

`func (o *GetStorageTransferStatus200Response) SetTransferStatus(v TransferStatus)`

SetTransferStatus sets TransferStatus field to given value.


### GetResponseId

`func (o *GetStorageTransferStatus200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetStorageTransferStatus200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetStorageTransferStatus200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


