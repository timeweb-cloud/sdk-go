# GetBranches200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Branches** | [**[]Branch**](Branch.md) |  | 

## Methods

### NewGetBranches200Response

`func NewGetBranches200Response(meta Meta, branches []Branch, ) *GetBranches200Response`

NewGetBranches200Response instantiates a new GetBranches200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBranches200ResponseWithDefaults

`func NewGetBranches200ResponseWithDefaults() *GetBranches200Response`

NewGetBranches200ResponseWithDefaults instantiates a new GetBranches200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBranches200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBranches200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBranches200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetBranches

`func (o *GetBranches200Response) GetBranches() []Branch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *GetBranches200Response) GetBranchesOk() (*[]Branch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *GetBranches200Response) SetBranches(v []Branch)`

SetBranches sets Branches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


