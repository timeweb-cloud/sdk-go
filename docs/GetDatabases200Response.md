# GetDatabases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Dbs** | [**[]Db**](Db.md) |  | 

## Methods

### NewGetDatabases200Response

`func NewGetDatabases200Response(meta Meta, dbs []Db, ) *GetDatabases200Response`

NewGetDatabases200Response instantiates a new GetDatabases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabases200ResponseWithDefaults

`func NewGetDatabases200ResponseWithDefaults() *GetDatabases200Response`

NewGetDatabases200ResponseWithDefaults instantiates a new GetDatabases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabases200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabases200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabases200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDbs

`func (o *GetDatabases200Response) GetDbs() []Db`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *GetDatabases200Response) GetDbsOk() (*[]Db, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *GetDatabases200Response) SetDbs(v []Db)`

SetDbs sets Dbs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


