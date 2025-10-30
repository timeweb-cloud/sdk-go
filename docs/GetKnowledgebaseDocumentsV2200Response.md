# GetKnowledgebaseDocumentsV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnowledgebaseDocuments** | [**[]Document**](Document.md) | Список документов | 
**Meta** | [**GetKnowledgebaseDocumentsV2200ResponseMeta**](GetKnowledgebaseDocumentsV2200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebaseDocumentsV2200Response

`func NewGetKnowledgebaseDocumentsV2200Response(knowledgebaseDocuments []Document, meta GetKnowledgebaseDocumentsV2200ResponseMeta, ) *GetKnowledgebaseDocumentsV2200Response`

NewGetKnowledgebaseDocumentsV2200Response instantiates a new GetKnowledgebaseDocumentsV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebaseDocumentsV2200ResponseWithDefaults

`func NewGetKnowledgebaseDocumentsV2200ResponseWithDefaults() *GetKnowledgebaseDocumentsV2200Response`

NewGetKnowledgebaseDocumentsV2200ResponseWithDefaults instantiates a new GetKnowledgebaseDocumentsV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgebaseDocuments

`func (o *GetKnowledgebaseDocumentsV2200Response) GetKnowledgebaseDocuments() []Document`

GetKnowledgebaseDocuments returns the KnowledgebaseDocuments field if non-nil, zero value otherwise.

### GetKnowledgebaseDocumentsOk

`func (o *GetKnowledgebaseDocumentsV2200Response) GetKnowledgebaseDocumentsOk() (*[]Document, bool)`

GetKnowledgebaseDocumentsOk returns a tuple with the KnowledgebaseDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebaseDocuments

`func (o *GetKnowledgebaseDocumentsV2200Response) SetKnowledgebaseDocuments(v []Document)`

SetKnowledgebaseDocuments sets KnowledgebaseDocuments field to given value.


### GetMeta

`func (o *GetKnowledgebaseDocumentsV2200Response) GetMeta() GetKnowledgebaseDocumentsV2200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebaseDocumentsV2200Response) GetMetaOk() (*GetKnowledgebaseDocumentsV2200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebaseDocumentsV2200Response) SetMeta(v GetKnowledgebaseDocumentsV2200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


