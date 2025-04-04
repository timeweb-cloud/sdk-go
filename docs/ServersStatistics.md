# ServersStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Имя статистики. | [optional] 
**List** | Pointer to [**[]ServersStatisticsListInner**](ServersStatisticsListInner.md) |  | [optional] 
**Meta** | Pointer to [**ServersStatisticsMeta**](ServersStatisticsMeta.md) |  | [optional] 

## Methods

### NewServersStatistics

`func NewServersStatistics() *ServersStatistics`

NewServersStatistics instantiates a new ServersStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersStatisticsWithDefaults

`func NewServersStatisticsWithDefaults() *ServersStatistics`

NewServersStatisticsWithDefaults instantiates a new ServersStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServersStatistics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServersStatistics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServersStatistics) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServersStatistics) HasName() bool`

HasName returns a boolean if a field has been set.

### GetList

`func (o *ServersStatistics) GetList() []ServersStatisticsListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ServersStatistics) GetListOk() (*[]ServersStatisticsListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ServersStatistics) SetList(v []ServersStatisticsListInner)`

SetList sets List field to given value.

### HasList

`func (o *ServersStatistics) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMeta

`func (o *ServersStatistics) GetMeta() ServersStatisticsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServersStatistics) GetMetaOk() (*ServersStatisticsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServersStatistics) SetMeta(v ServersStatisticsMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServersStatistics) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


