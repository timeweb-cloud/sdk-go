# CreateDatabaseCluster201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | [**DatabaseCluster**](DatabaseCluster.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateDatabaseCluster201Response

`func NewCreateDatabaseCluster201Response(db DatabaseCluster, responseId string, ) *CreateDatabaseCluster201Response`

NewCreateDatabaseCluster201Response instantiates a new CreateDatabaseCluster201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseCluster201ResponseWithDefaults

`func NewCreateDatabaseCluster201ResponseWithDefaults() *CreateDatabaseCluster201Response`

NewCreateDatabaseCluster201ResponseWithDefaults instantiates a new CreateDatabaseCluster201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *CreateDatabaseCluster201Response) GetDb() DatabaseCluster`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *CreateDatabaseCluster201Response) GetDbOk() (*DatabaseCluster, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *CreateDatabaseCluster201Response) SetDb(v DatabaseCluster)`

SetDb sets Db field to given value.


### GetResponseId

`func (o *CreateDatabaseCluster201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateDatabaseCluster201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateDatabaseCluster201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


