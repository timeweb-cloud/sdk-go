# AddBalancerToProject200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**ProjectResource**](ProjectResource.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewAddBalancerToProject200Response

`func NewAddBalancerToProject200Response(resource ProjectResource, responseId string, ) *AddBalancerToProject200Response`

NewAddBalancerToProject200Response instantiates a new AddBalancerToProject200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBalancerToProject200ResponseWithDefaults

`func NewAddBalancerToProject200ResponseWithDefaults() *AddBalancerToProject200Response`

NewAddBalancerToProject200ResponseWithDefaults instantiates a new AddBalancerToProject200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *AddBalancerToProject200Response) GetResource() ProjectResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AddBalancerToProject200Response) GetResourceOk() (*ProjectResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AddBalancerToProject200Response) SetResource(v ProjectResource)`

SetResource sets Resource field to given value.


### GetResponseId

`func (o *AddBalancerToProject200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddBalancerToProject200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddBalancerToProject200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


