# ImageUrlIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**URLType**](URLType.md) |  | [optional] 
**Filename** | Pointer to **string** | Имя файла для загрузки в облачное хранилище. | [optional] 
**Auth** | Pointer to [**ImageUrlAuth**](ImageUrlAuth.md) |  | [optional] 

## Methods

### NewImageUrlIn

`func NewImageUrlIn() *ImageUrlIn`

NewImageUrlIn instantiates a new ImageUrlIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageUrlInWithDefaults

`func NewImageUrlInWithDefaults() *ImageUrlIn`

NewImageUrlInWithDefaults instantiates a new ImageUrlIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImageUrlIn) GetType() URLType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageUrlIn) GetTypeOk() (*URLType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageUrlIn) SetType(v URLType)`

SetType sets Type field to given value.

### HasType

`func (o *ImageUrlIn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFilename

`func (o *ImageUrlIn) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ImageUrlIn) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ImageUrlIn) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ImageUrlIn) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetAuth

`func (o *ImageUrlIn) GetAuth() ImageUrlAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ImageUrlIn) GetAuthOk() (*ImageUrlAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ImageUrlIn) SetAuth(v ImageUrlAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ImageUrlIn) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


