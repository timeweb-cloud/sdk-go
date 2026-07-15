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

// checks if the Person type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Person{}

// Person struct for Person
type Person struct {
	// Уникальный идентификатор администратора домена.
	Id int32 `json:"id"`
	// Тип администратора: `person` — физическое лицо, `org` — организация, `ip` — индивидуальный предприниматель.
	Type string `json:"type"`
	// Имя администратора (ФИО или название организации).
	Name string `json:"name"`
	// Имя администратора в латинской транслитерации.
	NameEng string `json:"name_eng"`
	// Адрес электронной почты администратора.
	Email string `json:"email"`
	// Контактный телефон администратора.
	Phone string `json:"phone"`
	// Почтовый индекс.
	Postcode string `json:"postcode"`
	// Почтовый адрес.
	MailingAddress string `json:"mailing_address"`
	// Код страны администратора.
	CountryCode NullableString `json:"country_code"`
	// Это логическое значение, которое показывает, является ли администратор резидентом РФ.
	IsResident bool `json:"is_resident"`
	// Это логическое значение, которое показывает, заполнены ли данные администратора не полностью.
	IsBlank bool `json:"is_blank"`
	// Это логическое значение, которое показывает, закрыт ли администратор.
	IsClosed bool `json:"is_closed"`
	// Дата рождения. Только для типов `person` и `ip`.
	Birthdate *string `json:"birthdate,omitempty"`
	// Дата выдачи паспорта. Только для типов `person` и `ip`.
	PassportDate *string `json:"passport_date,omitempty"`
	// Номер паспорта. Только для типов `person` и `ip`.
	PassportNumber *string `json:"passport_number,omitempty"`
	// Кем выдан паспорт. Только для типов `person` и `ip`.
	PassportPlace *string `json:"passport_place,omitempty"`
	// Серия паспорта. Только для типов `person` и `ip`.
	PassportSeries *string `json:"passport_series,omitempty"`
	// ИНН. Только для типов `org` и `ip`.
	Inn *string `json:"inn,omitempty"`
	// КПП организации. Только для типа `org`.
	Kpp *string `json:"kpp,omitempty"`
	// Юридический адрес организации. Только для типа `org`.
	LegalAddress *string `json:"legal_address,omitempty"`
	// Контактное лицо организации. Только для типа `org`.
	ContactName *string `json:"contact_name,omitempty"`
}

// NewPerson instantiates a new Person object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerson(id int32, type_ string, name string, nameEng string, email string, phone string, postcode string, mailingAddress string, countryCode NullableString, isResident bool, isBlank bool, isClosed bool) *Person {
	this := Person{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.NameEng = nameEng
	this.Email = email
	this.Phone = phone
	this.Postcode = postcode
	this.MailingAddress = mailingAddress
	this.CountryCode = countryCode
	this.IsResident = isResident
	this.IsBlank = isBlank
	this.IsClosed = isClosed
	return &this
}

// NewPersonWithDefaults instantiates a new Person object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonWithDefaults() *Person {
	this := Person{}
	return &this
}

// GetId returns the Id field value
func (o *Person) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Person) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Person) SetId(v int32) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Person) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Person) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Person) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *Person) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Person) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Person) SetName(v string) {
	o.Name = v
}

// GetNameEng returns the NameEng field value
func (o *Person) GetNameEng() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameEng
}

// GetNameEngOk returns a tuple with the NameEng field value
// and a boolean to check if the value has been set.
func (o *Person) GetNameEngOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEng, true
}

// SetNameEng sets field value
func (o *Person) SetNameEng(v string) {
	o.NameEng = v
}

// GetEmail returns the Email field value
func (o *Person) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Person) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Person) SetEmail(v string) {
	o.Email = v
}

// GetPhone returns the Phone field value
func (o *Person) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *Person) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *Person) SetPhone(v string) {
	o.Phone = v
}

// GetPostcode returns the Postcode field value
func (o *Person) GetPostcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Postcode
}

// GetPostcodeOk returns a tuple with the Postcode field value
// and a boolean to check if the value has been set.
func (o *Person) GetPostcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Postcode, true
}

// SetPostcode sets field value
func (o *Person) SetPostcode(v string) {
	o.Postcode = v
}

// GetMailingAddress returns the MailingAddress field value
func (o *Person) GetMailingAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MailingAddress
}

// GetMailingAddressOk returns a tuple with the MailingAddress field value
// and a boolean to check if the value has been set.
func (o *Person) GetMailingAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MailingAddress, true
}

// SetMailingAddress sets field value
func (o *Person) SetMailingAddress(v string) {
	o.MailingAddress = v
}

// GetCountryCode returns the CountryCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Person) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Person) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// SetCountryCode sets field value
func (o *Person) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// GetIsResident returns the IsResident field value
func (o *Person) GetIsResident() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsResident
}

// GetIsResidentOk returns a tuple with the IsResident field value
// and a boolean to check if the value has been set.
func (o *Person) GetIsResidentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsResident, true
}

// SetIsResident sets field value
func (o *Person) SetIsResident(v bool) {
	o.IsResident = v
}

// GetIsBlank returns the IsBlank field value
func (o *Person) GetIsBlank() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBlank
}

// GetIsBlankOk returns a tuple with the IsBlank field value
// and a boolean to check if the value has been set.
func (o *Person) GetIsBlankOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBlank, true
}

// SetIsBlank sets field value
func (o *Person) SetIsBlank(v bool) {
	o.IsBlank = v
}

// GetIsClosed returns the IsClosed field value
func (o *Person) GetIsClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value
// and a boolean to check if the value has been set.
func (o *Person) GetIsClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClosed, true
}

// SetIsClosed sets field value
func (o *Person) SetIsClosed(v bool) {
	o.IsClosed = v
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *Person) GetBirthdate() string {
	if o == nil || IsNil(o.Birthdate) {
		var ret string
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetBirthdateOk() (*string, bool) {
	if o == nil || IsNil(o.Birthdate) {
		return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *Person) HasBirthdate() bool {
	if o != nil && !IsNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given string and assigns it to the Birthdate field.
func (o *Person) SetBirthdate(v string) {
	o.Birthdate = &v
}

// GetPassportDate returns the PassportDate field value if set, zero value otherwise.
func (o *Person) GetPassportDate() string {
	if o == nil || IsNil(o.PassportDate) {
		var ret string
		return ret
	}
	return *o.PassportDate
}

// GetPassportDateOk returns a tuple with the PassportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetPassportDateOk() (*string, bool) {
	if o == nil || IsNil(o.PassportDate) {
		return nil, false
	}
	return o.PassportDate, true
}

// HasPassportDate returns a boolean if a field has been set.
func (o *Person) HasPassportDate() bool {
	if o != nil && !IsNil(o.PassportDate) {
		return true
	}

	return false
}

// SetPassportDate gets a reference to the given string and assigns it to the PassportDate field.
func (o *Person) SetPassportDate(v string) {
	o.PassportDate = &v
}

// GetPassportNumber returns the PassportNumber field value if set, zero value otherwise.
func (o *Person) GetPassportNumber() string {
	if o == nil || IsNil(o.PassportNumber) {
		var ret string
		return ret
	}
	return *o.PassportNumber
}

// GetPassportNumberOk returns a tuple with the PassportNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetPassportNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PassportNumber) {
		return nil, false
	}
	return o.PassportNumber, true
}

// HasPassportNumber returns a boolean if a field has been set.
func (o *Person) HasPassportNumber() bool {
	if o != nil && !IsNil(o.PassportNumber) {
		return true
	}

	return false
}

// SetPassportNumber gets a reference to the given string and assigns it to the PassportNumber field.
func (o *Person) SetPassportNumber(v string) {
	o.PassportNumber = &v
}

// GetPassportPlace returns the PassportPlace field value if set, zero value otherwise.
func (o *Person) GetPassportPlace() string {
	if o == nil || IsNil(o.PassportPlace) {
		var ret string
		return ret
	}
	return *o.PassportPlace
}

// GetPassportPlaceOk returns a tuple with the PassportPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetPassportPlaceOk() (*string, bool) {
	if o == nil || IsNil(o.PassportPlace) {
		return nil, false
	}
	return o.PassportPlace, true
}

// HasPassportPlace returns a boolean if a field has been set.
func (o *Person) HasPassportPlace() bool {
	if o != nil && !IsNil(o.PassportPlace) {
		return true
	}

	return false
}

// SetPassportPlace gets a reference to the given string and assigns it to the PassportPlace field.
func (o *Person) SetPassportPlace(v string) {
	o.PassportPlace = &v
}

// GetPassportSeries returns the PassportSeries field value if set, zero value otherwise.
func (o *Person) GetPassportSeries() string {
	if o == nil || IsNil(o.PassportSeries) {
		var ret string
		return ret
	}
	return *o.PassportSeries
}

// GetPassportSeriesOk returns a tuple with the PassportSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetPassportSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.PassportSeries) {
		return nil, false
	}
	return o.PassportSeries, true
}

// HasPassportSeries returns a boolean if a field has been set.
func (o *Person) HasPassportSeries() bool {
	if o != nil && !IsNil(o.PassportSeries) {
		return true
	}

	return false
}

// SetPassportSeries gets a reference to the given string and assigns it to the PassportSeries field.
func (o *Person) SetPassportSeries(v string) {
	o.PassportSeries = &v
}

// GetInn returns the Inn field value if set, zero value otherwise.
func (o *Person) GetInn() string {
	if o == nil || IsNil(o.Inn) {
		var ret string
		return ret
	}
	return *o.Inn
}

// GetInnOk returns a tuple with the Inn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetInnOk() (*string, bool) {
	if o == nil || IsNil(o.Inn) {
		return nil, false
	}
	return o.Inn, true
}

// HasInn returns a boolean if a field has been set.
func (o *Person) HasInn() bool {
	if o != nil && !IsNil(o.Inn) {
		return true
	}

	return false
}

// SetInn gets a reference to the given string and assigns it to the Inn field.
func (o *Person) SetInn(v string) {
	o.Inn = &v
}

// GetKpp returns the Kpp field value if set, zero value otherwise.
func (o *Person) GetKpp() string {
	if o == nil || IsNil(o.Kpp) {
		var ret string
		return ret
	}
	return *o.Kpp
}

// GetKppOk returns a tuple with the Kpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetKppOk() (*string, bool) {
	if o == nil || IsNil(o.Kpp) {
		return nil, false
	}
	return o.Kpp, true
}

// HasKpp returns a boolean if a field has been set.
func (o *Person) HasKpp() bool {
	if o != nil && !IsNil(o.Kpp) {
		return true
	}

	return false
}

// SetKpp gets a reference to the given string and assigns it to the Kpp field.
func (o *Person) SetKpp(v string) {
	o.Kpp = &v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *Person) GetLegalAddress() string {
	if o == nil || IsNil(o.LegalAddress) {
		var ret string
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetLegalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.LegalAddress) {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *Person) HasLegalAddress() bool {
	if o != nil && !IsNil(o.LegalAddress) {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given string and assigns it to the LegalAddress field.
func (o *Person) SetLegalAddress(v string) {
	o.LegalAddress = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *Person) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *Person) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *Person) SetContactName(v string) {
	o.ContactName = &v
}

func (o Person) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Person) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["name_eng"] = o.NameEng
	toSerialize["email"] = o.Email
	toSerialize["phone"] = o.Phone
	toSerialize["postcode"] = o.Postcode
	toSerialize["mailing_address"] = o.MailingAddress
	toSerialize["country_code"] = o.CountryCode.Get()
	toSerialize["is_resident"] = o.IsResident
	toSerialize["is_blank"] = o.IsBlank
	toSerialize["is_closed"] = o.IsClosed
	if !IsNil(o.Birthdate) {
		toSerialize["birthdate"] = o.Birthdate
	}
	if !IsNil(o.PassportDate) {
		toSerialize["passport_date"] = o.PassportDate
	}
	if !IsNil(o.PassportNumber) {
		toSerialize["passport_number"] = o.PassportNumber
	}
	if !IsNil(o.PassportPlace) {
		toSerialize["passport_place"] = o.PassportPlace
	}
	if !IsNil(o.PassportSeries) {
		toSerialize["passport_series"] = o.PassportSeries
	}
	if !IsNil(o.Inn) {
		toSerialize["inn"] = o.Inn
	}
	if !IsNil(o.Kpp) {
		toSerialize["kpp"] = o.Kpp
	}
	if !IsNil(o.LegalAddress) {
		toSerialize["legal_address"] = o.LegalAddress
	}
	if !IsNil(o.ContactName) {
		toSerialize["contact_name"] = o.ContactName
	}
	return toSerialize, nil
}

type NullablePerson struct {
	value *Person
	isSet bool
}

func (v NullablePerson) Get() *Person {
	return v.value
}

func (v *NullablePerson) Set(val *Person) {
	v.value = val
	v.isSet = true
}

func (v NullablePerson) IsSet() bool {
	return v.isSet
}

func (v *NullablePerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerson(val *Person) *NullablePerson {
	return &NullablePerson{value: val, isSet: true}
}

func (v NullablePerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


