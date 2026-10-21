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

// checks if the Model type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model{}

// Model AI модель
type Model struct {
	// ID модели
	Id int32 `json:"id"`
	// ID провайдера модели
	ProviderId int32 `json:"provider_id"`
	// Семейство модели
	Name string `json:"name"`
	// Отображаемое название модели
	PublicName string `json:"public_name"`
	// Идентификатор модели с префиксом провайдера — передаётся в поле `model` запросов к AI Gateway
	ModelName string `json:"model_name"`
	// Тип модели: - `llm` — языковая модель - `hosted-llm` — языковая модель, размещённая в Timeweb Cloud - `embedding` — модель эмбеддингов - `image` — генерация изображений - `audio` — работа с аудио - `video` — генерация видео - `moderation` — модерация контента - `system-one` — служебная модель с собственным API
	Type string `json:"type"`
	// Версия модели
	Version string `json:"version"`
	// Модель устарела и недоступна для выбора в новых агентах
	IsDeprecated bool `json:"is_deprecated"`
	// Поддержка модели остановлена: запросы к ней не выполняются
	IsStopped bool `json:"is_stopped"`
	// Дата отключения модели у провайдера. `null`, если отключение не запланировано
	DeprecationDate NullableTime `json:"deprecation_date"`
	// Модель поддерживает режим рассуждений
	IsReasoning bool `json:"is_reasoning"`
	// Параметры генерации, которые можно настроить у агента с этой моделью. Ключ — имя параметра (`temperature`, `top_p`, `max_tokens`, `reasoning_effort` и т.д.). `null`, если параметры не заданы
	ParamsInfo map[string]ModelParamInfo `json:"params_info"`
}

// NewModel instantiates a new Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel(id int32, providerId int32, name string, publicName string, modelName string, type_ string, version string, isDeprecated bool, isStopped bool, deprecationDate NullableTime, isReasoning bool, paramsInfo map[string]ModelParamInfo) *Model {
	this := Model{}
	this.Id = id
	this.ProviderId = providerId
	this.Name = name
	this.PublicName = publicName
	this.ModelName = modelName
	this.Type = type_
	this.Version = version
	this.IsDeprecated = isDeprecated
	this.IsStopped = isStopped
	this.DeprecationDate = deprecationDate
	this.IsReasoning = isReasoning
	this.ParamsInfo = paramsInfo
	return &this
}

// NewModelWithDefaults instantiates a new Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWithDefaults() *Model {
	this := Model{}
	return &this
}

// GetId returns the Id field value
func (o *Model) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Model) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Model) SetId(v int32) {
	o.Id = v
}

// GetProviderId returns the ProviderId field value
func (o *Model) GetProviderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *Model) GetProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *Model) SetProviderId(v int32) {
	o.ProviderId = v
}

// GetName returns the Name field value
func (o *Model) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Model) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Model) SetName(v string) {
	o.Name = v
}

// GetPublicName returns the PublicName field value
func (o *Model) GetPublicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicName
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
func (o *Model) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicName, true
}

// SetPublicName sets field value
func (o *Model) SetPublicName(v string) {
	o.PublicName = v
}

// GetModelName returns the ModelName field value
func (o *Model) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *Model) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *Model) SetModelName(v string) {
	o.ModelName = v
}

// GetType returns the Type field value
func (o *Model) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Model) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Model) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *Model) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Model) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Model) SetVersion(v string) {
	o.Version = v
}

// GetIsDeprecated returns the IsDeprecated field value
func (o *Model) GetIsDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value
// and a boolean to check if the value has been set.
func (o *Model) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecated, true
}

// SetIsDeprecated sets field value
func (o *Model) SetIsDeprecated(v bool) {
	o.IsDeprecated = v
}

// GetIsStopped returns the IsStopped field value
func (o *Model) GetIsStopped() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsStopped
}

// GetIsStoppedOk returns a tuple with the IsStopped field value
// and a boolean to check if the value has been set.
func (o *Model) GetIsStoppedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsStopped, true
}

// SetIsStopped sets field value
func (o *Model) SetIsStopped(v bool) {
	o.IsStopped = v
}

// GetDeprecationDate returns the DeprecationDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Model) GetDeprecationDate() time.Time {
	if o == nil || o.DeprecationDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DeprecationDate.Get()
}

// GetDeprecationDateOk returns a tuple with the DeprecationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetDeprecationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeprecationDate.Get(), o.DeprecationDate.IsSet()
}

// SetDeprecationDate sets field value
func (o *Model) SetDeprecationDate(v time.Time) {
	o.DeprecationDate.Set(&v)
}

// GetIsReasoning returns the IsReasoning field value
func (o *Model) GetIsReasoning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReasoning
}

// GetIsReasoningOk returns a tuple with the IsReasoning field value
// and a boolean to check if the value has been set.
func (o *Model) GetIsReasoningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReasoning, true
}

// SetIsReasoning sets field value
func (o *Model) SetIsReasoning(v bool) {
	o.IsReasoning = v
}

// GetParamsInfo returns the ParamsInfo field value
// If the value is explicit nil, the zero value for map[string]ModelParamInfo will be returned
func (o *Model) GetParamsInfo() map[string]ModelParamInfo {
	if o == nil {
		var ret map[string]ModelParamInfo
		return ret
	}

	return o.ParamsInfo
}

// GetParamsInfoOk returns a tuple with the ParamsInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Model) GetParamsInfoOk() (*map[string]ModelParamInfo, bool) {
	if o == nil || IsNil(o.ParamsInfo) {
		return nil, false
	}
	return &o.ParamsInfo, true
}

// SetParamsInfo sets field value
func (o *Model) SetParamsInfo(v map[string]ModelParamInfo) {
	o.ParamsInfo = v
}

func (o Model) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["provider_id"] = o.ProviderId
	toSerialize["name"] = o.Name
	toSerialize["public_name"] = o.PublicName
	toSerialize["model_name"] = o.ModelName
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	toSerialize["is_deprecated"] = o.IsDeprecated
	toSerialize["is_stopped"] = o.IsStopped
	toSerialize["deprecation_date"] = o.DeprecationDate.Get()
	toSerialize["is_reasoning"] = o.IsReasoning
	if o.ParamsInfo != nil {
		toSerialize["params_info"] = o.ParamsInfo
	}
	return toSerialize, nil
}

type NullableModel struct {
	value *Model
	isSet bool
}

func (v NullableModel) Get() *Model {
	return v.value
}

func (v *NullableModel) Set(val *Model) {
	v.value = val
	v.isSet = true
}

func (v NullableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel(val *Model) *NullableModel {
	return &NullableModel{value: val, isSet: true}
}

func (v NullableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


