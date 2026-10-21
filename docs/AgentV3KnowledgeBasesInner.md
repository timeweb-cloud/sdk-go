# AgentV3KnowledgeBasesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID базы знаний | 
**RetrievalMode** | **string** | Режим поиска по базе знаний: &#x60;classic&#x60; — найденные фрагменты автоматически добавляются к каждому запросу; &#x60;agentic&#x60; — агент сам решает, когда искать в базе знаний. | 

## Methods

### NewAgentV3KnowledgeBasesInner

`func NewAgentV3KnowledgeBasesInner(id int32, retrievalMode string, ) *AgentV3KnowledgeBasesInner`

NewAgentV3KnowledgeBasesInner instantiates a new AgentV3KnowledgeBasesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentV3KnowledgeBasesInnerWithDefaults

`func NewAgentV3KnowledgeBasesInnerWithDefaults() *AgentV3KnowledgeBasesInner`

NewAgentV3KnowledgeBasesInnerWithDefaults instantiates a new AgentV3KnowledgeBasesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentV3KnowledgeBasesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentV3KnowledgeBasesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentV3KnowledgeBasesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetRetrievalMode

`func (o *AgentV3KnowledgeBasesInner) GetRetrievalMode() string`

GetRetrievalMode returns the RetrievalMode field if non-nil, zero value otherwise.

### GetRetrievalModeOk

`func (o *AgentV3KnowledgeBasesInner) GetRetrievalModeOk() (*string, bool)`

GetRetrievalModeOk returns a tuple with the RetrievalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievalMode

`func (o *AgentV3KnowledgeBasesInner) SetRetrievalMode(v string)`

SetRetrievalMode sets RetrievalMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


