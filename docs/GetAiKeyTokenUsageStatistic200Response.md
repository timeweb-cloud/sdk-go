# GetAiKeyTokenUsageStatistic200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiKeyTokenStatistics** | [**[]AiKeyTokenStatisticPoint**](AiKeyTokenStatisticPoint.md) |  | 
**Meta** | [**GetAiKeyTokenUsageStatistic200ResponseMeta**](GetAiKeyTokenUsageStatistic200ResponseMeta.md) |  | 

## Methods

### NewGetAiKeyTokenUsageStatistic200Response

`func NewGetAiKeyTokenUsageStatistic200Response(aiKeyTokenStatistics []AiKeyTokenStatisticPoint, meta GetAiKeyTokenUsageStatistic200ResponseMeta, ) *GetAiKeyTokenUsageStatistic200Response`

NewGetAiKeyTokenUsageStatistic200Response instantiates a new GetAiKeyTokenUsageStatistic200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiKeyTokenUsageStatistic200ResponseWithDefaults

`func NewGetAiKeyTokenUsageStatistic200ResponseWithDefaults() *GetAiKeyTokenUsageStatistic200Response`

NewGetAiKeyTokenUsageStatistic200ResponseWithDefaults instantiates a new GetAiKeyTokenUsageStatistic200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiKeyTokenStatistics

`func (o *GetAiKeyTokenUsageStatistic200Response) GetAiKeyTokenStatistics() []AiKeyTokenStatisticPoint`

GetAiKeyTokenStatistics returns the AiKeyTokenStatistics field if non-nil, zero value otherwise.

### GetAiKeyTokenStatisticsOk

`func (o *GetAiKeyTokenUsageStatistic200Response) GetAiKeyTokenStatisticsOk() (*[]AiKeyTokenStatisticPoint, bool)`

GetAiKeyTokenStatisticsOk returns a tuple with the AiKeyTokenStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiKeyTokenStatistics

`func (o *GetAiKeyTokenUsageStatistic200Response) SetAiKeyTokenStatistics(v []AiKeyTokenStatisticPoint)`

SetAiKeyTokenStatistics sets AiKeyTokenStatistics field to given value.


### GetMeta

`func (o *GetAiKeyTokenUsageStatistic200Response) GetMeta() GetAiKeyTokenUsageStatistic200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAiKeyTokenUsageStatistic200Response) GetMetaOk() (*GetAiKeyTokenUsageStatistic200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAiKeyTokenUsageStatistic200Response) SetMeta(v GetAiKeyTokenUsageStatistic200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


