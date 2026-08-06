# GetDatabasePrivileges200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privileges** | [**[]DbPrivilegeGroup**](DbPrivilegeGroup.md) | Список привилегий, доступных для кластера базы данных | 

## Methods

### NewGetDatabasePrivileges200Response

`func NewGetDatabasePrivileges200Response(privileges []DbPrivilegeGroup, ) *GetDatabasePrivileges200Response`

NewGetDatabasePrivileges200Response instantiates a new GetDatabasePrivileges200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasePrivileges200ResponseWithDefaults

`func NewGetDatabasePrivileges200ResponseWithDefaults() *GetDatabasePrivileges200Response`

NewGetDatabasePrivileges200ResponseWithDefaults instantiates a new GetDatabasePrivileges200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivileges

`func (o *GetDatabasePrivileges200Response) GetPrivileges() []DbPrivilegeGroup`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *GetDatabasePrivileges200Response) GetPrivilegesOk() (*[]DbPrivilegeGroup, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *GetDatabasePrivileges200Response) SetPrivileges(v []DbPrivilegeGroup)`

SetPrivileges sets Privileges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


