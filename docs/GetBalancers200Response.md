# GetBalancers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Balancers** | [**[]Balancer**](Balancer.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetBalancers200Response

`func NewGetBalancers200Response(meta Meta, balancers []Balancer, responseId string, ) *GetBalancers200Response`

NewGetBalancers200Response instantiates a new GetBalancers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalancers200ResponseWithDefaults

`func NewGetBalancers200ResponseWithDefaults() *GetBalancers200Response`

NewGetBalancers200ResponseWithDefaults instantiates a new GetBalancers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBalancers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBalancers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBalancers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetBalancers

`func (o *GetBalancers200Response) GetBalancers() []Balancer`

GetBalancers returns the Balancers field if non-nil, zero value otherwise.

### GetBalancersOk

`func (o *GetBalancers200Response) GetBalancersOk() (*[]Balancer, bool)`

GetBalancersOk returns a tuple with the Balancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancers

`func (o *GetBalancers200Response) SetBalancers(v []Balancer)`

SetBalancers sets Balancers field to given value.


### GetResponseId

`func (o *GetBalancers200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetBalancers200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetBalancers200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


