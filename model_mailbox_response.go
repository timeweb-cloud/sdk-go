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

// checks if the MailboxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailboxResponse{}

// MailboxResponse struct for MailboxResponse
type MailboxResponse struct {
	// IDN имя домена
	IdnName *string `json:"idn_name,omitempty"`
	// Сообщение автоответчика
	AutoreplyMessage *string `json:"autoreply_message,omitempty"`
	// Статус автоответчика
	AutoreplyStatus *bool `json:"autoreply_status,omitempty"`
	// Тема автоответчика
	AutoreplySubject *string `json:"autoreply_subject,omitempty"`
	// Комментарий
	Comment *string `json:"comment,omitempty"`
	// Действие фильтра спама
	FilterAction *string `json:"filter_action,omitempty"`
	// Статус фильтра спама
	FilterStatus *bool `json:"filter_status,omitempty"`
	// Список адресов для пересылки
	ForwardList []string `json:"forward_list,omitempty"`
	// Статус пересылки
	ForwardStatus *bool `json:"forward_status,omitempty"`
	// Контроль исходящей почты
	OutgoingControl *bool `json:"outgoing_control,omitempty"`
	// Email для исходящих писем
	OutgoingEmail *string `json:"outgoing_email,omitempty"`
	// Пароль (обычно пустая строка в ответе)
	Password *string `json:"password,omitempty"`
	// Белый список адресов
	WhiteList []string `json:"white_list,omitempty"`
	// Доступ к веб-почте
	Webmail *bool `json:"webmail,omitempty"`
	// Использование Dovecot
	Dovecot *bool `json:"dovecot,omitempty"`
	// Полное доменное имя
	Fqdn *string `json:"fqdn,omitempty"`
	// Оставлять копии писем при пересылке
	LeaveMessages *bool `json:"leave_messages,omitempty"`
	// Имя почтового ящика
	Mailbox *string `json:"mailbox,omitempty"`
	// ФИО владельца
	OwnerFullName *string `json:"owner_full_name,omitempty"`
}

// NewMailboxResponse instantiates a new MailboxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailboxResponse() *MailboxResponse {
	this := MailboxResponse{}
	return &this
}

// NewMailboxResponseWithDefaults instantiates a new MailboxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailboxResponseWithDefaults() *MailboxResponse {
	this := MailboxResponse{}
	return &this
}

// GetIdnName returns the IdnName field value if set, zero value otherwise.
func (o *MailboxResponse) GetIdnName() string {
	if o == nil || IsNil(o.IdnName) {
		var ret string
		return ret
	}
	return *o.IdnName
}

// GetIdnNameOk returns a tuple with the IdnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetIdnNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdnName) {
		return nil, false
	}
	return o.IdnName, true
}

// HasIdnName returns a boolean if a field has been set.
func (o *MailboxResponse) HasIdnName() bool {
	if o != nil && !IsNil(o.IdnName) {
		return true
	}

	return false
}

// SetIdnName gets a reference to the given string and assigns it to the IdnName field.
func (o *MailboxResponse) SetIdnName(v string) {
	o.IdnName = &v
}

// GetAutoreplyMessage returns the AutoreplyMessage field value if set, zero value otherwise.
func (o *MailboxResponse) GetAutoreplyMessage() string {
	if o == nil || IsNil(o.AutoreplyMessage) {
		var ret string
		return ret
	}
	return *o.AutoreplyMessage
}

// GetAutoreplyMessageOk returns a tuple with the AutoreplyMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetAutoreplyMessageOk() (*string, bool) {
	if o == nil || IsNil(o.AutoreplyMessage) {
		return nil, false
	}
	return o.AutoreplyMessage, true
}

// HasAutoreplyMessage returns a boolean if a field has been set.
func (o *MailboxResponse) HasAutoreplyMessage() bool {
	if o != nil && !IsNil(o.AutoreplyMessage) {
		return true
	}

	return false
}

// SetAutoreplyMessage gets a reference to the given string and assigns it to the AutoreplyMessage field.
func (o *MailboxResponse) SetAutoreplyMessage(v string) {
	o.AutoreplyMessage = &v
}

// GetAutoreplyStatus returns the AutoreplyStatus field value if set, zero value otherwise.
func (o *MailboxResponse) GetAutoreplyStatus() bool {
	if o == nil || IsNil(o.AutoreplyStatus) {
		var ret bool
		return ret
	}
	return *o.AutoreplyStatus
}

// GetAutoreplyStatusOk returns a tuple with the AutoreplyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetAutoreplyStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoreplyStatus) {
		return nil, false
	}
	return o.AutoreplyStatus, true
}

// HasAutoreplyStatus returns a boolean if a field has been set.
func (o *MailboxResponse) HasAutoreplyStatus() bool {
	if o != nil && !IsNil(o.AutoreplyStatus) {
		return true
	}

	return false
}

// SetAutoreplyStatus gets a reference to the given bool and assigns it to the AutoreplyStatus field.
func (o *MailboxResponse) SetAutoreplyStatus(v bool) {
	o.AutoreplyStatus = &v
}

// GetAutoreplySubject returns the AutoreplySubject field value if set, zero value otherwise.
func (o *MailboxResponse) GetAutoreplySubject() string {
	if o == nil || IsNil(o.AutoreplySubject) {
		var ret string
		return ret
	}
	return *o.AutoreplySubject
}

// GetAutoreplySubjectOk returns a tuple with the AutoreplySubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetAutoreplySubjectOk() (*string, bool) {
	if o == nil || IsNil(o.AutoreplySubject) {
		return nil, false
	}
	return o.AutoreplySubject, true
}

// HasAutoreplySubject returns a boolean if a field has been set.
func (o *MailboxResponse) HasAutoreplySubject() bool {
	if o != nil && !IsNil(o.AutoreplySubject) {
		return true
	}

	return false
}

// SetAutoreplySubject gets a reference to the given string and assigns it to the AutoreplySubject field.
func (o *MailboxResponse) SetAutoreplySubject(v string) {
	o.AutoreplySubject = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MailboxResponse) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MailboxResponse) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MailboxResponse) SetComment(v string) {
	o.Comment = &v
}

// GetFilterAction returns the FilterAction field value if set, zero value otherwise.
func (o *MailboxResponse) GetFilterAction() string {
	if o == nil || IsNil(o.FilterAction) {
		var ret string
		return ret
	}
	return *o.FilterAction
}

// GetFilterActionOk returns a tuple with the FilterAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetFilterActionOk() (*string, bool) {
	if o == nil || IsNil(o.FilterAction) {
		return nil, false
	}
	return o.FilterAction, true
}

// HasFilterAction returns a boolean if a field has been set.
func (o *MailboxResponse) HasFilterAction() bool {
	if o != nil && !IsNil(o.FilterAction) {
		return true
	}

	return false
}

// SetFilterAction gets a reference to the given string and assigns it to the FilterAction field.
func (o *MailboxResponse) SetFilterAction(v string) {
	o.FilterAction = &v
}

// GetFilterStatus returns the FilterStatus field value if set, zero value otherwise.
func (o *MailboxResponse) GetFilterStatus() bool {
	if o == nil || IsNil(o.FilterStatus) {
		var ret bool
		return ret
	}
	return *o.FilterStatus
}

// GetFilterStatusOk returns a tuple with the FilterStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetFilterStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.FilterStatus) {
		return nil, false
	}
	return o.FilterStatus, true
}

// HasFilterStatus returns a boolean if a field has been set.
func (o *MailboxResponse) HasFilterStatus() bool {
	if o != nil && !IsNil(o.FilterStatus) {
		return true
	}

	return false
}

// SetFilterStatus gets a reference to the given bool and assigns it to the FilterStatus field.
func (o *MailboxResponse) SetFilterStatus(v bool) {
	o.FilterStatus = &v
}

// GetForwardList returns the ForwardList field value if set, zero value otherwise.
func (o *MailboxResponse) GetForwardList() []string {
	if o == nil || IsNil(o.ForwardList) {
		var ret []string
		return ret
	}
	return o.ForwardList
}

// GetForwardListOk returns a tuple with the ForwardList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetForwardListOk() ([]string, bool) {
	if o == nil || IsNil(o.ForwardList) {
		return nil, false
	}
	return o.ForwardList, true
}

// HasForwardList returns a boolean if a field has been set.
func (o *MailboxResponse) HasForwardList() bool {
	if o != nil && !IsNil(o.ForwardList) {
		return true
	}

	return false
}

// SetForwardList gets a reference to the given []string and assigns it to the ForwardList field.
func (o *MailboxResponse) SetForwardList(v []string) {
	o.ForwardList = v
}

// GetForwardStatus returns the ForwardStatus field value if set, zero value otherwise.
func (o *MailboxResponse) GetForwardStatus() bool {
	if o == nil || IsNil(o.ForwardStatus) {
		var ret bool
		return ret
	}
	return *o.ForwardStatus
}

// GetForwardStatusOk returns a tuple with the ForwardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetForwardStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.ForwardStatus) {
		return nil, false
	}
	return o.ForwardStatus, true
}

// HasForwardStatus returns a boolean if a field has been set.
func (o *MailboxResponse) HasForwardStatus() bool {
	if o != nil && !IsNil(o.ForwardStatus) {
		return true
	}

	return false
}

// SetForwardStatus gets a reference to the given bool and assigns it to the ForwardStatus field.
func (o *MailboxResponse) SetForwardStatus(v bool) {
	o.ForwardStatus = &v
}

// GetOutgoingControl returns the OutgoingControl field value if set, zero value otherwise.
func (o *MailboxResponse) GetOutgoingControl() bool {
	if o == nil || IsNil(o.OutgoingControl) {
		var ret bool
		return ret
	}
	return *o.OutgoingControl
}

// GetOutgoingControlOk returns a tuple with the OutgoingControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetOutgoingControlOk() (*bool, bool) {
	if o == nil || IsNil(o.OutgoingControl) {
		return nil, false
	}
	return o.OutgoingControl, true
}

// HasOutgoingControl returns a boolean if a field has been set.
func (o *MailboxResponse) HasOutgoingControl() bool {
	if o != nil && !IsNil(o.OutgoingControl) {
		return true
	}

	return false
}

// SetOutgoingControl gets a reference to the given bool and assigns it to the OutgoingControl field.
func (o *MailboxResponse) SetOutgoingControl(v bool) {
	o.OutgoingControl = &v
}

// GetOutgoingEmail returns the OutgoingEmail field value if set, zero value otherwise.
func (o *MailboxResponse) GetOutgoingEmail() string {
	if o == nil || IsNil(o.OutgoingEmail) {
		var ret string
		return ret
	}
	return *o.OutgoingEmail
}

// GetOutgoingEmailOk returns a tuple with the OutgoingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetOutgoingEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OutgoingEmail) {
		return nil, false
	}
	return o.OutgoingEmail, true
}

// HasOutgoingEmail returns a boolean if a field has been set.
func (o *MailboxResponse) HasOutgoingEmail() bool {
	if o != nil && !IsNil(o.OutgoingEmail) {
		return true
	}

	return false
}

// SetOutgoingEmail gets a reference to the given string and assigns it to the OutgoingEmail field.
func (o *MailboxResponse) SetOutgoingEmail(v string) {
	o.OutgoingEmail = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MailboxResponse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MailboxResponse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MailboxResponse) SetPassword(v string) {
	o.Password = &v
}

// GetWhiteList returns the WhiteList field value if set, zero value otherwise.
func (o *MailboxResponse) GetWhiteList() []string {
	if o == nil || IsNil(o.WhiteList) {
		var ret []string
		return ret
	}
	return o.WhiteList
}

// GetWhiteListOk returns a tuple with the WhiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetWhiteListOk() ([]string, bool) {
	if o == nil || IsNil(o.WhiteList) {
		return nil, false
	}
	return o.WhiteList, true
}

// HasWhiteList returns a boolean if a field has been set.
func (o *MailboxResponse) HasWhiteList() bool {
	if o != nil && !IsNil(o.WhiteList) {
		return true
	}

	return false
}

// SetWhiteList gets a reference to the given []string and assigns it to the WhiteList field.
func (o *MailboxResponse) SetWhiteList(v []string) {
	o.WhiteList = v
}

// GetWebmail returns the Webmail field value if set, zero value otherwise.
func (o *MailboxResponse) GetWebmail() bool {
	if o == nil || IsNil(o.Webmail) {
		var ret bool
		return ret
	}
	return *o.Webmail
}

// GetWebmailOk returns a tuple with the Webmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetWebmailOk() (*bool, bool) {
	if o == nil || IsNil(o.Webmail) {
		return nil, false
	}
	return o.Webmail, true
}

// HasWebmail returns a boolean if a field has been set.
func (o *MailboxResponse) HasWebmail() bool {
	if o != nil && !IsNil(o.Webmail) {
		return true
	}

	return false
}

// SetWebmail gets a reference to the given bool and assigns it to the Webmail field.
func (o *MailboxResponse) SetWebmail(v bool) {
	o.Webmail = &v
}

// GetDovecot returns the Dovecot field value if set, zero value otherwise.
func (o *MailboxResponse) GetDovecot() bool {
	if o == nil || IsNil(o.Dovecot) {
		var ret bool
		return ret
	}
	return *o.Dovecot
}

// GetDovecotOk returns a tuple with the Dovecot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetDovecotOk() (*bool, bool) {
	if o == nil || IsNil(o.Dovecot) {
		return nil, false
	}
	return o.Dovecot, true
}

// HasDovecot returns a boolean if a field has been set.
func (o *MailboxResponse) HasDovecot() bool {
	if o != nil && !IsNil(o.Dovecot) {
		return true
	}

	return false
}

// SetDovecot gets a reference to the given bool and assigns it to the Dovecot field.
func (o *MailboxResponse) SetDovecot(v bool) {
	o.Dovecot = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *MailboxResponse) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *MailboxResponse) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *MailboxResponse) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetLeaveMessages returns the LeaveMessages field value if set, zero value otherwise.
func (o *MailboxResponse) GetLeaveMessages() bool {
	if o == nil || IsNil(o.LeaveMessages) {
		var ret bool
		return ret
	}
	return *o.LeaveMessages
}

// GetLeaveMessagesOk returns a tuple with the LeaveMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetLeaveMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.LeaveMessages) {
		return nil, false
	}
	return o.LeaveMessages, true
}

// HasLeaveMessages returns a boolean if a field has been set.
func (o *MailboxResponse) HasLeaveMessages() bool {
	if o != nil && !IsNil(o.LeaveMessages) {
		return true
	}

	return false
}

// SetLeaveMessages gets a reference to the given bool and assigns it to the LeaveMessages field.
func (o *MailboxResponse) SetLeaveMessages(v bool) {
	o.LeaveMessages = &v
}

// GetMailbox returns the Mailbox field value if set, zero value otherwise.
func (o *MailboxResponse) GetMailbox() string {
	if o == nil || IsNil(o.Mailbox) {
		var ret string
		return ret
	}
	return *o.Mailbox
}

// GetMailboxOk returns a tuple with the Mailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetMailboxOk() (*string, bool) {
	if o == nil || IsNil(o.Mailbox) {
		return nil, false
	}
	return o.Mailbox, true
}

// HasMailbox returns a boolean if a field has been set.
func (o *MailboxResponse) HasMailbox() bool {
	if o != nil && !IsNil(o.Mailbox) {
		return true
	}

	return false
}

// SetMailbox gets a reference to the given string and assigns it to the Mailbox field.
func (o *MailboxResponse) SetMailbox(v string) {
	o.Mailbox = &v
}

// GetOwnerFullName returns the OwnerFullName field value if set, zero value otherwise.
func (o *MailboxResponse) GetOwnerFullName() string {
	if o == nil || IsNil(o.OwnerFullName) {
		var ret string
		return ret
	}
	return *o.OwnerFullName
}

// GetOwnerFullNameOk returns a tuple with the OwnerFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailboxResponse) GetOwnerFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerFullName) {
		return nil, false
	}
	return o.OwnerFullName, true
}

// HasOwnerFullName returns a boolean if a field has been set.
func (o *MailboxResponse) HasOwnerFullName() bool {
	if o != nil && !IsNil(o.OwnerFullName) {
		return true
	}

	return false
}

// SetOwnerFullName gets a reference to the given string and assigns it to the OwnerFullName field.
func (o *MailboxResponse) SetOwnerFullName(v string) {
	o.OwnerFullName = &v
}

func (o MailboxResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailboxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdnName) {
		toSerialize["idn_name"] = o.IdnName
	}
	if !IsNil(o.AutoreplyMessage) {
		toSerialize["autoreply_message"] = o.AutoreplyMessage
	}
	if !IsNil(o.AutoreplyStatus) {
		toSerialize["autoreply_status"] = o.AutoreplyStatus
	}
	if !IsNil(o.AutoreplySubject) {
		toSerialize["autoreply_subject"] = o.AutoreplySubject
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.FilterAction) {
		toSerialize["filter_action"] = o.FilterAction
	}
	if !IsNil(o.FilterStatus) {
		toSerialize["filter_status"] = o.FilterStatus
	}
	if !IsNil(o.ForwardList) {
		toSerialize["forward_list"] = o.ForwardList
	}
	if !IsNil(o.ForwardStatus) {
		toSerialize["forward_status"] = o.ForwardStatus
	}
	if !IsNil(o.OutgoingControl) {
		toSerialize["outgoing_control"] = o.OutgoingControl
	}
	if !IsNil(o.OutgoingEmail) {
		toSerialize["outgoing_email"] = o.OutgoingEmail
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.WhiteList) {
		toSerialize["white_list"] = o.WhiteList
	}
	if !IsNil(o.Webmail) {
		toSerialize["webmail"] = o.Webmail
	}
	if !IsNil(o.Dovecot) {
		toSerialize["dovecot"] = o.Dovecot
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.LeaveMessages) {
		toSerialize["leave_messages"] = o.LeaveMessages
	}
	if !IsNil(o.Mailbox) {
		toSerialize["mailbox"] = o.Mailbox
	}
	if !IsNil(o.OwnerFullName) {
		toSerialize["owner_full_name"] = o.OwnerFullName
	}
	return toSerialize, nil
}

type NullableMailboxResponse struct {
	value *MailboxResponse
	isSet bool
}

func (v NullableMailboxResponse) Get() *MailboxResponse {
	return v.value
}

func (v *NullableMailboxResponse) Set(val *MailboxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMailboxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMailboxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailboxResponse(val *MailboxResponse) *NullableMailboxResponse {
	return &NullableMailboxResponse{value: val, isSet: true}
}

func (v NullableMailboxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailboxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


