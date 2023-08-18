# GetMailQuota200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | [**Quota**](Quota.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetMailQuota200Response

`func NewGetMailQuota200Response(quota Quota, responseId string, ) *GetMailQuota200Response`

NewGetMailQuota200Response instantiates a new GetMailQuota200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMailQuota200ResponseWithDefaults

`func NewGetMailQuota200ResponseWithDefaults() *GetMailQuota200Response`

NewGetMailQuota200ResponseWithDefaults instantiates a new GetMailQuota200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *GetMailQuota200Response) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetMailQuota200Response) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetMailQuota200Response) SetQuota(v Quota)`

SetQuota sets Quota field to given value.


### GetResponseId

`func (o *GetMailQuota200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetMailQuota200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetMailQuota200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


