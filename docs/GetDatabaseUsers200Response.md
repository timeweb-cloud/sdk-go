# GetDatabaseUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Admins** | [**[]DatabaseAdmin**](DatabaseAdmin.md) |  | 

## Methods

### NewGetDatabaseUsers200Response

`func NewGetDatabaseUsers200Response(meta Meta, admins []DatabaseAdmin, ) *GetDatabaseUsers200Response`

NewGetDatabaseUsers200Response instantiates a new GetDatabaseUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseUsers200ResponseWithDefaults

`func NewGetDatabaseUsers200ResponseWithDefaults() *GetDatabaseUsers200Response`

NewGetDatabaseUsers200ResponseWithDefaults instantiates a new GetDatabaseUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseUsers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseUsers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseUsers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetAdmins

`func (o *GetDatabaseUsers200Response) GetAdmins() []DatabaseAdmin`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *GetDatabaseUsers200Response) GetAdminsOk() (*[]DatabaseAdmin, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *GetDatabaseUsers200Response) SetAdmins(v []DatabaseAdmin)`

SetAdmins sets Admins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


