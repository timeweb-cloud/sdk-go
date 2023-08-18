# GetFinances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finances** | [**Finances**](Finances.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetFinances200Response

`func NewGetFinances200Response(finances Finances, responseId string, ) *GetFinances200Response`

NewGetFinances200Response instantiates a new GetFinances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinances200ResponseWithDefaults

`func NewGetFinances200ResponseWithDefaults() *GetFinances200Response`

NewGetFinances200ResponseWithDefaults instantiates a new GetFinances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinances

`func (o *GetFinances200Response) GetFinances() Finances`

GetFinances returns the Finances field if non-nil, zero value otherwise.

### GetFinancesOk

`func (o *GetFinances200Response) GetFinancesOk() (*Finances, bool)`

GetFinancesOk returns a tuple with the Finances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinances

`func (o *GetFinances200Response) SetFinances(v Finances)`

SetFinances sets Finances field to given value.


### GetResponseId

`func (o *GetFinances200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetFinances200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetFinances200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


