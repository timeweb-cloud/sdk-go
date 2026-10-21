# GetKnowledgebasesV3200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnowledgeBases** | [**[]KnowledgebaseV3**](KnowledgebaseV3.md) |  | 
**Meta** | [**GetKnowledgebasesV3200ResponseMeta**](GetKnowledgebasesV3200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebasesV3200Response

`func NewGetKnowledgebasesV3200Response(knowledgeBases []KnowledgebaseV3, meta GetKnowledgebasesV3200ResponseMeta, ) *GetKnowledgebasesV3200Response`

NewGetKnowledgebasesV3200Response instantiates a new GetKnowledgebasesV3200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebasesV3200ResponseWithDefaults

`func NewGetKnowledgebasesV3200ResponseWithDefaults() *GetKnowledgebasesV3200Response`

NewGetKnowledgebasesV3200ResponseWithDefaults instantiates a new GetKnowledgebasesV3200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgeBases

`func (o *GetKnowledgebasesV3200Response) GetKnowledgeBases() []KnowledgebaseV3`

GetKnowledgeBases returns the KnowledgeBases field if non-nil, zero value otherwise.

### GetKnowledgeBasesOk

`func (o *GetKnowledgebasesV3200Response) GetKnowledgeBasesOk() (*[]KnowledgebaseV3, bool)`

GetKnowledgeBasesOk returns a tuple with the KnowledgeBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBases

`func (o *GetKnowledgebasesV3200Response) SetKnowledgeBases(v []KnowledgebaseV3)`

SetKnowledgeBases sets KnowledgeBases field to given value.


### GetMeta

`func (o *GetKnowledgebasesV3200Response) GetMeta() GetKnowledgebasesV3200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebasesV3200Response) GetMetaOk() (*GetKnowledgebasesV3200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebasesV3200Response) SetMeta(v GetKnowledgebasesV3200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


