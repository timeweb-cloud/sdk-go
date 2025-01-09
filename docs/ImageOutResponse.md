# ImageOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Image** | [**Image**](Image.md) |  | 

## Methods

### NewImageOutResponse

`func NewImageOutResponse(image Image, ) *ImageOutResponse`

NewImageOutResponse instantiates a new ImageOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageOutResponseWithDefaults

`func NewImageOutResponseWithDefaults() *ImageOutResponse`

NewImageOutResponseWithDefaults instantiates a new ImageOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *ImageOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *ImageOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *ImageOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *ImageOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetImage

`func (o *ImageOutResponse) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageOutResponse) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageOutResponse) SetImage(v Image)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


