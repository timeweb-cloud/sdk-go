# GetCommits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Commits** | [**[]Commit**](Commit.md) |  | 

## Methods

### NewGetCommits200Response

`func NewGetCommits200Response(meta Meta, commits []Commit, ) *GetCommits200Response`

NewGetCommits200Response instantiates a new GetCommits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommits200ResponseWithDefaults

`func NewGetCommits200ResponseWithDefaults() *GetCommits200Response`

NewGetCommits200ResponseWithDefaults instantiates a new GetCommits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetCommits200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCommits200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCommits200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetCommits

`func (o *GetCommits200Response) GetCommits() []Commit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GetCommits200Response) GetCommitsOk() (*[]Commit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GetCommits200Response) SetCommits(v []Commit)`

SetCommits sets Commits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


