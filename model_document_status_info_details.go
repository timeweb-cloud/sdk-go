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

// checks if the DocumentStatusInfoDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentStatusInfoDetails{}

// DocumentStatusInfoDetails Подробности, набор полей зависит от состояния. Могут присутствовать и другие поля
type DocumentStatusInfoDetails struct {
	// `indexing`: время начала индексации
	StartedAt *time.Time `json:"started_at,omitempty"`
	// `indexing`: текущий этап
	CurrentStage *string `json:"current_stage,omitempty"`
	// `indexing`: прогресс в процентах
	Progress *float32 `json:"progress,omitempty"`
	// `indexed`: длительность индексации в миллисекундах
	IndexingDurationMs *float32 `json:"indexing_duration_ms,omitempty"`
	// `indexed`: израсходовано токенов
	TokensUsed *float32 `json:"tokens_used,omitempty"`
	// `indexed`: оценка количества токенов перед индексацией
	TokensEstimated *float32 `json:"tokens_estimated,omitempty"`
	// `indexed`: количество фрагментов, на которые разбит документ
	ChunksCount *float32 `json:"chunks_count,omitempty"`
	// `INSUFFICIENT_TOKENS`: сколько токенов нужно
	Needed *float32 `json:"needed,omitempty"`
	// `INSUFFICIENT_TOKENS`: сколько токенов доступно
	Available *float32 `json:"available,omitempty"`
	// `INSUFFICIENT_TOKENS`: оценка без запаса
	Estimated *float32 `json:"estimated,omitempty"`
	// `CRITICAL_BALANCE`, `NEGATIVE_BALANCE`: текущий остаток токенов
	Balance *float32 `json:"balance,omitempty"`
	// `CRITICAL_BALANCE`: минимально необходимый остаток
	MinRequired *float32 `json:"min_required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DocumentStatusInfoDetails DocumentStatusInfoDetails

// NewDocumentStatusInfoDetails instantiates a new DocumentStatusInfoDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentStatusInfoDetails() *DocumentStatusInfoDetails {
	this := DocumentStatusInfoDetails{}
	return &this
}

// NewDocumentStatusInfoDetailsWithDefaults instantiates a new DocumentStatusInfoDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentStatusInfoDetailsWithDefaults() *DocumentStatusInfoDetails {
	this := DocumentStatusInfoDetails{}
	return &this
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DocumentStatusInfoDetails) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCurrentStage returns the CurrentStage field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetCurrentStage() string {
	if o == nil || IsNil(o.CurrentStage) {
		var ret string
		return ret
	}
	return *o.CurrentStage
}

// GetCurrentStageOk returns a tuple with the CurrentStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetCurrentStageOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentStage) {
		return nil, false
	}
	return o.CurrentStage, true
}

// HasCurrentStage returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasCurrentStage() bool {
	if o != nil && !IsNil(o.CurrentStage) {
		return true
	}

	return false
}

// SetCurrentStage gets a reference to the given string and assigns it to the CurrentStage field.
func (o *DocumentStatusInfoDetails) SetCurrentStage(v string) {
	o.CurrentStage = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetProgress() float32 {
	if o == nil || IsNil(o.Progress) {
		var ret float32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetProgressOk() (*float32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float32 and assigns it to the Progress field.
func (o *DocumentStatusInfoDetails) SetProgress(v float32) {
	o.Progress = &v
}

// GetIndexingDurationMs returns the IndexingDurationMs field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetIndexingDurationMs() float32 {
	if o == nil || IsNil(o.IndexingDurationMs) {
		var ret float32
		return ret
	}
	return *o.IndexingDurationMs
}

// GetIndexingDurationMsOk returns a tuple with the IndexingDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetIndexingDurationMsOk() (*float32, bool) {
	if o == nil || IsNil(o.IndexingDurationMs) {
		return nil, false
	}
	return o.IndexingDurationMs, true
}

// HasIndexingDurationMs returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasIndexingDurationMs() bool {
	if o != nil && !IsNil(o.IndexingDurationMs) {
		return true
	}

	return false
}

// SetIndexingDurationMs gets a reference to the given float32 and assigns it to the IndexingDurationMs field.
func (o *DocumentStatusInfoDetails) SetIndexingDurationMs(v float32) {
	o.IndexingDurationMs = &v
}

// GetTokensUsed returns the TokensUsed field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetTokensUsed() float32 {
	if o == nil || IsNil(o.TokensUsed) {
		var ret float32
		return ret
	}
	return *o.TokensUsed
}

// GetTokensUsedOk returns a tuple with the TokensUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetTokensUsedOk() (*float32, bool) {
	if o == nil || IsNil(o.TokensUsed) {
		return nil, false
	}
	return o.TokensUsed, true
}

// HasTokensUsed returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasTokensUsed() bool {
	if o != nil && !IsNil(o.TokensUsed) {
		return true
	}

	return false
}

// SetTokensUsed gets a reference to the given float32 and assigns it to the TokensUsed field.
func (o *DocumentStatusInfoDetails) SetTokensUsed(v float32) {
	o.TokensUsed = &v
}

// GetTokensEstimated returns the TokensEstimated field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetTokensEstimated() float32 {
	if o == nil || IsNil(o.TokensEstimated) {
		var ret float32
		return ret
	}
	return *o.TokensEstimated
}

// GetTokensEstimatedOk returns a tuple with the TokensEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetTokensEstimatedOk() (*float32, bool) {
	if o == nil || IsNil(o.TokensEstimated) {
		return nil, false
	}
	return o.TokensEstimated, true
}

// HasTokensEstimated returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasTokensEstimated() bool {
	if o != nil && !IsNil(o.TokensEstimated) {
		return true
	}

	return false
}

// SetTokensEstimated gets a reference to the given float32 and assigns it to the TokensEstimated field.
func (o *DocumentStatusInfoDetails) SetTokensEstimated(v float32) {
	o.TokensEstimated = &v
}

// GetChunksCount returns the ChunksCount field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetChunksCount() float32 {
	if o == nil || IsNil(o.ChunksCount) {
		var ret float32
		return ret
	}
	return *o.ChunksCount
}

// GetChunksCountOk returns a tuple with the ChunksCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetChunksCountOk() (*float32, bool) {
	if o == nil || IsNil(o.ChunksCount) {
		return nil, false
	}
	return o.ChunksCount, true
}

// HasChunksCount returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasChunksCount() bool {
	if o != nil && !IsNil(o.ChunksCount) {
		return true
	}

	return false
}

// SetChunksCount gets a reference to the given float32 and assigns it to the ChunksCount field.
func (o *DocumentStatusInfoDetails) SetChunksCount(v float32) {
	o.ChunksCount = &v
}

// GetNeeded returns the Needed field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetNeeded() float32 {
	if o == nil || IsNil(o.Needed) {
		var ret float32
		return ret
	}
	return *o.Needed
}

// GetNeededOk returns a tuple with the Needed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetNeededOk() (*float32, bool) {
	if o == nil || IsNil(o.Needed) {
		return nil, false
	}
	return o.Needed, true
}

// HasNeeded returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasNeeded() bool {
	if o != nil && !IsNil(o.Needed) {
		return true
	}

	return false
}

// SetNeeded gets a reference to the given float32 and assigns it to the Needed field.
func (o *DocumentStatusInfoDetails) SetNeeded(v float32) {
	o.Needed = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetAvailable() float32 {
	if o == nil || IsNil(o.Available) {
		var ret float32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetAvailableOk() (*float32, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given float32 and assigns it to the Available field.
func (o *DocumentStatusInfoDetails) SetAvailable(v float32) {
	o.Available = &v
}

// GetEstimated returns the Estimated field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetEstimated() float32 {
	if o == nil || IsNil(o.Estimated) {
		var ret float32
		return ret
	}
	return *o.Estimated
}

// GetEstimatedOk returns a tuple with the Estimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetEstimatedOk() (*float32, bool) {
	if o == nil || IsNil(o.Estimated) {
		return nil, false
	}
	return o.Estimated, true
}

// HasEstimated returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasEstimated() bool {
	if o != nil && !IsNil(o.Estimated) {
		return true
	}

	return false
}

// SetEstimated gets a reference to the given float32 and assigns it to the Estimated field.
func (o *DocumentStatusInfoDetails) SetEstimated(v float32) {
	o.Estimated = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *DocumentStatusInfoDetails) SetBalance(v float32) {
	o.Balance = &v
}

// GetMinRequired returns the MinRequired field value if set, zero value otherwise.
func (o *DocumentStatusInfoDetails) GetMinRequired() float32 {
	if o == nil || IsNil(o.MinRequired) {
		var ret float32
		return ret
	}
	return *o.MinRequired
}

// GetMinRequiredOk returns a tuple with the MinRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatusInfoDetails) GetMinRequiredOk() (*float32, bool) {
	if o == nil || IsNil(o.MinRequired) {
		return nil, false
	}
	return o.MinRequired, true
}

// HasMinRequired returns a boolean if a field has been set.
func (o *DocumentStatusInfoDetails) HasMinRequired() bool {
	if o != nil && !IsNil(o.MinRequired) {
		return true
	}

	return false
}

// SetMinRequired gets a reference to the given float32 and assigns it to the MinRequired field.
func (o *DocumentStatusInfoDetails) SetMinRequired(v float32) {
	o.MinRequired = &v
}

func (o DocumentStatusInfoDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentStatusInfoDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.CurrentStage) {
		toSerialize["current_stage"] = o.CurrentStage
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.IndexingDurationMs) {
		toSerialize["indexing_duration_ms"] = o.IndexingDurationMs
	}
	if !IsNil(o.TokensUsed) {
		toSerialize["tokens_used"] = o.TokensUsed
	}
	if !IsNil(o.TokensEstimated) {
		toSerialize["tokens_estimated"] = o.TokensEstimated
	}
	if !IsNil(o.ChunksCount) {
		toSerialize["chunks_count"] = o.ChunksCount
	}
	if !IsNil(o.Needed) {
		toSerialize["needed"] = o.Needed
	}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Estimated) {
		toSerialize["estimated"] = o.Estimated
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.MinRequired) {
		toSerialize["min_required"] = o.MinRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DocumentStatusInfoDetails) UnmarshalJSON(bytes []byte) (err error) {
	varDocumentStatusInfoDetails := _DocumentStatusInfoDetails{}

	if err = json.Unmarshal(bytes, &varDocumentStatusInfoDetails); err == nil {
		*o = DocumentStatusInfoDetails(varDocumentStatusInfoDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "started_at")
		delete(additionalProperties, "current_stage")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "indexing_duration_ms")
		delete(additionalProperties, "tokens_used")
		delete(additionalProperties, "tokens_estimated")
		delete(additionalProperties, "chunks_count")
		delete(additionalProperties, "needed")
		delete(additionalProperties, "available")
		delete(additionalProperties, "estimated")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "min_required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentStatusInfoDetails struct {
	value *DocumentStatusInfoDetails
	isSet bool
}

func (v NullableDocumentStatusInfoDetails) Get() *DocumentStatusInfoDetails {
	return v.value
}

func (v *NullableDocumentStatusInfoDetails) Set(val *DocumentStatusInfoDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentStatusInfoDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentStatusInfoDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentStatusInfoDetails(val *DocumentStatusInfoDetails) *NullableDocumentStatusInfoDetails {
	return &NullableDocumentStatusInfoDetails{value: val, isSet: true}
}

func (v NullableDocumentStatusInfoDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentStatusInfoDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


