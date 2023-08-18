# AddCountriesToAllowedList201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | [**AddCountries**](AddCountries.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewAddCountriesToAllowedList201Response

`func NewAddCountriesToAllowedList201Response(countries AddCountries, responseId string, ) *AddCountriesToAllowedList201Response`

NewAddCountriesToAllowedList201Response instantiates a new AddCountriesToAllowedList201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCountriesToAllowedList201ResponseWithDefaults

`func NewAddCountriesToAllowedList201ResponseWithDefaults() *AddCountriesToAllowedList201Response`

NewAddCountriesToAllowedList201ResponseWithDefaults instantiates a new AddCountriesToAllowedList201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *AddCountriesToAllowedList201Response) GetCountries() AddCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *AddCountriesToAllowedList201Response) GetCountriesOk() (*AddCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *AddCountriesToAllowedList201Response) SetCountries(v AddCountries)`

SetCountries sets Countries field to given value.


### GetResponseId

`func (o *AddCountriesToAllowedList201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddCountriesToAllowedList201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddCountriesToAllowedList201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


