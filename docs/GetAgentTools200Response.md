# GetAgentTools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | [**[]Tool**](Tool.md) |  | 
**Meta** | [**GetAgentTools200ResponseMeta**](GetAgentTools200ResponseMeta.md) |  | 

## Methods

### NewGetAgentTools200Response

`func NewGetAgentTools200Response(tools []Tool, meta GetAgentTools200ResponseMeta, ) *GetAgentTools200Response`

NewGetAgentTools200Response instantiates a new GetAgentTools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentTools200ResponseWithDefaults

`func NewGetAgentTools200ResponseWithDefaults() *GetAgentTools200Response`

NewGetAgentTools200ResponseWithDefaults instantiates a new GetAgentTools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *GetAgentTools200Response) GetTools() []Tool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *GetAgentTools200Response) GetToolsOk() (*[]Tool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *GetAgentTools200Response) SetTools(v []Tool)`

SetTools sets Tools field to given value.


### GetMeta

`func (o *GetAgentTools200Response) GetMeta() GetAgentTools200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentTools200Response) GetMetaOk() (*GetAgentTools200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentTools200Response) SetMeta(v GetAgentTools200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


