# AddServerIPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип IP-адреса | 
**Ptr** | Pointer to **string** | PTR-запись IP-адреса | [optional] 

## Methods

### NewAddServerIPRequest

`func NewAddServerIPRequest(type_ string, ) *AddServerIPRequest`

NewAddServerIPRequest instantiates a new AddServerIPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServerIPRequestWithDefaults

`func NewAddServerIPRequestWithDefaults() *AddServerIPRequest`

NewAddServerIPRequestWithDefaults instantiates a new AddServerIPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddServerIPRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddServerIPRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddServerIPRequest) SetType(v string)`

SetType sets Type field to given value.


### GetPtr

`func (o *AddServerIPRequest) GetPtr() string`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *AddServerIPRequest) GetPtrOk() (*string, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *AddServerIPRequest) SetPtr(v string)`

SetPtr sets Ptr field to given value.

### HasPtr

`func (o *AddServerIPRequest) HasPtr() bool`

HasPtr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


