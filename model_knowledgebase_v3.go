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

// checks if the KnowledgebaseV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KnowledgebaseV3{}

// KnowledgebaseV3 База знаний (v3): количество документов и привязанные агенты с режимом поиска
type KnowledgebaseV3 struct {
	// ID базы знаний
	Id int32 `json:"id"`
	// Название базы знаний
	Name string `json:"name"`
	// Описание базы знаний
	Description NullableString `json:"description"`
	// ID проекта, в котором находится база знаний. Поля нет, если проект не определён
	ProjectId *int32 `json:"project_id,omitempty"`
	// ID кластера базы данных OpenSearch, который создан для базы знаний в вашем аккаунте
	DbaasId int32 `json:"dbaas_id"`
	// Статус базы знаний: - `creating` — создаётся кластер базы данных, документы будут проиндексированы после перехода в `active` - `active` — база знаний работает - `blocked` — заблокирована: закончились токены или не удалось продлить подписку - `deleted` — удалена
	Status string `json:"status"`
	// Дата и время последней синхронизации. `null`, если синхронизации не было
	LastSync NullableTime `json:"last_sync"`
	// Количество токенов, доступных в текущем периоде подписки (пакет подписки и докупленные пакеты)
	TotalTokens int32 `json:"total_tokens"`
	// Количество израсходованных токенов
	UsedTokens int32 `json:"used_tokens"`
	// Количество оставшихся токенов
	RemainingTokens int32 `json:"remaining_tokens"`
	// ID пакета токенов подписки базы знаний
	TokenPackageId int32 `json:"token_package_id"`
	// Дата следующего продления подписки (списания стоимости пакета токенов)
	SubscriptionRenewalDate NullableTime `json:"subscription_renewal_date"`
	// Количество документов в базе знаний
	DocumentsCount int32 `json:"documents_count"`
	// Агенты, к которым привязана база знаний, и режим поиска для каждого
	Agents []RetrievalModeRef `json:"agents"`
	// Дата и время создания базы знаний
	CreatedAt time.Time `json:"created_at"`
}

// NewKnowledgebaseV3 instantiates a new KnowledgebaseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnowledgebaseV3(id int32, name string, description NullableString, dbaasId int32, status string, lastSync NullableTime, totalTokens int32, usedTokens int32, remainingTokens int32, tokenPackageId int32, subscriptionRenewalDate NullableTime, documentsCount int32, agents []RetrievalModeRef, createdAt time.Time) *KnowledgebaseV3 {
	this := KnowledgebaseV3{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.DbaasId = dbaasId
	this.Status = status
	this.LastSync = lastSync
	this.TotalTokens = totalTokens
	this.UsedTokens = usedTokens
	this.RemainingTokens = remainingTokens
	this.TokenPackageId = tokenPackageId
	this.SubscriptionRenewalDate = subscriptionRenewalDate
	this.DocumentsCount = documentsCount
	this.Agents = agents
	this.CreatedAt = createdAt
	return &this
}

// NewKnowledgebaseV3WithDefaults instantiates a new KnowledgebaseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnowledgebaseV3WithDefaults() *KnowledgebaseV3 {
	this := KnowledgebaseV3{}
	return &this
}

// GetId returns the Id field value
func (o *KnowledgebaseV3) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KnowledgebaseV3) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KnowledgebaseV3) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KnowledgebaseV3) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KnowledgebaseV3) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgebaseV3) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *KnowledgebaseV3) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *KnowledgebaseV3) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *KnowledgebaseV3) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *KnowledgebaseV3) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetDbaasId returns the DbaasId field value
func (o *KnowledgebaseV3) GetDbaasId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DbaasId
}

// GetDbaasIdOk returns a tuple with the DbaasId field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetDbaasIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbaasId, true
}

// SetDbaasId sets field value
func (o *KnowledgebaseV3) SetDbaasId(v int32) {
	o.DbaasId = v
}

// GetStatus returns the Status field value
func (o *KnowledgebaseV3) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *KnowledgebaseV3) SetStatus(v string) {
	o.Status = v
}

// GetLastSync returns the LastSync field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *KnowledgebaseV3) GetLastSync() time.Time {
	if o == nil || o.LastSync.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastSync.Get()
}

// GetLastSyncOk returns a tuple with the LastSync field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgebaseV3) GetLastSyncOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSync.Get(), o.LastSync.IsSet()
}

// SetLastSync sets field value
func (o *KnowledgebaseV3) SetLastSync(v time.Time) {
	o.LastSync.Set(&v)
}

// GetTotalTokens returns the TotalTokens field value
func (o *KnowledgebaseV3) GetTotalTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetTotalTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTokens, true
}

// SetTotalTokens sets field value
func (o *KnowledgebaseV3) SetTotalTokens(v int32) {
	o.TotalTokens = v
}

// GetUsedTokens returns the UsedTokens field value
func (o *KnowledgebaseV3) GetUsedTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsedTokens
}

// GetUsedTokensOk returns a tuple with the UsedTokens field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetUsedTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedTokens, true
}

// SetUsedTokens sets field value
func (o *KnowledgebaseV3) SetUsedTokens(v int32) {
	o.UsedTokens = v
}

// GetRemainingTokens returns the RemainingTokens field value
func (o *KnowledgebaseV3) GetRemainingTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemainingTokens
}

// GetRemainingTokensOk returns a tuple with the RemainingTokens field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetRemainingTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingTokens, true
}

// SetRemainingTokens sets field value
func (o *KnowledgebaseV3) SetRemainingTokens(v int32) {
	o.RemainingTokens = v
}

// GetTokenPackageId returns the TokenPackageId field value
func (o *KnowledgebaseV3) GetTokenPackageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenPackageId
}

// GetTokenPackageIdOk returns a tuple with the TokenPackageId field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetTokenPackageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenPackageId, true
}

// SetTokenPackageId sets field value
func (o *KnowledgebaseV3) SetTokenPackageId(v int32) {
	o.TokenPackageId = v
}

// GetSubscriptionRenewalDate returns the SubscriptionRenewalDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *KnowledgebaseV3) GetSubscriptionRenewalDate() time.Time {
	if o == nil || o.SubscriptionRenewalDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.SubscriptionRenewalDate.Get()
}

// GetSubscriptionRenewalDateOk returns a tuple with the SubscriptionRenewalDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgebaseV3) GetSubscriptionRenewalDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionRenewalDate.Get(), o.SubscriptionRenewalDate.IsSet()
}

// SetSubscriptionRenewalDate sets field value
func (o *KnowledgebaseV3) SetSubscriptionRenewalDate(v time.Time) {
	o.SubscriptionRenewalDate.Set(&v)
}

// GetDocumentsCount returns the DocumentsCount field value
func (o *KnowledgebaseV3) GetDocumentsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DocumentsCount
}

// GetDocumentsCountOk returns a tuple with the DocumentsCount field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetDocumentsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentsCount, true
}

// SetDocumentsCount sets field value
func (o *KnowledgebaseV3) SetDocumentsCount(v int32) {
	o.DocumentsCount = v
}

// GetAgents returns the Agents field value
func (o *KnowledgebaseV3) GetAgents() []RetrievalModeRef {
	if o == nil {
		var ret []RetrievalModeRef
		return ret
	}

	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetAgentsOk() ([]RetrievalModeRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.Agents, true
}

// SetAgents sets field value
func (o *KnowledgebaseV3) SetAgents(v []RetrievalModeRef) {
	o.Agents = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *KnowledgebaseV3) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *KnowledgebaseV3) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *KnowledgebaseV3) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o KnowledgebaseV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KnowledgebaseV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	toSerialize["dbaas_id"] = o.DbaasId
	toSerialize["status"] = o.Status
	toSerialize["last_sync"] = o.LastSync.Get()
	toSerialize["total_tokens"] = o.TotalTokens
	toSerialize["used_tokens"] = o.UsedTokens
	toSerialize["remaining_tokens"] = o.RemainingTokens
	toSerialize["token_package_id"] = o.TokenPackageId
	toSerialize["subscription_renewal_date"] = o.SubscriptionRenewalDate.Get()
	toSerialize["documents_count"] = o.DocumentsCount
	toSerialize["agents"] = o.Agents
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableKnowledgebaseV3 struct {
	value *KnowledgebaseV3
	isSet bool
}

func (v NullableKnowledgebaseV3) Get() *KnowledgebaseV3 {
	return v.value
}

func (v *NullableKnowledgebaseV3) Set(val *KnowledgebaseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableKnowledgebaseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableKnowledgebaseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnowledgebaseV3(val *KnowledgebaseV3) *NullableKnowledgebaseV3 {
	return &NullableKnowledgebaseV3{value: val, isSet: true}
}

func (v NullableKnowledgebaseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnowledgebaseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


