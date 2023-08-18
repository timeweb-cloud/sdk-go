# UpdateMailQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Общее количество места на почте (в Мб). | [optional] 

## Methods

### NewUpdateMailQuotaRequest

`func NewUpdateMailQuotaRequest() *UpdateMailQuotaRequest`

NewUpdateMailQuotaRequest instantiates a new UpdateMailQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailQuotaRequestWithDefaults

`func NewUpdateMailQuotaRequestWithDefaults() *UpdateMailQuotaRequest`

NewUpdateMailQuotaRequestWithDefaults instantiates a new UpdateMailQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *UpdateMailQuotaRequest) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UpdateMailQuotaRequest) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UpdateMailQuotaRequest) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UpdateMailQuotaRequest) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


