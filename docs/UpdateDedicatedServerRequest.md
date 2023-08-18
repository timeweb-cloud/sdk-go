# UpdateDedicatedServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Удобочитаемое имя выделенного сервера. Максимальная длина — 255 символов, имя должно быть уникальным. | [optional] 
**Comment** | Pointer to **string** | Комментарий к выделенному серверу. Максимальная длина — 255 символов. | [optional] 

## Methods

### NewUpdateDedicatedServerRequest

`func NewUpdateDedicatedServerRequest() *UpdateDedicatedServerRequest`

NewUpdateDedicatedServerRequest instantiates a new UpdateDedicatedServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDedicatedServerRequestWithDefaults

`func NewUpdateDedicatedServerRequestWithDefaults() *UpdateDedicatedServerRequest`

NewUpdateDedicatedServerRequestWithDefaults instantiates a new UpdateDedicatedServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDedicatedServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDedicatedServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDedicatedServerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDedicatedServerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *UpdateDedicatedServerRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateDedicatedServerRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateDedicatedServerRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateDedicatedServerRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


