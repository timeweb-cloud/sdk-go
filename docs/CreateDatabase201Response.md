# CreateDatabase201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | [**Db**](Db.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateDatabase201Response

`func NewCreateDatabase201Response(db Db, responseId string, ) *CreateDatabase201Response`

NewCreateDatabase201Response instantiates a new CreateDatabase201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabase201ResponseWithDefaults

`func NewCreateDatabase201ResponseWithDefaults() *CreateDatabase201Response`

NewCreateDatabase201ResponseWithDefaults instantiates a new CreateDatabase201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *CreateDatabase201Response) GetDb() Db`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *CreateDatabase201Response) GetDbOk() (*Db, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *CreateDatabase201Response) SetDb(v Db)`

SetDb sets Db field to given value.


### GetResponseId

`func (o *CreateDatabase201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateDatabase201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateDatabase201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


