# UpdateDatabaseCluster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | [**DatabaseCluster**](DatabaseCluster.md) |  | 
**Warning** | Pointer to **string** | Предупреждение о частично примененных изменениях. Появляется, если в &#x60;config_parameters&#x60; были переданы значения, не прошедшие валидацию: такие параметры отбрасываются, остальные изменения применяются. | [optional] 

## Methods

### NewUpdateDatabaseCluster200Response

`func NewUpdateDatabaseCluster200Response(db DatabaseCluster, ) *UpdateDatabaseCluster200Response`

NewUpdateDatabaseCluster200Response instantiates a new UpdateDatabaseCluster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatabaseCluster200ResponseWithDefaults

`func NewUpdateDatabaseCluster200ResponseWithDefaults() *UpdateDatabaseCluster200Response`

NewUpdateDatabaseCluster200ResponseWithDefaults instantiates a new UpdateDatabaseCluster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *UpdateDatabaseCluster200Response) GetDb() DatabaseCluster`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *UpdateDatabaseCluster200Response) GetDbOk() (*DatabaseCluster, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *UpdateDatabaseCluster200Response) SetDb(v DatabaseCluster)`

SetDb sets Db field to given value.


### GetWarning

`func (o *UpdateDatabaseCluster200Response) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *UpdateDatabaseCluster200Response) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *UpdateDatabaseCluster200Response) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *UpdateDatabaseCluster200Response) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


