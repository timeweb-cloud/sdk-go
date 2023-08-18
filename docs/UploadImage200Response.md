# UploadImage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**UploadSuccessful** | [**UploadSuccessful**](UploadSuccessful.md) |  | 

## Methods

### NewUploadImage200Response

`func NewUploadImage200Response(responseId string, uploadSuccessful UploadSuccessful, ) *UploadImage200Response`

NewUploadImage200Response instantiates a new UploadImage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadImage200ResponseWithDefaults

`func NewUploadImage200ResponseWithDefaults() *UploadImage200Response`

NewUploadImage200ResponseWithDefaults instantiates a new UploadImage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *UploadImage200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *UploadImage200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *UploadImage200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetUploadSuccessful

`func (o *UploadImage200Response) GetUploadSuccessful() UploadSuccessful`

GetUploadSuccessful returns the UploadSuccessful field if non-nil, zero value otherwise.

### GetUploadSuccessfulOk

`func (o *UploadImage200Response) GetUploadSuccessfulOk() (*UploadSuccessful, bool)`

GetUploadSuccessfulOk returns a tuple with the UploadSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSuccessful

`func (o *UploadImage200Response) SetUploadSuccessful(v UploadSuccessful)`

SetUploadSuccessful sets UploadSuccessful field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


