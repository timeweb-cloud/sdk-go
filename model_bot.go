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
	"time"
)

// checks if the Bot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bot{}

// Bot Бот, подключенный к AI агенту (Telegram или MAX)
type Bot struct {
	// ID бота в Timeweb Cloud
	Id string `json:"id"`
	// Логин аккаунта, которому принадлежит бот
	AccountId string `json:"account_id"`
	// ID агента, который отвечает от имени бота
	AgentId int32 `json:"agent_id"`
	// Username бота в мессенджере (без `@`)
	Username string `json:"username"`
	// Имя бота в мессенджере
	FirstName NullableString `json:"first_name"`
	// Комментарий к боту. Пустая строка, если не задан
	Description string `json:"description"`
	// Статус подключения: `pending` — ожидает подключения, `active` — работает, `error` — ошибка подключения
	Status string `json:"status"`
	// Включено ли ограничение доступа по списку username. У MAX-ботов всегда `false`
	WhitelistEnabled bool `json:"whitelist_enabled"`
	// Username пользователей, которым разрешено писать боту: в нижнем регистре, без `@`. У MAX-ботов всегда пустой
	WhitelistUsernames []string `json:"whitelist_usernames"`
	// Дата и время создания
	CreatedAt time.Time `json:"created_at"`
	// Дата и время последнего обновления
	UpdatedAt time.Time `json:"updated_at"`
}

// NewBot instantiates a new Bot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBot(id string, accountId string, agentId int32, username string, firstName NullableString, description string, status string, whitelistEnabled bool, whitelistUsernames []string, createdAt time.Time, updatedAt time.Time) *Bot {
	this := Bot{}
	this.Id = id
	this.AccountId = accountId
	this.AgentId = agentId
	this.Username = username
	this.FirstName = firstName
	this.Description = description
	this.Status = status
	this.WhitelistEnabled = whitelistEnabled
	this.WhitelistUsernames = whitelistUsernames
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewBotWithDefaults instantiates a new Bot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotWithDefaults() *Bot {
	this := Bot{}
	return &this
}

// GetId returns the Id field value
func (o *Bot) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bot) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bot) SetId(v string) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *Bot) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Bot) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Bot) SetAccountId(v string) {
	o.AccountId = v
}

// GetAgentId returns the AgentId field value
func (o *Bot) GetAgentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *Bot) GetAgentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *Bot) SetAgentId(v int32) {
	o.AgentId = v
}

// GetUsername returns the Username field value
func (o *Bot) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Bot) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Bot) SetUsername(v string) {
	o.Username = v
}

// GetFirstName returns the FirstName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bot) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bot) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// SetFirstName sets field value
func (o *Bot) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// GetDescription returns the Description field value
func (o *Bot) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Bot) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Bot) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value
func (o *Bot) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Bot) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Bot) SetStatus(v string) {
	o.Status = v
}

// GetWhitelistEnabled returns the WhitelistEnabled field value
func (o *Bot) GetWhitelistEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WhitelistEnabled
}

// GetWhitelistEnabledOk returns a tuple with the WhitelistEnabled field value
// and a boolean to check if the value has been set.
func (o *Bot) GetWhitelistEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhitelistEnabled, true
}

// SetWhitelistEnabled sets field value
func (o *Bot) SetWhitelistEnabled(v bool) {
	o.WhitelistEnabled = v
}

// GetWhitelistUsernames returns the WhitelistUsernames field value
func (o *Bot) GetWhitelistUsernames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WhitelistUsernames
}

// GetWhitelistUsernamesOk returns a tuple with the WhitelistUsernames field value
// and a boolean to check if the value has been set.
func (o *Bot) GetWhitelistUsernamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhitelistUsernames, true
}

// SetWhitelistUsernames sets field value
func (o *Bot) SetWhitelistUsernames(v []string) {
	o.WhitelistUsernames = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Bot) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Bot) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Bot) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Bot) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Bot) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Bot) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Bot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["account_id"] = o.AccountId
	toSerialize["agent_id"] = o.AgentId
	toSerialize["username"] = o.Username
	toSerialize["first_name"] = o.FirstName.Get()
	toSerialize["description"] = o.Description
	toSerialize["status"] = o.Status
	toSerialize["whitelist_enabled"] = o.WhitelistEnabled
	toSerialize["whitelist_usernames"] = o.WhitelistUsernames
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableBot struct {
	value *Bot
	isSet bool
}

func (v NullableBot) Get() *Bot {
	return v.value
}

func (v *NullableBot) Set(val *Bot) {
	v.value = val
	v.isSet = true
}

func (v NullableBot) IsSet() bool {
	return v.isSet
}

func (v *NullableBot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBot(val *Bot) *NullableBot {
	return &NullableBot{value: val, isSet: true}
}

func (v NullableBot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


