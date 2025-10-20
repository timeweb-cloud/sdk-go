# GetAgentStatistics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentStatistics** | [**[]TokenStatistic**](TokenStatistic.md) |  | 
**Meta** | [**GetAgentStatistics200ResponseMeta**](GetAgentStatistics200ResponseMeta.md) |  | 

## Methods

### NewGetAgentStatistics200Response

`func NewGetAgentStatistics200Response(agentStatistics []TokenStatistic, meta GetAgentStatistics200ResponseMeta, ) *GetAgentStatistics200Response`

NewGetAgentStatistics200Response instantiates a new GetAgentStatistics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentStatistics200ResponseWithDefaults

`func NewGetAgentStatistics200ResponseWithDefaults() *GetAgentStatistics200Response`

NewGetAgentStatistics200ResponseWithDefaults instantiates a new GetAgentStatistics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentStatistics

`func (o *GetAgentStatistics200Response) GetAgentStatistics() []TokenStatistic`

GetAgentStatistics returns the AgentStatistics field if non-nil, zero value otherwise.

### GetAgentStatisticsOk

`func (o *GetAgentStatistics200Response) GetAgentStatisticsOk() (*[]TokenStatistic, bool)`

GetAgentStatisticsOk returns a tuple with the AgentStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatistics

`func (o *GetAgentStatistics200Response) SetAgentStatistics(v []TokenStatistic)`

SetAgentStatistics sets AgentStatistics field to given value.


### GetMeta

`func (o *GetAgentStatistics200Response) GetMeta() GetAgentStatistics200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentStatistics200Response) GetMetaOk() (*GetAgentStatistics200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentStatistics200Response) SetMeta(v GetAgentStatistics200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


