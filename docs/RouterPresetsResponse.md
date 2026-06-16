# RouterPresetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**RouterPresets** | [**[]RouterPreset**](RouterPreset.md) | Тарифы роутеров | 
**Meta** | [**ComponentsSchemasMeta**](ComponentsSchemasMeta.md) |  | 

## Methods

### NewRouterPresetsResponse

`func NewRouterPresetsResponse(routerPresets []RouterPreset, meta ComponentsSchemasMeta, ) *RouterPresetsResponse`

NewRouterPresetsResponse instantiates a new RouterPresetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterPresetsResponseWithDefaults

`func NewRouterPresetsResponseWithDefaults() *RouterPresetsResponse`

NewRouterPresetsResponseWithDefaults instantiates a new RouterPresetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *RouterPresetsResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *RouterPresetsResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *RouterPresetsResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *RouterPresetsResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetRouterPresets

`func (o *RouterPresetsResponse) GetRouterPresets() []RouterPreset`

GetRouterPresets returns the RouterPresets field if non-nil, zero value otherwise.

### GetRouterPresetsOk

`func (o *RouterPresetsResponse) GetRouterPresetsOk() (*[]RouterPreset, bool)`

GetRouterPresetsOk returns a tuple with the RouterPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPresets

`func (o *RouterPresetsResponse) SetRouterPresets(v []RouterPreset)`

SetRouterPresets sets RouterPresets field to given value.


### GetMeta

`func (o *RouterPresetsResponse) GetMeta() ComponentsSchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RouterPresetsResponse) GetMetaOk() (*ComponentsSchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RouterPresetsResponse) SetMeta(v ComponentsSchemasMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


