# GetKnowledgebasesTokenPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenPackages** | [**[]TokenPackage**](TokenPackage.md) |  | 
**Meta** | [**GetKnowledgebasesTokenPackages200ResponseMeta**](GetKnowledgebasesTokenPackages200ResponseMeta.md) |  | 

## Methods

### NewGetKnowledgebasesTokenPackages200Response

`func NewGetKnowledgebasesTokenPackages200Response(tokenPackages []TokenPackage, meta GetKnowledgebasesTokenPackages200ResponseMeta, ) *GetKnowledgebasesTokenPackages200Response`

NewGetKnowledgebasesTokenPackages200Response instantiates a new GetKnowledgebasesTokenPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKnowledgebasesTokenPackages200ResponseWithDefaults

`func NewGetKnowledgebasesTokenPackages200ResponseWithDefaults() *GetKnowledgebasesTokenPackages200Response`

NewGetKnowledgebasesTokenPackages200ResponseWithDefaults instantiates a new GetKnowledgebasesTokenPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenPackages

`func (o *GetKnowledgebasesTokenPackages200Response) GetTokenPackages() []TokenPackage`

GetTokenPackages returns the TokenPackages field if non-nil, zero value otherwise.

### GetTokenPackagesOk

`func (o *GetKnowledgebasesTokenPackages200Response) GetTokenPackagesOk() (*[]TokenPackage, bool)`

GetTokenPackagesOk returns a tuple with the TokenPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackages

`func (o *GetKnowledgebasesTokenPackages200Response) SetTokenPackages(v []TokenPackage)`

SetTokenPackages sets TokenPackages field to given value.


### GetMeta

`func (o *GetKnowledgebasesTokenPackages200Response) GetMeta() GetKnowledgebasesTokenPackages200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKnowledgebasesTokenPackages200Response) GetMetaOk() (*GetKnowledgebasesTokenPackages200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKnowledgebasesTokenPackages200Response) SetMeta(v GetKnowledgebasesTokenPackages200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


