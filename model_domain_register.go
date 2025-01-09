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

// checks if the DomainRegister type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainRegister{}

// DomainRegister Заявка на регистрацию домена
type DomainRegister struct {
	// Тип создаваемой заявки.
	Action string `json:"action"`
	// Полное имя домена.
	Fqdn string `json:"fqdn"`
	// Это логическое значение, которое показывает, включено ли автопродление домена.
	IsAutoprolongEnabled *bool `json:"is_autoprolong_enabled,omitempty"`
	// Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Опция недоступна для доменов в зонах .ru и .рф.
	IsWhoisPrivacyEnabled *bool `json:"is_whois_privacy_enabled,omitempty"`
	Period *DomainPaymentPeriod `json:"period,omitempty"`
	// ID администратора, на которого регистрируется домен.
	PersonId float32 `json:"person_id"`
}

// NewDomainRegister instantiates a new DomainRegister object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRegister(action string, fqdn string, personId float32) *DomainRegister {
	this := DomainRegister{}
	this.Action = action
	this.Fqdn = fqdn
	this.PersonId = personId
	return &this
}

// NewDomainRegisterWithDefaults instantiates a new DomainRegister object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRegisterWithDefaults() *DomainRegister {
	this := DomainRegister{}
	return &this
}

// GetAction returns the Action field value
func (o *DomainRegister) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DomainRegister) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *DomainRegister) SetAction(v string) {
	o.Action = v
}

// GetFqdn returns the Fqdn field value
func (o *DomainRegister) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *DomainRegister) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *DomainRegister) SetFqdn(v string) {
	o.Fqdn = v
}

// GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field value if set, zero value otherwise.
func (o *DomainRegister) GetIsAutoprolongEnabled() bool {
	if o == nil || IsNil(o.IsAutoprolongEnabled) {
		var ret bool
		return ret
	}
	return *o.IsAutoprolongEnabled
}

// GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRegister) GetIsAutoprolongEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoprolongEnabled) {
		return nil, false
	}
	return o.IsAutoprolongEnabled, true
}

// HasIsAutoprolongEnabled returns a boolean if a field has been set.
func (o *DomainRegister) HasIsAutoprolongEnabled() bool {
	if o != nil && !IsNil(o.IsAutoprolongEnabled) {
		return true
	}

	return false
}

// SetIsAutoprolongEnabled gets a reference to the given bool and assigns it to the IsAutoprolongEnabled field.
func (o *DomainRegister) SetIsAutoprolongEnabled(v bool) {
	o.IsAutoprolongEnabled = &v
}

// GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field value if set, zero value otherwise.
func (o *DomainRegister) GetIsWhoisPrivacyEnabled() bool {
	if o == nil || IsNil(o.IsWhoisPrivacyEnabled) {
		var ret bool
		return ret
	}
	return *o.IsWhoisPrivacyEnabled
}

// GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRegister) GetIsWhoisPrivacyEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWhoisPrivacyEnabled) {
		return nil, false
	}
	return o.IsWhoisPrivacyEnabled, true
}

// HasIsWhoisPrivacyEnabled returns a boolean if a field has been set.
func (o *DomainRegister) HasIsWhoisPrivacyEnabled() bool {
	if o != nil && !IsNil(o.IsWhoisPrivacyEnabled) {
		return true
	}

	return false
}

// SetIsWhoisPrivacyEnabled gets a reference to the given bool and assigns it to the IsWhoisPrivacyEnabled field.
func (o *DomainRegister) SetIsWhoisPrivacyEnabled(v bool) {
	o.IsWhoisPrivacyEnabled = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DomainRegister) GetPeriod() DomainPaymentPeriod {
	if o == nil || IsNil(o.Period) {
		var ret DomainPaymentPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRegister) GetPeriodOk() (*DomainPaymentPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DomainRegister) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given DomainPaymentPeriod and assigns it to the Period field.
func (o *DomainRegister) SetPeriod(v DomainPaymentPeriod) {
	o.Period = &v
}

// GetPersonId returns the PersonId field value
func (o *DomainRegister) GetPersonId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value
// and a boolean to check if the value has been set.
func (o *DomainRegister) GetPersonIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PersonId, true
}

// SetPersonId sets field value
func (o *DomainRegister) SetPersonId(v float32) {
	o.PersonId = v
}

func (o DomainRegister) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainRegister) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["fqdn"] = o.Fqdn
	if !IsNil(o.IsAutoprolongEnabled) {
		toSerialize["is_autoprolong_enabled"] = o.IsAutoprolongEnabled
	}
	if !IsNil(o.IsWhoisPrivacyEnabled) {
		toSerialize["is_whois_privacy_enabled"] = o.IsWhoisPrivacyEnabled
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	toSerialize["person_id"] = o.PersonId
	return toSerialize, nil
}

type NullableDomainRegister struct {
	value *DomainRegister
	isSet bool
}

func (v NullableDomainRegister) Get() *DomainRegister {
	return v.value
}

func (v *NullableDomainRegister) Set(val *DomainRegister) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRegister) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRegister) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRegister(val *DomainRegister) *NullableDomainRegister {
	return &NullableDomainRegister{value: val, isSet: true}
}

func (v NullableDomainRegister) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRegister) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


