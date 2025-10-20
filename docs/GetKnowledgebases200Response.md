# GetKnowledgebases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Knowledgebases** | [**[]Knowledgebase**](Knowledgebase.md) |  | 
**Meta** | [**GetKnowledgebases200ResponseMeta**](GetKnowledgebases200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebases200Response

`func NewGetKnowledgebases200Response(knowledgebases []Knowledgebase, meta GetKnowledgebases200ResponseMeta, ) *GetKnowledgebases200Response`

NewGetKnowledgebases200Response instantiates a new GetKnowledgebases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebases200ResponseWithDefaults

`func NewGetKnowledgebases200ResponseWithDefaults() *GetKnowledgebases200Response`

NewGetKnowledgebases200ResponseWithDefaults instantiates a new GetKnowledgebases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgebases

`func (o *GetKnowledgebases200Response) GetKnowledgebases() []Knowledgebase`

GetKnowledgebases returns the Knowledgebases field if non-nil, zero value otherwise.

### GetKnowledgebasesOk

`func (o *GetKnowledgebases200Response) GetKnowledgebasesOk() (*[]Knowledgebase, bool)`

GetKnowledgebasesOk returns a tuple with the Knowledgebases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebases

`func (o *GetKnowledgebases200Response) SetKnowledgebases(v []Knowledgebase)`

SetKnowledgebases sets Knowledgebases field to given value.


### GetMeta

`func (o *GetKnowledgebases200Response) GetMeta() GetKnowledgebases200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebases200Response) GetMetaOk() (*GetKnowledgebases200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebases200Response) SetMeta(v GetKnowledgebases200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


