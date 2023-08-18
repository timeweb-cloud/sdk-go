# CreateStorage201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | [**Bucket**](Bucket.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateStorage201Response

`func NewCreateStorage201Response(bucket Bucket, responseId string, ) *CreateStorage201Response`

NewCreateStorage201Response instantiates a new CreateStorage201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorage201ResponseWithDefaults

`func NewCreateStorage201ResponseWithDefaults() *CreateStorage201Response`

NewCreateStorage201ResponseWithDefaults instantiates a new CreateStorage201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *CreateStorage201Response) GetBucket() Bucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CreateStorage201Response) GetBucketOk() (*Bucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CreateStorage201Response) SetBucket(v Bucket)`

SetBucket sets Bucket field to given value.


### GetResponseId

`func (o *CreateStorage201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateStorage201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateStorage201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


