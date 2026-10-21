# GetAiKeyCallsStatistic200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiKeyCallsStatistics** | [**[]AiKeyCallsStatisticPoint**](AiKeyCallsStatisticPoint.md) |  | 
**Meta** | [**GetAiKeyTokenUsageStatistic200ResponseMeta**](GetAiKeyTokenUsageStatistic200ResponseMeta.md) |  | 

## Methods

### NewGetAiKeyCallsStatistic200Response

`func NewGetAiKeyCallsStatistic200Response(aiKeyCallsStatistics []AiKeyCallsStatisticPoint, meta GetAiKeyTokenUsageStatistic200ResponseMeta, ) *GetAiKeyCallsStatistic200Response`

NewGetAiKeyCallsStatistic200Response instantiates a new GetAiKeyCallsStatistic200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiKeyCallsStatistic200ResponseWithDefaults

`func NewGetAiKeyCallsStatistic200ResponseWithDefaults() *GetAiKeyCallsStatistic200Response`

NewGetAiKeyCallsStatistic200ResponseWithDefaults instantiates a new GetAiKeyCallsStatistic200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiKeyCallsStatistics

`func (o *GetAiKeyCallsStatistic200Response) GetAiKeyCallsStatistics() []AiKeyCallsStatisticPoint`

GetAiKeyCallsStatistics returns the AiKeyCallsStatistics field if non-nil, zero value otherwise.

### GetAiKeyCallsStatisticsOk

`func (o *GetAiKeyCallsStatistic200Response) GetAiKeyCallsStatisticsOk() (*[]AiKeyCallsStatisticPoint, bool)`

GetAiKeyCallsStatisticsOk returns a tuple with the AiKeyCallsStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeyCallsStatistics

`func (o *GetAiKeyCallsStatistic200Response) SetAiKeyCallsStatistics(v []AiKeyCallsStatisticPoint)`

SetAiKeyCallsStatistics sets AiKeyCallsStatistics field to given value.


### GetMeta

`func (o *GetAiKeyCallsStatistic200Response) GetMeta() GetAiKeyTokenUsageStatistic200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAiKeyCallsStatistic200Response) GetMetaOk() (*GetAiKeyTokenUsageStatistic200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAiKeyCallsStatistic200Response) SetMeta(v GetAiKeyTokenUsageStatistic200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


