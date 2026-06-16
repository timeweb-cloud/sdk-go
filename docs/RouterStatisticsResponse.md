# RouterStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Statistics** | [**[]RouterStatistic**](RouterStatistic.md) | Статистика | 

## Methods

### NewRouterStatisticsResponse

`func NewRouterStatisticsResponse(statistics []RouterStatistic, ) *RouterStatisticsResponse`

NewRouterStatisticsResponse instantiates a new RouterStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterStatisticsResponseWithDefaults

`func NewRouterStatisticsResponseWithDefaults() *RouterStatisticsResponse`

NewRouterStatisticsResponseWithDefaults instantiates a new RouterStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *RouterStatisticsResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *RouterStatisticsResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *RouterStatisticsResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *RouterStatisticsResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetStatistics

`func (o *RouterStatisticsResponse) GetStatistics() []RouterStatistic`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *RouterStatisticsResponse) GetStatisticsOk() (*[]RouterStatistic, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *RouterStatisticsResponse) SetStatistics(v []RouterStatistic)`

SetStatistics sets Statistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


