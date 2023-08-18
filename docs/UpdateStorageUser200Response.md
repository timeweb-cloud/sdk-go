# UpdateStorageUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**BucketUser**](BucketUser.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewUpdateStorageUser200Response

`func NewUpdateStorageUser200Response(user BucketUser, responseId string, ) *UpdateStorageUser200Response`

NewUpdateStorageUser200Response instantiates a new UpdateStorageUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageUser200ResponseWithDefaults

`func NewUpdateStorageUser200ResponseWithDefaults() *UpdateStorageUser200Response`

NewUpdateStorageUser200ResponseWithDefaults instantiates a new UpdateStorageUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UpdateStorageUser200Response) GetUser() BucketUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateStorageUser200Response) GetUserOk() (*BucketUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateStorageUser200Response) SetUser(v BucketUser)`

SetUser sets User field to given value.


### GetResponseId

`func (o *UpdateStorageUser200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *UpdateStorageUser200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *UpdateStorageUser200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


