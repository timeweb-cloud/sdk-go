# GetKnowledgebases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnowledgeBases** | [**[]Knowledgebase**](Knowledgebase.md) |  | 
**Meta** | [**GetKnowledgebasesV3200ResponseMeta**](GetKnowledgebasesV3200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebases200Response

`func NewGetKnowledgebases200Response(knowledgeBases []Knowledgebase, meta GetKnowledgebasesV3200ResponseMeta, ) *GetKnowledgebases200Response`

NewGetKnowledgebases200Response instantiates a new GetKnowledgebases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebases200ResponseWithDefaults

`func NewGetKnowledgebases200ResponseWithDefaults() *GetKnowledgebases200Response`

NewGetKnowledgebases200ResponseWithDefaults instantiates a new GetKnowledgebases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgeBases

`func (o *GetKnowledgebases200Response) GetKnowledgeBases() []Knowledgebase`

GetKnowledgeBases returns the KnowledgeBases field if non-nil, zero value otherwise.

### GetKnowledgeBasesOk

`func (o *GetKnowledgebases200Response) GetKnowledgeBasesOk() (*[]Knowledgebase, bool)`

GetKnowledgeBasesOk returns a tuple with the KnowledgeBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBases

`func (o *GetKnowledgebases200Response) SetKnowledgeBases(v []Knowledgebase)`

SetKnowledgeBases sets KnowledgeBases field to given value.


### GetMeta

`func (o *GetKnowledgebases200Response) GetMeta() GetKnowledgebasesV3200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebases200Response) GetMetaOk() (*GetKnowledgebasesV3200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebases200Response) SetMeta(v GetKnowledgebasesV3200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


