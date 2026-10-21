# GetAgentChatMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatListMessages** | [**[]ChatMessage**](ChatMessage.md) |  | 
**Meta** | [**GetAgentChatMessages200ResponseMeta**](GetAgentChatMessages200ResponseMeta.md) |  | 

## Methods

### NewGetAgentChatMessages200Response

`func NewGetAgentChatMessages200Response(chatListMessages []ChatMessage, meta GetAgentChatMessages200ResponseMeta, ) *GetAgentChatMessages200Response`

NewGetAgentChatMessages200Response instantiates a new GetAgentChatMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentChatMessages200ResponseWithDefaults

`func NewGetAgentChatMessages200ResponseWithDefaults() *GetAgentChatMessages200Response`

NewGetAgentChatMessages200ResponseWithDefaults instantiates a new GetAgentChatMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatListMessages

`func (o *GetAgentChatMessages200Response) GetChatListMessages() []ChatMessage`

GetChatListMessages returns the ChatListMessages field if non-nil, zero value otherwise.

### GetChatListMessagesOk

`func (o *GetAgentChatMessages200Response) GetChatListMessagesOk() (*[]ChatMessage, bool)`

GetChatListMessagesOk returns a tuple with the ChatListMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatListMessages

`func (o *GetAgentChatMessages200Response) SetChatListMessages(v []ChatMessage)`

SetChatListMessages sets ChatListMessages field to given value.


### GetMeta

`func (o *GetAgentChatMessages200Response) GetMeta() GetAgentChatMessages200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentChatMessages200Response) GetMetaOk() (*GetAgentChatMessages200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentChatMessages200Response) SetMeta(v GetAgentChatMessages200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


