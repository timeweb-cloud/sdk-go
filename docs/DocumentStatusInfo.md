# DocumentStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип статуса документа | 
**Timestamp** | **time.Time** | Время записи информации (ISO string) | 
**ErrorCode** | Pointer to **string** | Код ошибки (только для type: &#39;error&#39;) | [optional] 
**Details** | Pointer to [**DocumentStatusInfoDetails**](DocumentStatusInfoDetails.md) |  | [optional] 

## Methods

### NewDocumentStatusInfo

`func NewDocumentStatusInfo(type_ string, timestamp time.Time, ) *DocumentStatusInfo`

NewDocumentStatusInfo instantiates a new DocumentStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentStatusInfoWithDefaults

`func NewDocumentStatusInfoWithDefaults() *DocumentStatusInfo`

NewDocumentStatusInfoWithDefaults instantiates a new DocumentStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DocumentStatusInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentStatusInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentStatusInfo) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *DocumentStatusInfo) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DocumentStatusInfo) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DocumentStatusInfo) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetErrorCode

`func (o *DocumentStatusInfo) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DocumentStatusInfo) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DocumentStatusInfo) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DocumentStatusInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetDetails

`func (o *DocumentStatusInfo) GetDetails() DocumentStatusInfoDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DocumentStatusInfo) GetDetailsOk() (*DocumentStatusInfoDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DocumentStatusInfo) SetDetails(v DocumentStatusInfoDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DocumentStatusInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


