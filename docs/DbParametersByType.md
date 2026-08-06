# DbParametersByType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mysql5** | **[]string** | Параметры, доступные для кластеров типа &#x60;mysql5&#x60;. | 
**Mysql** | **[]string** | Параметры, доступные для кластеров типа &#x60;mysql&#x60;. | 
**Mysql84** | **[]string** | Параметры, доступные для кластеров типа &#x60;mysql8_4&#x60;. | 
**Postgres** | **[]string** | Параметры, доступные для кластеров типа &#x60;postgres&#x60; (PostgreSQL 13). | 
**Postgres14** | **[]string** | Параметры, доступные для кластеров типа &#x60;postgres14&#x60;. | 
**Postgres15** | **[]string** | Параметры, доступные для кластеров типа &#x60;postgres15&#x60;. | 
**Postgres16** | **[]string** | Параметры, доступные для кластеров типа &#x60;postgres16&#x60;. | 
**Postgres17** | **[]string** | Параметры, доступные для кластеров типа &#x60;postgres17&#x60;. | 
**Postgres18** | **[]string** | Параметры, доступные для кластеров типа &#x60;postgres18&#x60;. Набор отличается от предыдущих версий PostgreSQL — например, добавлены &#x60;io_method&#x60; и &#x60;io_workers&#x60;. | 
**Redis** | **[]string** | Параметры, доступные для кластеров типа &#x60;redis&#x60;. | 
**Redis7** | **[]string** | Параметры, доступные для кластеров типа &#x60;redis7&#x60;. | 
**Redis81** | **[]string** | Параметры, доступные для кластеров типа &#x60;redis8_1&#x60;. | 
**Valkey** | **[]string** | Параметры, доступные для кластеров типа &#x60;valkey&#x60;. | 
**Valkey7** | **[]string** | Параметры, доступные для кластеров типа &#x60;valkey7&#x60;. | 
**Valkey81** | **[]string** | Параметры, доступные для кластеров типа &#x60;valkey8_1&#x60;. | 
**Valkey91** | **[]string** | Параметры, доступные для кластеров типа &#x60;valkey9_1&#x60;. | 
**Mongodb4** | **[]string** | Для кластеров типа &#x60;mongodb4&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Mongodb** | **[]string** | Для кластеров типа &#x60;mongodb&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Mongodb6** | **[]string** | Для кластеров типа &#x60;mongodb6&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Mongodb7** | **[]string** | Для кластеров типа &#x60;mongodb7&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Mongodb80** | **[]string** | Для кластеров типа &#x60;mongodb8_0&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Opensearch** | **[]string** | Для кластеров типа &#x60;opensearch&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Opensearch219** | **[]string** | Для кластеров типа &#x60;opensearch2_19&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Clickhouse** | **[]string** | Для кластеров типа &#x60;clickhouse&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Clickhouse24** | **[]string** | Для кластеров типа &#x60;clickhouse24&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Clickhouse25** | **[]string** | Для кластеров типа &#x60;clickhouse25&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Kafka** | **[]string** | Для кластеров типа &#x60;kafka&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Rabbitmq** | **[]string** | Для кластеров типа &#x60;rabbitmq&#x60; настраиваемых параметров нет — всегда пустой массив. | 
**Rabbitmq40** | **[]string** | Для кластеров типа &#x60;rabbitmq4_0&#x60; настраиваемых параметров нет — всегда пустой массив. | 

## Methods

### NewDbParametersByType

`func NewDbParametersByType(mysql5 []string, mysql []string, mysql84 []string, postgres []string, postgres14 []string, postgres15 []string, postgres16 []string, postgres17 []string, postgres18 []string, redis []string, redis7 []string, redis81 []string, valkey []string, valkey7 []string, valkey81 []string, valkey91 []string, mongodb4 []string, mongodb []string, mongodb6 []string, mongodb7 []string, mongodb80 []string, opensearch []string, opensearch219 []string, clickhouse []string, clickhouse24 []string, clickhouse25 []string, kafka []string, rabbitmq []string, rabbitmq40 []string, ) *DbParametersByType`

NewDbParametersByType instantiates a new DbParametersByType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbParametersByTypeWithDefaults

`func NewDbParametersByTypeWithDefaults() *DbParametersByType`

NewDbParametersByTypeWithDefaults instantiates a new DbParametersByType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMysql5

`func (o *DbParametersByType) GetMysql5() []string`

GetMysql5 returns the Mysql5 field if non-nil, zero value otherwise.

### GetMysql5Ok

`func (o *DbParametersByType) GetMysql5Ok() (*[]string, bool)`

GetMysql5Ok returns a tuple with the Mysql5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql5

`func (o *DbParametersByType) SetMysql5(v []string)`

SetMysql5 sets Mysql5 field to given value.


### GetMysql

`func (o *DbParametersByType) GetMysql() []string`

GetMysql returns the Mysql field if non-nil, zero value otherwise.

### GetMysqlOk

`func (o *DbParametersByType) GetMysqlOk() (*[]string, bool)`

GetMysqlOk returns a tuple with the Mysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql

`func (o *DbParametersByType) SetMysql(v []string)`

SetMysql sets Mysql field to given value.


### GetMysql84

`func (o *DbParametersByType) GetMysql84() []string`

GetMysql84 returns the Mysql84 field if non-nil, zero value otherwise.

### GetMysql84Ok

`func (o *DbParametersByType) GetMysql84Ok() (*[]string, bool)`

GetMysql84Ok returns a tuple with the Mysql84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql84

`func (o *DbParametersByType) SetMysql84(v []string)`

SetMysql84 sets Mysql84 field to given value.


### GetPostgres

`func (o *DbParametersByType) GetPostgres() []string`

GetPostgres returns the Postgres field if non-nil, zero value otherwise.

### GetPostgresOk

`func (o *DbParametersByType) GetPostgresOk() (*[]string, bool)`

GetPostgresOk returns a tuple with the Postgres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgres

`func (o *DbParametersByType) SetPostgres(v []string)`

SetPostgres sets Postgres field to given value.


### GetPostgres14

`func (o *DbParametersByType) GetPostgres14() []string`

GetPostgres14 returns the Postgres14 field if non-nil, zero value otherwise.

### GetPostgres14Ok

`func (o *DbParametersByType) GetPostgres14Ok() (*[]string, bool)`

GetPostgres14Ok returns a tuple with the Postgres14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgres14

`func (o *DbParametersByType) SetPostgres14(v []string)`

SetPostgres14 sets Postgres14 field to given value.


### GetPostgres15

`func (o *DbParametersByType) GetPostgres15() []string`

GetPostgres15 returns the Postgres15 field if non-nil, zero value otherwise.

### GetPostgres15Ok

`func (o *DbParametersByType) GetPostgres15Ok() (*[]string, bool)`

GetPostgres15Ok returns a tuple with the Postgres15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgres15

`func (o *DbParametersByType) SetPostgres15(v []string)`

SetPostgres15 sets Postgres15 field to given value.


### GetPostgres16

`func (o *DbParametersByType) GetPostgres16() []string`

GetPostgres16 returns the Postgres16 field if non-nil, zero value otherwise.

### GetPostgres16Ok

`func (o *DbParametersByType) GetPostgres16Ok() (*[]string, bool)`

GetPostgres16Ok returns a tuple with the Postgres16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgres16

`func (o *DbParametersByType) SetPostgres16(v []string)`

SetPostgres16 sets Postgres16 field to given value.


### GetPostgres17

`func (o *DbParametersByType) GetPostgres17() []string`

GetPostgres17 returns the Postgres17 field if non-nil, zero value otherwise.

### GetPostgres17Ok

`func (o *DbParametersByType) GetPostgres17Ok() (*[]string, bool)`

GetPostgres17Ok returns a tuple with the Postgres17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgres17

`func (o *DbParametersByType) SetPostgres17(v []string)`

SetPostgres17 sets Postgres17 field to given value.


### GetPostgres18

`func (o *DbParametersByType) GetPostgres18() []string`

GetPostgres18 returns the Postgres18 field if non-nil, zero value otherwise.

### GetPostgres18Ok

`func (o *DbParametersByType) GetPostgres18Ok() (*[]string, bool)`

GetPostgres18Ok returns a tuple with the Postgres18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgres18

`func (o *DbParametersByType) SetPostgres18(v []string)`

SetPostgres18 sets Postgres18 field to given value.


### GetRedis

`func (o *DbParametersByType) GetRedis() []string`

GetRedis returns the Redis field if non-nil, zero value otherwise.

### GetRedisOk

`func (o *DbParametersByType) GetRedisOk() (*[]string, bool)`

GetRedisOk returns a tuple with the Redis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis

`func (o *DbParametersByType) SetRedis(v []string)`

SetRedis sets Redis field to given value.


### GetRedis7

`func (o *DbParametersByType) GetRedis7() []string`

GetRedis7 returns the Redis7 field if non-nil, zero value otherwise.

### GetRedis7Ok

`func (o *DbParametersByType) GetRedis7Ok() (*[]string, bool)`

GetRedis7Ok returns a tuple with the Redis7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis7

`func (o *DbParametersByType) SetRedis7(v []string)`

SetRedis7 sets Redis7 field to given value.


### GetRedis81

`func (o *DbParametersByType) GetRedis81() []string`

GetRedis81 returns the Redis81 field if non-nil, zero value otherwise.

### GetRedis81Ok

`func (o *DbParametersByType) GetRedis81Ok() (*[]string, bool)`

GetRedis81Ok returns a tuple with the Redis81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis81

`func (o *DbParametersByType) SetRedis81(v []string)`

SetRedis81 sets Redis81 field to given value.


### GetValkey

`func (o *DbParametersByType) GetValkey() []string`

GetValkey returns the Valkey field if non-nil, zero value otherwise.

### GetValkeyOk

`func (o *DbParametersByType) GetValkeyOk() (*[]string, bool)`

GetValkeyOk returns a tuple with the Valkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValkey

`func (o *DbParametersByType) SetValkey(v []string)`

SetValkey sets Valkey field to given value.


### GetValkey7

`func (o *DbParametersByType) GetValkey7() []string`

GetValkey7 returns the Valkey7 field if non-nil, zero value otherwise.

### GetValkey7Ok

`func (o *DbParametersByType) GetValkey7Ok() (*[]string, bool)`

GetValkey7Ok returns a tuple with the Valkey7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValkey7

`func (o *DbParametersByType) SetValkey7(v []string)`

SetValkey7 sets Valkey7 field to given value.


### GetValkey81

`func (o *DbParametersByType) GetValkey81() []string`

GetValkey81 returns the Valkey81 field if non-nil, zero value otherwise.

### GetValkey81Ok

`func (o *DbParametersByType) GetValkey81Ok() (*[]string, bool)`

GetValkey81Ok returns a tuple with the Valkey81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValkey81

`func (o *DbParametersByType) SetValkey81(v []string)`

SetValkey81 sets Valkey81 field to given value.


### GetValkey91

`func (o *DbParametersByType) GetValkey91() []string`

GetValkey91 returns the Valkey91 field if non-nil, zero value otherwise.

### GetValkey91Ok

`func (o *DbParametersByType) GetValkey91Ok() (*[]string, bool)`

GetValkey91Ok returns a tuple with the Valkey91 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValkey91

`func (o *DbParametersByType) SetValkey91(v []string)`

SetValkey91 sets Valkey91 field to given value.


### GetMongodb4

`func (o *DbParametersByType) GetMongodb4() []string`

GetMongodb4 returns the Mongodb4 field if non-nil, zero value otherwise.

### GetMongodb4Ok

`func (o *DbParametersByType) GetMongodb4Ok() (*[]string, bool)`

GetMongodb4Ok returns a tuple with the Mongodb4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb4

`func (o *DbParametersByType) SetMongodb4(v []string)`

SetMongodb4 sets Mongodb4 field to given value.


### GetMongodb

`func (o *DbParametersByType) GetMongodb() []string`

GetMongodb returns the Mongodb field if non-nil, zero value otherwise.

### GetMongodbOk

`func (o *DbParametersByType) GetMongodbOk() (*[]string, bool)`

GetMongodbOk returns a tuple with the Mongodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb

`func (o *DbParametersByType) SetMongodb(v []string)`

SetMongodb sets Mongodb field to given value.


### GetMongodb6

`func (o *DbParametersByType) GetMongodb6() []string`

GetMongodb6 returns the Mongodb6 field if non-nil, zero value otherwise.

### GetMongodb6Ok

`func (o *DbParametersByType) GetMongodb6Ok() (*[]string, bool)`

GetMongodb6Ok returns a tuple with the Mongodb6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb6

`func (o *DbParametersByType) SetMongodb6(v []string)`

SetMongodb6 sets Mongodb6 field to given value.


### GetMongodb7

`func (o *DbParametersByType) GetMongodb7() []string`

GetMongodb7 returns the Mongodb7 field if non-nil, zero value otherwise.

### GetMongodb7Ok

`func (o *DbParametersByType) GetMongodb7Ok() (*[]string, bool)`

GetMongodb7Ok returns a tuple with the Mongodb7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb7

`func (o *DbParametersByType) SetMongodb7(v []string)`

SetMongodb7 sets Mongodb7 field to given value.


### GetMongodb80

`func (o *DbParametersByType) GetMongodb80() []string`

GetMongodb80 returns the Mongodb80 field if non-nil, zero value otherwise.

### GetMongodb80Ok

`func (o *DbParametersByType) GetMongodb80Ok() (*[]string, bool)`

GetMongodb80Ok returns a tuple with the Mongodb80 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb80

`func (o *DbParametersByType) SetMongodb80(v []string)`

SetMongodb80 sets Mongodb80 field to given value.


### GetOpensearch

`func (o *DbParametersByType) GetOpensearch() []string`

GetOpensearch returns the Opensearch field if non-nil, zero value otherwise.

### GetOpensearchOk

`func (o *DbParametersByType) GetOpensearchOk() (*[]string, bool)`

GetOpensearchOk returns a tuple with the Opensearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensearch

`func (o *DbParametersByType) SetOpensearch(v []string)`

SetOpensearch sets Opensearch field to given value.


### GetOpensearch219

`func (o *DbParametersByType) GetOpensearch219() []string`

GetOpensearch219 returns the Opensearch219 field if non-nil, zero value otherwise.

### GetOpensearch219Ok

`func (o *DbParametersByType) GetOpensearch219Ok() (*[]string, bool)`

GetOpensearch219Ok returns a tuple with the Opensearch219 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensearch219

`func (o *DbParametersByType) SetOpensearch219(v []string)`

SetOpensearch219 sets Opensearch219 field to given value.


### GetClickhouse

`func (o *DbParametersByType) GetClickhouse() []string`

GetClickhouse returns the Clickhouse field if non-nil, zero value otherwise.

### GetClickhouseOk

`func (o *DbParametersByType) GetClickhouseOk() (*[]string, bool)`

GetClickhouseOk returns a tuple with the Clickhouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouse

`func (o *DbParametersByType) SetClickhouse(v []string)`

SetClickhouse sets Clickhouse field to given value.


### GetClickhouse24

`func (o *DbParametersByType) GetClickhouse24() []string`

GetClickhouse24 returns the Clickhouse24 field if non-nil, zero value otherwise.

### GetClickhouse24Ok

`func (o *DbParametersByType) GetClickhouse24Ok() (*[]string, bool)`

GetClickhouse24Ok returns a tuple with the Clickhouse24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouse24

`func (o *DbParametersByType) SetClickhouse24(v []string)`

SetClickhouse24 sets Clickhouse24 field to given value.


### GetClickhouse25

`func (o *DbParametersByType) GetClickhouse25() []string`

GetClickhouse25 returns the Clickhouse25 field if non-nil, zero value otherwise.

### GetClickhouse25Ok

`func (o *DbParametersByType) GetClickhouse25Ok() (*[]string, bool)`

GetClickhouse25Ok returns a tuple with the Clickhouse25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouse25

`func (o *DbParametersByType) SetClickhouse25(v []string)`

SetClickhouse25 sets Clickhouse25 field to given value.


### GetKafka

`func (o *DbParametersByType) GetKafka() []string`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *DbParametersByType) GetKafkaOk() (*[]string, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *DbParametersByType) SetKafka(v []string)`

SetKafka sets Kafka field to given value.


### GetRabbitmq

`func (o *DbParametersByType) GetRabbitmq() []string`

GetRabbitmq returns the Rabbitmq field if non-nil, zero value otherwise.

### GetRabbitmqOk

`func (o *DbParametersByType) GetRabbitmqOk() (*[]string, bool)`

GetRabbitmqOk returns a tuple with the Rabbitmq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmq

`func (o *DbParametersByType) SetRabbitmq(v []string)`

SetRabbitmq sets Rabbitmq field to given value.


### GetRabbitmq40

`func (o *DbParametersByType) GetRabbitmq40() []string`

GetRabbitmq40 returns the Rabbitmq40 field if non-nil, zero value otherwise.

### GetRabbitmq40Ok

`func (o *DbParametersByType) GetRabbitmq40Ok() (*[]string, bool)`

GetRabbitmq40Ok returns a tuple with the Rabbitmq40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbitmq40

`func (o *DbParametersByType) SetRabbitmq40(v []string)`

SetRabbitmq40 sets Rabbitmq40 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


