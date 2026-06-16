# RoutersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Routers** | [**[]RouterOut**](RouterOut.md) | Роутеры | 
**Meta** | [**ComponentsSchemasMeta**](ComponentsSchemasMeta.md) |  | 

## Methods

### NewRoutersResponse

`func NewRoutersResponse(routers []RouterOut, meta ComponentsSchemasMeta, ) *RoutersResponse`

NewRoutersResponse instantiates a new RoutersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutersResponseWithDefaults

`func NewRoutersResponseWithDefaults() *RoutersResponse`

NewRoutersResponseWithDefaults instantiates a new RoutersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *RoutersResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *RoutersResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *RoutersResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *RoutersResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetRouters

`func (o *RoutersResponse) GetRouters() []RouterOut`

GetRouters returns the Routers field if non-nil, zero value otherwise.

### GetRoutersOk

`func (o *RoutersResponse) GetRoutersOk() (*[]RouterOut, bool)`

GetRoutersOk returns a tuple with the Routers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouters

`func (o *RoutersResponse) SetRouters(v []RouterOut)`

SetRouters sets Routers field to given value.


### GetMeta

`func (o *RoutersResponse) GetMeta() ComponentsSchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RoutersResponse) GetMetaOk() (*ComponentsSchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RoutersResponse) SetMeta(v ComponentsSchemasMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


