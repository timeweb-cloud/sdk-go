# GetAllProjectResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]Vds**](Vds.md) |  | 
**Balancers** | [**[]Balancer**](Balancer.md) |  | 
**Buckets** | [**[]Bucket**](Bucket.md) |  | 
**Clusters** | [**[]Clusterk8s**](Clusterk8s.md) |  | 
**Databases** | [**[]Db**](Db.md) |  | 
**DedicatedServers** | [**[]DedicatedServer**](DedicatedServer.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetAllProjectResources200Response

`func NewGetAllProjectResources200Response(servers []Vds, balancers []Balancer, buckets []Bucket, clusters []Clusterk8s, databases []Db, dedicatedServers []DedicatedServer, meta Meta, ) *GetAllProjectResources200Response`

NewGetAllProjectResources200Response instantiates a new GetAllProjectResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllProjectResources200ResponseWithDefaults

`func NewGetAllProjectResources200ResponseWithDefaults() *GetAllProjectResources200Response`

NewGetAllProjectResources200ResponseWithDefaults instantiates a new GetAllProjectResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *GetAllProjectResources200Response) GetServers() []Vds`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetAllProjectResources200Response) GetServersOk() (*[]Vds, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetAllProjectResources200Response) SetServers(v []Vds)`

SetServers sets Servers field to given value.


### GetBalancers

`func (o *GetAllProjectResources200Response) GetBalancers() []Balancer`

GetBalancers returns the Balancers field if non-nil, zero value otherwise.

### GetBalancersOk

`func (o *GetAllProjectResources200Response) GetBalancersOk() (*[]Balancer, bool)`

GetBalancersOk returns a tuple with the Balancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancers

`func (o *GetAllProjectResources200Response) SetBalancers(v []Balancer)`

SetBalancers sets Balancers field to given value.


### GetBuckets

`func (o *GetAllProjectResources200Response) GetBuckets() []Bucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GetAllProjectResources200Response) GetBucketsOk() (*[]Bucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GetAllProjectResources200Response) SetBuckets(v []Bucket)`

SetBuckets sets Buckets field to given value.


### GetClusters

`func (o *GetAllProjectResources200Response) GetClusters() []Clusterk8s`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GetAllProjectResources200Response) GetClustersOk() (*[]Clusterk8s, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GetAllProjectResources200Response) SetClusters(v []Clusterk8s)`

SetClusters sets Clusters field to given value.


### GetDatabases

`func (o *GetAllProjectResources200Response) GetDatabases() []Db`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *GetAllProjectResources200Response) GetDatabasesOk() (*[]Db, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *GetAllProjectResources200Response) SetDatabases(v []Db)`

SetDatabases sets Databases field to given value.


### GetDedicatedServers

`func (o *GetAllProjectResources200Response) GetDedicatedServers() []DedicatedServer`

GetDedicatedServers returns the DedicatedServers field if non-nil, zero value otherwise.

### GetDedicatedServersOk

`func (o *GetAllProjectResources200Response) GetDedicatedServersOk() (*[]DedicatedServer, bool)`

GetDedicatedServersOk returns a tuple with the DedicatedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedServers

`func (o *GetAllProjectResources200Response) SetDedicatedServers(v []DedicatedServer)`

SetDedicatedServers sets DedicatedServers field to given value.


### GetMeta

`func (o *GetAllProjectResources200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAllProjectResources200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAllProjectResources200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


