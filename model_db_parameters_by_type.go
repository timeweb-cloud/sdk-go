/*
Timeweb Cloud API

# Введение API Timeweb Cloud позволяет вам управлять ресурсами в облаке программным способом с использованием обычных HTTP-запросов.  Множество функций, которые доступны в панели управления Timeweb Cloud, также доступны через API, что позволяет вам автоматизировать ваши собственные сценарии.  В этой документации сперва будет описан общий дизайн и принципы работы API, а после этого конкретные конечные точки. Также будут приведены примеры запросов к ним.   ## Запросы Запросы должны выполняться по протоколу `HTTPS`, чтобы гарантировать шифрование транзакций. Поддерживаются следующие методы запроса: |Метод|Применение| |--- |--- | |GET|Извлекает данные о коллекциях и отдельных ресурсах.| |POST|Для коллекций создает новый ресурс этого типа. Также используется для выполнения действий с конкретным ресурсом.| |PUT|Обновляет существующий ресурс.| |PATCH|Некоторые ресурсы поддерживают частичное обновление, то есть обновление только части атрибутов ресурса, в этом случае вместо метода PUT будет использован PATCH.| |DELETE|Удаляет ресурс.|  Методы `POST`, `PUT` и `PATCH` могут включать объект в тело запроса с типом содержимого `application/json`.  ### Параметры в запросах Некоторые коллекции поддерживают пагинацию, поиск или сортировку в запросах. В параметрах запроса требуется передать: - `limit` — обозначает количество записей, которое необходимо вернуть  - `offset` — указывает на смещение, относительно начала списка  - `search` — позволяет указать набор символов для поиска  - `sort` — можно задать правило сортировки коллекции  ## Ответы Запросы вернут один из следующих кодов состояния ответа HTTP:  |Статус|Описание| |--- |--- | |200 OK|Действие с ресурсом было выполнено успешно.| |201 Created|Ресурс был успешно создан. При этом ресурс может быть как уже готовым к использованию, так и находиться в процессе запуска.| |204 No Content|Действие с ресурсом было выполнено успешно, и ответ не содержит дополнительной информации в теле.| |400 Bad Request|Был отправлен неверный запрос, например, в нем отсутствуют обязательные параметры и т. д. Тело ответа будет содержать дополнительную информацию об ошибке.| |401 Unauthorized|Ошибка аутентификации.| |403 Forbidden|Аутентификация прошла успешно, но недостаточно прав для выполнения действия.| |404 Not Found|Запрашиваемый ресурс не найден.| |409 Conflict|Запрос конфликтует с текущим состоянием.| |423 Locked|Ресурс из запроса заблокирован от применения к нему указанного метода.| |429 Too Many Requests|Был достигнут лимит по количеству запросов в единицу времени.| |500 Internal Server Error|При выполнении запроса произошла какая-то внутренняя ошибка. Чтобы решить эту проблему, лучше всего создать тикет в панели управления.|  ### Структура успешного ответа Все конечные точки будут возвращать данные в формате `JSON`. Ответы на `GET`-запросы будут иметь на верхнем уровне следующую структуру атрибутов:  |Название поля|Тип|Описание| |--- |--- |--- | |[entity_name]|object, object[], string[], number[], boolean|Динамическое поле, которое будет меняться в зависимости от запрашиваемого ресурса и будет содержать все атрибуты, необходимые для описания этого ресурса. Например, при запросе списка баз данных будет возвращаться поле `dbs`, а при запросе конкретного облачного сервера `server`. Для некоторых конечных точек в ответе может возвращаться сразу несколько ресурсов.| |meta|object|Опционально. Объект, который содержит вспомогательную информацию о ресурсе. Чаще всего будет встречаться при запросе коллекций и содержать поле `total`, которое будет указывать на количество элементов в коллекции.| |response_id|string|Опционально. В большинстве случаев в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID— так мы сможем найти ответ на него намного быстрее. Также вы можете использовать этот ID, чтобы убедиться, что это новый ответ на запрос и результат не был получен из кэша.|  Пример запроса на получение списка SSH-ключей: ```     HTTP/2.0 200 OK     {       \"ssh_keys\":[           {             \"body\":\"ssh-rsa AAAAB3NzaC1sdfghjkOAsBwWhs= example@device.local\",             \"created_at\":\"2021-09-15T19:52:27Z\",             \"expired_at\":null,             \"id\":5297,             \"is_default\":false,             \"name\":\"example@device.local\",             \"used_at\":null,             \"used_by\":[]           }       ],       \"meta\":{           \"total\":1       },       \"response_id\":\"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ### Структура ответа с ошибкой |Название поля|Тип|Описание| |--- |--- |--- | |status_code|number|Короткий числовой идентификатор ошибки.| |error_code|string|Короткий текстовый идентификатор ошибки, который уточняет числовой идентификатор и удобен для программной обработки. Самый простой пример — это код `not_found` для ошибки 404.| |message|string, string[]|Опционально. В большинстве случаев в ответе будет содержаться человекочитаемое подробное описание ошибки или ошибок, которые помогут понять, что нужно исправить.| |response_id|string|Опционально. В большинстве случае в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID — так мы сможем найти ответ на него намного быстрее.|  Пример: ```     HTTP/2.0 403 Forbidden     {       \"status_code\": 403,       \"error_code\":  \"forbidden\",       \"message\":     \"You do not have access for the attempted action\",       \"response_id\": \"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ## Статусы ресурсов Важно учесть, что при создании большинства ресурсов внутри платформы вам будет сразу возвращен ответ от сервера со статусом `200 OK` или `201 Created` и ID созданного ресурса в теле ответа, но при этом этот ресурс может быть ещё в *состоянии запуска*.  Для того чтобы понять, в каком состоянии сейчас находится ваш ресурс, мы добавили поле `status` в ответ на получение информации о ресурсе.  Список статусов будет отличаться в зависимости от типа ресурса. Увидеть поддерживаемый список статусов вы сможете в описании каждого конкретного ресурса.     ## Ограничение скорости запросов (Rate Limiting) Чтобы обеспечить стабильность для всех пользователей, Timeweb Cloud защищает API от всплесков входящего трафика, анализируя количество запросов c каждого аккаунта к каждой конечной точке.  Если ваше приложение отправляет более 20 запросов в секунду на одну конечную точку, то для этого запроса API может вернуть код состояния HTTP `429 Too Many Requests`.   ## Аутентификация Доступ к API осуществляется с помощью JWT-токена. Токенами можно управлять внутри панели управления Timeweb Cloud в разделе *API и Terraform*.  Токен необходимо передавать в заголовке каждого запроса в формате: ```   Authorization: Bearer $TIMEWEB_CLOUD_TOKEN ```  ## Формат примеров API Примеры в этой документации описаны с помощью `curl`, HTTP-клиента командной строки. На компьютерах `Linux` и `macOS` обычно по умолчанию установлен `curl`, и он доступен для загрузки на всех популярных платформах, включая `Windows`.  Каждый пример разделен на несколько строк символом `\\`, который совместим с `bash`. Типичный пример выглядит так: ```   curl -X PATCH      -H \"Content-Type: application/json\"      -H \"Authorization: Bearer $TIMEWEB_CLOUD_TOKEN\"      -d '{\"name\":\"Cute Corvus\",\"comment\":\"Development Server\"}'      \"https://api.timeweb.cloud/api/v1/dedicated/1051\" ``` - Параметр `-X` задает метод запроса. Для согласованности метод будет указан во всех примерах, даже если он явно не требуется для методов `GET`. - Строки `-H` задают требуемые HTTP-заголовки. - Примеры, для которых требуется объект JSON в теле запроса, передают требуемые данные через параметр `-d`.  Чтобы использовать приведенные примеры, не подставляя каждый раз в них свой токен, вы можете добавить токен один раз в переменные окружения в вашей консоли. Например, на `Linux` это можно сделать с помощью команды:  ``` TIMEWEB_CLOUD_TOKEN=\"token\" ```  После этого токен будет автоматически подставляться в ваши запросы.  Обратите внимание, что все значения в этой документации являются примерами. Не полагайтесь на IDы операционных систем, тарифов и т.д., используемые в примерах. Используйте соответствующую конечную точку для получения значений перед созданием ресурсов.   ## Версионирование API построено согласно принципам [семантического версионирования](https://semver.org/lang/ru). Это значит, что мы гарантируем обратную совместимость всех изменений в пределах одной мажорной версии.  Мажорная версия каждой конечной точки обозначается в пути запроса, например, запрос `/api/v1/servers` указывает, что этот метод имеет версию 1.

API version: 1.0.0
Contact: info@timeweb.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DbParametersByType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbParametersByType{}

// DbParametersByType Имена параметров конфигурации, доступных для каждого типа кластера базы данных. Ключ объекта — тип кластера (значение поля `type` при создании кластера), значение — массив имён параметров, которые можно передать в `config_parameters` для кластера этого типа. Наборы параметров различаются между версиями одной СУБД. Значения параметров этот метод не возвращает — рекомендуемые значения можно получить в `GET /api/v1/dbs/default-parameters`.
type DbParametersByType struct {
	// Параметры, доступные для кластеров типа `mysql5`.
	Mysql5 []string `json:"mysql5"`
	// Параметры, доступные для кластеров типа `mysql`.
	Mysql []string `json:"mysql"`
	// Параметры, доступные для кластеров типа `mysql8_4`.
	Mysql84 []string `json:"mysql8_4"`
	// Параметры, доступные для кластеров типа `postgres` (PostgreSQL 13).
	Postgres []string `json:"postgres"`
	// Параметры, доступные для кластеров типа `postgres14`.
	Postgres14 []string `json:"postgres14"`
	// Параметры, доступные для кластеров типа `postgres15`.
	Postgres15 []string `json:"postgres15"`
	// Параметры, доступные для кластеров типа `postgres16`.
	Postgres16 []string `json:"postgres16"`
	// Параметры, доступные для кластеров типа `postgres17`.
	Postgres17 []string `json:"postgres17"`
	// Параметры, доступные для кластеров типа `postgres18`. Набор отличается от предыдущих версий PostgreSQL — например, добавлены `io_method` и `io_workers`.
	Postgres18 []string `json:"postgres18"`
	// Параметры, доступные для кластеров типа `redis`.
	Redis []string `json:"redis"`
	// Параметры, доступные для кластеров типа `redis7`.
	Redis7 []string `json:"redis7"`
	// Параметры, доступные для кластеров типа `redis8_1`.
	Redis81 []string `json:"redis8_1"`
	// Параметры, доступные для кластеров типа `valkey`.
	Valkey []string `json:"valkey"`
	// Параметры, доступные для кластеров типа `valkey7`.
	Valkey7 []string `json:"valkey7"`
	// Параметры, доступные для кластеров типа `valkey8_1`.
	Valkey81 []string `json:"valkey8_1"`
	// Параметры, доступные для кластеров типа `valkey9_1`.
	Valkey91 []string `json:"valkey9_1"`
	// Для кластеров типа `mongodb4` настраиваемых параметров нет — всегда пустой массив.
	Mongodb4 []string `json:"mongodb4"`
	// Для кластеров типа `mongodb` настраиваемых параметров нет — всегда пустой массив.
	Mongodb []string `json:"mongodb"`
	// Для кластеров типа `mongodb6` настраиваемых параметров нет — всегда пустой массив.
	Mongodb6 []string `json:"mongodb6"`
	// Для кластеров типа `mongodb7` настраиваемых параметров нет — всегда пустой массив.
	Mongodb7 []string `json:"mongodb7"`
	// Для кластеров типа `mongodb8_0` настраиваемых параметров нет — всегда пустой массив.
	Mongodb80 []string `json:"mongodb8_0"`
	// Для кластеров типа `opensearch` настраиваемых параметров нет — всегда пустой массив.
	Opensearch []string `json:"opensearch"`
	// Для кластеров типа `opensearch2_19` настраиваемых параметров нет — всегда пустой массив.
	Opensearch219 []string `json:"opensearch2_19"`
	// Для кластеров типа `clickhouse` настраиваемых параметров нет — всегда пустой массив.
	Clickhouse []string `json:"clickhouse"`
	// Для кластеров типа `clickhouse24` настраиваемых параметров нет — всегда пустой массив.
	Clickhouse24 []string `json:"clickhouse24"`
	// Для кластеров типа `clickhouse25` настраиваемых параметров нет — всегда пустой массив.
	Clickhouse25 []string `json:"clickhouse25"`
	// Для кластеров типа `kafka` настраиваемых параметров нет — всегда пустой массив.
	Kafka []string `json:"kafka"`
	// Для кластеров типа `rabbitmq` настраиваемых параметров нет — всегда пустой массив.
	Rabbitmq []string `json:"rabbitmq"`
	// Для кластеров типа `rabbitmq4_0` настраиваемых параметров нет — всегда пустой массив.
	Rabbitmq40 []string `json:"rabbitmq4_0"`
}

// NewDbParametersByType instantiates a new DbParametersByType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbParametersByType(mysql5 []string, mysql []string, mysql84 []string, postgres []string, postgres14 []string, postgres15 []string, postgres16 []string, postgres17 []string, postgres18 []string, redis []string, redis7 []string, redis81 []string, valkey []string, valkey7 []string, valkey81 []string, valkey91 []string, mongodb4 []string, mongodb []string, mongodb6 []string, mongodb7 []string, mongodb80 []string, opensearch []string, opensearch219 []string, clickhouse []string, clickhouse24 []string, clickhouse25 []string, kafka []string, rabbitmq []string, rabbitmq40 []string) *DbParametersByType {
	this := DbParametersByType{}
	this.Mysql5 = mysql5
	this.Mysql = mysql
	this.Mysql84 = mysql84
	this.Postgres = postgres
	this.Postgres14 = postgres14
	this.Postgres15 = postgres15
	this.Postgres16 = postgres16
	this.Postgres17 = postgres17
	this.Postgres18 = postgres18
	this.Redis = redis
	this.Redis7 = redis7
	this.Redis81 = redis81
	this.Valkey = valkey
	this.Valkey7 = valkey7
	this.Valkey81 = valkey81
	this.Valkey91 = valkey91
	this.Mongodb4 = mongodb4
	this.Mongodb = mongodb
	this.Mongodb6 = mongodb6
	this.Mongodb7 = mongodb7
	this.Mongodb80 = mongodb80
	this.Opensearch = opensearch
	this.Opensearch219 = opensearch219
	this.Clickhouse = clickhouse
	this.Clickhouse24 = clickhouse24
	this.Clickhouse25 = clickhouse25
	this.Kafka = kafka
	this.Rabbitmq = rabbitmq
	this.Rabbitmq40 = rabbitmq40
	return &this
}

// NewDbParametersByTypeWithDefaults instantiates a new DbParametersByType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbParametersByTypeWithDefaults() *DbParametersByType {
	this := DbParametersByType{}
	return &this
}

// GetMysql5 returns the Mysql5 field value
func (o *DbParametersByType) GetMysql5() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mysql5
}

// GetMysql5Ok returns a tuple with the Mysql5 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMysql5Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mysql5, true
}

// SetMysql5 sets field value
func (o *DbParametersByType) SetMysql5(v []string) {
	o.Mysql5 = v
}

// GetMysql returns the Mysql field value
func (o *DbParametersByType) GetMysql() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mysql
}

// GetMysqlOk returns a tuple with the Mysql field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMysqlOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mysql, true
}

// SetMysql sets field value
func (o *DbParametersByType) SetMysql(v []string) {
	o.Mysql = v
}

// GetMysql84 returns the Mysql84 field value
func (o *DbParametersByType) GetMysql84() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mysql84
}

// GetMysql84Ok returns a tuple with the Mysql84 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMysql84Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mysql84, true
}

// SetMysql84 sets field value
func (o *DbParametersByType) SetMysql84(v []string) {
	o.Mysql84 = v
}

// GetPostgres returns the Postgres field value
func (o *DbParametersByType) GetPostgres() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Postgres
}

// GetPostgresOk returns a tuple with the Postgres field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetPostgresOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postgres, true
}

// SetPostgres sets field value
func (o *DbParametersByType) SetPostgres(v []string) {
	o.Postgres = v
}

// GetPostgres14 returns the Postgres14 field value
func (o *DbParametersByType) GetPostgres14() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Postgres14
}

// GetPostgres14Ok returns a tuple with the Postgres14 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetPostgres14Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postgres14, true
}

// SetPostgres14 sets field value
func (o *DbParametersByType) SetPostgres14(v []string) {
	o.Postgres14 = v
}

// GetPostgres15 returns the Postgres15 field value
func (o *DbParametersByType) GetPostgres15() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Postgres15
}

// GetPostgres15Ok returns a tuple with the Postgres15 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetPostgres15Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postgres15, true
}

// SetPostgres15 sets field value
func (o *DbParametersByType) SetPostgres15(v []string) {
	o.Postgres15 = v
}

// GetPostgres16 returns the Postgres16 field value
func (o *DbParametersByType) GetPostgres16() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Postgres16
}

// GetPostgres16Ok returns a tuple with the Postgres16 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetPostgres16Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postgres16, true
}

// SetPostgres16 sets field value
func (o *DbParametersByType) SetPostgres16(v []string) {
	o.Postgres16 = v
}

// GetPostgres17 returns the Postgres17 field value
func (o *DbParametersByType) GetPostgres17() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Postgres17
}

// GetPostgres17Ok returns a tuple with the Postgres17 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetPostgres17Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postgres17, true
}

// SetPostgres17 sets field value
func (o *DbParametersByType) SetPostgres17(v []string) {
	o.Postgres17 = v
}

// GetPostgres18 returns the Postgres18 field value
func (o *DbParametersByType) GetPostgres18() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Postgres18
}

// GetPostgres18Ok returns a tuple with the Postgres18 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetPostgres18Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postgres18, true
}

// SetPostgres18 sets field value
func (o *DbParametersByType) SetPostgres18(v []string) {
	o.Postgres18 = v
}

// GetRedis returns the Redis field value
func (o *DbParametersByType) GetRedis() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Redis
}

// GetRedisOk returns a tuple with the Redis field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetRedisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Redis, true
}

// SetRedis sets field value
func (o *DbParametersByType) SetRedis(v []string) {
	o.Redis = v
}

// GetRedis7 returns the Redis7 field value
func (o *DbParametersByType) GetRedis7() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Redis7
}

// GetRedis7Ok returns a tuple with the Redis7 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetRedis7Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Redis7, true
}

// SetRedis7 sets field value
func (o *DbParametersByType) SetRedis7(v []string) {
	o.Redis7 = v
}

// GetRedis81 returns the Redis81 field value
func (o *DbParametersByType) GetRedis81() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Redis81
}

// GetRedis81Ok returns a tuple with the Redis81 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetRedis81Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Redis81, true
}

// SetRedis81 sets field value
func (o *DbParametersByType) SetRedis81(v []string) {
	o.Redis81 = v
}

// GetValkey returns the Valkey field value
func (o *DbParametersByType) GetValkey() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Valkey
}

// GetValkeyOk returns a tuple with the Valkey field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetValkeyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Valkey, true
}

// SetValkey sets field value
func (o *DbParametersByType) SetValkey(v []string) {
	o.Valkey = v
}

// GetValkey7 returns the Valkey7 field value
func (o *DbParametersByType) GetValkey7() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Valkey7
}

// GetValkey7Ok returns a tuple with the Valkey7 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetValkey7Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Valkey7, true
}

// SetValkey7 sets field value
func (o *DbParametersByType) SetValkey7(v []string) {
	o.Valkey7 = v
}

// GetValkey81 returns the Valkey81 field value
func (o *DbParametersByType) GetValkey81() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Valkey81
}

// GetValkey81Ok returns a tuple with the Valkey81 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetValkey81Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Valkey81, true
}

// SetValkey81 sets field value
func (o *DbParametersByType) SetValkey81(v []string) {
	o.Valkey81 = v
}

// GetValkey91 returns the Valkey91 field value
func (o *DbParametersByType) GetValkey91() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Valkey91
}

// GetValkey91Ok returns a tuple with the Valkey91 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetValkey91Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Valkey91, true
}

// SetValkey91 sets field value
func (o *DbParametersByType) SetValkey91(v []string) {
	o.Valkey91 = v
}

// GetMongodb4 returns the Mongodb4 field value
func (o *DbParametersByType) GetMongodb4() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mongodb4
}

// GetMongodb4Ok returns a tuple with the Mongodb4 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMongodb4Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mongodb4, true
}

// SetMongodb4 sets field value
func (o *DbParametersByType) SetMongodb4(v []string) {
	o.Mongodb4 = v
}

// GetMongodb returns the Mongodb field value
func (o *DbParametersByType) GetMongodb() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mongodb
}

// GetMongodbOk returns a tuple with the Mongodb field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMongodbOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mongodb, true
}

// SetMongodb sets field value
func (o *DbParametersByType) SetMongodb(v []string) {
	o.Mongodb = v
}

// GetMongodb6 returns the Mongodb6 field value
func (o *DbParametersByType) GetMongodb6() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mongodb6
}

// GetMongodb6Ok returns a tuple with the Mongodb6 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMongodb6Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mongodb6, true
}

// SetMongodb6 sets field value
func (o *DbParametersByType) SetMongodb6(v []string) {
	o.Mongodb6 = v
}

// GetMongodb7 returns the Mongodb7 field value
func (o *DbParametersByType) GetMongodb7() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mongodb7
}

// GetMongodb7Ok returns a tuple with the Mongodb7 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMongodb7Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mongodb7, true
}

// SetMongodb7 sets field value
func (o *DbParametersByType) SetMongodb7(v []string) {
	o.Mongodb7 = v
}

// GetMongodb80 returns the Mongodb80 field value
func (o *DbParametersByType) GetMongodb80() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mongodb80
}

// GetMongodb80Ok returns a tuple with the Mongodb80 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetMongodb80Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mongodb80, true
}

// SetMongodb80 sets field value
func (o *DbParametersByType) SetMongodb80(v []string) {
	o.Mongodb80 = v
}

// GetOpensearch returns the Opensearch field value
func (o *DbParametersByType) GetOpensearch() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Opensearch
}

// GetOpensearchOk returns a tuple with the Opensearch field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetOpensearchOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Opensearch, true
}

// SetOpensearch sets field value
func (o *DbParametersByType) SetOpensearch(v []string) {
	o.Opensearch = v
}

// GetOpensearch219 returns the Opensearch219 field value
func (o *DbParametersByType) GetOpensearch219() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Opensearch219
}

// GetOpensearch219Ok returns a tuple with the Opensearch219 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetOpensearch219Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Opensearch219, true
}

// SetOpensearch219 sets field value
func (o *DbParametersByType) SetOpensearch219(v []string) {
	o.Opensearch219 = v
}

// GetClickhouse returns the Clickhouse field value
func (o *DbParametersByType) GetClickhouse() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Clickhouse
}

// GetClickhouseOk returns a tuple with the Clickhouse field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetClickhouseOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clickhouse, true
}

// SetClickhouse sets field value
func (o *DbParametersByType) SetClickhouse(v []string) {
	o.Clickhouse = v
}

// GetClickhouse24 returns the Clickhouse24 field value
func (o *DbParametersByType) GetClickhouse24() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Clickhouse24
}

// GetClickhouse24Ok returns a tuple with the Clickhouse24 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetClickhouse24Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clickhouse24, true
}

// SetClickhouse24 sets field value
func (o *DbParametersByType) SetClickhouse24(v []string) {
	o.Clickhouse24 = v
}

// GetClickhouse25 returns the Clickhouse25 field value
func (o *DbParametersByType) GetClickhouse25() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Clickhouse25
}

// GetClickhouse25Ok returns a tuple with the Clickhouse25 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetClickhouse25Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clickhouse25, true
}

// SetClickhouse25 sets field value
func (o *DbParametersByType) SetClickhouse25(v []string) {
	o.Clickhouse25 = v
}

// GetKafka returns the Kafka field value
func (o *DbParametersByType) GetKafka() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Kafka
}

// GetKafkaOk returns a tuple with the Kafka field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetKafkaOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kafka, true
}

// SetKafka sets field value
func (o *DbParametersByType) SetKafka(v []string) {
	o.Kafka = v
}

// GetRabbitmq returns the Rabbitmq field value
func (o *DbParametersByType) GetRabbitmq() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Rabbitmq
}

// GetRabbitmqOk returns a tuple with the Rabbitmq field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetRabbitmqOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rabbitmq, true
}

// SetRabbitmq sets field value
func (o *DbParametersByType) SetRabbitmq(v []string) {
	o.Rabbitmq = v
}

// GetRabbitmq40 returns the Rabbitmq40 field value
func (o *DbParametersByType) GetRabbitmq40() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Rabbitmq40
}

// GetRabbitmq40Ok returns a tuple with the Rabbitmq40 field value
// and a boolean to check if the value has been set.
func (o *DbParametersByType) GetRabbitmq40Ok() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rabbitmq40, true
}

// SetRabbitmq40 sets field value
func (o *DbParametersByType) SetRabbitmq40(v []string) {
	o.Rabbitmq40 = v
}

func (o DbParametersByType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbParametersByType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mysql5"] = o.Mysql5
	toSerialize["mysql"] = o.Mysql
	toSerialize["mysql8_4"] = o.Mysql84
	toSerialize["postgres"] = o.Postgres
	toSerialize["postgres14"] = o.Postgres14
	toSerialize["postgres15"] = o.Postgres15
	toSerialize["postgres16"] = o.Postgres16
	toSerialize["postgres17"] = o.Postgres17
	toSerialize["postgres18"] = o.Postgres18
	toSerialize["redis"] = o.Redis
	toSerialize["redis7"] = o.Redis7
	toSerialize["redis8_1"] = o.Redis81
	toSerialize["valkey"] = o.Valkey
	toSerialize["valkey7"] = o.Valkey7
	toSerialize["valkey8_1"] = o.Valkey81
	toSerialize["valkey9_1"] = o.Valkey91
	toSerialize["mongodb4"] = o.Mongodb4
	toSerialize["mongodb"] = o.Mongodb
	toSerialize["mongodb6"] = o.Mongodb6
	toSerialize["mongodb7"] = o.Mongodb7
	toSerialize["mongodb8_0"] = o.Mongodb80
	toSerialize["opensearch"] = o.Opensearch
	toSerialize["opensearch2_19"] = o.Opensearch219
	toSerialize["clickhouse"] = o.Clickhouse
	toSerialize["clickhouse24"] = o.Clickhouse24
	toSerialize["clickhouse25"] = o.Clickhouse25
	toSerialize["kafka"] = o.Kafka
	toSerialize["rabbitmq"] = o.Rabbitmq
	toSerialize["rabbitmq4_0"] = o.Rabbitmq40
	return toSerialize, nil
}

type NullableDbParametersByType struct {
	value *DbParametersByType
	isSet bool
}

func (v NullableDbParametersByType) Get() *DbParametersByType {
	return v.value
}

func (v *NullableDbParametersByType) Set(val *DbParametersByType) {
	v.value = val
	v.isSet = true
}

func (v NullableDbParametersByType) IsSet() bool {
	return v.isSet
}

func (v *NullableDbParametersByType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbParametersByType(val *DbParametersByType) *NullableDbParametersByType {
	return &NullableDbParametersByType{value: val, isSet: true}
}

func (v NullableDbParametersByType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbParametersByType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


