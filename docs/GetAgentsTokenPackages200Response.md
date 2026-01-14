# GetAgentsTokenPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenPackages** | [**[]TokenPackage**](TokenPackage.md) |  | 
**Meta** | [**GetAgentsTokenPackages200ResponseMeta**](GetAgentsTokenPackages200ResponseMeta.md) |  | 

## Methods

### NewGetAgentsTokenPackages200Response

`func NewGetAgentsTokenPackages200Response(tokenPackages []TokenPackage, meta GetAgentsTokenPackages200ResponseMeta, ) *GetAgentsTokenPackages200Response`

NewGetAgentsTokenPackages200Response instantiates a new GetAgentsTokenPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgentsTokenPackages200ResponseWithDefaults

`func NewGetAgentsTokenPackages200ResponseWithDefaults() *GetAgentsTokenPackages200Response`

NewGetAgentsTokenPackages200ResponseWithDefaults instantiates a new GetAgentsTokenPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenPackages

`func (o *GetAgentsTokenPackages200Response) GetTokenPackages() []TokenPackage`

GetTokenPackages returns the TokenPackages field if non-nil, zero value otherwise.

### GetTokenPackagesOk

`func (o *GetAgentsTokenPackages200Response) GetTokenPackagesOk() (*[]TokenPackage, bool)`

GetTokenPackagesOk returns a tuple with the TokenPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPackages

`func (o *GetAgentsTokenPackages200Response) SetTokenPackages(v []TokenPackage)`

SetTokenPackages sets TokenPackages field to given value.


### GetMeta

`func (o *GetAgentsTokenPackages200Response) GetMeta() GetAgentsTokenPackages200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgentsTokenPackages200Response) GetMetaOk() (*GetAgentsTokenPackages200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgentsTokenPackages200Response) SetMeta(v GetAgentsTokenPackages200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


