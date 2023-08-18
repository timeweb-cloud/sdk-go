# GetTLD200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopLevelDomain** | [**TopLevelDomain**](TopLevelDomain.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetTLD200Response

`func NewGetTLD200Response(topLevelDomain TopLevelDomain, responseId string, ) *GetTLD200Response`

NewGetTLD200Response instantiates a new GetTLD200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTLD200ResponseWithDefaults

`func NewGetTLD200ResponseWithDefaults() *GetTLD200Response`

NewGetTLD200ResponseWithDefaults instantiates a new GetTLD200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopLevelDomain

`func (o *GetTLD200Response) GetTopLevelDomain() TopLevelDomain`

GetTopLevelDomain returns the TopLevelDomain field if non-nil, zero value otherwise.

### GetTopLevelDomainOk

`func (o *GetTLD200Response) GetTopLevelDomainOk() (*TopLevelDomain, bool)`

GetTopLevelDomainOk returns a tuple with the TopLevelDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelDomain

`func (o *GetTLD200Response) SetTopLevelDomain(v TopLevelDomain)`

SetTopLevelDomain sets TopLevelDomain field to given value.


### GetResponseId

`func (o *GetTLD200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetTLD200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetTLD200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


