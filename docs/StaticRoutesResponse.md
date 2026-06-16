# StaticRoutesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**StaticRoutes** | [**[]StaticRouteOut**](StaticRouteOut.md) | Статические маршруты | 
**Meta** | [**ComponentsSchemasMeta**](ComponentsSchemasMeta.md) |  | 

## Methods

### NewStaticRoutesResponse

`func NewStaticRoutesResponse(staticRoutes []StaticRouteOut, meta ComponentsSchemasMeta, ) *StaticRoutesResponse`

NewStaticRoutesResponse instantiates a new StaticRoutesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutesResponseWithDefaults

`func NewStaticRoutesResponseWithDefaults() *StaticRoutesResponse`

NewStaticRoutesResponseWithDefaults instantiates a new StaticRoutesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *StaticRoutesResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *StaticRoutesResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *StaticRoutesResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *StaticRoutesResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *StaticRoutesResponse) GetStaticRoutes() []StaticRouteOut`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *StaticRoutesResponse) GetStaticRoutesOk() (*[]StaticRouteOut, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *StaticRoutesResponse) SetStaticRoutes(v []StaticRouteOut)`

SetStaticRoutes sets StaticRoutes field to given value.


### GetMeta

`func (o *StaticRoutesResponse) GetMeta() ComponentsSchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StaticRoutesResponse) GetMetaOk() (*ComponentsSchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StaticRoutesResponse) SetMeta(v ComponentsSchemasMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


