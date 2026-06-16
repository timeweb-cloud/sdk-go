# RouterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Router** | [**RouterOut**](RouterOut.md) |  | 

## Methods

### NewRouterResponse

`func NewRouterResponse(router RouterOut, ) *RouterResponse`

NewRouterResponse instantiates a new RouterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterResponseWithDefaults

`func NewRouterResponseWithDefaults() *RouterResponse`

NewRouterResponseWithDefaults instantiates a new RouterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *RouterResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *RouterResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *RouterResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *RouterResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetRouter

`func (o *RouterResponse) GetRouter() RouterOut`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *RouterResponse) GetRouterOk() (*RouterOut, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *RouterResponse) SetRouter(v RouterOut)`

SetRouter sets Router field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


