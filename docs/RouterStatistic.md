# RouterStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя метрики | 
**List** | [**[]RouterStatisticListInner**](RouterStatisticListInner.md) | Значения метрики | 
**Meta** | [**RouterStatisticMeta**](RouterStatisticMeta.md) |  | 

## Methods

### NewRouterStatistic

`func NewRouterStatistic(name string, list []RouterStatisticListInner, meta RouterStatisticMeta, ) *RouterStatistic`

NewRouterStatistic instantiates a new RouterStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterStatisticWithDefaults

`func NewRouterStatisticWithDefaults() *RouterStatistic`

NewRouterStatisticWithDefaults instantiates a new RouterStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouterStatistic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouterStatistic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouterStatistic) SetName(v string)`

SetName sets Name field to given value.


### GetList

`func (o *RouterStatistic) GetList() []RouterStatisticListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *RouterStatistic) GetListOk() (*[]RouterStatisticListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *RouterStatistic) SetList(v []RouterStatisticListInner)`

SetList sets List field to given value.


### GetMeta

`func (o *RouterStatistic) GetMeta() RouterStatisticMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RouterStatistic) GetMetaOk() (*RouterStatisticMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RouterStatistic) SetMeta(v RouterStatisticMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


