# GetServerDisks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**ServerDisks** | [**[]ServerDisk**](ServerDisk.md) |  | 

## Methods

### NewGetServerDisks200Response

`func NewGetServerDisks200Response(meta Meta, serverDisks []ServerDisk, ) *GetServerDisks200Response`

NewGetServerDisks200Response instantiates a new GetServerDisks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerDisks200ResponseWithDefaults

`func NewGetServerDisks200ResponseWithDefaults() *GetServerDisks200Response`

NewGetServerDisks200ResponseWithDefaults instantiates a new GetServerDisks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetServerDisks200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServerDisks200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServerDisks200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetServerDisks

`func (o *GetServerDisks200Response) GetServerDisks() []ServerDisk`

GetServerDisks returns the ServerDisks field if non-nil, zero value otherwise.

### GetServerDisksOk

`func (o *GetServerDisks200Response) GetServerDisksOk() (*[]ServerDisk, bool)`

GetServerDisksOk returns a tuple with the ServerDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDisks

`func (o *GetServerDisks200Response) SetServerDisks(v []ServerDisk)`

SetServerDisks sets ServerDisks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


