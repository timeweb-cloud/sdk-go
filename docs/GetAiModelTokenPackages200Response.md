# GetAiModelTokenPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenPackages** | [**[]TokenPackage**](TokenPackage.md) |  | 
**Meta** | [**GetAiModelTokenPackages200ResponseMeta**](GetAiModelTokenPackages200ResponseMeta.md) |  | 

## Methods

### NewGetAiModelTokenPackages200Response

`func NewGetAiModelTokenPackages200Response(tokenPackages []TokenPackage, meta GetAiModelTokenPackages200ResponseMeta, ) *GetAiModelTokenPackages200Response`

NewGetAiModelTokenPackages200Response instantiates a new GetAiModelTokenPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiModelTokenPackages200ResponseWithDefaults

`func NewGetAiModelTokenPackages200ResponseWithDefaults() *GetAiModelTokenPackages200Response`

NewGetAiModelTokenPackages200ResponseWithDefaults instantiates a new GetAiModelTokenPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenPackages

`func (o *GetAiModelTokenPackages200Response) GetTokenPackages() []TokenPackage`

GetTokenPackages returns the TokenPackages field if non-nil, zero value otherwise.

### GetTokenPackagesOk

`func (o *GetAiModelTokenPackages200Response) GetTokenPackagesOk() (*[]TokenPackage, bool)`

GetTokenPackagesOk returns a tuple with the TokenPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackages

`func (o *GetAiModelTokenPackages200Response) SetTokenPackages(v []TokenPackage)`

SetTokenPackages sets TokenPackages field to given value.


### GetMeta

`func (o *GetAiModelTokenPackages200Response) GetMeta() GetAiModelTokenPackages200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAiModelTokenPackages200Response) GetMetaOk() (*GetAiModelTokenPackages200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAiModelTokenPackages200Response) SetMeta(v GetAiModelTokenPackages200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


