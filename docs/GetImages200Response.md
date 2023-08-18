# GetImages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Images** | [**[]ImageOutAPI**](ImageOutAPI.md) | Массив объектов Образ | 

## Methods

### NewGetImages200Response

`func NewGetImages200Response(responseId string, meta Meta, images []ImageOutAPI, ) *GetImages200Response`

NewGetImages200Response instantiates a new GetImages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImages200ResponseWithDefaults

`func NewGetImages200ResponseWithDefaults() *GetImages200Response`

NewGetImages200ResponseWithDefaults instantiates a new GetImages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetImages200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetImages200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetImages200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetImages200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetImages200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetImages200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetImages

`func (o *GetImages200Response) GetImages() []ImageOutAPI`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetImages200Response) GetImagesOk() (*[]ImageOutAPI, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetImages200Response) SetImages(v []ImageOutAPI)`

SetImages sets Images field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


