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

// checks if the MailboxV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailboxV2{}

// MailboxV2 Почтовый ящик (API v2)
type MailboxV2 struct {
	// IDN домен почтового ящика
	IdnName string `json:"idn_name"`
	// Сообщение автоответчика на входящие письма
	AutoreplyMessage string `json:"autoreply_message"`
	// Включен ли автоответчик на входящие письма
	AutoreplyStatus bool `json:"autoreply_status"`
	// Тема сообщения автоответчика на входящие письма
	AutoreplySubject string `json:"autoreply_subject"`
	// Комментарий к почтовому ящику
	Comment string `json:"comment"`
	// Что делать с письмами, которые попадают в спам
	FilterAction string `json:"filter_action"`
	// Включен ли спам-фильтр
	FilterStatus bool `json:"filter_status"`
	// Список адресов для пересылки входящих писем
	ForwardList []string `json:"forward_list"`
	// Включена ли пересылка входящих писем
	ForwardStatus bool `json:"forward_status"`
	// Включена ли пересылка исходящих писем
	OutgoingControl bool `json:"outgoing_control"`
	// Адрес для пересылки исходящих писем
	OutgoingEmail string `json:"outgoing_email"`
	// Пароль почтового ящика (всегда возвращается пустой строкой)
	Password string `json:"password"`
	// Адрес для пересылки спама при выбранном действии forward
	Spambox string `json:"spambox"`
	// Белый список адресов от которых письма не будут попадать в спам
	WhiteList []string `json:"white_list"`
	// Доступен ли Webmail
	Webmail bool `json:"webmail"`
	// Есть ли доступ через dovecot
	Dovecot bool `json:"dovecot"`
	// Домен почты
	Fqdn string `json:"fqdn"`
	// Оставлять ли сообщения на сервере при пересылке
	LeaveMessages bool `json:"leave_messages"`
	// Название почтового ящика
	Mailbox string `json:"mailbox"`
	// ФИО владельца почтового ящика
	OwnerFullName string `json:"owner_full_name"`
}

// NewMailboxV2 instantiates a new MailboxV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailboxV2(idnName string, autoreplyMessage string, autoreplyStatus bool, autoreplySubject string, comment string, filterAction string, filterStatus bool, forwardList []string, forwardStatus bool, outgoingControl bool, outgoingEmail string, password string, spambox string, whiteList []string, webmail bool, dovecot bool, fqdn string, leaveMessages bool, mailbox string, ownerFullName string) *MailboxV2 {
	this := MailboxV2{}
	this.IdnName = idnName
	this.AutoreplyMessage = autoreplyMessage
	this.AutoreplyStatus = autoreplyStatus
	this.AutoreplySubject = autoreplySubject
	this.Comment = comment
	this.FilterAction = filterAction
	this.FilterStatus = filterStatus
	this.ForwardList = forwardList
	this.ForwardStatus = forwardStatus
	this.OutgoingControl = outgoingControl
	this.OutgoingEmail = outgoingEmail
	this.Password = password
	this.Spambox = spambox
	this.WhiteList = whiteList
	this.Webmail = webmail
	this.Dovecot = dovecot
	this.Fqdn = fqdn
	this.LeaveMessages = leaveMessages
	this.Mailbox = mailbox
	this.OwnerFullName = ownerFullName
	return &this
}

// NewMailboxV2WithDefaults instantiates a new MailboxV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailboxV2WithDefaults() *MailboxV2 {
	this := MailboxV2{}
	return &this
}

// GetIdnName returns the IdnName field value
func (o *MailboxV2) GetIdnName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdnName
}

// GetIdnNameOk returns a tuple with the IdnName field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetIdnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdnName, true
}

// SetIdnName sets field value
func (o *MailboxV2) SetIdnName(v string) {
	o.IdnName = v
}

// GetAutoreplyMessage returns the AutoreplyMessage field value
func (o *MailboxV2) GetAutoreplyMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoreplyMessage
}

// GetAutoreplyMessageOk returns a tuple with the AutoreplyMessage field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetAutoreplyMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoreplyMessage, true
}

// SetAutoreplyMessage sets field value
func (o *MailboxV2) SetAutoreplyMessage(v string) {
	o.AutoreplyMessage = v
}

// GetAutoreplyStatus returns the AutoreplyStatus field value
func (o *MailboxV2) GetAutoreplyStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoreplyStatus
}

// GetAutoreplyStatusOk returns a tuple with the AutoreplyStatus field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetAutoreplyStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoreplyStatus, true
}

// SetAutoreplyStatus sets field value
func (o *MailboxV2) SetAutoreplyStatus(v bool) {
	o.AutoreplyStatus = v
}

// GetAutoreplySubject returns the AutoreplySubject field value
func (o *MailboxV2) GetAutoreplySubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoreplySubject
}

// GetAutoreplySubjectOk returns a tuple with the AutoreplySubject field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetAutoreplySubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoreplySubject, true
}

// SetAutoreplySubject sets field value
func (o *MailboxV2) SetAutoreplySubject(v string) {
	o.AutoreplySubject = v
}

// GetComment returns the Comment field value
func (o *MailboxV2) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *MailboxV2) SetComment(v string) {
	o.Comment = v
}

// GetFilterAction returns the FilterAction field value
func (o *MailboxV2) GetFilterAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterAction
}

// GetFilterActionOk returns a tuple with the FilterAction field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetFilterActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterAction, true
}

// SetFilterAction sets field value
func (o *MailboxV2) SetFilterAction(v string) {
	o.FilterAction = v
}

// GetFilterStatus returns the FilterStatus field value
func (o *MailboxV2) GetFilterStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FilterStatus
}

// GetFilterStatusOk returns a tuple with the FilterStatus field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetFilterStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterStatus, true
}

// SetFilterStatus sets field value
func (o *MailboxV2) SetFilterStatus(v bool) {
	o.FilterStatus = v
}

// GetForwardList returns the ForwardList field value
func (o *MailboxV2) GetForwardList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ForwardList
}

// GetForwardListOk returns a tuple with the ForwardList field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetForwardListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForwardList, true
}

// SetForwardList sets field value
func (o *MailboxV2) SetForwardList(v []string) {
	o.ForwardList = v
}

// GetForwardStatus returns the ForwardStatus field value
func (o *MailboxV2) GetForwardStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForwardStatus
}

// GetForwardStatusOk returns a tuple with the ForwardStatus field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetForwardStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwardStatus, true
}

// SetForwardStatus sets field value
func (o *MailboxV2) SetForwardStatus(v bool) {
	o.ForwardStatus = v
}

// GetOutgoingControl returns the OutgoingControl field value
func (o *MailboxV2) GetOutgoingControl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OutgoingControl
}

// GetOutgoingControlOk returns a tuple with the OutgoingControl field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetOutgoingControlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutgoingControl, true
}

// SetOutgoingControl sets field value
func (o *MailboxV2) SetOutgoingControl(v bool) {
	o.OutgoingControl = v
}

// GetOutgoingEmail returns the OutgoingEmail field value
func (o *MailboxV2) GetOutgoingEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutgoingEmail
}

// GetOutgoingEmailOk returns a tuple with the OutgoingEmail field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetOutgoingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutgoingEmail, true
}

// SetOutgoingEmail sets field value
func (o *MailboxV2) SetOutgoingEmail(v string) {
	o.OutgoingEmail = v
}

// GetPassword returns the Password field value
func (o *MailboxV2) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *MailboxV2) SetPassword(v string) {
	o.Password = v
}

// GetSpambox returns the Spambox field value
func (o *MailboxV2) GetSpambox() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Spambox
}

// GetSpamboxOk returns a tuple with the Spambox field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetSpamboxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spambox, true
}

// SetSpambox sets field value
func (o *MailboxV2) SetSpambox(v string) {
	o.Spambox = v
}

// GetWhiteList returns the WhiteList field value
func (o *MailboxV2) GetWhiteList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WhiteList
}

// GetWhiteListOk returns a tuple with the WhiteList field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetWhiteListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhiteList, true
}

// SetWhiteList sets field value
func (o *MailboxV2) SetWhiteList(v []string) {
	o.WhiteList = v
}

// GetWebmail returns the Webmail field value
func (o *MailboxV2) GetWebmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Webmail
}

// GetWebmailOk returns a tuple with the Webmail field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetWebmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Webmail, true
}

// SetWebmail sets field value
func (o *MailboxV2) SetWebmail(v bool) {
	o.Webmail = v
}

// GetDovecot returns the Dovecot field value
func (o *MailboxV2) GetDovecot() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dovecot
}

// GetDovecotOk returns a tuple with the Dovecot field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetDovecotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dovecot, true
}

// SetDovecot sets field value
func (o *MailboxV2) SetDovecot(v bool) {
	o.Dovecot = v
}

// GetFqdn returns the Fqdn field value
func (o *MailboxV2) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *MailboxV2) SetFqdn(v string) {
	o.Fqdn = v
}

// GetLeaveMessages returns the LeaveMessages field value
func (o *MailboxV2) GetLeaveMessages() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LeaveMessages
}

// GetLeaveMessagesOk returns a tuple with the LeaveMessages field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetLeaveMessagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeaveMessages, true
}

// SetLeaveMessages sets field value
func (o *MailboxV2) SetLeaveMessages(v bool) {
	o.LeaveMessages = v
}

// GetMailbox returns the Mailbox field value
func (o *MailboxV2) GetMailbox() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mailbox
}

// GetMailboxOk returns a tuple with the Mailbox field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetMailboxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mailbox, true
}

// SetMailbox sets field value
func (o *MailboxV2) SetMailbox(v string) {
	o.Mailbox = v
}

// GetOwnerFullName returns the OwnerFullName field value
func (o *MailboxV2) GetOwnerFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerFullName
}

// GetOwnerFullNameOk returns a tuple with the OwnerFullName field value
// and a boolean to check if the value has been set.
func (o *MailboxV2) GetOwnerFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerFullName, true
}

// SetOwnerFullName sets field value
func (o *MailboxV2) SetOwnerFullName(v string) {
	o.OwnerFullName = v
}

func (o MailboxV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailboxV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["idn_name"] = o.IdnName
	toSerialize["autoreply_message"] = o.AutoreplyMessage
	toSerialize["autoreply_status"] = o.AutoreplyStatus
	toSerialize["autoreply_subject"] = o.AutoreplySubject
	toSerialize["comment"] = o.Comment
	toSerialize["filter_action"] = o.FilterAction
	toSerialize["filter_status"] = o.FilterStatus
	toSerialize["forward_list"] = o.ForwardList
	toSerialize["forward_status"] = o.ForwardStatus
	toSerialize["outgoing_control"] = o.OutgoingControl
	toSerialize["outgoing_email"] = o.OutgoingEmail
	toSerialize["password"] = o.Password
	toSerialize["spambox"] = o.Spambox
	toSerialize["white_list"] = o.WhiteList
	toSerialize["webmail"] = o.Webmail
	toSerialize["dovecot"] = o.Dovecot
	toSerialize["fqdn"] = o.Fqdn
	toSerialize["leave_messages"] = o.LeaveMessages
	toSerialize["mailbox"] = o.Mailbox
	toSerialize["owner_full_name"] = o.OwnerFullName
	return toSerialize, nil
}

type NullableMailboxV2 struct {
	value *MailboxV2
	isSet bool
}

func (v NullableMailboxV2) Get() *MailboxV2 {
	return v.value
}

func (v *NullableMailboxV2) Set(val *MailboxV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMailboxV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMailboxV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailboxV2(val *MailboxV2) *NullableMailboxV2 {
	return &NullableMailboxV2{value: val, isSet: true}
}

func (v NullableMailboxV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailboxV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


