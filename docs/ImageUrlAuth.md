# ImageUrlAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Токен доступа к API облачного хранилища | 
**RefreshToken** | Pointer to **string** | Токен обновления доступов к API | [optional] 
**Expiry** | Pointer to **time.Time** | Время истечения работы токена доступа | [optional] 
**TokenType** | Pointer to **string** | Тип токена доступа | [optional] [default to "Bearer"]

## Methods

### NewImageUrlAuth

`func NewImageUrlAuth(accessToken string, ) *ImageUrlAuth`

NewImageUrlAuth instantiates a new ImageUrlAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageUrlAuthWithDefaults

`func NewImageUrlAuthWithDefaults() *ImageUrlAuth`

NewImageUrlAuthWithDefaults instantiates a new ImageUrlAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ImageUrlAuth) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ImageUrlAuth) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ImageUrlAuth) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *ImageUrlAuth) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ImageUrlAuth) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ImageUrlAuth) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ImageUrlAuth) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiry

`func (o *ImageUrlAuth) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ImageUrlAuth) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ImageUrlAuth) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ImageUrlAuth) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetTokenType

`func (o *ImageUrlAuth) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ImageUrlAuth) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ImageUrlAuth) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *ImageUrlAuth) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


