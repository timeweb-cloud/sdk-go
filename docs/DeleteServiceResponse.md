# DeleteServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Хеш, который совместно с кодом авторизации надо будет отправить для удаления. | 

## Methods

### NewDeleteServiceResponse

`func NewDeleteServiceResponse(hash string, ) *DeleteServiceResponse`

NewDeleteServiceResponse instantiates a new DeleteServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServiceResponseWithDefaults

`func NewDeleteServiceResponseWithDefaults() *DeleteServiceResponse`

NewDeleteServiceResponseWithDefaults instantiates a new DeleteServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *DeleteServiceResponse) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DeleteServiceResponse) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DeleteServiceResponse) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


