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

// checks if the ChatListMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatListMetrics{}

// ChatListMetrics Сводка по чатам агента с учётом фильтра по периоду (`from`/`to`)
type ChatListMetrics struct {
	// Количество чатов с учётом фильтра `response_time_ranges` (без учёта `reaction_type`)
	TotalReactions int32 `json:"total_reactions"`
	// Количество чатов с учётом фильтра `reaction_type` (без учёта `response_time_ranges`)
	TotalTime int32 `json:"total_time"`
	// Чатов хотя бы с одной положительной оценкой
	WithLikes int32 `json:"with_likes"`
	// Чатов хотя бы с одной отрицательной оценкой
	WithDislikes int32 `json:"with_dislikes"`
	// Чатов хотя бы с одной оценкой
	WithAny int32 `json:"with_any"`
	// Количество чатов по среднему времени ответа (в секундах) с учётом фильтра `response_time_ranges`
	ResponseTimeGroupsReactions map[string]int32 `json:"response_time_groups_reactions"`
	// Количество чатов по среднему времени ответа (в секундах) с учётом фильтра `reaction_type`
	ResponseTimeGroupsTime map[string]int32 `json:"response_time_groups_time"`
}

// NewChatListMetrics instantiates a new ChatListMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatListMetrics(totalReactions int32, totalTime int32, withLikes int32, withDislikes int32, withAny int32, responseTimeGroupsReactions map[string]int32, responseTimeGroupsTime map[string]int32) *ChatListMetrics {
	this := ChatListMetrics{}
	this.TotalReactions = totalReactions
	this.TotalTime = totalTime
	this.WithLikes = withLikes
	this.WithDislikes = withDislikes
	this.WithAny = withAny
	this.ResponseTimeGroupsReactions = responseTimeGroupsReactions
	this.ResponseTimeGroupsTime = responseTimeGroupsTime
	return &this
}

// NewChatListMetricsWithDefaults instantiates a new ChatListMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatListMetricsWithDefaults() *ChatListMetrics {
	this := ChatListMetrics{}
	return &this
}

// GetTotalReactions returns the TotalReactions field value
func (o *ChatListMetrics) GetTotalReactions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalReactions
}

// GetTotalReactionsOk returns a tuple with the TotalReactions field value
// and a boolean to check if the value has been set.
func (o *ChatListMetrics) GetTotalReactionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalReactions, true
}

// SetTotalReactions sets field value
func (o *ChatListMetrics) SetTotalReactions(v int32) {
	o.TotalReactions = v
}

// GetTotalTime returns the TotalTime field value
func (o *ChatListMetrics) GetTotalTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value
// and a boolean to check if the value has been set.
func (o *ChatListMetrics) GetTotalTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTime, true
}

// SetTotalTime sets field value
func (o *ChatListMetrics) SetTotalTime(v int32) {
	o.TotalTime = v
}

// GetWithLikes returns the WithLikes field value
func (o *ChatListMetrics) GetWithLikes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WithLikes
}

// GetWithLikesOk returns a tuple with the WithLikes field value
// and a boolean to check if the value has been set.
func (o *ChatListMetrics) GetWithLikesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithLikes, true
}

// SetWithLikes sets field value
func (o *ChatListMetrics) SetWithLikes(v int32) {
	o.WithLikes = v
}

// GetWithDislikes returns the WithDislikes field value
func (o *ChatListMetrics) GetWithDislikes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WithDislikes
}

// GetWithDislikesOk returns a tuple with the WithDislikes field value
// and a boolean to check if the value has been set.
func (o *ChatListMetrics) GetWithDislikesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithDislikes, true
}

// SetWithDislikes sets field value
func (o *ChatListMetrics) SetWithDislikes(v int32) {
	o.WithDislikes = v
}

// GetWithAny returns the WithAny field value
func (o *ChatListMetrics) GetWithAny() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WithAny
}

// GetWithAnyOk returns a tuple with the WithAny field value
// and a boolean to check if the value has been set.
func (o *ChatListMetrics) GetWithAnyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithAny, true
}

// SetWithAny sets field value
func (o *ChatListMetrics) SetWithAny(v int32) {
	o.WithAny = v
}

// GetResponseTimeGroupsReactions returns the ResponseTimeGroupsReactions field value
func (o *ChatListMetrics) GetResponseTimeGroupsReactions() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.ResponseTimeGroupsReactions
}

// GetResponseTimeGroupsReactionsOk returns a tuple with the ResponseTimeGroupsReactions field value
// and a boolean to check if the value has been set.
func (o *ChatListMetrics) GetResponseTimeGroupsReactionsOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTimeGroupsReactions, true
}

// SetResponseTimeGroupsReactions sets field value
func (o *ChatListMetrics) SetResponseTimeGroupsReactions(v map[string]int32) {
	o.ResponseTimeGroupsReactions = v
}

// GetResponseTimeGroupsTime returns the ResponseTimeGroupsTime field value
func (o *ChatListMetrics) GetResponseTimeGroupsTime() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.ResponseTimeGroupsTime
}

// GetResponseTimeGroupsTimeOk returns a tuple with the ResponseTimeGroupsTime field value
// and a boolean to check if the value has been set.
func (o *ChatListMetrics) GetResponseTimeGroupsTimeOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTimeGroupsTime, true
}

// SetResponseTimeGroupsTime sets field value
func (o *ChatListMetrics) SetResponseTimeGroupsTime(v map[string]int32) {
	o.ResponseTimeGroupsTime = v
}

func (o ChatListMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatListMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_reactions"] = o.TotalReactions
	toSerialize["total_time"] = o.TotalTime
	toSerialize["with_likes"] = o.WithLikes
	toSerialize["with_dislikes"] = o.WithDislikes
	toSerialize["with_any"] = o.WithAny
	toSerialize["response_time_groups_reactions"] = o.ResponseTimeGroupsReactions
	toSerialize["response_time_groups_time"] = o.ResponseTimeGroupsTime
	return toSerialize, nil
}

type NullableChatListMetrics struct {
	value *ChatListMetrics
	isSet bool
}

func (v NullableChatListMetrics) Get() *ChatListMetrics {
	return v.value
}

func (v *NullableChatListMetrics) Set(val *ChatListMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableChatListMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableChatListMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatListMetrics(val *ChatListMetrics) *NullableChatListMetrics {
	return &NullableChatListMetrics{value: val, isSet: true}
}

func (v NullableChatListMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatListMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


