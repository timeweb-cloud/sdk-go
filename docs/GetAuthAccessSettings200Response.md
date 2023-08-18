# GetAuthAccessSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsIpRestrictionsEnabled** | **bool** | Это логическое значение, которое показывает, включено ли ограничение доступа по IP-адресу. | 
**IsCountryRestrictionsEnabled** | **bool** | Это логическое значение, которое показывает, включено ли ограничение доступа по стране. | 
**WhiteList** | [**GetAuthAccessSettings200ResponseAllOfWhiteList**](GetAuthAccessSettings200ResponseAllOfWhiteList.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetAuthAccessSettings200Response

`func NewGetAuthAccessSettings200Response(isIpRestrictionsEnabled bool, isCountryRestrictionsEnabled bool, whiteList GetAuthAccessSettings200ResponseAllOfWhiteList, responseId string, ) *GetAuthAccessSettings200Response`

NewGetAuthAccessSettings200Response instantiates a new GetAuthAccessSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthAccessSettings200ResponseWithDefaults

`func NewGetAuthAccessSettings200ResponseWithDefaults() *GetAuthAccessSettings200Response`

NewGetAuthAccessSettings200ResponseWithDefaults instantiates a new GetAuthAccessSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsIpRestrictionsEnabled

`func (o *GetAuthAccessSettings200Response) GetIsIpRestrictionsEnabled() bool`

GetIsIpRestrictionsEnabled returns the IsIpRestrictionsEnabled field if non-nil, zero value otherwise.

### GetIsIpRestrictionsEnabledOk

`func (o *GetAuthAccessSettings200Response) GetIsIpRestrictionsEnabledOk() (*bool, bool)`

GetIsIpRestrictionsEnabledOk returns a tuple with the IsIpRestrictionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIpRestrictionsEnabled

`func (o *GetAuthAccessSettings200Response) SetIsIpRestrictionsEnabled(v bool)`

SetIsIpRestrictionsEnabled sets IsIpRestrictionsEnabled field to given value.


### GetIsCountryRestrictionsEnabled

`func (o *GetAuthAccessSettings200Response) GetIsCountryRestrictionsEnabled() bool`

GetIsCountryRestrictionsEnabled returns the IsCountryRestrictionsEnabled field if non-nil, zero value otherwise.

### GetIsCountryRestrictionsEnabledOk

`func (o *GetAuthAccessSettings200Response) GetIsCountryRestrictionsEnabledOk() (*bool, bool)`

GetIsCountryRestrictionsEnabledOk returns a tuple with the IsCountryRestrictionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCountryRestrictionsEnabled

`func (o *GetAuthAccessSettings200Response) SetIsCountryRestrictionsEnabled(v bool)`

SetIsCountryRestrictionsEnabled sets IsCountryRestrictionsEnabled field to given value.


### GetWhiteList

`func (o *GetAuthAccessSettings200Response) GetWhiteList() GetAuthAccessSettings200ResponseAllOfWhiteList`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *GetAuthAccessSettings200Response) GetWhiteListOk() (*GetAuthAccessSettings200ResponseAllOfWhiteList, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *GetAuthAccessSettings200Response) SetWhiteList(v GetAuthAccessSettings200ResponseAllOfWhiteList)`

SetWhiteList sets WhiteList field to given value.


### GetResponseId

`func (o *GetAuthAccessSettings200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetAuthAccessSettings200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetAuthAccessSettings200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


