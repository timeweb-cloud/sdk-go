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

// checks if the Status type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Status{}

// Status Статус аккаунта
type Status struct {
	// Это логическое значение, которое показывает, заблокирован ли аккаунт.
	IsBlocked bool `json:"is_blocked"`
	// Это логическое значение, которое показывает, заблокирован ли аккаунт навсегда.
	IsPermanentBlocked bool `json:"is_permanent_blocked"`
	// Это логическое значение, которое показывает, требуется ли отправлять счета на почту.
	IsSendBillLetters bool `json:"is_send_bill_letters"`
	CompanyInfo StatusCompanyInfo `json:"company_info"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда последний раз изменялся пароль.
	LastPasswordChangedAt string `json:"last_password_changed_at"`
	// ID аккаунта для яндекс метрики.
	YmClientId NullableString `json:"ym_client_id"`
}

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus(isBlocked bool, isPermanentBlocked bool, isSendBillLetters bool, companyInfo StatusCompanyInfo, lastPasswordChangedAt string, ymClientId NullableString) *Status {
	this := Status{}
	this.IsBlocked = isBlocked
	this.IsPermanentBlocked = isPermanentBlocked
	this.IsSendBillLetters = isSendBillLetters
	this.CompanyInfo = companyInfo
	this.LastPasswordChangedAt = lastPasswordChangedAt
	this.YmClientId = ymClientId
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetIsBlocked returns the IsBlocked field value
func (o *Status) GetIsBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBlocked
}

// GetIsBlockedOk returns a tuple with the IsBlocked field value
// and a boolean to check if the value has been set.
func (o *Status) GetIsBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBlocked, true
}

// SetIsBlocked sets field value
func (o *Status) SetIsBlocked(v bool) {
	o.IsBlocked = v
}

// GetIsPermanentBlocked returns the IsPermanentBlocked field value
func (o *Status) GetIsPermanentBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPermanentBlocked
}

// GetIsPermanentBlockedOk returns a tuple with the IsPermanentBlocked field value
// and a boolean to check if the value has been set.
func (o *Status) GetIsPermanentBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPermanentBlocked, true
}

// SetIsPermanentBlocked sets field value
func (o *Status) SetIsPermanentBlocked(v bool) {
	o.IsPermanentBlocked = v
}

// GetIsSendBillLetters returns the IsSendBillLetters field value
func (o *Status) GetIsSendBillLetters() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSendBillLetters
}

// GetIsSendBillLettersOk returns a tuple with the IsSendBillLetters field value
// and a boolean to check if the value has been set.
func (o *Status) GetIsSendBillLettersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSendBillLetters, true
}

// SetIsSendBillLetters sets field value
func (o *Status) SetIsSendBillLetters(v bool) {
	o.IsSendBillLetters = v
}

// GetCompanyInfo returns the CompanyInfo field value
func (o *Status) GetCompanyInfo() StatusCompanyInfo {
	if o == nil {
		var ret StatusCompanyInfo
		return ret
	}

	return o.CompanyInfo
}

// GetCompanyInfoOk returns a tuple with the CompanyInfo field value
// and a boolean to check if the value has been set.
func (o *Status) GetCompanyInfoOk() (*StatusCompanyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyInfo, true
}

// SetCompanyInfo sets field value
func (o *Status) SetCompanyInfo(v StatusCompanyInfo) {
	o.CompanyInfo = v
}

// GetLastPasswordChangedAt returns the LastPasswordChangedAt field value
func (o *Status) GetLastPasswordChangedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPasswordChangedAt
}

// GetLastPasswordChangedAtOk returns a tuple with the LastPasswordChangedAt field value
// and a boolean to check if the value has been set.
func (o *Status) GetLastPasswordChangedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPasswordChangedAt, true
}

// SetLastPasswordChangedAt sets field value
func (o *Status) SetLastPasswordChangedAt(v string) {
	o.LastPasswordChangedAt = v
}

// GetYmClientId returns the YmClientId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Status) GetYmClientId() string {
	if o == nil || o.YmClientId.Get() == nil {
		var ret string
		return ret
	}

	return *o.YmClientId.Get()
}

// GetYmClientIdOk returns a tuple with the YmClientId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Status) GetYmClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.YmClientId.Get(), o.YmClientId.IsSet()
}

// SetYmClientId sets field value
func (o *Status) SetYmClientId(v string) {
	o.YmClientId.Set(&v)
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Status) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_blocked"] = o.IsBlocked
	toSerialize["is_permanent_blocked"] = o.IsPermanentBlocked
	toSerialize["is_send_bill_letters"] = o.IsSendBillLetters
	toSerialize["company_info"] = o.CompanyInfo
	toSerialize["last_password_changed_at"] = o.LastPasswordChangedAt
	toSerialize["ym_client_id"] = o.YmClientId.Get()
	return toSerialize, nil
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


