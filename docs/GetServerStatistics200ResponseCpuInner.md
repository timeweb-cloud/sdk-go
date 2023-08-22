# GetServerStatistics200ResponseCpuInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedAt** | **time.Time** | Дата события в формате ISO 8061 | 
**Load** | **float32** | Нагрузка нагрузка на процессор. Возвращает значение от 0 до 1, где 1 это 100% | 

## Methods

### NewGetServerStatistics200ResponseCpuInner

`func NewGetServerStatistics200ResponseCpuInner(loggedAt time.Time, load float32, ) *GetServerStatistics200ResponseCpuInner`

NewGetServerStatistics200ResponseCpuInner instantiates a new GetServerStatistics200ResponseCpuInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerStatistics200ResponseCpuInnerWithDefaults

`func NewGetServerStatistics200ResponseCpuInnerWithDefaults() *GetServerStatistics200ResponseCpuInner`

NewGetServerStatistics200ResponseCpuInnerWithDefaults instantiates a new GetServerStatistics200ResponseCpuInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedAt

`func (o *GetServerStatistics200ResponseCpuInner) GetLoggedAt() time.Time`

GetLoggedAt returns the LoggedAt field if non-nil, zero value otherwise.

### GetLoggedAtOk

`func (o *GetServerStatistics200ResponseCpuInner) GetLoggedAtOk() (*time.Time, bool)`

GetLoggedAtOk returns a tuple with the LoggedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedAt

`func (o *GetServerStatistics200ResponseCpuInner) SetLoggedAt(v time.Time)`

SetLoggedAt sets LoggedAt field to given value.


### GetLoad

`func (o *GetServerStatistics200ResponseCpuInner) GetLoad() float32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *GetServerStatistics200ResponseCpuInner) GetLoadOk() (*float32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *GetServerStatistics200ResponseCpuInner) SetLoad(v float32)`

SetLoad sets Load field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


