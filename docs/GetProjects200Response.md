# GetProjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | [**[]Project**](Project.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetProjects200Response

`func NewGetProjects200Response(projects []Project, meta Meta, ) *GetProjects200Response`

NewGetProjects200Response instantiates a new GetProjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjects200ResponseWithDefaults

`func NewGetProjects200ResponseWithDefaults() *GetProjects200Response`

NewGetProjects200ResponseWithDefaults instantiates a new GetProjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *GetProjects200Response) GetProjects() []Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetProjects200Response) GetProjectsOk() (*[]Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetProjects200Response) SetProjects(v []Project)`

SetProjects sets Projects field to given value.


### GetMeta

`func (o *GetProjects200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProjects200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProjects200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


