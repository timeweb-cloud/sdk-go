# DeleteCountriesFromAllowedList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | [**RemoveCountries**](RemoveCountries.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewDeleteCountriesFromAllowedList200Response

`func NewDeleteCountriesFromAllowedList200Response(countries RemoveCountries, responseId string, ) *DeleteCountriesFromAllowedList200Response`

NewDeleteCountriesFromAllowedList200Response instantiates a new DeleteCountriesFromAllowedList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCountriesFromAllowedList200ResponseWithDefaults

`func NewDeleteCountriesFromAllowedList200ResponseWithDefaults() *DeleteCountriesFromAllowedList200Response`

NewDeleteCountriesFromAllowedList200ResponseWithDefaults instantiates a new DeleteCountriesFromAllowedList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *DeleteCountriesFromAllowedList200Response) GetCountries() RemoveCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *DeleteCountriesFromAllowedList200Response) GetCountriesOk() (*RemoveCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *DeleteCountriesFromAllowedList200Response) SetCountries(v RemoveCountries)`

SetCountries sets Countries field to given value.


### GetResponseId

`func (o *DeleteCountriesFromAllowedList200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *DeleteCountriesFromAllowedList200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *DeleteCountriesFromAllowedList200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


