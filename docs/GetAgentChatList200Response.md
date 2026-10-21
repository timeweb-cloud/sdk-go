# GetAgentChatList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatList** | [**[]Chat**](Chat.md) |  | 
**Meta** | [**GetAgentChatList200ResponseMeta**](GetAgentChatList200ResponseMeta.md) |  | 

## Methods

### NewGetAgentChatList200Response

`func NewGetAgentChatList200Response(chatList []Chat, meta GetAgentChatList200ResponseMeta, ) *GetAgentChatList200Response`

NewGetAgentChatList200Response instantiates a new GetAgentChatList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentChatList200ResponseWithDefaults

`func NewGetAgentChatList200ResponseWithDefaults() *GetAgentChatList200Response`

NewGetAgentChatList200ResponseWithDefaults instantiates a new GetAgentChatList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatList

`func (o *GetAgentChatList200Response) GetChatList() []Chat`

GetChatList returns the ChatList field if non-nil, zero value otherwise.

### GetChatListOk

`func (o *GetAgentChatList200Response) GetChatListOk() (*[]Chat, bool)`

GetChatListOk returns a tuple with the ChatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatList

`func (o *GetAgentChatList200Response) SetChatList(v []Chat)`

SetChatList sets ChatList field to given value.


### GetMeta

`func (o *GetAgentChatList200Response) GetMeta() GetAgentChatList200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentChatList200Response) GetMetaOk() (*GetAgentChatList200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentChatList200Response) SetMeta(v GetAgentChatList200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


