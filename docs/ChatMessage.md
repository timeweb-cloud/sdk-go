# ChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID сообщения | 
**MessageId** | Pointer to **string** | Строковый идентификатор сообщения. Есть не у всех сообщений. | [optional] 
**Role** | **string** | Автор сообщения | 
**CreatedAt** | **time.Time** | Дата создания сообщения | 
**Text** | **string** | Текст сообщения | 
**IsGood** | **NullableBool** | Оценка ответа: &#x60;true&#x60; — понравился, &#x60;false&#x60; — не понравился, &#x60;null&#x60; — без оценки | 
**UptimeMs** | **NullableInt32** | Время до начала ответа агента в миллисекундах (для ответов агента) | 
**File** | [**NullableChatMessageFile**](ChatMessageFile.md) |  | 

## Methods

### NewChatMessage

`func NewChatMessage(id int32, role string, createdAt time.Time, text string, isGood NullableBool, uptimeMs NullableInt32, file NullableChatMessageFile, ) *ChatMessage`

NewChatMessage instantiates a new ChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageWithDefaults

`func NewChatMessageWithDefaults() *ChatMessage`

NewChatMessageWithDefaults instantiates a new ChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatMessage) SetId(v int32)`

SetId sets Id field to given value.


### GetMessageId

`func (o *ChatMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ChatMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ChatMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ChatMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetRole

`func (o *ChatMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *ChatMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetText

`func (o *ChatMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChatMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChatMessage) SetText(v string)`

SetText sets Text field to given value.


### GetIsGood

`func (o *ChatMessage) GetIsGood() bool`

GetIsGood returns the IsGood field if non-nil, zero value otherwise.

### GetIsGoodOk

`func (o *ChatMessage) GetIsGoodOk() (*bool, bool)`

GetIsGoodOk returns a tuple with the IsGood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGood

`func (o *ChatMessage) SetIsGood(v bool)`

SetIsGood sets IsGood field to given value.


### SetIsGoodNil

`func (o *ChatMessage) SetIsGoodNil(b bool)`

 SetIsGoodNil sets the value for IsGood to be an explicit nil

### UnsetIsGood
`func (o *ChatMessage) UnsetIsGood()`

UnsetIsGood ensures that no value is present for IsGood, not even an explicit nil
### GetUptimeMs

`func (o *ChatMessage) GetUptimeMs() int32`

GetUptimeMs returns the UptimeMs field if non-nil, zero value otherwise.

### GetUptimeMsOk

`func (o *ChatMessage) GetUptimeMsOk() (*int32, bool)`

GetUptimeMsOk returns a tuple with the UptimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeMs

`func (o *ChatMessage) SetUptimeMs(v int32)`

SetUptimeMs sets UptimeMs field to given value.


### SetUptimeMsNil

`func (o *ChatMessage) SetUptimeMsNil(b bool)`

 SetUptimeMsNil sets the value for UptimeMs to be an explicit nil

### UnsetUptimeMs
`func (o *ChatMessage) UnsetUptimeMs()`

UnsetUptimeMs ensures that no value is present for UptimeMs, not even an explicit nil
### GetFile

`func (o *ChatMessage) GetFile() ChatMessageFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ChatMessage) GetFileOk() (*ChatMessageFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ChatMessage) SetFile(v ChatMessageFile)`

SetFile sets File field to given value.


### SetFileNil

`func (o *ChatMessage) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *ChatMessage) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


