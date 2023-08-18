# GetDatabaseInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**Instances** | [**[]DatabaseInstance**](DatabaseInstance.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetDatabaseInstances200Response

`func NewGetDatabaseInstances200Response(meta Meta, instances []DatabaseInstance, responseId string, ) *GetDatabaseInstances200Response`

NewGetDatabaseInstances200Response instantiates a new GetDatabaseInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabaseInstances200ResponseWithDefaults

`func NewGetDatabaseInstances200ResponseWithDefaults() *GetDatabaseInstances200Response`

NewGetDatabaseInstances200ResponseWithDefaults instantiates a new GetDatabaseInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDatabaseInstances200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDatabaseInstances200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDatabaseInstances200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetInstances

`func (o *GetDatabaseInstances200Response) GetInstances() []DatabaseInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetDatabaseInstances200Response) GetInstancesOk() (*[]DatabaseInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetDatabaseInstances200Response) SetInstances(v []DatabaseInstance)`

SetInstances sets Instances field to given value.


### GetResponseId

`func (o *GetDatabaseInstances200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetDatabaseInstances200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetDatabaseInstances200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


