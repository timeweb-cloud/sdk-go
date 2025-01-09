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

// checks if the ClusterOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterOut{}

// ClusterOut struct for ClusterOut
type ClusterOut struct {
	// ID кластера
	Id int32 `json:"id"`
	// Название
	Name string `json:"name"`
	// Дата и время создания кластера в формате ISO8601
	CreatedAt time.Time `json:"created_at"`
	// Статус
	Status string `json:"status"`
	// Описание
	Description string `json:"description"`
	// Версия Kubernetes
	K8sVersion string `json:"k8s_version"`
	// Используемый сетевой драйвер
	NetworkDriver string `json:"network_driver"`
	// Логическое значение, показывающее, включен ли Ingress
	Ingress bool `json:"ingress"`
	// ID тарифа мастер-ноды
	PresetId int32 `json:"preset_id"`
	// Общее количество ядер
	Cpu *int32 `json:"cpu,omitempty"`
	// Общее количество памяти
	Ram *int32 `json:"ram,omitempty"`
	// Общее дисковое пространство
	Disk *int32 `json:"disk,omitempty"`
	// Зона доступности
	AvailabilityZone *string `json:"availability_zone,omitempty"`
	// ID проекта
	ProjectId *int32 `json:"project_id,omitempty"`
}

// NewClusterOut instantiates a new ClusterOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterOut(id int32, name string, createdAt time.Time, status string, description string, k8sVersion string, networkDriver string, ingress bool, presetId int32) *ClusterOut {
	this := ClusterOut{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.Status = status
	this.Description = description
	this.K8sVersion = k8sVersion
	this.NetworkDriver = networkDriver
	this.Ingress = ingress
	this.PresetId = presetId
	var cpu int32 = 0
	this.Cpu = &cpu
	var ram int32 = 0
	this.Ram = &ram
	var disk int32 = 0
	this.Disk = &disk
	return &this
}

// NewClusterOutWithDefaults instantiates a new ClusterOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterOutWithDefaults() *ClusterOut {
	this := ClusterOut{}
	var cpu int32 = 0
	this.Cpu = &cpu
	var ram int32 = 0
	this.Ram = &ram
	var disk int32 = 0
	this.Disk = &disk
	return &this
}

// GetId returns the Id field value
func (o *ClusterOut) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterOut) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ClusterOut) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterOut) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ClusterOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ClusterOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStatus returns the Status field value
func (o *ClusterOut) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ClusterOut) SetStatus(v string) {
	o.Status = v
}

// GetDescription returns the Description field value
func (o *ClusterOut) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClusterOut) SetDescription(v string) {
	o.Description = v
}

// GetK8sVersion returns the K8sVersion field value
func (o *ClusterOut) GetK8sVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.K8sVersion
}

// GetK8sVersionOk returns a tuple with the K8sVersion field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetK8sVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.K8sVersion, true
}

// SetK8sVersion sets field value
func (o *ClusterOut) SetK8sVersion(v string) {
	o.K8sVersion = v
}

// GetNetworkDriver returns the NetworkDriver field value
func (o *ClusterOut) GetNetworkDriver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkDriver
}

// GetNetworkDriverOk returns a tuple with the NetworkDriver field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetNetworkDriverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkDriver, true
}

// SetNetworkDriver sets field value
func (o *ClusterOut) SetNetworkDriver(v string) {
	o.NetworkDriver = v
}

// GetIngress returns the Ingress field value
func (o *ClusterOut) GetIngress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ingress
}

// GetIngressOk returns a tuple with the Ingress field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetIngressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ingress, true
}

// SetIngress sets field value
func (o *ClusterOut) SetIngress(v bool) {
	o.Ingress = v
}

// GetPresetId returns the PresetId field value
func (o *ClusterOut) GetPresetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetPresetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *ClusterOut) SetPresetId(v int32) {
	o.PresetId = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ClusterOut) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ClusterOut) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ClusterOut) SetCpu(v int32) {
	o.Cpu = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *ClusterOut) GetRam() int32 {
	if o == nil || IsNil(o.Ram) {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetRamOk() (*int32, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *ClusterOut) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *ClusterOut) SetRam(v int32) {
	o.Ram = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *ClusterOut) GetDisk() int32 {
	if o == nil || IsNil(o.Disk) {
		var ret int32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetDiskOk() (*int32, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *ClusterOut) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given int32 and assigns it to the Disk field.
func (o *ClusterOut) SetDisk(v int32) {
	o.Disk = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *ClusterOut) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *ClusterOut) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *ClusterOut) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ClusterOut) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOut) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ClusterOut) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ClusterOut) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o ClusterOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["status"] = o.Status
	toSerialize["description"] = o.Description
	toSerialize["k8s_version"] = o.K8sVersion
	toSerialize["network_driver"] = o.NetworkDriver
	toSerialize["ingress"] = o.Ingress
	toSerialize["preset_id"] = o.PresetId
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availability_zone"] = o.AvailabilityZone
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableClusterOut struct {
	value *ClusterOut
	isSet bool
}

func (v NullableClusterOut) Get() *ClusterOut {
	return v.value
}

func (v *NullableClusterOut) Set(val *ClusterOut) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterOut) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterOut(val *ClusterOut) *NullableClusterOut {
	return &NullableClusterOut{value: val, isSet: true}
}

func (v NullableClusterOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


