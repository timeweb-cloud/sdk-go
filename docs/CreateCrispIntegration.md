# CreateCrispIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** | Website ID сайта в Crisp (Settings → Website Settings) | 
**TokenIdentifier** | **string** | Identifier API-токена Crisp (Settings → Workspace Settings → Advanced configuration → API Token) | 
**TokenKey** | **string** | Key API-токена Crisp | 

## Methods

### NewCreateCrispIntegration

`func NewCreateCrispIntegration(websiteId string, tokenIdentifier string, tokenKey string, ) *CreateCrispIntegration`

NewCreateCrispIntegration instantiates a new CreateCrispIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCrispIntegrationWithDefaults

`func NewCreateCrispIntegrationWithDefaults() *CreateCrispIntegration`

NewCreateCrispIntegrationWithDefaults instantiates a new CreateCrispIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *CreateCrispIntegration) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *CreateCrispIntegration) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *CreateCrispIntegration) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetTokenIdentifier

`func (o *CreateCrispIntegration) GetTokenIdentifier() string`

GetTokenIdentifier returns the TokenIdentifier field if non-nil, zero value otherwise.

### GetTokenIdentifierOk

`func (o *CreateCrispIntegration) GetTokenIdentifierOk() (*string, bool)`

GetTokenIdentifierOk returns a tuple with the TokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdentifier

`func (o *CreateCrispIntegration) SetTokenIdentifier(v string)`

SetTokenIdentifier sets TokenIdentifier field to given value.


### GetTokenKey

`func (o *CreateCrispIntegration) GetTokenKey() string`

GetTokenKey returns the TokenKey field if non-nil, zero value otherwise.

### GetTokenKeyOk

`func (o *CreateCrispIntegration) GetTokenKeyOk() (*string, bool)`

GetTokenKeyOk returns a tuple with the TokenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenKey

`func (o *CreateCrispIntegration) SetTokenKey(v string)`

SetTokenKey sets TokenKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


