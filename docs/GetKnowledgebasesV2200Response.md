# GetKnowledgebasesV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Knowledgebases** | [**[]KnowledgebaseV2**](KnowledgebaseV2.md) |  | 
**Meta** | [**GetKnowledgebasesV2200ResponseMeta**](GetKnowledgebasesV2200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebasesV2200Response

`func NewGetKnowledgebasesV2200Response(knowledgebases []KnowledgebaseV2, meta GetKnowledgebasesV2200ResponseMeta, ) *GetKnowledgebasesV2200Response`

NewGetKnowledgebasesV2200Response instantiates a new GetKnowledgebasesV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebasesV2200ResponseWithDefaults

`func NewGetKnowledgebasesV2200ResponseWithDefaults() *GetKnowledgebasesV2200Response`

NewGetKnowledgebasesV2200ResponseWithDefaults instantiates a new GetKnowledgebasesV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgebases

`func (o *GetKnowledgebasesV2200Response) GetKnowledgebases() []KnowledgebaseV2`

GetKnowledgebases returns the Knowledgebases field if non-nil, zero value otherwise.

### GetKnowledgebasesOk

`func (o *GetKnowledgebasesV2200Response) GetKnowledgebasesOk() (*[]KnowledgebaseV2, bool)`

GetKnowledgebasesOk returns a tuple with the Knowledgebases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebases

`func (o *GetKnowledgebasesV2200Response) SetKnowledgebases(v []KnowledgebaseV2)`

SetKnowledgebases sets Knowledgebases field to given value.


### GetMeta

`func (o *GetKnowledgebasesV2200Response) GetMeta() GetKnowledgebasesV2200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebasesV2200Response) GetMetaOk() (*GetKnowledgebasesV2200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebasesV2200Response) SetMeta(v GetKnowledgebasesV2200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


