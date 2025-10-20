# GetKnowledgebaseStatistics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnowledgebaseStatistics** | [**[]TokenStatistic**](TokenStatistic.md) |  | 
**Meta** | [**GetAgentStatistics200ResponseMeta**](GetAgentStatistics200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebaseStatistics200Response

`func NewGetKnowledgebaseStatistics200Response(knowledgebaseStatistics []TokenStatistic, meta GetAgentStatistics200ResponseMeta, ) *GetKnowledgebaseStatistics200Response`

NewGetKnowledgebaseStatistics200Response instantiates a new GetKnowledgebaseStatistics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebaseStatistics200ResponseWithDefaults

`func NewGetKnowledgebaseStatistics200ResponseWithDefaults() *GetKnowledgebaseStatistics200Response`

NewGetKnowledgebaseStatistics200ResponseWithDefaults instantiates a new GetKnowledgebaseStatistics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgebaseStatistics

`func (o *GetKnowledgebaseStatistics200Response) GetKnowledgebaseStatistics() []TokenStatistic`

GetKnowledgebaseStatistics returns the KnowledgebaseStatistics field if non-nil, zero value otherwise.

### GetKnowledgebaseStatisticsOk

`func (o *GetKnowledgebaseStatistics200Response) GetKnowledgebaseStatisticsOk() (*[]TokenStatistic, bool)`

GetKnowledgebaseStatisticsOk returns a tuple with the KnowledgebaseStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebaseStatistics

`func (o *GetKnowledgebaseStatistics200Response) SetKnowledgebaseStatistics(v []TokenStatistic)`

SetKnowledgebaseStatistics sets KnowledgebaseStatistics field to given value.


### GetMeta

`func (o *GetKnowledgebaseStatistics200Response) GetMeta() GetAgentStatistics200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebaseStatistics200Response) GetMetaOk() (*GetAgentStatistics200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebaseStatistics200Response) SetMeta(v GetAgentStatistics200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


