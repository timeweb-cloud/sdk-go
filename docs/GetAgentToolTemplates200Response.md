# GetAgentToolTemplates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolTemplates** | [**[]ToolTemplate**](ToolTemplate.md) |  | 
**Meta** | [**GetAgentToolTemplates200ResponseMeta**](GetAgentToolTemplates200ResponseMeta.md) |  | 

## Methods

### NewGetAgentToolTemplates200Response

`func NewGetAgentToolTemplates200Response(toolTemplates []ToolTemplate, meta GetAgentToolTemplates200ResponseMeta, ) *GetAgentToolTemplates200Response`

NewGetAgentToolTemplates200Response instantiates a new GetAgentToolTemplates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentToolTemplates200ResponseWithDefaults

`func NewGetAgentToolTemplates200ResponseWithDefaults() *GetAgentToolTemplates200Response`

NewGetAgentToolTemplates200ResponseWithDefaults instantiates a new GetAgentToolTemplates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolTemplates

`func (o *GetAgentToolTemplates200Response) GetToolTemplates() []ToolTemplate`

GetToolTemplates returns the ToolTemplates field if non-nil, zero value otherwise.

### GetToolTemplatesOk

`func (o *GetAgentToolTemplates200Response) GetToolTemplatesOk() (*[]ToolTemplate, bool)`

GetToolTemplatesOk returns a tuple with the ToolTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolTemplates

`func (o *GetAgentToolTemplates200Response) SetToolTemplates(v []ToolTemplate)`

SetToolTemplates sets ToolTemplates field to given value.


### GetMeta

`func (o *GetAgentToolTemplates200Response) GetMeta() GetAgentToolTemplates200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentToolTemplates200Response) GetMetaOk() (*GetAgentToolTemplates200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentToolTemplates200Response) SetMeta(v GetAgentToolTemplates200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


