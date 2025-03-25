# DeleteServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | Хеш, который совместно с кодом авторизации надо будет отправить для удаления. | [optional] 
**IsMovedInQuarantine** | Pointer to **bool** | Флаг, указывающий на то, что сервис был перемещен в карантин или был удален немедленно. | [optional] 

## Methods

### NewDeleteServiceResponse

`func NewDeleteServiceResponse() *DeleteServiceResponse`

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

### HasHash

`func (o *DeleteServiceResponse) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIsMovedInQuarantine

`func (o *DeleteServiceResponse) GetIsMovedInQuarantine() bool`

GetIsMovedInQuarantine returns the IsMovedInQuarantine field if non-nil, zero value otherwise.

### GetIsMovedInQuarantineOk

`func (o *DeleteServiceResponse) GetIsMovedInQuarantineOk() (*bool, bool)`

GetIsMovedInQuarantineOk returns a tuple with the IsMovedInQuarantine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMovedInQuarantine

`func (o *DeleteServiceResponse) SetIsMovedInQuarantine(v bool)`

SetIsMovedInQuarantine sets IsMovedInQuarantine field to given value.

### HasIsMovedInQuarantine

`func (o *DeleteServiceResponse) HasIsMovedInQuarantine() bool`

HasIsMovedInQuarantine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


