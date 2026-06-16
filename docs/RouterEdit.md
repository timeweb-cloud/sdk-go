# RouterEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Новое имя роутера | [optional] 
**Comment** | Pointer to **string** | Новое описание роутера | [optional] 

## Methods

### NewRouterEdit

`func NewRouterEdit() *RouterEdit`

NewRouterEdit instantiates a new RouterEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterEditWithDefaults

`func NewRouterEditWithDefaults() *RouterEdit`

NewRouterEditWithDefaults instantiates a new RouterEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouterEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouterEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouterEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouterEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *RouterEdit) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RouterEdit) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RouterEdit) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RouterEdit) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


