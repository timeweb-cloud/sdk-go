# BucketWebsiteConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Включено ли сайтовое хранилище. | [optional] 
**IndexPage** | Pointer to **string** | Страница сайта. | [optional] 
**ErrorPages** | Pointer to [**[]BucketWebsiteConfigErrorPagesInner**](BucketWebsiteConfigErrorPagesInner.md) | Страницы ошибок. | [optional] 

## Methods

### NewBucketWebsiteConfig

`func NewBucketWebsiteConfig() *BucketWebsiteConfig`

NewBucketWebsiteConfig instantiates a new BucketWebsiteConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWebsiteConfigWithDefaults

`func NewBucketWebsiteConfigWithDefaults() *BucketWebsiteConfig`

NewBucketWebsiteConfigWithDefaults instantiates a new BucketWebsiteConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BucketWebsiteConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BucketWebsiteConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BucketWebsiteConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BucketWebsiteConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIndexPage

`func (o *BucketWebsiteConfig) GetIndexPage() string`

GetIndexPage returns the IndexPage field if non-nil, zero value otherwise.

### GetIndexPageOk

`func (o *BucketWebsiteConfig) GetIndexPageOk() (*string, bool)`

GetIndexPageOk returns a tuple with the IndexPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPage

`func (o *BucketWebsiteConfig) SetIndexPage(v string)`

SetIndexPage sets IndexPage field to given value.

### HasIndexPage

`func (o *BucketWebsiteConfig) HasIndexPage() bool`

HasIndexPage returns a boolean if a field has been set.

### GetErrorPages

`func (o *BucketWebsiteConfig) GetErrorPages() []BucketWebsiteConfigErrorPagesInner`

GetErrorPages returns the ErrorPages field if non-nil, zero value otherwise.

### GetErrorPagesOk

`func (o *BucketWebsiteConfig) GetErrorPagesOk() (*[]BucketWebsiteConfigErrorPagesInner, bool)`

GetErrorPagesOk returns a tuple with the ErrorPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPages

`func (o *BucketWebsiteConfig) SetErrorPages(v []BucketWebsiteConfigErrorPagesInner)`

SetErrorPages sets ErrorPages field to given value.

### HasErrorPages

`func (o *BucketWebsiteConfig) HasErrorPages() bool`

HasErrorPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


