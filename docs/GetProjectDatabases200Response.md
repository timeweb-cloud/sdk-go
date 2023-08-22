# GetProjectDatabases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Databases** | [**[]Db**](Db.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetProjectDatabases200Response

`func NewGetProjectDatabases200Response(databases []Db, meta Meta, ) *GetProjectDatabases200Response`

NewGetProjectDatabases200Response instantiates a new GetProjectDatabases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectDatabases200ResponseWithDefaults

`func NewGetProjectDatabases200ResponseWithDefaults() *GetProjectDatabases200Response`

NewGetProjectDatabases200ResponseWithDefaults instantiates a new GetProjectDatabases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabases

`func (o *GetProjectDatabases200Response) GetDatabases() []Db`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *GetProjectDatabases200Response) GetDatabasesOk() (*[]Db, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *GetProjectDatabases200Response) SetDatabases(v []Db)`

SetDatabases sets Databases field to given value.


### GetMeta

`func (o *GetProjectDatabases200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProjectDatabases200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProjectDatabases200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


