# GetAgentSkillCatalog200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillCatalogItems** | [**[]SkillCatalogItem**](SkillCatalogItem.md) |  | 
**Meta** | [**GetAgentSkillCatalog200ResponseMeta**](GetAgentSkillCatalog200ResponseMeta.md) |  | 

## Methods

### NewGetAgentSkillCatalog200Response

`func NewGetAgentSkillCatalog200Response(skillCatalogItems []SkillCatalogItem, meta GetAgentSkillCatalog200ResponseMeta, ) *GetAgentSkillCatalog200Response`

NewGetAgentSkillCatalog200Response instantiates a new GetAgentSkillCatalog200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentSkillCatalog200ResponseWithDefaults

`func NewGetAgentSkillCatalog200ResponseWithDefaults() *GetAgentSkillCatalog200Response`

NewGetAgentSkillCatalog200ResponseWithDefaults instantiates a new GetAgentSkillCatalog200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillCatalogItems

`func (o *GetAgentSkillCatalog200Response) GetSkillCatalogItems() []SkillCatalogItem`

GetSkillCatalogItems returns the SkillCatalogItems field if non-nil, zero value otherwise.

### GetSkillCatalogItemsOk

`func (o *GetAgentSkillCatalog200Response) GetSkillCatalogItemsOk() (*[]SkillCatalogItem, bool)`

GetSkillCatalogItemsOk returns a tuple with the SkillCatalogItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillCatalogItems

`func (o *GetAgentSkillCatalog200Response) SetSkillCatalogItems(v []SkillCatalogItem)`

SetSkillCatalogItems sets SkillCatalogItems field to given value.


### GetMeta

`func (o *GetAgentSkillCatalog200Response) GetMeta() GetAgentSkillCatalog200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentSkillCatalog200Response) GetMetaOk() (*GetAgentSkillCatalog200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentSkillCatalog200Response) SetMeta(v GetAgentSkillCatalog200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


