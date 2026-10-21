# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID чата | 
**AgentId** | **int32** | ID агента | 
**CreatedAt** | **time.Time** | Дата первого сообщения чата | 
**FirstMessage** | [**ChatFirstMessage**](ChatFirstMessage.md) |  | 
**LikesCount** | **int32** | Количество положительных оценок в чате | 
**DislikesCount** | **int32** | Количество отрицательных оценок в чате | 
**UptimeAvg** | **NullableFloat32** | Среднее время до начала ответа агента в чате, в миллисекундах | 

## Methods

### NewChat

`func NewChat(id string, agentId int32, createdAt time.Time, firstMessage ChatFirstMessage, likesCount int32, dislikesCount int32, uptimeAvg NullableFloat32, ) *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Chat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Chat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Chat) SetId(v string)`

SetId sets Id field to given value.


### GetAgentId

`func (o *Chat) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *Chat) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *Chat) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


### GetCreatedAt

`func (o *Chat) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Chat) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Chat) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFirstMessage

`func (o *Chat) GetFirstMessage() ChatFirstMessage`

GetFirstMessage returns the FirstMessage field if non-nil, zero value otherwise.

### GetFirstMessageOk

`func (o *Chat) GetFirstMessageOk() (*ChatFirstMessage, bool)`

GetFirstMessageOk returns a tuple with the FirstMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessage

`func (o *Chat) SetFirstMessage(v ChatFirstMessage)`

SetFirstMessage sets FirstMessage field to given value.


### GetLikesCount

`func (o *Chat) GetLikesCount() int32`

GetLikesCount returns the LikesCount field if non-nil, zero value otherwise.

### GetLikesCountOk

`func (o *Chat) GetLikesCountOk() (*int32, bool)`

GetLikesCountOk returns a tuple with the LikesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikesCount

`func (o *Chat) SetLikesCount(v int32)`

SetLikesCount sets LikesCount field to given value.


### GetDislikesCount

`func (o *Chat) GetDislikesCount() int32`

GetDislikesCount returns the DislikesCount field if non-nil, zero value otherwise.

### GetDislikesCountOk

`func (o *Chat) GetDislikesCountOk() (*int32, bool)`

GetDislikesCountOk returns a tuple with the DislikesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikesCount

`func (o *Chat) SetDislikesCount(v int32)`

SetDislikesCount sets DislikesCount field to given value.


### GetUptimeAvg

`func (o *Chat) GetUptimeAvg() float32`

GetUptimeAvg returns the UptimeAvg field if non-nil, zero value otherwise.

### GetUptimeAvgOk

`func (o *Chat) GetUptimeAvgOk() (*float32, bool)`

GetUptimeAvgOk returns a tuple with the UptimeAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeAvg

`func (o *Chat) SetUptimeAvg(v float32)`

SetUptimeAvg sets UptimeAvg field to given value.


### SetUptimeAvgNil

`func (o *Chat) SetUptimeAvgNil(b bool)`

 SetUptimeAvgNil sets the value for UptimeAvg to be an explicit nil

### UnsetUptimeAvg
`func (o *Chat) UnsetUptimeAvg()`

UnsetUptimeAvg ensures that no value is present for UptimeAvg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


