# GetLinkCardPayment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmationUrl** | **string** | URL для подтверждения оплаты | 
**ResponseId** | **string** | ID запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetLinkCardPayment200Response

`func NewGetLinkCardPayment200Response(confirmationUrl string, responseId string, ) *GetLinkCardPayment200Response`

NewGetLinkCardPayment200Response instantiates a new GetLinkCardPayment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkCardPayment200ResponseWithDefaults

`func NewGetLinkCardPayment200ResponseWithDefaults() *GetLinkCardPayment200Response`

NewGetLinkCardPayment200ResponseWithDefaults instantiates a new GetLinkCardPayment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmationUrl

`func (o *GetLinkCardPayment200Response) GetConfirmationUrl() string`

GetConfirmationUrl returns the ConfirmationUrl field if non-nil, zero value otherwise.

### GetConfirmationUrlOk

`func (o *GetLinkCardPayment200Response) GetConfirmationUrlOk() (*string, bool)`

GetConfirmationUrlOk returns a tuple with the ConfirmationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationUrl

`func (o *GetLinkCardPayment200Response) SetConfirmationUrl(v string)`

SetConfirmationUrl sets ConfirmationUrl field to given value.


### GetResponseId

`func (o *GetLinkCardPayment200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetLinkCardPayment200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetLinkCardPayment200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


