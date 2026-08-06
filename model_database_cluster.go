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

// checks if the DatabaseCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseCluster{}

// DatabaseCluster Кластер базы данных
type DatabaseCluster struct {
	// ID для каждого экземпляра базы данных. Автоматически генерируется при создании.
	Id float32 `json:"id"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда была создана база данных.
	CreatedAt string `json:"created_at"`
	// Локация сервера.
	Location NullableString `json:"location"`
	// Название кластера базы данных.
	Name string `json:"name"`
	// Описание кластера базы данных.
	Description string `json:"description"`
	// Список сетей кластера базы данных.
	Networks []DatabaseClusterNetworksInner `json:"networks"`
	// Использование публичного IPv6-адреса.
	IsEnabledPublicIpv6 bool `json:"is_enabled_public_ipv6"`
	// Тип базы данных. Список возможных значений шире, чем список типов, доступных при создании нового кластера.
	Type NullableString `json:"type"`
	// Тип хеширования кластера базы данных (mysql5 | mysql | postgres).
	HashType NullableString `json:"hash_type"`
	// Ссылка на аватар для базы данных.
	AvatarLink NullableString `json:"avatar_link"`
	// Порт
	Port NullableInt32 `json:"port"`
	// Текущий статус кластера базы данных. Значение `read_only` означает, что запись в кластер заблокирована из-за переполнения диска — чтобы снять блокировку, освободите место или увеличьте размер диска.
	Status string `json:"status"`
	// ID тарифа. Равен `null` у кластеров, созданных через конфигуратор — в этом случае заполнен `configurator_id`.
	PresetId NullableInt32 `json:"preset_id"`
	// ID конфигуратора. Равен `null` у кластеров, созданных по тарифу.
	ConfiguratorId NullableInt32 `json:"configurator_id"`
	// Количество ядер процессора.
	Cpu NullableInt32 `json:"cpu"`
	// Частота процессора.
	CpuFrequency NullableString `json:"cpu_frequency"`
	// Используются ли выделенные ядра процессора.
	IsDedicatedCpu bool `json:"is_dedicated_cpu"`
	// Объем оперативной памяти (в Мб).
	Ram NullableInt32 `json:"ram"`
	Disk NullableDatabaseClusterDisk `json:"disk"`
	// Подключен ли к кластеру дополнительный диск.
	HasAdditionalDisk bool `json:"has_additional_disk"`
	DiskAutoscaling NullableDatabaseClusterDiskAutoscaling `json:"disk_autoscaling"`
	ConfigParameters Mysql `json:"config_parameters"`
	// Доступность публичного IP-адреса
	IsEnabledPublicNetwork bool `json:"is_enabled_public_network"`
	// Включено ли защищенное подключение к кластеру базы данных.
	IsSecureConnectionEnabled bool `json:"is_secure_connection_enabled"`
	// Включены ли автоматические резервные копии кластера базы данных.
	IsAutobackupsEnabled bool `json:"is_autobackups_enabled"`
	// Включено ли расписание резервного копирования кластера базы данных.
	IsBackupScheduleEnabled bool `json:"is_backup_schedule_enabled"`
	AvailabilityZone AvailabilityZone `json:"availability_zone"`
	// ID проекта, в котором находится кластер базы данных.
	ProjectId *int32 `json:"project_id,omitempty"`
	// Список реплик кластера базы данных.
	ReplicaList []DatabaseClusterReplicaListInner `json:"replica_list"`
	// Список доменов кластера базы данных. Если публичная сеть отключена (`is_enabled_public_network: false`), список всегда пустой.
	Domains []DatabaseClusterDomainsInner `json:"domains"`
	// Список дочерних сервисов кластера базы данных.
	ChildServices []DatabaseClusterChildServicesInner `json:"child_services"`
	// Список родительских сервисов кластера базы данных.
	ParentServices []DatabaseClusterParentServicesInner `json:"parent_services"`
	MaintenanceSlot DatabaseClusterMaintenanceSlot `json:"maintenance_slot"`
}

// NewDatabaseCluster instantiates a new DatabaseCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseCluster(id float32, createdAt string, location NullableString, name string, description string, networks []DatabaseClusterNetworksInner, isEnabledPublicIpv6 bool, type_ NullableString, hashType NullableString, avatarLink NullableString, port NullableInt32, status string, presetId NullableInt32, configuratorId NullableInt32, cpu NullableInt32, cpuFrequency NullableString, isDedicatedCpu bool, ram NullableInt32, disk NullableDatabaseClusterDisk, hasAdditionalDisk bool, diskAutoscaling NullableDatabaseClusterDiskAutoscaling, configParameters Mysql, isEnabledPublicNetwork bool, isSecureConnectionEnabled bool, isAutobackupsEnabled bool, isBackupScheduleEnabled bool, availabilityZone AvailabilityZone, replicaList []DatabaseClusterReplicaListInner, domains []DatabaseClusterDomainsInner, childServices []DatabaseClusterChildServicesInner, parentServices []DatabaseClusterParentServicesInner, maintenanceSlot DatabaseClusterMaintenanceSlot) *DatabaseCluster {
	this := DatabaseCluster{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Location = location
	this.Name = name
	this.Description = description
	this.Networks = networks
	this.IsEnabledPublicIpv6 = isEnabledPublicIpv6
	this.Type = type_
	this.HashType = hashType
	this.AvatarLink = avatarLink
	this.Port = port
	this.Status = status
	this.PresetId = presetId
	this.ConfiguratorId = configuratorId
	this.Cpu = cpu
	this.CpuFrequency = cpuFrequency
	this.IsDedicatedCpu = isDedicatedCpu
	this.Ram = ram
	this.Disk = disk
	this.HasAdditionalDisk = hasAdditionalDisk
	this.DiskAutoscaling = diskAutoscaling
	this.ConfigParameters = configParameters
	this.IsEnabledPublicNetwork = isEnabledPublicNetwork
	this.IsSecureConnectionEnabled = isSecureConnectionEnabled
	this.IsAutobackupsEnabled = isAutobackupsEnabled
	this.IsBackupScheduleEnabled = isBackupScheduleEnabled
	this.AvailabilityZone = availabilityZone
	this.ReplicaList = replicaList
	this.Domains = domains
	this.ChildServices = childServices
	this.ParentServices = parentServices
	this.MaintenanceSlot = maintenanceSlot
	return &this
}

// NewDatabaseClusterWithDefaults instantiates a new DatabaseCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterWithDefaults() *DatabaseCluster {
	this := DatabaseCluster{}
	return &this
}

// GetId returns the Id field value
func (o *DatabaseCluster) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DatabaseCluster) SetId(v float32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DatabaseCluster) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DatabaseCluster) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatabaseCluster) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}

	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// SetLocation sets field value
func (o *DatabaseCluster) SetLocation(v string) {
	o.Location.Set(&v)
}

// GetName returns the Name field value
func (o *DatabaseCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseCluster) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *DatabaseCluster) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DatabaseCluster) SetDescription(v string) {
	o.Description = v
}

// GetNetworks returns the Networks field value
func (o *DatabaseCluster) GetNetworks() []DatabaseClusterNetworksInner {
	if o == nil {
		var ret []DatabaseClusterNetworksInner
		return ret
	}

	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetNetworksOk() ([]DatabaseClusterNetworksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Networks, true
}

// SetNetworks sets field value
func (o *DatabaseCluster) SetNetworks(v []DatabaseClusterNetworksInner) {
	o.Networks = v
}

// GetIsEnabledPublicIpv6 returns the IsEnabledPublicIpv6 field value
func (o *DatabaseCluster) GetIsEnabledPublicIpv6() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabledPublicIpv6
}

// GetIsEnabledPublicIpv6Ok returns a tuple with the IsEnabledPublicIpv6 field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetIsEnabledPublicIpv6Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabledPublicIpv6, true
}

// SetIsEnabledPublicIpv6 sets field value
func (o *DatabaseCluster) SetIsEnabledPublicIpv6(v bool) {
	o.IsEnabledPublicIpv6 = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatabaseCluster) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *DatabaseCluster) SetType(v string) {
	o.Type.Set(&v)
}

// GetHashType returns the HashType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatabaseCluster) GetHashType() string {
	if o == nil || o.HashType.Get() == nil {
		var ret string
		return ret
	}

	return *o.HashType.Get()
}

// GetHashTypeOk returns a tuple with the HashType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetHashTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HashType.Get(), o.HashType.IsSet()
}

// SetHashType sets field value
func (o *DatabaseCluster) SetHashType(v string) {
	o.HashType.Set(&v)
}

// GetAvatarLink returns the AvatarLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatabaseCluster) GetAvatarLink() string {
	if o == nil || o.AvatarLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.AvatarLink.Get()
}

// GetAvatarLinkOk returns a tuple with the AvatarLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetAvatarLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarLink.Get(), o.AvatarLink.IsSet()
}

// SetAvatarLink sets field value
func (o *DatabaseCluster) SetAvatarLink(v string) {
	o.AvatarLink.Set(&v)
}

// GetPort returns the Port field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DatabaseCluster) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// SetPort sets field value
func (o *DatabaseCluster) SetPort(v int32) {
	o.Port.Set(&v)
}

// GetStatus returns the Status field value
func (o *DatabaseCluster) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DatabaseCluster) SetStatus(v string) {
	o.Status = v
}

// GetPresetId returns the PresetId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DatabaseCluster) GetPresetId() int32 {
	if o == nil || o.PresetId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PresetId.Get()
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetPresetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresetId.Get(), o.PresetId.IsSet()
}

// SetPresetId sets field value
func (o *DatabaseCluster) SetPresetId(v int32) {
	o.PresetId.Set(&v)
}

// GetConfiguratorId returns the ConfiguratorId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DatabaseCluster) GetConfiguratorId() int32 {
	if o == nil || o.ConfiguratorId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ConfiguratorId.Get()
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetConfiguratorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfiguratorId.Get(), o.ConfiguratorId.IsSet()
}

// SetConfiguratorId sets field value
func (o *DatabaseCluster) SetConfiguratorId(v int32) {
	o.ConfiguratorId.Set(&v)
}

// GetCpu returns the Cpu field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DatabaseCluster) GetCpu() int32 {
	if o == nil || o.Cpu.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// SetCpu sets field value
func (o *DatabaseCluster) SetCpu(v int32) {
	o.Cpu.Set(&v)
}

// GetCpuFrequency returns the CpuFrequency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatabaseCluster) GetCpuFrequency() string {
	if o == nil || o.CpuFrequency.Get() == nil {
		var ret string
		return ret
	}

	return *o.CpuFrequency.Get()
}

// GetCpuFrequencyOk returns a tuple with the CpuFrequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetCpuFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuFrequency.Get(), o.CpuFrequency.IsSet()
}

// SetCpuFrequency sets field value
func (o *DatabaseCluster) SetCpuFrequency(v string) {
	o.CpuFrequency.Set(&v)
}

// GetIsDedicatedCpu returns the IsDedicatedCpu field value
func (o *DatabaseCluster) GetIsDedicatedCpu() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDedicatedCpu
}

// GetIsDedicatedCpuOk returns a tuple with the IsDedicatedCpu field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetIsDedicatedCpuOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDedicatedCpu, true
}

// SetIsDedicatedCpu sets field value
func (o *DatabaseCluster) SetIsDedicatedCpu(v bool) {
	o.IsDedicatedCpu = v
}

// GetRam returns the Ram field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DatabaseCluster) GetRam() int32 {
	if o == nil || o.Ram.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Ram.Get()
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ram.Get(), o.Ram.IsSet()
}

// SetRam sets field value
func (o *DatabaseCluster) SetRam(v int32) {
	o.Ram.Set(&v)
}

// GetDisk returns the Disk field value
// If the value is explicit nil, the zero value for DatabaseClusterDisk will be returned
func (o *DatabaseCluster) GetDisk() DatabaseClusterDisk {
	if o == nil || o.Disk.Get() == nil {
		var ret DatabaseClusterDisk
		return ret
	}

	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetDiskOk() (*DatabaseClusterDisk, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// SetDisk sets field value
func (o *DatabaseCluster) SetDisk(v DatabaseClusterDisk) {
	o.Disk.Set(&v)
}

// GetHasAdditionalDisk returns the HasAdditionalDisk field value
func (o *DatabaseCluster) GetHasAdditionalDisk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasAdditionalDisk
}

// GetHasAdditionalDiskOk returns a tuple with the HasAdditionalDisk field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetHasAdditionalDiskOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAdditionalDisk, true
}

// SetHasAdditionalDisk sets field value
func (o *DatabaseCluster) SetHasAdditionalDisk(v bool) {
	o.HasAdditionalDisk = v
}

// GetDiskAutoscaling returns the DiskAutoscaling field value
// If the value is explicit nil, the zero value for DatabaseClusterDiskAutoscaling will be returned
func (o *DatabaseCluster) GetDiskAutoscaling() DatabaseClusterDiskAutoscaling {
	if o == nil || o.DiskAutoscaling.Get() == nil {
		var ret DatabaseClusterDiskAutoscaling
		return ret
	}

	return *o.DiskAutoscaling.Get()
}

// GetDiskAutoscalingOk returns a tuple with the DiskAutoscaling field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseCluster) GetDiskAutoscalingOk() (*DatabaseClusterDiskAutoscaling, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskAutoscaling.Get(), o.DiskAutoscaling.IsSet()
}

// SetDiskAutoscaling sets field value
func (o *DatabaseCluster) SetDiskAutoscaling(v DatabaseClusterDiskAutoscaling) {
	o.DiskAutoscaling.Set(&v)
}

// GetConfigParameters returns the ConfigParameters field value
func (o *DatabaseCluster) GetConfigParameters() Mysql {
	if o == nil {
		var ret Mysql
		return ret
	}

	return o.ConfigParameters
}

// GetConfigParametersOk returns a tuple with the ConfigParameters field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetConfigParametersOk() (*Mysql, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigParameters, true
}

// SetConfigParameters sets field value
func (o *DatabaseCluster) SetConfigParameters(v Mysql) {
	o.ConfigParameters = v
}

// GetIsEnabledPublicNetwork returns the IsEnabledPublicNetwork field value
func (o *DatabaseCluster) GetIsEnabledPublicNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabledPublicNetwork
}

// GetIsEnabledPublicNetworkOk returns a tuple with the IsEnabledPublicNetwork field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetIsEnabledPublicNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabledPublicNetwork, true
}

// SetIsEnabledPublicNetwork sets field value
func (o *DatabaseCluster) SetIsEnabledPublicNetwork(v bool) {
	o.IsEnabledPublicNetwork = v
}

// GetIsSecureConnectionEnabled returns the IsSecureConnectionEnabled field value
func (o *DatabaseCluster) GetIsSecureConnectionEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSecureConnectionEnabled
}

// GetIsSecureConnectionEnabledOk returns a tuple with the IsSecureConnectionEnabled field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetIsSecureConnectionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSecureConnectionEnabled, true
}

// SetIsSecureConnectionEnabled sets field value
func (o *DatabaseCluster) SetIsSecureConnectionEnabled(v bool) {
	o.IsSecureConnectionEnabled = v
}

// GetIsAutobackupsEnabled returns the IsAutobackupsEnabled field value
func (o *DatabaseCluster) GetIsAutobackupsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutobackupsEnabled
}

// GetIsAutobackupsEnabledOk returns a tuple with the IsAutobackupsEnabled field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetIsAutobackupsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutobackupsEnabled, true
}

// SetIsAutobackupsEnabled sets field value
func (o *DatabaseCluster) SetIsAutobackupsEnabled(v bool) {
	o.IsAutobackupsEnabled = v
}

// GetIsBackupScheduleEnabled returns the IsBackupScheduleEnabled field value
func (o *DatabaseCluster) GetIsBackupScheduleEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBackupScheduleEnabled
}

// GetIsBackupScheduleEnabledOk returns a tuple with the IsBackupScheduleEnabled field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetIsBackupScheduleEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBackupScheduleEnabled, true
}

// SetIsBackupScheduleEnabled sets field value
func (o *DatabaseCluster) SetIsBackupScheduleEnabled(v bool) {
	o.IsBackupScheduleEnabled = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *DatabaseCluster) GetAvailabilityZone() AvailabilityZone {
	if o == nil {
		var ret AvailabilityZone
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *DatabaseCluster) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DatabaseCluster) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DatabaseCluster) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DatabaseCluster) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetReplicaList returns the ReplicaList field value
func (o *DatabaseCluster) GetReplicaList() []DatabaseClusterReplicaListInner {
	if o == nil {
		var ret []DatabaseClusterReplicaListInner
		return ret
	}

	return o.ReplicaList
}

// GetReplicaListOk returns a tuple with the ReplicaList field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetReplicaListOk() ([]DatabaseClusterReplicaListInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicaList, true
}

// SetReplicaList sets field value
func (o *DatabaseCluster) SetReplicaList(v []DatabaseClusterReplicaListInner) {
	o.ReplicaList = v
}

// GetDomains returns the Domains field value
func (o *DatabaseCluster) GetDomains() []DatabaseClusterDomainsInner {
	if o == nil {
		var ret []DatabaseClusterDomainsInner
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetDomainsOk() ([]DatabaseClusterDomainsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *DatabaseCluster) SetDomains(v []DatabaseClusterDomainsInner) {
	o.Domains = v
}

// GetChildServices returns the ChildServices field value
func (o *DatabaseCluster) GetChildServices() []DatabaseClusterChildServicesInner {
	if o == nil {
		var ret []DatabaseClusterChildServicesInner
		return ret
	}

	return o.ChildServices
}

// GetChildServicesOk returns a tuple with the ChildServices field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetChildServicesOk() ([]DatabaseClusterChildServicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChildServices, true
}

// SetChildServices sets field value
func (o *DatabaseCluster) SetChildServices(v []DatabaseClusterChildServicesInner) {
	o.ChildServices = v
}

// GetParentServices returns the ParentServices field value
func (o *DatabaseCluster) GetParentServices() []DatabaseClusterParentServicesInner {
	if o == nil {
		var ret []DatabaseClusterParentServicesInner
		return ret
	}

	return o.ParentServices
}

// GetParentServicesOk returns a tuple with the ParentServices field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetParentServicesOk() ([]DatabaseClusterParentServicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentServices, true
}

// SetParentServices sets field value
func (o *DatabaseCluster) SetParentServices(v []DatabaseClusterParentServicesInner) {
	o.ParentServices = v
}

// GetMaintenanceSlot returns the MaintenanceSlot field value
func (o *DatabaseCluster) GetMaintenanceSlot() DatabaseClusterMaintenanceSlot {
	if o == nil {
		var ret DatabaseClusterMaintenanceSlot
		return ret
	}

	return o.MaintenanceSlot
}

// GetMaintenanceSlotOk returns a tuple with the MaintenanceSlot field value
// and a boolean to check if the value has been set.
func (o *DatabaseCluster) GetMaintenanceSlotOk() (*DatabaseClusterMaintenanceSlot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaintenanceSlot, true
}

// SetMaintenanceSlot sets field value
func (o *DatabaseCluster) SetMaintenanceSlot(v DatabaseClusterMaintenanceSlot) {
	o.MaintenanceSlot = v
}

func (o DatabaseCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["location"] = o.Location.Get()
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["networks"] = o.Networks
	toSerialize["is_enabled_public_ipv6"] = o.IsEnabledPublicIpv6
	toSerialize["type"] = o.Type.Get()
	toSerialize["hash_type"] = o.HashType.Get()
	toSerialize["avatar_link"] = o.AvatarLink.Get()
	toSerialize["port"] = o.Port.Get()
	toSerialize["status"] = o.Status
	toSerialize["preset_id"] = o.PresetId.Get()
	toSerialize["configurator_id"] = o.ConfiguratorId.Get()
	toSerialize["cpu"] = o.Cpu.Get()
	toSerialize["cpu_frequency"] = o.CpuFrequency.Get()
	toSerialize["is_dedicated_cpu"] = o.IsDedicatedCpu
	toSerialize["ram"] = o.Ram.Get()
	toSerialize["disk"] = o.Disk.Get()
	toSerialize["has_additional_disk"] = o.HasAdditionalDisk
	toSerialize["disk_autoscaling"] = o.DiskAutoscaling.Get()
	toSerialize["config_parameters"] = o.ConfigParameters
	toSerialize["is_enabled_public_network"] = o.IsEnabledPublicNetwork
	toSerialize["is_secure_connection_enabled"] = o.IsSecureConnectionEnabled
	toSerialize["is_autobackups_enabled"] = o.IsAutobackupsEnabled
	toSerialize["is_backup_schedule_enabled"] = o.IsBackupScheduleEnabled
	toSerialize["availability_zone"] = o.AvailabilityZone
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	toSerialize["replica_list"] = o.ReplicaList
	toSerialize["domains"] = o.Domains
	toSerialize["child_services"] = o.ChildServices
	toSerialize["parent_services"] = o.ParentServices
	toSerialize["maintenance_slot"] = o.MaintenanceSlot
	return toSerialize, nil
}

type NullableDatabaseCluster struct {
	value *DatabaseCluster
	isSet bool
}

func (v NullableDatabaseCluster) Get() *DatabaseCluster {
	return v.value
}

func (v *NullableDatabaseCluster) Set(val *DatabaseCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseCluster(val *DatabaseCluster) *NullableDatabaseCluster {
	return &NullableDatabaseCluster{value: val, isSet: true}
}

func (v NullableDatabaseCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


