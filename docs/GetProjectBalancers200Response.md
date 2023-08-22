# GetProjectBalancers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balancers** | [**[]Balancer**](Balancer.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetProjectBalancers200Response

`func NewGetProjectBalancers200Response(balancers []Balancer, meta Meta, ) *GetProjectBalancers200Response`

NewGetProjectBalancers200Response instantiates a new GetProjectBalancers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectBalancers200ResponseWithDefaults

`func NewGetProjectBalancers200ResponseWithDefaults() *GetProjectBalancers200Response`

NewGetProjectBalancers200ResponseWithDefaults instantiates a new GetProjectBalancers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancers

`func (o *GetProjectBalancers200Response) GetBalancers() []Balancer`

GetBalancers returns the Balancers field if non-nil, zero value otherwise.

### GetBalancersOk

`func (o *GetProjectBalancers200Response) GetBalancersOk() (*[]Balancer, bool)`

GetBalancersOk returns a tuple with the Balancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancers

`func (o *GetProjectBalancers200Response) SetBalancers(v []Balancer)`

SetBalancers sets Balancers field to given value.


### GetMeta

`func (o *GetProjectBalancers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProjectBalancers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProjectBalancers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


