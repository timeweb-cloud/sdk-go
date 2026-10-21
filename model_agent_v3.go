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

// checks if the AgentV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentV3{}

// AgentV3 AI агент (API v3)
type AgentV3 struct {
	// ID агента
	Id int32 `json:"id"`
	// Название агента
	Name string `json:"name"`
	// Описание агента
	Description NullableString `json:"description"`
	// ID проекта, в котором находится агент
	ProjectId *int32 `json:"project_id,omitempty"`
	// ID основной модели агента
	ModelId int32 `json:"model_id"`
	// ID провайдера модели
	ProviderId int32 `json:"provider_id"`
	Settings AgentSettings `json:"settings"`
	// Статус агента: - `active` — работает; - `suspended` — приостановлен владельцем; - `blocked` — заблокирован (для агентов с пакетом токенов — закончились токены или подписка); - `deleted` — удалён.
	Status string `json:"status"`
	// Тип доступа: `public` — агент отвечает анонимным посетителям сайта через виджет; `private` — только по токену аккаунта.
	AccessType string `json:"access_type"`
	// Идентификатор доступа агента. Используется для вызова агента из своих приложений через OpenAI-совместимое API (см. https://agent.timeweb.cloud/docs) и для встраивания виджета.
	AccessId string `json:"access_id"`
	// Устарело. Дневной лимит агента в токенах без привязки к модели, если он задан; `0` — такого лимита нет. Лимиты в деньгах, по моделям и за другие периоды здесь не отражаются — используйте лимиты агента (`/api/v1/cloud-ai/agents/{id}/entity-limits`).
	// Deprecated
	TokenLimit int32 `json:"token_limit"`
	// Всего токенов в пакете. Только для агентов с пакетом токенов (устаревшая схема оплаты).
	TotalTokens *int32 `json:"total_tokens,omitempty"`
	// Использовано токенов пакета. Для агентов с поресурсной оплатой — `null`.
	UsedTokens NullableInt32 `json:"used_tokens"`
	// Осталось токенов в пакете. Только для агентов с пакетом токенов.
	RemainingTokens *int32 `json:"remaining_tokens,omitempty"`
	// ID пакета токенов. Только для агентов с пакетом токенов; у агентов с поресурсной оплатой поле отсутствует.
	TokenPackageId *int32 `json:"token_package_id,omitempty"`
	// Дата продления подписки на пакет токенов. Только для агентов с пакетом токенов.
	SubscriptionRenewalDate *time.Time `json:"subscription_renewal_date,omitempty"`
	// Включена ли автоматическая докупка пакетов токенов
	AutoBuyAdditionalTokens bool `json:"auto_buy_additional_tokens"`
	// Инструменты, подключённые к агенту
	Tools []AgentToolsInner `json:"tools"`
	// ID дополнительных моделей агента
	AdditionalModelIds []int32 `json:"additional_model_ids"`
	// Дополнительные модели агента с настройками
	AdditionalModels []AgentAdditionalModel `json:"additional_models"`
	// Включён ли веб-поиск
	IsWebSearchEnabled bool `json:"is_web_search_enabled"`
	// ID ключа AI Gateway, к которому привязан агент. `null` — агента можно привязать к ключу (см. обновление агента); поле отсутствует — привязка недоступна.
	AiGatewayKeyId NullableString `json:"ai_gateway_key_id,omitempty"`
	// Дата создания агента
	CreatedAt time.Time `json:"created_at"`
	// Базы знаний, подключённые к агенту, с режимом поиска
	KnowledgeBases []AgentV3KnowledgeBasesInner `json:"knowledge_bases"`
}

// NewAgentV3 instantiates a new AgentV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentV3(id int32, name string, description NullableString, modelId int32, providerId int32, settings AgentSettings, status string, accessType string, accessId string, tokenLimit int32, usedTokens NullableInt32, autoBuyAdditionalTokens bool, tools []AgentToolsInner, additionalModelIds []int32, additionalModels []AgentAdditionalModel, isWebSearchEnabled bool, createdAt time.Time, knowledgeBases []AgentV3KnowledgeBasesInner) *AgentV3 {
	this := AgentV3{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.ModelId = modelId
	this.ProviderId = providerId
	this.Settings = settings
	this.Status = status
	this.AccessType = accessType
	this.AccessId = accessId
	this.TokenLimit = tokenLimit
	this.UsedTokens = usedTokens
	this.AutoBuyAdditionalTokens = autoBuyAdditionalTokens
	this.Tools = tools
	this.AdditionalModelIds = additionalModelIds
	this.AdditionalModels = additionalModels
	this.IsWebSearchEnabled = isWebSearchEnabled
	this.CreatedAt = createdAt
	this.KnowledgeBases = knowledgeBases
	return &this
}

// NewAgentV3WithDefaults instantiates a new AgentV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentV3WithDefaults() *AgentV3 {
	this := AgentV3{}
	return &this
}

// GetId returns the Id field value
func (o *AgentV3) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentV3) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AgentV3) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AgentV3) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AgentV3) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentV3) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *AgentV3) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AgentV3) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentV3) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AgentV3) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *AgentV3) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetModelId returns the ModelId field value
func (o *AgentV3) GetModelId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetModelIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *AgentV3) SetModelId(v int32) {
	o.ModelId = v
}

// GetProviderId returns the ProviderId field value
func (o *AgentV3) GetProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *AgentV3) SetProviderId(v int32) {
	o.ProviderId = v
}

// GetSettings returns the Settings field value
func (o *AgentV3) GetSettings() AgentSettings {
	if o == nil {
		var ret AgentSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetSettingsOk() (*AgentSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *AgentV3) SetSettings(v AgentSettings) {
	o.Settings = v
}

// GetStatus returns the Status field value
func (o *AgentV3) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AgentV3) SetStatus(v string) {
	o.Status = v
}

// GetAccessType returns the AccessType field value
func (o *AgentV3) GetAccessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetAccessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *AgentV3) SetAccessType(v string) {
	o.AccessType = v
}

// GetAccessId returns the AccessId field value
func (o *AgentV3) GetAccessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetAccessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessId, true
}

// SetAccessId sets field value
func (o *AgentV3) SetAccessId(v string) {
	o.AccessId = v
}

// GetTokenLimit returns the TokenLimit field value
// Deprecated
func (o *AgentV3) GetTokenLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenLimit
}

// GetTokenLimitOk returns a tuple with the TokenLimit field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *AgentV3) GetTokenLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenLimit, true
}

// SetTokenLimit sets field value
// Deprecated
func (o *AgentV3) SetTokenLimit(v int32) {
	o.TokenLimit = v
}

// GetTotalTokens returns the TotalTokens field value if set, zero value otherwise.
func (o *AgentV3) GetTotalTokens() int32 {
	if o == nil || IsNil(o.TotalTokens) {
		var ret int32
		return ret
	}
	return *o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentV3) GetTotalTokensOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalTokens) {
		return nil, false
	}
	return o.TotalTokens, true
}

// HasTotalTokens returns a boolean if a field has been set.
func (o *AgentV3) HasTotalTokens() bool {
	if o != nil && !IsNil(o.TotalTokens) {
		return true
	}

	return false
}

// SetTotalTokens gets a reference to the given int32 and assigns it to the TotalTokens field.
func (o *AgentV3) SetTotalTokens(v int32) {
	o.TotalTokens = &v
}

// GetUsedTokens returns the UsedTokens field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AgentV3) GetUsedTokens() int32 {
	if o == nil || o.UsedTokens.Get() == nil {
		var ret int32
		return ret
	}

	return *o.UsedTokens.Get()
}

// GetUsedTokensOk returns a tuple with the UsedTokens field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentV3) GetUsedTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedTokens.Get(), o.UsedTokens.IsSet()
}

// SetUsedTokens sets field value
func (o *AgentV3) SetUsedTokens(v int32) {
	o.UsedTokens.Set(&v)
}

// GetRemainingTokens returns the RemainingTokens field value if set, zero value otherwise.
func (o *AgentV3) GetRemainingTokens() int32 {
	if o == nil || IsNil(o.RemainingTokens) {
		var ret int32
		return ret
	}
	return *o.RemainingTokens
}

// GetRemainingTokensOk returns a tuple with the RemainingTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentV3) GetRemainingTokensOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainingTokens) {
		return nil, false
	}
	return o.RemainingTokens, true
}

// HasRemainingTokens returns a boolean if a field has been set.
func (o *AgentV3) HasRemainingTokens() bool {
	if o != nil && !IsNil(o.RemainingTokens) {
		return true
	}

	return false
}

// SetRemainingTokens gets a reference to the given int32 and assigns it to the RemainingTokens field.
func (o *AgentV3) SetRemainingTokens(v int32) {
	o.RemainingTokens = &v
}

// GetTokenPackageId returns the TokenPackageId field value if set, zero value otherwise.
func (o *AgentV3) GetTokenPackageId() int32 {
	if o == nil || IsNil(o.TokenPackageId) {
		var ret int32
		return ret
	}
	return *o.TokenPackageId
}

// GetTokenPackageIdOk returns a tuple with the TokenPackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentV3) GetTokenPackageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TokenPackageId) {
		return nil, false
	}
	return o.TokenPackageId, true
}

// HasTokenPackageId returns a boolean if a field has been set.
func (o *AgentV3) HasTokenPackageId() bool {
	if o != nil && !IsNil(o.TokenPackageId) {
		return true
	}

	return false
}

// SetTokenPackageId gets a reference to the given int32 and assigns it to the TokenPackageId field.
func (o *AgentV3) SetTokenPackageId(v int32) {
	o.TokenPackageId = &v
}

// GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field value if set, zero value otherwise.
func (o *AgentV3) GetSubscriptionRenewalDate() time.Time {
	if o == nil || IsNil(o.SubscriptionRenewalDate) {
		var ret time.Time
		return ret
	}
	return *o.SubscriptionRenewalDate
}

// GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentV3) GetSubscriptionRenewalDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubscriptionRenewalDate) {
		return nil, false
	}
	return o.SubscriptionRenewalDate, true
}

// HasSubscriptionRenewalDate returns a boolean if a field has been set.
func (o *AgentV3) HasSubscriptionRenewalDate() bool {
	if o != nil && !IsNil(o.SubscriptionRenewalDate) {
		return true
	}

	return false
}

// SetSubscriptionRenewalDate gets a reference to the given time.Time and assigns it to the SubscriptionRenewalDate field.
func (o *AgentV3) SetSubscriptionRenewalDate(v time.Time) {
	o.SubscriptionRenewalDate = &v
}

// GetAutoBuyAdditionalTokens returns the AutoBuyAdditionalTokens field value
func (o *AgentV3) GetAutoBuyAdditionalTokens() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoBuyAdditionalTokens
}

// GetAutoBuyAdditionalTokensOk returns a tuple with the AutoBuyAdditionalTokens field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetAutoBuyAdditionalTokensOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoBuyAdditionalTokens, true
}

// SetAutoBuyAdditionalTokens sets field value
func (o *AgentV3) SetAutoBuyAdditionalTokens(v bool) {
	o.AutoBuyAdditionalTokens = v
}

// GetTools returns the Tools field value
func (o *AgentV3) GetTools() []AgentToolsInner {
	if o == nil {
		var ret []AgentToolsInner
		return ret
	}

	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetToolsOk() ([]AgentToolsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tools, true
}

// SetTools sets field value
func (o *AgentV3) SetTools(v []AgentToolsInner) {
	o.Tools = v
}

// GetAdditionalModelIds returns the AdditionalModelIds field value
func (o *AgentV3) GetAdditionalModelIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AdditionalModelIds
}

// GetAdditionalModelIdsOk returns a tuple with the AdditionalModelIds field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetAdditionalModelIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalModelIds, true
}

// SetAdditionalModelIds sets field value
func (o *AgentV3) SetAdditionalModelIds(v []int32) {
	o.AdditionalModelIds = v
}

// GetAdditionalModels returns the AdditionalModels field value
func (o *AgentV3) GetAdditionalModels() []AgentAdditionalModel {
	if o == nil {
		var ret []AgentAdditionalModel
		return ret
	}

	return o.AdditionalModels
}

// GetAdditionalModelsOk returns a tuple with the AdditionalModels field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetAdditionalModelsOk() ([]AgentAdditionalModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalModels, true
}

// SetAdditionalModels sets field value
func (o *AgentV3) SetAdditionalModels(v []AgentAdditionalModel) {
	o.AdditionalModels = v
}

// GetIsWebSearchEnabled returns the IsWebSearchEnabled field value
func (o *AgentV3) GetIsWebSearchEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWebSearchEnabled
}

// GetIsWebSearchEnabledOk returns a tuple with the IsWebSearchEnabled field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetIsWebSearchEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWebSearchEnabled, true
}

// SetIsWebSearchEnabled sets field value
func (o *AgentV3) SetIsWebSearchEnabled(v bool) {
	o.IsWebSearchEnabled = v
}

// GetAiGatewayKeyId returns the AiGatewayKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentV3) GetAiGatewayKeyId() string {
	if o == nil || IsNil(o.AiGatewayKeyId.Get()) {
		var ret string
		return ret
	}
	return *o.AiGatewayKeyId.Get()
}

// GetAiGatewayKeyIdOk returns a tuple with the AiGatewayKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentV3) GetAiGatewayKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AiGatewayKeyId.Get(), o.AiGatewayKeyId.IsSet()
}

// HasAiGatewayKeyId returns a boolean if a field has been set.
func (o *AgentV3) HasAiGatewayKeyId() bool {
	if o != nil && o.AiGatewayKeyId.IsSet() {
		return true
	}

	return false
}

// SetAiGatewayKeyId gets a reference to the given NullableString and assigns it to the AiGatewayKeyId field.
func (o *AgentV3) SetAiGatewayKeyId(v string) {
	o.AiGatewayKeyId.Set(&v)
}
// SetAiGatewayKeyIdNil sets the value for AiGatewayKeyId to be an explicit nil
func (o *AgentV3) SetAiGatewayKeyIdNil() {
	o.AiGatewayKeyId.Set(nil)
}

// UnsetAiGatewayKeyId ensures that no value is present for AiGatewayKeyId, not even an explicit nil
func (o *AgentV3) UnsetAiGatewayKeyId() {
	o.AiGatewayKeyId.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AgentV3) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AgentV3) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetKnowledgeBases returns the KnowledgeBases field value
func (o *AgentV3) GetKnowledgeBases() []AgentV3KnowledgeBasesInner {
	if o == nil {
		var ret []AgentV3KnowledgeBasesInner
		return ret
	}

	return o.KnowledgeBases
}

// GetKnowledgeBasesOk returns a tuple with the KnowledgeBases field value
// and a boolean to check if the value has been set.
func (o *AgentV3) GetKnowledgeBasesOk() ([]AgentV3KnowledgeBasesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.KnowledgeBases, true
}

// SetKnowledgeBases sets field value
func (o *AgentV3) SetKnowledgeBases(v []AgentV3KnowledgeBasesInner) {
	o.KnowledgeBases = v
}

func (o AgentV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	toSerialize["model_id"] = o.ModelId
	toSerialize["provider_id"] = o.ProviderId
	toSerialize["settings"] = o.Settings
	toSerialize["status"] = o.Status
	toSerialize["access_type"] = o.AccessType
	toSerialize["access_id"] = o.AccessId
	toSerialize["token_limit"] = o.TokenLimit
	if !IsNil(o.TotalTokens) {
		toSerialize["total_tokens"] = o.TotalTokens
	}
	toSerialize["used_tokens"] = o.UsedTokens.Get()
	if !IsNil(o.RemainingTokens) {
		toSerialize["remaining_tokens"] = o.RemainingTokens
	}
	if !IsNil(o.TokenPackageId) {
		toSerialize["token_package_id"] = o.TokenPackageId
	}
	if !IsNil(o.SubscriptionRenewalDate) {
		toSerialize["subscription_renewal_date"] = o.SubscriptionRenewalDate
	}
	toSerialize["auto_buy_additional_tokens"] = o.AutoBuyAdditionalTokens
	toSerialize["tools"] = o.Tools
	toSerialize["additional_model_ids"] = o.AdditionalModelIds
	toSerialize["additional_models"] = o.AdditionalModels
	toSerialize["is_web_search_enabled"] = o.IsWebSearchEnabled
	if o.AiGatewayKeyId.IsSet() {
		toSerialize["ai_gateway_key_id"] = o.AiGatewayKeyId.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["knowledge_bases"] = o.KnowledgeBases
	return toSerialize, nil
}

type NullableAgentV3 struct {
	value *AgentV3
	isSet bool
}

func (v NullableAgentV3) Get() *AgentV3 {
	return v.value
}

func (v *NullableAgentV3) Set(val *AgentV3) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentV3) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentV3(val *AgentV3) *NullableAgentV3 {
	return &NullableAgentV3{value: val, isSet: true}
}

func (v NullableAgentV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


