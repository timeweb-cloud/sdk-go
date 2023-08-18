# CreateDatabaseInstance201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**DatabaseInstance**](DatabaseInstance.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateDatabaseInstance201Response

`func NewCreateDatabaseInstance201Response(instance DatabaseInstance, responseId string, ) *CreateDatabaseInstance201Response`

NewCreateDatabaseInstance201Response instantiates a new CreateDatabaseInstance201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseInstance201ResponseWithDefaults

`func NewCreateDatabaseInstance201ResponseWithDefaults() *CreateDatabaseInstance201Response`

NewCreateDatabaseInstance201ResponseWithDefaults instantiates a new CreateDatabaseInstance201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *CreateDatabaseInstance201Response) GetInstance() DatabaseInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CreateDatabaseInstance201Response) GetInstanceOk() (*DatabaseInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CreateDatabaseInstance201Response) SetInstance(v DatabaseInstance)`

SetInstance sets Instance field to given value.


### GetResponseId

`func (o *CreateDatabaseInstance201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateDatabaseInstance201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateDatabaseInstance201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


