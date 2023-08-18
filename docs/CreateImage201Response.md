# CreateImage201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Image** | [**ImageOutAPI**](ImageOutAPI.md) |  | 

## Methods

### NewCreateImage201Response

`func NewCreateImage201Response(responseId string, image ImageOutAPI, ) *CreateImage201Response`

NewCreateImage201Response instantiates a new CreateImage201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImage201ResponseWithDefaults

`func NewCreateImage201ResponseWithDefaults() *CreateImage201Response`

NewCreateImage201ResponseWithDefaults instantiates a new CreateImage201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *CreateImage201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateImage201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateImage201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetImage

`func (o *CreateImage201Response) GetImage() ImageOutAPI`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateImage201Response) GetImageOk() (*ImageOutAPI, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateImage201Response) SetImage(v ImageOutAPI)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


