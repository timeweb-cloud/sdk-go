# AiSolutionChatStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | **bool** | Готов ли чат принимать сообщения | 
**Reason** | Pointer to **string** | Почему чат ещё не готов. Отсутствует, если чат готов. &#x60;unreachable&#x60; — до сервера не достучаться (разворачивается, выключен или нет публичного IP); &#x60;software_installing&#x60; — устанавливается ПО агента; &#x60;configuring&#x60; — агент настраивается; &#x60;install_failed&#x60; — установка ПО не удалась, пересоздайте решение; &#x60;machine_error&#x60; — сервер в состоянии ошибки, пересоздайте решение; &#x60;deleted&#x60; — решение удалено; &#x60;outdated&#x60; — сервер создан старой версией решения, пересоздайте решение | [optional] 

## Methods

### NewAiSolutionChatStatus

`func NewAiSolutionChatStatus(ready bool, ) *AiSolutionChatStatus`

NewAiSolutionChatStatus instantiates a new AiSolutionChatStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSolutionChatStatusWithDefaults

`func NewAiSolutionChatStatusWithDefaults() *AiSolutionChatStatus`

NewAiSolutionChatStatusWithDefaults instantiates a new AiSolutionChatStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *AiSolutionChatStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *AiSolutionChatStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *AiSolutionChatStatus) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetReason

`func (o *AiSolutionChatStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AiSolutionChatStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AiSolutionChatStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AiSolutionChatStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


