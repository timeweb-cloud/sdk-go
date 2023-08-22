# GetProjectStorages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]Bucket**](Bucket.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetProjectStorages200Response

`func NewGetProjectStorages200Response(buckets []Bucket, meta Meta, ) *GetProjectStorages200Response`

NewGetProjectStorages200Response instantiates a new GetProjectStorages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectStorages200ResponseWithDefaults

`func NewGetProjectStorages200ResponseWithDefaults() *GetProjectStorages200Response`

NewGetProjectStorages200ResponseWithDefaults instantiates a new GetProjectStorages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *GetProjectStorages200Response) GetBuckets() []Bucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GetProjectStorages200Response) GetBucketsOk() (*[]Bucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GetProjectStorages200Response) SetBuckets(v []Bucket)`

SetBuckets sets Buckets field to given value.


### GetMeta

`func (o *GetProjectStorages200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetProjectStorages200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetProjectStorages200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


