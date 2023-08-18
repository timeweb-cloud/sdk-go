# CheckDomain200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDomainAvailable** | **bool** | Это логическое значение, которое показывает, доступен ли домен для регистрации. | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewCheckDomain200Response

`func NewCheckDomain200Response(isDomainAvailable bool, responseId string, ) *CheckDomain200Response`

NewCheckDomain200Response instantiates a new CheckDomain200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDomain200ResponseWithDefaults

`func NewCheckDomain200ResponseWithDefaults() *CheckDomain200Response`

NewCheckDomain200ResponseWithDefaults instantiates a new CheckDomain200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDomainAvailable

`func (o *CheckDomain200Response) GetIsDomainAvailable() bool`

GetIsDomainAvailable returns the IsDomainAvailable field if non-nil, zero value otherwise.

### GetIsDomainAvailableOk

`func (o *CheckDomain200Response) GetIsDomainAvailableOk() (*bool, bool)`

GetIsDomainAvailableOk returns a tuple with the IsDomainAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomainAvailable

`func (o *CheckDomain200Response) SetIsDomainAvailable(v bool)`

SetIsDomainAvailable sets IsDomainAvailable field to given value.


### GetResponseId

`func (o *CheckDomain200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CheckDomain200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CheckDomain200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


