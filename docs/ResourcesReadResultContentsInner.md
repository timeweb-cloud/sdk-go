# ResourcesReadResultContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | URI ресурса | 
**MimeType** | Pointer to **string** | Тип содержимого | [optional] 
**Text** | **string** | Содержимое ресурса — JSON, сериализованный в строку | 

## Methods

### NewResourcesReadResultContentsInner

`func NewResourcesReadResultContentsInner(uri string, text string, ) *ResourcesReadResultContentsInner`

NewResourcesReadResultContentsInner instantiates a new ResourcesReadResultContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesReadResultContentsInnerWithDefaults

`func NewResourcesReadResultContentsInnerWithDefaults() *ResourcesReadResultContentsInner`

NewResourcesReadResultContentsInnerWithDefaults instantiates a new ResourcesReadResultContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *ResourcesReadResultContentsInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ResourcesReadResultContentsInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ResourcesReadResultContentsInner) SetUri(v string)`

SetUri sets Uri field to given value.


### GetMimeType

`func (o *ResourcesReadResultContentsInner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResourcesReadResultContentsInner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResourcesReadResultContentsInner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResourcesReadResultContentsInner) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetText

`func (o *ResourcesReadResultContentsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ResourcesReadResultContentsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ResourcesReadResultContentsInner) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


