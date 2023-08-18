# CreateBalancerRule200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | [**Rule**](Rule.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateBalancerRule200Response

`func NewCreateBalancerRule200Response(rule Rule, responseId string, ) *CreateBalancerRule200Response`

NewCreateBalancerRule200Response instantiates a new CreateBalancerRule200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBalancerRule200ResponseWithDefaults

`func NewCreateBalancerRule200ResponseWithDefaults() *CreateBalancerRule200Response`

NewCreateBalancerRule200ResponseWithDefaults instantiates a new CreateBalancerRule200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *CreateBalancerRule200Response) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateBalancerRule200Response) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateBalancerRule200Response) SetRule(v Rule)`

SetRule sets Rule field to given value.


### GetResponseId

`func (o *CreateBalancerRule200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateBalancerRule200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateBalancerRule200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


