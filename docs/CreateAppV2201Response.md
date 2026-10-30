# CreateAppV2201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**AppV2**](AppV2.md) |  | 
**ResponseId** | **string** | ID запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCreateAppV2201Response

`func NewCreateAppV2201Response(app AppV2, responseId string, ) *CreateAppV2201Response`

NewCreateAppV2201Response instantiates a new CreateAppV2201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppV2201ResponseWithDefaults

`func NewCreateAppV2201ResponseWithDefaults() *CreateAppV2201Response`

NewCreateAppV2201ResponseWithDefaults instantiates a new CreateAppV2201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *CreateAppV2201Response) GetApp() AppV2`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CreateAppV2201Response) GetAppOk() (*AppV2, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CreateAppV2201Response) SetApp(v AppV2)`

SetApp sets App field to given value.


### GetResponseId

`func (o *CreateAppV2201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateAppV2201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateAppV2201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


