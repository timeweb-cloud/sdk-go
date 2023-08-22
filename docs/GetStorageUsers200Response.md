# GetStorageUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]BucketUser**](BucketUser.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetStorageUsers200Response

`func NewGetStorageUsers200Response(users []BucketUser, meta Meta, ) *GetStorageUsers200Response`

NewGetStorageUsers200Response instantiates a new GetStorageUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageUsers200ResponseWithDefaults

`func NewGetStorageUsers200ResponseWithDefaults() *GetStorageUsers200Response`

NewGetStorageUsers200ResponseWithDefaults instantiates a new GetStorageUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetStorageUsers200Response) GetUsers() []BucketUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetStorageUsers200Response) GetUsersOk() (*[]BucketUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetStorageUsers200Response) SetUsers(v []BucketUser)`

SetUsers sets Users field to given value.


### GetMeta

`func (o *GetStorageUsers200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetStorageUsers200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetStorageUsers200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


