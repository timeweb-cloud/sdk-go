# Tool2Annotations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnlyHint** | Pointer to **bool** | &#x60;true&#x60; у инструментов, которые только читают данные. Инструменты без этой аннотации могут что-то изменить и требуют подтверждения | [optional] 

## Methods

### NewTool2Annotations

`func NewTool2Annotations() *Tool2Annotations`

NewTool2Annotations instantiates a new Tool2Annotations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTool2AnnotationsWithDefaults

`func NewTool2AnnotationsWithDefaults() *Tool2Annotations`

NewTool2AnnotationsWithDefaults instantiates a new Tool2Annotations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnlyHint

`func (o *Tool2Annotations) GetReadOnlyHint() bool`

GetReadOnlyHint returns the ReadOnlyHint field if non-nil, zero value otherwise.

### GetReadOnlyHintOk

`func (o *Tool2Annotations) GetReadOnlyHintOk() (*bool, bool)`

GetReadOnlyHintOk returns a tuple with the ReadOnlyHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyHint

`func (o *Tool2Annotations) SetReadOnlyHint(v bool)`

SetReadOnlyHint sets ReadOnlyHint field to given value.

### HasReadOnlyHint

`func (o *Tool2Annotations) HasReadOnlyHint() bool`

HasReadOnlyHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


