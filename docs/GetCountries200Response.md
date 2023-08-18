# GetCountries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | **map[string]interface{}** | Список стран, приходит в виде объекта, где ключ - код страны в формате Alpha-2 ISO 3166-1, а значение - название страны в удобочитаемом формате. | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetCountries200Response

`func NewGetCountries200Response(countries map[string]interface{}, responseId string, ) *GetCountries200Response`

NewGetCountries200Response instantiates a new GetCountries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCountries200ResponseWithDefaults

`func NewGetCountries200ResponseWithDefaults() *GetCountries200Response`

NewGetCountries200ResponseWithDefaults instantiates a new GetCountries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *GetCountries200Response) GetCountries() map[string]interface{}`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GetCountries200Response) GetCountriesOk() (*map[string]interface{}, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GetCountries200Response) SetCountries(v map[string]interface{})`

SetCountries sets Countries field to given value.


### GetResponseId

`func (o *GetCountries200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetCountries200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetCountries200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


