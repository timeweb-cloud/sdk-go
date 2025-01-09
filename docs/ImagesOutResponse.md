# ImagesOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Meta** | [**Meta**](Meta.md) |  | 
**Images** | [**[]Image**](Image.md) |  | 

## Methods

### NewImagesOutResponse

`func NewImagesOutResponse(meta Meta, images []Image, ) *ImagesOutResponse`

NewImagesOutResponse instantiates a new ImagesOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagesOutResponseWithDefaults

`func NewImagesOutResponseWithDefaults() *ImagesOutResponse`

NewImagesOutResponseWithDefaults instantiates a new ImagesOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *ImagesOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *ImagesOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *ImagesOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *ImagesOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *ImagesOutResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ImagesOutResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ImagesOutResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetImages

`func (o *ImagesOutResponse) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ImagesOutResponse) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ImagesOutResponse) SetImages(v []Image)`

SetImages sets Images field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


