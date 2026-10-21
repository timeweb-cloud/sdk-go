# GetKnowledgebasesV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnowledgeBases** | [**[]KnowledgebaseV2**](KnowledgebaseV2.md) |  | 
**Meta** | [**GetKnowledgebasesV3200ResponseMeta**](GetKnowledgebasesV3200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebasesV2200Response

`func NewGetKnowledgebasesV2200Response(knowledgeBases []KnowledgebaseV2, meta GetKnowledgebasesV3200ResponseMeta, ) *GetKnowledgebasesV2200Response`

NewGetKnowledgebasesV2200Response instantiates a new GetKnowledgebasesV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebasesV2200ResponseWithDefaults

`func NewGetKnowledgebasesV2200ResponseWithDefaults() *GetKnowledgebasesV2200Response`

NewGetKnowledgebasesV2200ResponseWithDefaults instantiates a new GetKnowledgebasesV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgeBases

`func (o *GetKnowledgebasesV2200Response) GetKnowledgeBases() []KnowledgebaseV2`

GetKnowledgeBases returns the KnowledgeBases field if non-nil, zero value otherwise.

### GetKnowledgeBasesOk

`func (o *GetKnowledgebasesV2200Response) GetKnowledgeBasesOk() (*[]KnowledgebaseV2, bool)`

GetKnowledgeBasesOk returns a tuple with the KnowledgeBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBases

`func (o *GetKnowledgebasesV2200Response) SetKnowledgeBases(v []KnowledgebaseV2)`

SetKnowledgeBases sets KnowledgeBases field to given value.


### GetMeta

`func (o *GetKnowledgebasesV2200Response) GetMeta() GetKnowledgebasesV3200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebasesV2200Response) GetMetaOk() (*GetKnowledgebasesV3200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebasesV2200Response) SetMeta(v GetKnowledgebasesV3200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


