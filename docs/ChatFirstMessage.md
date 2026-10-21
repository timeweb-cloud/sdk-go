# ChatFirstMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID сообщения | 
**Role** | **string** | Автор сообщения | 
**CreatedAt** | **time.Time** | Дата создания сообщения | 
**Text** | **string** | Текст сообщения | 

## Methods

### NewChatFirstMessage

`func NewChatFirstMessage(id int32, role string, createdAt time.Time, text string, ) *ChatFirstMessage`

NewChatFirstMessage instantiates a new ChatFirstMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatFirstMessageWithDefaults

`func NewChatFirstMessageWithDefaults() *ChatFirstMessage`

NewChatFirstMessageWithDefaults instantiates a new ChatFirstMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatFirstMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatFirstMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatFirstMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetRole

`func (o *ChatFirstMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatFirstMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatFirstMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *ChatFirstMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatFirstMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatFirstMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetText

`func (o *ChatFirstMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChatFirstMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChatFirstMessage) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


