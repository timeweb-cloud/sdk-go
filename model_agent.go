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

// checks if the Agent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Agent{}

// Agent AI Agent
type Agent struct {
	// Уникальный идентификатор агента
	Id float32 `json:"id"`
	// Название агента
	Name string `json:"name"`
	// Описание агента
	Description string `json:"description"`
	// ID модели
	ModelId float32 `json:"model_id"`
	// ID провайдера
	ProviderId float32 `json:"provider_id"`
	Settings AgentSettings `json:"settings"`
	// Статус агента
	Status string `json:"status"`
	// Тип доступа к агенту
	AccessType string `json:"access_type"`
	// Всего токенов выделено агенту
	TotalTokens float32 `json:"total_tokens"`
	// Использовано токенов
	UsedTokens float32 `json:"used_tokens"`
	// Осталось токенов
	RemainingTokens float32 `json:"remaining_tokens"`
	// ID пакета токенов
	TokenPackageId float32 `json:"token_package_id"`
	// Дата обновления подписки
	SubscriptionRenewalDate time.Time `json:"subscription_renewal_date"`
	// ID баз знаний, связанных с агентом
	KnowledgeBasesIds []float32 `json:"knowledge_bases_ids"`
	// ID доступа
	AccessId float32 `json:"access_id"`
	// Дата создания агента
	CreatedAt time.Time `json:"created_at"`
}

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent(id float32, name string, description string, modelId float32, providerId float32, settings AgentSettings, status string, accessType string, totalTokens float32, usedTokens float32, remainingTokens float32, tokenPackageId float32, subscriptionRenewalDate time.Time, knowledgeBasesIds []float32, accessId float32, createdAt time.Time) *Agent {
	this := Agent{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.ModelId = modelId
	this.ProviderId = providerId
	this.Settings = settings
	this.Status = status
	this.AccessType = accessType
	this.TotalTokens = totalTokens
	this.UsedTokens = usedTokens
	this.RemainingTokens = remainingTokens
	this.TokenPackageId = tokenPackageId
	this.SubscriptionRenewalDate = subscriptionRenewalDate
	this.KnowledgeBasesIds = knowledgeBasesIds
	this.AccessId = accessId
	this.CreatedAt = createdAt
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	return &this
}

// GetId returns the Id field value
func (o *Agent) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Agent) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Agent) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Agent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Agent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Agent) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Agent) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Agent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Agent) SetDescription(v string) {
	o.Description = v
}

// GetModelId returns the ModelId field value
func (o *Agent) GetModelId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetModelIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *Agent) SetModelId(v float32) {
	o.ModelId = v
}

// GetProviderId returns the ProviderId field value
func (o *Agent) GetProviderId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetProviderIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *Agent) SetProviderId(v float32) {
	o.ProviderId = v
}

// GetSettings returns the Settings field value
func (o *Agent) GetSettings() AgentSettings {
	if o == nil {
		var ret AgentSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *Agent) GetSettingsOk() (*AgentSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *Agent) SetSettings(v AgentSettings) {
	o.Settings = v
}

// GetStatus returns the Status field value
func (o *Agent) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Agent) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Agent) SetStatus(v string) {
	o.Status = v
}

// GetAccessType returns the AccessType field value
func (o *Agent) GetAccessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *Agent) GetAccessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *Agent) SetAccessType(v string) {
	o.AccessType = v
}

// GetTotalTokens returns the TotalTokens field value
func (o *Agent) GetTotalTokens() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value
// and a boolean to check if the value has been set.
func (o *Agent) GetTotalTokensOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTokens, true
}

// SetTotalTokens sets field value
func (o *Agent) SetTotalTokens(v float32) {
	o.TotalTokens = v
}

// GetUsedTokens returns the UsedTokens field value
func (o *Agent) GetUsedTokens() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UsedTokens
}

// GetUsedTokensOk returns a tuple with the UsedTokens field value
// and a boolean to check if the value has been set.
func (o *Agent) GetUsedTokensOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedTokens, true
}

// SetUsedTokens sets field value
func (o *Agent) SetUsedTokens(v float32) {
	o.UsedTokens = v
}

// GetRemainingTokens returns the RemainingTokens field value
func (o *Agent) GetRemainingTokens() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RemainingTokens
}

// GetRemainingTokensOk returns a tuple with the RemainingTokens field value
// and a boolean to check if the value has been set.
func (o *Agent) GetRemainingTokensOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingTokens, true
}

// SetRemainingTokens sets field value
func (o *Agent) SetRemainingTokens(v float32) {
	o.RemainingTokens = v
}

// GetTokenPackageId returns the TokenPackageId field value
func (o *Agent) GetTokenPackageId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TokenPackageId
}

// GetTokenPackageIdOk returns a tuple with the TokenPackageId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetTokenPackageIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenPackageId, true
}

// SetTokenPackageId sets field value
func (o *Agent) SetTokenPackageId(v float32) {
	o.TokenPackageId = v
}

// GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field value
func (o *Agent) GetSubscriptionRenewalDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SubscriptionRenewalDate
}

// GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field value
// and a boolean to check if the value has been set.
func (o *Agent) GetSubscriptionRenewalDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionRenewalDate, true
}

// SetSubscriptionRenewalDate sets field value
func (o *Agent) SetSubscriptionRenewalDate(v time.Time) {
	o.SubscriptionRenewalDate = v
}

// GetKnowledgeBasesIds returns the KnowledgeBasesIds field value
func (o *Agent) GetKnowledgeBasesIds() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.KnowledgeBasesIds
}

// GetKnowledgeBasesIdsOk returns a tuple with the KnowledgeBasesIds field value
// and a boolean to check if the value has been set.
func (o *Agent) GetKnowledgeBasesIdsOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.KnowledgeBasesIds, true
}

// SetKnowledgeBasesIds sets field value
func (o *Agent) SetKnowledgeBasesIds(v []float32) {
	o.KnowledgeBasesIds = v
}

// GetAccessId returns the AccessId field value
func (o *Agent) GetAccessId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccessId
}

// GetAccessIdOk returns a tuple with the AccessId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetAccessIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessId, true
}

// SetAccessId sets field value
func (o *Agent) SetAccessId(v float32) {
	o.AccessId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Agent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Agent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Agent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Agent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["model_id"] = o.ModelId
	toSerialize["provider_id"] = o.ProviderId
	toSerialize["settings"] = o.Settings
	toSerialize["status"] = o.Status
	toSerialize["access_type"] = o.AccessType
	toSerialize["total_tokens"] = o.TotalTokens
	toSerialize["used_tokens"] = o.UsedTokens
	toSerialize["remaining_tokens"] = o.RemainingTokens
	toSerialize["token_package_id"] = o.TokenPackageId
	toSerialize["subscription_renewal_date"] = o.SubscriptionRenewalDate
	toSerialize["knowledge_bases_ids"] = o.KnowledgeBasesIds
	toSerialize["access_id"] = o.AccessId
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


