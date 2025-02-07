# Go API client for openapi

# Введение
API Timeweb Cloud позволяет вам управлять ресурсами в облаке программным способом с использованием обычных HTTP-запросов.

Множество функций, которые доступны в панели управления Timeweb Cloud, также доступны через API, что позволяет вам автоматизировать ваши собственные сценарии.

В этой документации сперва будет описан общий дизайн и принципы работы API, а после этого конкретные конечные точки. Также будут приведены примеры запросов к ним.


## Запросы
Запросы должны выполняться по протоколу `HTTPS`, чтобы гарантировать шифрование транзакций. Поддерживаются следующие методы запроса:
|Метод|Применение|
|--- |--- |
|GET|Извлекает данные о коллекциях и отдельных ресурсах.|
|POST|Для коллекций создает новый ресурс этого типа. Также используется для выполнения действий с конкретным ресурсом.|
|PUT|Обновляет существующий ресурс.|
|PATCH|Некоторые ресурсы поддерживают частичное обновление, то есть обновление только части атрибутов ресурса, в этом случае вместо метода PUT будет использован PATCH.|
|DELETE|Удаляет ресурс.|

Методы `POST`, `PUT` и `PATCH` могут включать объект в тело запроса с типом содержимого `application/json`.

### Параметры в запросах
Некоторые коллекции поддерживают пагинацию, поиск или сортировку в запросах. В параметрах запроса требуется передать:
- `limit` — обозначает количество записей, которое необходимо вернуть
 - `offset` — указывает на смещение, относительно начала списка
 - `search` — позволяет указать набор символов для поиска
 - `sort` — можно задать правило сортировки коллекции

## Ответы
Запросы вернут один из следующих кодов состояния ответа HTTP:

|Статус|Описание|
|--- |--- |
|200 OK|Действие с ресурсом было выполнено успешно.|
|201 Created|Ресурс был успешно создан. При этом ресурс может быть как уже готовым к использованию, так и находиться в процессе запуска.|
|204 No Content|Действие с ресурсом было выполнено успешно, и ответ не содержит дополнительной информации в теле.|
|400 Bad Request|Был отправлен неверный запрос, например, в нем отсутствуют обязательные параметры и т. д. Тело ответа будет содержать дополнительную информацию об ошибке.|
|401 Unauthorized|Ошибка аутентификации.|
|403 Forbidden|Аутентификация прошла успешно, но недостаточно прав для выполнения действия.|
|404 Not Found|Запрашиваемый ресурс не найден.|
|409 Conflict|Запрос конфликтует с текущим состоянием.|
|423 Locked|Ресурс из запроса заблокирован от применения к нему указанного метода.|
|429 Too Many Requests|Был достигнут лимит по количеству запросов в единицу времени.|
|500 Internal Server Error|При выполнении запроса произошла какая-то внутренняя ошибка. Чтобы решить эту проблему, лучше всего создать тикет в панели управления.|

### Структура успешного ответа
Все конечные точки будут возвращать данные в формате `JSON`. Ответы на `GET`-запросы будут иметь на верхнем уровне следующую структуру атрибутов: 
|Название поля|Тип|Описание|
|--- |--- |--- |
|[entity_name]|object, object[], string[], number[], boolean|Динамическое поле, которое будет меняться в зависимости от запрашиваемого ресурса и будет содержать все атрибуты, необходимые для описания этого ресурса. Например, при запросе списка баз данных будет возвращаться поле `dbs`, а при запросе конкретного облачного сервера `server`. Для некоторых конечных точек в ответе может возвращаться сразу несколько ресурсов.|
|meta|object|Опционально. Объект, который содержит вспомогательную информацию о ресурсе. Чаще всего будет встречаться при запросе коллекций и содержать поле `total`, которое будет указывать на количество элементов в коллекции.|
|response_id|string|Опционально. В большинстве случаев в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID— так мы сможем найти ответ на него намного быстрее. Также вы можете использовать этот ID, чтобы убедиться, что это новый ответ на запрос и результат не был получен из кэша.|

Пример запроса на получение списка SSH-ключей:
```
    HTTP/2.0 200 OK
    {
      \"ssh_keys\":[
          {
            \"body\":\"ssh-rsa AAAAB3NzaC1sdfghjkOAsBwWhs= example@device.local\",
            \"created_at\":\"2021-09-15T19:52:27Z\",
            \"expired_at\":null,
            \"id\":5297,
            \"is_default\":false,
            \"name\":\"example@device.local\",
            \"used_at\":null,
            \"used_by\":[]
          }
      ],
      \"meta\":{
          \"total\":1
      },
      \"response_id\":\"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"
    }
```

### Структура ответа с ошибкой
|Название поля|Тип|Описание|
|--- |--- |--- |
|status_code|number|Короткий числовой идентификатор ошибки.|
|error_code|string|Короткий текстовый идентификатор ошибки, который уточняет числовой идентификатор и удобен для программной обработки. Самый простой пример — это код `not_found` для ошибки 404.|
|message|string, string[]|Опционально. В большинстве случаев в ответе будет содержаться человекочитаемое подробное описание ошибки или ошибок, которые помогут понять, что нужно исправить.|
|response_id|string|Опционально. В большинстве случае в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID — так мы сможем найти ответ на него намного быстрее.|

Пример:
```
    HTTP/2.0 403 Forbidden
    {
      \"status_code\": 403,
      \"error_code\":  \"forbidden\",
      \"message\":     \"You do not have access for the attempted action\",
      \"response_id\": \"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"
    }
```

## Статусы ресурсов
Важно учесть, что при создании большинства ресурсов внутри платформы вам будет сразу возвращен ответ от сервера со статусом `200 OK` или `201 Created` и ID созданного ресурса в теле ответа, но при этом этот ресурс может быть ещё в *состоянии запуска*.

Для того чтобы понять, в каком состоянии сейчас находится ваш ресурс, мы добавили поле `status` в ответ на получение информации о ресурсе.

Список статусов будет отличаться в зависимости от типа ресурса. Увидеть поддерживаемый список статусов вы сможете в описании каждого конкретного ресурса.

 

## Ограничение скорости запросов (Rate Limiting)
Чтобы обеспечить стабильность для всех пользователей, Timeweb Cloud защищает API от всплесков входящего трафика, анализируя количество запросов c каждого аккаунта к каждой конечной точке.

Если ваше приложение отправляет более 20 запросов в секунду на одну конечную точку, то для этого запроса API может вернуть код состояния HTTP `429 Too Many Requests`.


## Аутентификация
Доступ к API осуществляется с помощью JWT-токена. Токенами можно управлять внутри панели управления Timeweb Cloud в разделе *API и Terraform*.

Токен необходимо передавать в заголовке каждого запроса в формате:
```
  Authorization: Bearer $TIMEWEB_CLOUD_TOKEN
```

## Формат примеров API
Примеры в этой документации описаны с помощью `curl`, HTTP-клиента командной строки. На компьютерах `Linux` и `macOS` обычно по умолчанию установлен `curl`, и он доступен для загрузки на всех популярных платформах, включая `Windows`.

Каждый пример разделен на несколько строк символом `\\`, который совместим с `bash`. Типичный пример выглядит так:
```
  curl -X PATCH 
    -H \"Content-Type: application/json\" 
    -H \"Authorization: Bearer $TIMEWEB_CLOUD_TOKEN\" 
    -d '{\"name\":\"Cute Corvus\",\"comment\":\"Development Server\"}' 
    \"https://api.timeweb.cloud/api/v1/dedicated/1051\"
```
- Параметр `-X` задает метод запроса. Для согласованности метод будет указан во всех примерах, даже если он явно не требуется для методов `GET`.
- Строки `-H` задают требуемые HTTP-заголовки.
- Примеры, для которых требуется объект JSON в теле запроса, передают требуемые данные через параметр `-d`.

Чтобы использовать приведенные примеры, не подставляя каждый раз в них свой токен, вы можете добавить токен один раз в переменные окружения в вашей консоли. Например, на `Linux` это можно сделать с помощью команды:

```
TIMEWEB_CLOUD_TOKEN=\"token\"
```

После этого токен будет автоматически подставляться в ваши запросы.

Обратите внимание, что все значения в этой документации являются примерами. Не полагайтесь на IDы операционных систем, тарифов и т.д., используемые в примерах. Используйте соответствующую конечную точку для получения значений перед созданием ресурсов.


## Версионирование
API построено согласно принципам [семантического версионирования](https://semver.org/lang/ru). Это значит, что мы гарантируем обратную совместимость всех изменений в пределах одной мажорной версии.

Мажорная версия каждой конечной точки обозначается в пути запроса, например, запрос `/api/v1/servers` указывает, что этот метод имеет версию 1.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.timeweb.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIKeysAPI* | [**CreateToken**](docs/APIKeysAPI.md#createtoken) | **Post** /api/v1/auth/api-keys | Создание токена
*APIKeysAPI* | [**DeleteToken**](docs/APIKeysAPI.md#deletetoken) | **Delete** /api/v1/auth/api-keys/{token_id} | Удалить токен
*APIKeysAPI* | [**GetTokens**](docs/APIKeysAPI.md#gettokens) | **Get** /api/v1/auth/api-keys | Получение списка выпущенных токенов
*APIKeysAPI* | [**ReissueToken**](docs/APIKeysAPI.md#reissuetoken) | **Put** /api/v1/auth/api-keys/{token_id} | Перевыпустить токен
*APIKeysAPI* | [**UpdateToken**](docs/APIKeysAPI.md#updatetoken) | **Patch** /api/v1/auth/api-keys/{token_id} | Изменить токен
*AccountAPI* | [**AddCountriesToAllowedList**](docs/AccountAPI.md#addcountriestoallowedlist) | **Post** /api/v1/auth/access/countries | Добавление стран в список разрешенных
*AccountAPI* | [**AddIPsToAllowedList**](docs/AccountAPI.md#addipstoallowedlist) | **Post** /api/v1/auth/access/ips | Добавление IP-адресов в список разрешенных
*AccountAPI* | [**DeleteCountriesFromAllowedList**](docs/AccountAPI.md#deletecountriesfromallowedlist) | **Delete** /api/v1/auth/access/countries | Удаление стран из списка разрешенных
*AccountAPI* | [**DeleteIPsFromAllowedList**](docs/AccountAPI.md#deleteipsfromallowedlist) | **Delete** /api/v1/auth/access/ips | Удаление IP-адресов из списка разрешенных
*AccountAPI* | [**GetAccountStatus**](docs/AccountAPI.md#getaccountstatus) | **Get** /api/v1/account/status | Получение статуса аккаунта
*AccountAPI* | [**GetAuthAccessSettings**](docs/AccountAPI.md#getauthaccesssettings) | **Get** /api/v1/auth/access | Получить информацию о ограничениях авторизации пользователя
*AccountAPI* | [**GetCountries**](docs/AccountAPI.md#getcountries) | **Get** /api/v1/auth/access/countries | Получение списка стран
*AccountAPI* | [**GetFinances**](docs/AccountAPI.md#getfinances) | **Get** /api/v1/account/finances | Получение платежной информации
*AccountAPI* | [**GetNotificationSettings**](docs/AccountAPI.md#getnotificationsettings) | **Get** /api/v1/account/notification-settings | Получение настроек уведомлений аккаунта
*AccountAPI* | [**UpdateAuthRestrictionsByCountries**](docs/AccountAPI.md#updateauthrestrictionsbycountries) | **Post** /api/v1/auth/access/countries/enabled | Включение/отключение ограничений по стране
*AccountAPI* | [**UpdateAuthRestrictionsByIP**](docs/AccountAPI.md#updateauthrestrictionsbyip) | **Post** /api/v1/auth/access/ips/enabled | Включение/отключение ограничений по IP-адресу
*AccountAPI* | [**UpdateNotificationSettings**](docs/AccountAPI.md#updatenotificationsettings) | **Patch** /api/v1/account/notification-settings | Изменение настроек уведомлений аккаунта
*AppsAPI* | [**AddProvider**](docs/AppsAPI.md#addprovider) | **Post** /api/v1/vcs-provider | Привязка vcs провайдера
*AppsAPI* | [**CreateApp**](docs/AppsAPI.md#createapp) | **Post** /api/v1/apps | Создание приложения
*AppsAPI* | [**CreateDeploy**](docs/AppsAPI.md#createdeploy) | **Post** /api/v1/apps/{app_id}/deploy | Запуск деплоя приложения
*AppsAPI* | [**DeleteApp**](docs/AppsAPI.md#deleteapp) | **Delete** /api/v1/apps/{app_id} | Удаление приложения
*AppsAPI* | [**DeleteProvider**](docs/AppsAPI.md#deleteprovider) | **Delete** /api/v1/vcs-provider/{provider_id} | Отвязка vcs провайдера от аккаунта
*AppsAPI* | [**DeployAction**](docs/AppsAPI.md#deployaction) | **Post** /api/v1/apps/{app_id}/deploy/{deploy_id}/stop | Остановка деплоя приложения
*AppsAPI* | [**GetApp**](docs/AppsAPI.md#getapp) | **Get** /api/v1/apps/{app_id} | Получение приложения по id
*AppsAPI* | [**GetAppDeploys**](docs/AppsAPI.md#getappdeploys) | **Get** /api/v1/apps/{app_id}/deploys | Получение списка деплоев приложения
*AppsAPI* | [**GetAppLogs**](docs/AppsAPI.md#getapplogs) | **Get** /api/v1/apps/{app_id}/logs | Получение логов приложения
*AppsAPI* | [**GetAppStatistics**](docs/AppsAPI.md#getappstatistics) | **Get** /api/v1/apps/{app_id}/statistics | Получение статистики приложения
*AppsAPI* | [**GetApps**](docs/AppsAPI.md#getapps) | **Get** /api/v1/apps | Получение списка приложений
*AppsAPI* | [**GetAppsPresets**](docs/AppsAPI.md#getappspresets) | **Get** /api/v1/presets/apps | Получение списка доступных тарифов для приложения
*AppsAPI* | [**GetBranches**](docs/AppsAPI.md#getbranches) | **Get** /api/v1/vcs-provider/{provider_id}/repository/{repository_id} | Получение списка веток репозитория
*AppsAPI* | [**GetCommits**](docs/AppsAPI.md#getcommits) | **Get** /api/v1/vcs-provider/{provider_id}/repository/{repository_id}/branch | Получение списка коммитов ветки репозитория
*AppsAPI* | [**GetDeployLogs**](docs/AppsAPI.md#getdeploylogs) | **Get** /api/v1/apps/{app_id}/deploy/{deploy_id}/logs | Получение логов деплоя приложения
*AppsAPI* | [**GetDeploySettings**](docs/AppsAPI.md#getdeploysettings) | **Get** /api/v1/deploy-settings/apps | Получение списка дефолтных настроек деплоя для приложения
*AppsAPI* | [**GetFrameworks**](docs/AppsAPI.md#getframeworks) | **Get** /api/v1/frameworks/apps | Получение списка доступных фреймворков для приложения
*AppsAPI* | [**GetProviders**](docs/AppsAPI.md#getproviders) | **Get** /api/v1/vcs-provider | Получение списка vcs провайдеров
*AppsAPI* | [**GetRepositories**](docs/AppsAPI.md#getrepositories) | **Get** /api/v1/vcs-provider/{provider_id} | Получение списка репозиториев vcs провайдера
*AppsAPI* | [**UpdateAppSettings**](docs/AppsAPI.md#updateappsettings) | **Patch** /api/v1/apps/{app_id} | Изменение настроек приложения
*AppsAPI* | [**UpdateAppState**](docs/AppsAPI.md#updateappstate) | **Patch** /api/v1/apps/{app_id}/action/{action} | Изменение состояния приложения
*BalancersAPI* | [**AddIPsToBalancer**](docs/BalancersAPI.md#addipstobalancer) | **Post** /api/v1/balancers/{balancer_id}/ips | Добавление IP-адресов к балансировщику
*BalancersAPI* | [**CreateBalancer**](docs/BalancersAPI.md#createbalancer) | **Post** /api/v1/balancers | Создание бaлансировщика
*BalancersAPI* | [**CreateBalancerRule**](docs/BalancersAPI.md#createbalancerrule) | **Post** /api/v1/balancers/{balancer_id}/rules | Создание правила для балансировщика
*BalancersAPI* | [**DeleteBalancer**](docs/BalancersAPI.md#deletebalancer) | **Delete** /api/v1/balancers/{balancer_id} | Удаление балансировщика
*BalancersAPI* | [**DeleteBalancerRule**](docs/BalancersAPI.md#deletebalancerrule) | **Delete** /api/v1/balancers/{balancer_id}/rules/{rule_id} | Удаление правила для балансировщика
*BalancersAPI* | [**DeleteIPsFromBalancer**](docs/BalancersAPI.md#deleteipsfrombalancer) | **Delete** /api/v1/balancers/{balancer_id}/ips | Удаление IP-адресов из балансировщика
*BalancersAPI* | [**GetBalancer**](docs/BalancersAPI.md#getbalancer) | **Get** /api/v1/balancers/{balancer_id} | Получение бaлансировщика
*BalancersAPI* | [**GetBalancerIPs**](docs/BalancersAPI.md#getbalancerips) | **Get** /api/v1/balancers/{balancer_id}/ips | Получение списка IP-адресов балансировщика
*BalancersAPI* | [**GetBalancerRules**](docs/BalancersAPI.md#getbalancerrules) | **Get** /api/v1/balancers/{balancer_id}/rules | Получение правил балансировщика
*BalancersAPI* | [**GetBalancers**](docs/BalancersAPI.md#getbalancers) | **Get** /api/v1/balancers | Получение списка всех бaлансировщиков
*BalancersAPI* | [**GetBalancersPresets**](docs/BalancersAPI.md#getbalancerspresets) | **Get** /api/v1/presets/balancers | Получение списка тарифов для балансировщика
*BalancersAPI* | [**UpdateBalancer**](docs/BalancersAPI.md#updatebalancer) | **Patch** /api/v1/balancers/{balancer_id} | Обновление балансировщика
*BalancersAPI* | [**UpdateBalancerRule**](docs/BalancersAPI.md#updatebalancerrule) | **Patch** /api/v1/balancers/{balancer_id}/rules/{rule_id} | Обновление правила для балансировщика
*DatabasesAPI* | [**CreateDatabase**](docs/DatabasesAPI.md#createdatabase) | **Post** /api/v1/dbs | Создание базы данных
*DatabasesAPI* | [**CreateDatabaseBackup**](docs/DatabasesAPI.md#createdatabasebackup) | **Post** /api/v1/dbs/{db_id}/backups | Создание бэкапа базы данных
*DatabasesAPI* | [**CreateDatabaseCluster**](docs/DatabasesAPI.md#createdatabasecluster) | **Post** /api/v1/databases | Создание кластера базы данных
*DatabasesAPI* | [**CreateDatabaseInstance**](docs/DatabasesAPI.md#createdatabaseinstance) | **Post** /api/v1/databases/{db_cluster_id}/instances | Создание инстанса базы данных
*DatabasesAPI* | [**CreateDatabaseUser**](docs/DatabasesAPI.md#createdatabaseuser) | **Post** /api/v1/databases/{db_cluster_id}/admins | Создание пользователя базы данных
*DatabasesAPI* | [**DeleteDatabase**](docs/DatabasesAPI.md#deletedatabase) | **Delete** /api/v1/dbs/{db_id} | Удаление базы данных
*DatabasesAPI* | [**DeleteDatabaseBackup**](docs/DatabasesAPI.md#deletedatabasebackup) | **Delete** /api/v1/dbs/{db_id}/backups/{backup_id} | Удаление бэкапа базы данных
*DatabasesAPI* | [**DeleteDatabaseCluster**](docs/DatabasesAPI.md#deletedatabasecluster) | **Delete** /api/v1/databases/{db_cluster_id} | Удаление кластера базы данных
*DatabasesAPI* | [**DeleteDatabaseInstance**](docs/DatabasesAPI.md#deletedatabaseinstance) | **Delete** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Удаление инстанса базы данных
*DatabasesAPI* | [**DeleteDatabaseUser**](docs/DatabasesAPI.md#deletedatabaseuser) | **Delete** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Удаление пользователя базы данных
*DatabasesAPI* | [**GetDatabase**](docs/DatabasesAPI.md#getdatabase) | **Get** /api/v1/dbs/{db_id} | Получение базы данных
*DatabasesAPI* | [**GetDatabaseAutoBackupsSettings**](docs/DatabasesAPI.md#getdatabaseautobackupssettings) | **Get** /api/v1/dbs/{db_id}/auto-backups | Получение настроек автобэкапов базы данных
*DatabasesAPI* | [**GetDatabaseBackup**](docs/DatabasesAPI.md#getdatabasebackup) | **Get** /api/v1/dbs/{db_id}/backups/{backup_id} | Получение бэкапа базы данных
*DatabasesAPI* | [**GetDatabaseBackups**](docs/DatabasesAPI.md#getdatabasebackups) | **Get** /api/v1/dbs/{db_id}/backups | Список бэкапов базы данных
*DatabasesAPI* | [**GetDatabaseCluster**](docs/DatabasesAPI.md#getdatabasecluster) | **Get** /api/v1/databases/{db_cluster_id} | Получение кластера базы данных
*DatabasesAPI* | [**GetDatabaseClusterTypes**](docs/DatabasesAPI.md#getdatabaseclustertypes) | **Get** /api/v1/database-types | Получение списка типов кластеров баз данных
*DatabasesAPI* | [**GetDatabaseClusters**](docs/DatabasesAPI.md#getdatabaseclusters) | **Get** /api/v1/databases | Получение списка кластеров баз данных
*DatabasesAPI* | [**GetDatabaseInstance**](docs/DatabasesAPI.md#getdatabaseinstance) | **Get** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Получение инстанса базы данных
*DatabasesAPI* | [**GetDatabaseInstances**](docs/DatabasesAPI.md#getdatabaseinstances) | **Get** /api/v1/databases/{db_cluster_id}/instances | Получение списка инстансов баз данных
*DatabasesAPI* | [**GetDatabaseParameters**](docs/DatabasesAPI.md#getdatabaseparameters) | **Get** /api/v1/dbs/parameters | Получение списка параметров баз данных
*DatabasesAPI* | [**GetDatabaseUser**](docs/DatabasesAPI.md#getdatabaseuser) | **Get** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Получение пользователя базы данных
*DatabasesAPI* | [**GetDatabaseUsers**](docs/DatabasesAPI.md#getdatabaseusers) | **Get** /api/v1/databases/{db_cluster_id}/admins | Получение списка пользователей базы данных
*DatabasesAPI* | [**GetDatabases**](docs/DatabasesAPI.md#getdatabases) | **Get** /api/v1/dbs | Получение списка всех баз данных
*DatabasesAPI* | [**GetDatabasesPresets**](docs/DatabasesAPI.md#getdatabasespresets) | **Get** /api/v1/presets/dbs | Получение списка тарифов для баз данных
*DatabasesAPI* | [**RestoreDatabaseFromBackup**](docs/DatabasesAPI.md#restoredatabasefrombackup) | **Put** /api/v1/dbs/{db_id}/backups/{backup_id} | Восстановление базы данных из бэкапа
*DatabasesAPI* | [**UpdateDatabase**](docs/DatabasesAPI.md#updatedatabase) | **Patch** /api/v1/dbs/{db_id} | Обновление базы данных
*DatabasesAPI* | [**UpdateDatabaseAutoBackupsSettings**](docs/DatabasesAPI.md#updatedatabaseautobackupssettings) | **Patch** /api/v1/dbs/{db_id}/auto-backups | Изменение настроек автобэкапов базы данных
*DatabasesAPI* | [**UpdateDatabaseCluster**](docs/DatabasesAPI.md#updatedatabasecluster) | **Patch** /api/v1/databases/{db_cluster_id} | Изменение кластера базы данных
*DatabasesAPI* | [**UpdateDatabaseInstance**](docs/DatabasesAPI.md#updatedatabaseinstance) | **Patch** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Изменение инстанса базы данных
*DatabasesAPI* | [**UpdateDatabaseUser**](docs/DatabasesAPI.md#updatedatabaseuser) | **Patch** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Изменение пользователя базы данных
*DedicatedServersAPI* | [**CreateDedicatedServer**](docs/DedicatedServersAPI.md#creatededicatedserver) | **Post** /api/v1/dedicated-servers | Создание выделенного сервера
*DedicatedServersAPI* | [**DeleteDedicatedServer**](docs/DedicatedServersAPI.md#deletededicatedserver) | **Delete** /api/v1/dedicated-servers/{dedicated_id} | Удаление выделенного сервера
*DedicatedServersAPI* | [**GetDedicatedServer**](docs/DedicatedServersAPI.md#getdedicatedserver) | **Get** /api/v1/dedicated-servers/{dedicated_id} | Получение выделенного сервера
*DedicatedServersAPI* | [**GetDedicatedServerPresetAdditionalServices**](docs/DedicatedServersAPI.md#getdedicatedserverpresetadditionalservices) | **Get** /api/v1/presets/dedicated-servers/{preset_id}/additional-services | Получение дополнительных услуг для выделенного сервера
*DedicatedServersAPI* | [**GetDedicatedServers**](docs/DedicatedServersAPI.md#getdedicatedservers) | **Get** /api/v1/dedicated-servers | Получение списка выделенных серверов
*DedicatedServersAPI* | [**GetDedicatedServersPresets**](docs/DedicatedServersAPI.md#getdedicatedserverspresets) | **Get** /api/v1/presets/dedicated-servers | Получение списка тарифов для выделенного сервера
*DedicatedServersAPI* | [**UpdateDedicatedServer**](docs/DedicatedServersAPI.md#updatededicatedserver) | **Patch** /api/v1/dedicated-servers/{dedicated_id} | Обновление выделенного сервера
*DomainsAPI* | [**AddDomain**](docs/DomainsAPI.md#adddomain) | **Post** /api/v1/add-domain/{fqdn} | Добавление домена на аккаунт
*DomainsAPI* | [**AddSubdomain**](docs/DomainsAPI.md#addsubdomain) | **Post** /api/v1/domains/{fqdn}/subdomains/{subdomain_fqdn} | Добавление поддомена
*DomainsAPI* | [**CheckDomain**](docs/DomainsAPI.md#checkdomain) | **Get** /api/v1/check-domain/{fqdn} | Проверить, доступен ли домен для регистрации
*DomainsAPI* | [**CreateDomainDNSRecord**](docs/DomainsAPI.md#createdomaindnsrecord) | **Post** /api/v1/domains/{fqdn}/dns-records | Добавить информацию о DNS-записи для домена или поддомена
*DomainsAPI* | [**CreateDomainRequest**](docs/DomainsAPI.md#createdomainrequest) | **Post** /api/v1/domains-requests | Создание заявки на регистрацию/продление/трансфер домена
*DomainsAPI* | [**DeleteDomain**](docs/DomainsAPI.md#deletedomain) | **Delete** /api/v1/domains/{fqdn} | Удаление домена
*DomainsAPI* | [**DeleteDomainDNSRecord**](docs/DomainsAPI.md#deletedomaindnsrecord) | **Delete** /api/v1/domains/{fqdn}/dns-records/{record_id} | Удалить информацию о DNS-записи для домена или поддомена
*DomainsAPI* | [**DeleteSubdomain**](docs/DomainsAPI.md#deletesubdomain) | **Delete** /api/v1/domains/{fqdn}/subdomains/{subdomain_fqdn} | Удаление поддомена
*DomainsAPI* | [**GetDomain**](docs/DomainsAPI.md#getdomain) | **Get** /api/v1/domains/{fqdn} | Получение информации о домене
*DomainsAPI* | [**GetDomainDNSRecords**](docs/DomainsAPI.md#getdomaindnsrecords) | **Get** /api/v1/domains/{fqdn}/dns-records | Получить информацию обо всех пользовательских DNS-записях домена или поддомена
*DomainsAPI* | [**GetDomainDefaultDNSRecords**](docs/DomainsAPI.md#getdomaindefaultdnsrecords) | **Get** /api/v1/domains/{fqdn}/default-dns-records | Получить информацию обо всех DNS-записях по умолчанию домена или поддомена
*DomainsAPI* | [**GetDomainNameServers**](docs/DomainsAPI.md#getdomainnameservers) | **Get** /api/v1/domains/{fqdn}/name-servers | Получение списка name-серверов домена
*DomainsAPI* | [**GetDomainRequest**](docs/DomainsAPI.md#getdomainrequest) | **Get** /api/v1/domains-requests/{request_id} | Получение заявки на регистрацию/продление/трансфер домена
*DomainsAPI* | [**GetDomainRequests**](docs/DomainsAPI.md#getdomainrequests) | **Get** /api/v1/domains-requests | Получение списка заявок на регистрацию/продление/трансфер домена
*DomainsAPI* | [**GetDomains**](docs/DomainsAPI.md#getdomains) | **Get** /api/v1/domains | Получение списка всех доменов
*DomainsAPI* | [**GetTLD**](docs/DomainsAPI.md#gettld) | **Get** /api/v1/tlds/{tld_id} | Получить информацию о доменной зоне по ID
*DomainsAPI* | [**GetTLDs**](docs/DomainsAPI.md#gettlds) | **Get** /api/v1/tlds | Получить информацию о доменных зонах
*DomainsAPI* | [**UpdateDomainAutoProlongation**](docs/DomainsAPI.md#updatedomainautoprolongation) | **Patch** /api/v1/domains/{fqdn} | Включение/выключение автопродления домена
*DomainsAPI* | [**UpdateDomainDNSRecord**](docs/DomainsAPI.md#updatedomaindnsrecord) | **Patch** /api/v1/domains/{fqdn}/dns-records/{record_id} | Обновить информацию о DNS-записи домена или поддомена
*DomainsAPI* | [**UpdateDomainNameServers**](docs/DomainsAPI.md#updatedomainnameservers) | **Put** /api/v1/domains/{fqdn}/name-servers | Изменение name-серверов домена
*DomainsAPI* | [**UpdateDomainRequest**](docs/DomainsAPI.md#updatedomainrequest) | **Patch** /api/v1/domains-requests/{request_id} | Оплата/обновление заявки на регистрацию/продление/трансфер домена
*FirewallAPI* | [**AddResourceToGroup**](docs/FirewallAPI.md#addresourcetogroup) | **Post** /api/v1/firewall/groups/{group_id}/resources/{resource_id} | Линковка ресурса в firewall group
*FirewallAPI* | [**CreateGroup**](docs/FirewallAPI.md#creategroup) | **Post** /api/v1/firewall/groups | Создание группы правил
*FirewallAPI* | [**CreateGroupRule**](docs/FirewallAPI.md#creategrouprule) | **Post** /api/v1/firewall/groups/{group_id}/rules | Создание firewall правила
*FirewallAPI* | [**DeleteGroup**](docs/FirewallAPI.md#deletegroup) | **Delete** /api/v1/firewall/groups/{group_id} | Удаление группы правил
*FirewallAPI* | [**DeleteGroupRule**](docs/FirewallAPI.md#deletegrouprule) | **Delete** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Удаление firewall правила
*FirewallAPI* | [**DeleteResourceFromGroup**](docs/FirewallAPI.md#deleteresourcefromgroup) | **Delete** /api/v1/firewall/groups/{group_id}/resources/{resource_id} | Отлинковка ресурса из firewall group
*FirewallAPI* | [**GetGroup**](docs/FirewallAPI.md#getgroup) | **Get** /api/v1/firewall/groups/{group_id} | Получение информации о группе правил
*FirewallAPI* | [**GetGroupResources**](docs/FirewallAPI.md#getgroupresources) | **Get** /api/v1/firewall/groups/{group_id}/resources | Получение слинкованных ресурсов
*FirewallAPI* | [**GetGroupRule**](docs/FirewallAPI.md#getgrouprule) | **Get** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Получение информации о правиле
*FirewallAPI* | [**GetGroupRules**](docs/FirewallAPI.md#getgrouprules) | **Get** /api/v1/firewall/groups/{group_id}/rules | Получение списка правил
*FirewallAPI* | [**GetGroups**](docs/FirewallAPI.md#getgroups) | **Get** /api/v1/firewall/groups | Получение групп правил
*FirewallAPI* | [**GetRulesForResource**](docs/FirewallAPI.md#getrulesforresource) | **Get** /api/v1/firewall/service/{resource_type}/{resource_id} | Получение групп правил для ресурса
*FirewallAPI* | [**UpdateGroup**](docs/FirewallAPI.md#updategroup) | **Patch** /api/v1/firewall/groups/{group_id} | Обновление группы правил
*FirewallAPI* | [**UpdateGroupRule**](docs/FirewallAPI.md#updategrouprule) | **Patch** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Обновление firewall правила
*FloatingIPAPI* | [**BindFloatingIp**](docs/FloatingIPAPI.md#bindfloatingip) | **Post** /api/v1/floating-ips/{floating_ip_id}/bind | Привязать IP к сервису
*FloatingIPAPI* | [**CreateFloatingIp**](docs/FloatingIPAPI.md#createfloatingip) | **Post** /api/v1/floating-ips | Создание плавающего IP
*FloatingIPAPI* | [**DeleteFloatingIP**](docs/FloatingIPAPI.md#deletefloatingip) | **Delete** /api/v1/floating-ips/{floating_ip_id} | Удаление плавающего IP по ID
*FloatingIPAPI* | [**GetFloatingIp**](docs/FloatingIPAPI.md#getfloatingip) | **Get** /api/v1/floating-ips/{floating_ip_id} | Получение плавающего IP
*FloatingIPAPI* | [**GetFloatingIps**](docs/FloatingIPAPI.md#getfloatingips) | **Get** /api/v1/floating-ips | Получение списка плавающих IP
*FloatingIPAPI* | [**UnbindFloatingIp**](docs/FloatingIPAPI.md#unbindfloatingip) | **Post** /api/v1/floating-ips/{floating_ip_id}/unbind | Отвязать IP от сервиса
*FloatingIPAPI* | [**UpdateFloatingIP**](docs/FloatingIPAPI.md#updatefloatingip) | **Patch** /api/v1/floating-ips/{floating_ip_id} | Изменение плавающего IP по ID
*ImagesAPI* | [**CreateImage**](docs/ImagesAPI.md#createimage) | **Post** /api/v1/images | Создание образа
*ImagesAPI* | [**CreateImageDownloadUrl**](docs/ImagesAPI.md#createimagedownloadurl) | **Post** /api/v1/images/{image_id}/download-url | Создание ссылки на скачивание образа
*ImagesAPI* | [**DeleteImage**](docs/ImagesAPI.md#deleteimage) | **Delete** /api/v1/images/{image_id} | Удаление образа
*ImagesAPI* | [**DeleteImageDownloadURL**](docs/ImagesAPI.md#deleteimagedownloadurl) | **Delete** /api/v1/images/{image_id}/download-url/{image_url_id} | Удаление ссылки на образ
*ImagesAPI* | [**GetImage**](docs/ImagesAPI.md#getimage) | **Get** /api/v1/images/{image_id} | Получение информации о образе
*ImagesAPI* | [**GetImageDownloadURL**](docs/ImagesAPI.md#getimagedownloadurl) | **Get** /api/v1/images/{image_id}/download-url/{image_url_id} | Получение информации о ссылке на скачивание образа
*ImagesAPI* | [**GetImageDownloadURLs**](docs/ImagesAPI.md#getimagedownloadurls) | **Get** /api/v1/images/{image_id}/download-url | Получение информации о ссылках на скачивание образов
*ImagesAPI* | [**GetImages**](docs/ImagesAPI.md#getimages) | **Get** /api/v1/images | Получение списка образов
*ImagesAPI* | [**UpdateImage**](docs/ImagesAPI.md#updateimage) | **Patch** /api/v1/images/{image_id} | Обновление информации о образе
*ImagesAPI* | [**UploadImage**](docs/ImagesAPI.md#uploadimage) | **Post** /api/v1/images/{image_id} | Загрузка образа
*KubernetesAPI* | [**CreateCluster**](docs/KubernetesAPI.md#createcluster) | **Post** /api/v1/k8s/clusters | Создание кластера
*KubernetesAPI* | [**CreateClusterNodeGroup**](docs/KubernetesAPI.md#createclusternodegroup) | **Post** /api/v1/k8s/clusters/{cluster_id}/groups | Создание группы нод
*KubernetesAPI* | [**DeleteCluster**](docs/KubernetesAPI.md#deletecluster) | **Delete** /api/v1/k8s/clusters/{cluster_id} | Удаление кластера
*KubernetesAPI* | [**DeleteClusterNode**](docs/KubernetesAPI.md#deleteclusternode) | **Delete** /api/v1/k8s/clusters/{cluster_id}/nodes/{node_id} | Удаление ноды
*KubernetesAPI* | [**DeleteClusterNodeGroup**](docs/KubernetesAPI.md#deleteclusternodegroup) | **Delete** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id} | Удаление группы нод
*KubernetesAPI* | [**GetCluster**](docs/KubernetesAPI.md#getcluster) | **Get** /api/v1/k8s/clusters/{cluster_id} | Получение информации о кластере
*KubernetesAPI* | [**GetClusterKubeconfig**](docs/KubernetesAPI.md#getclusterkubeconfig) | **Get** /api/v1/k8s/clusters/{cluster_id}/kubeconfig | Получение файла kubeconfig
*KubernetesAPI* | [**GetClusterNodeGroup**](docs/KubernetesAPI.md#getclusternodegroup) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id} | Получение информации о группе нод
*KubernetesAPI* | [**GetClusterNodeGroups**](docs/KubernetesAPI.md#getclusternodegroups) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups | Получение групп нод кластера
*KubernetesAPI* | [**GetClusterNodes**](docs/KubernetesAPI.md#getclusternodes) | **Get** /api/v1/k8s/clusters/{cluster_id}/nodes | Получение списка нод
*KubernetesAPI* | [**GetClusterNodesFromGroup**](docs/KubernetesAPI.md#getclusternodesfromgroup) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Получение списка нод, принадлежащих группе
*KubernetesAPI* | [**GetClusterResources**](docs/KubernetesAPI.md#getclusterresources) | **Get** /api/v1/k8s/clusters/{cluster_id}/resources | Получение ресурсов кластера
*KubernetesAPI* | [**GetClusters**](docs/KubernetesAPI.md#getclusters) | **Get** /api/v1/k8s/clusters | Получение списка кластеров
*KubernetesAPI* | [**GetK8SNetworkDrivers**](docs/KubernetesAPI.md#getk8snetworkdrivers) | **Get** /api/v1/k8s/network-drivers | Получение списка сетевых драйверов k8s
*KubernetesAPI* | [**GetK8SVersions**](docs/KubernetesAPI.md#getk8sversions) | **Get** /api/v1/k8s/k8s-versions | Получение списка версий k8s
*KubernetesAPI* | [**GetKubernetesPresets**](docs/KubernetesAPI.md#getkubernetespresets) | **Get** /api/v1/presets/k8s | Получение списка тарифов
*KubernetesAPI* | [**IncreaseCountOfNodesInGroup**](docs/KubernetesAPI.md#increasecountofnodesingroup) | **Post** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Увеличение количества нод в группе на указанное количество
*KubernetesAPI* | [**ReduceCountOfNodesInGroup**](docs/KubernetesAPI.md#reducecountofnodesingroup) | **Delete** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Уменьшение количества нод в группе на указанное количество
*KubernetesAPI* | [**UpdateCluster**](docs/KubernetesAPI.md#updatecluster) | **Patch** /api/v1/k8s/clusters/{cluster_id} | Обновление информации о кластере
*LocationsAPI* | [**GetLocations**](docs/LocationsAPI.md#getlocations) | **Get** /api/v2/locations | Получение списка локаций
*MailAPI* | [**CreateDomainMailbox**](docs/MailAPI.md#createdomainmailbox) | **Post** /api/v1/mail/domains/{domain} | Создание почтового ящика
*MailAPI* | [**CreateMultipleDomainMailboxes**](docs/MailAPI.md#createmultipledomainmailboxes) | **Post** /api/v1/mail/domains/{domain}/batch | Множественное создание почтовых ящиков
*MailAPI* | [**DeleteMailbox**](docs/MailAPI.md#deletemailbox) | **Delete** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Удаление почтового ящика
*MailAPI* | [**GetDomainMailInfo**](docs/MailAPI.md#getdomainmailinfo) | **Get** /api/v1/mail/domains/{domain}/info | Получение почтовой информации о домене
*MailAPI* | [**GetDomainMailboxes**](docs/MailAPI.md#getdomainmailboxes) | **Get** /api/v1/mail/domains/{domain} | Получение списка почтовых ящиков домена
*MailAPI* | [**GetMailQuota**](docs/MailAPI.md#getmailquota) | **Get** /api/v1/mail/quota | Получение квоты почты аккаунта
*MailAPI* | [**GetMailbox**](docs/MailAPI.md#getmailbox) | **Get** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Получение почтового ящика
*MailAPI* | [**GetMailboxes**](docs/MailAPI.md#getmailboxes) | **Get** /api/v1/mail | Получение списка почтовых ящиков аккаунта
*MailAPI* | [**UpdateDomainMailInfo**](docs/MailAPI.md#updatedomainmailinfo) | **Patch** /api/v1/mail/domains/{domain}/info | Изменение почтовой информации о домене
*MailAPI* | [**UpdateMailQuota**](docs/MailAPI.md#updatemailquota) | **Patch** /api/v1/mail/quota | Изменение квоты почты аккаунта
*MailAPI* | [**UpdateMailbox**](docs/MailAPI.md#updatemailbox) | **Patch** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Изменение почтового ящика
*NetworkDrivesAPI* | [**CreateNetworkDrive**](docs/NetworkDrivesAPI.md#createnetworkdrive) | **Post** /api/v1/network-drives | Создание сетевого диска
*NetworkDrivesAPI* | [**DeleteNetworkDrive**](docs/NetworkDrivesAPI.md#deletenetworkdrive) | **Delete** /api/v1/network-drives/{network_drive_id} | Удаление сетевого диска по идентификатору
*NetworkDrivesAPI* | [**GetNetworkDrive**](docs/NetworkDrivesAPI.md#getnetworkdrive) | **Get** /api/v1/network-drives/{network_drive_id} | Получение сетевого диска
*NetworkDrivesAPI* | [**GetNetworkDrives**](docs/NetworkDrivesAPI.md#getnetworkdrives) | **Get** /api/v1/network-drives | Получение списка cетевых дисков
*NetworkDrivesAPI* | [**GetNetworkDrivesAvailableResources**](docs/NetworkDrivesAPI.md#getnetworkdrivesavailableresources) | **Get** /api/v1/network-drives/available-resources | Получение списка сервисов доступных для подключения диска
*NetworkDrivesAPI* | [**GetNetworkDrivesPresets**](docs/NetworkDrivesAPI.md#getnetworkdrivespresets) | **Get** /api/v1/presets/network-drives | Получение списка доступных тарифов для сетевого диска
*NetworkDrivesAPI* | [**MountNetworkDrive**](docs/NetworkDrivesAPI.md#mountnetworkdrive) | **Post** /api/v1/network-drives/{network_drive_id}/mount | Подключить сетевой диск к сервису
*NetworkDrivesAPI* | [**UnmountNetworkDrive**](docs/NetworkDrivesAPI.md#unmountnetworkdrive) | **Post** /api/v1/network-drives/{network_drive_id}/unmount | Отключить сетевой диск от сервиса
*NetworkDrivesAPI* | [**UpdateNetworkDrive**](docs/NetworkDrivesAPI.md#updatenetworkdrive) | **Patch** /api/v1/network-drives/{network_drive_id} | Изменение сетевого диска по ID
*ProjectsAPI* | [**AddBalancerToProject**](docs/ProjectsAPI.md#addbalancertoproject) | **Post** /api/v1/projects/{project_id}/resources/balancers | Добавление балансировщика в проект
*ProjectsAPI* | [**AddClusterToProject**](docs/ProjectsAPI.md#addclustertoproject) | **Post** /api/v1/projects/{project_id}/resources/clusters | Добавление кластера в проект
*ProjectsAPI* | [**AddDatabaseToProject**](docs/ProjectsAPI.md#adddatabasetoproject) | **Post** /api/v1/projects/{project_id}/resources/databases | Добавление базы данных в проект
*ProjectsAPI* | [**AddDedicatedServerToProject**](docs/ProjectsAPI.md#adddedicatedservertoproject) | **Post** /api/v1/projects/{project_id}/resources/dedicated | Добавление выделенного сервера в проект
*ProjectsAPI* | [**AddServerToProject**](docs/ProjectsAPI.md#addservertoproject) | **Post** /api/v1/projects/{project_id}/resources/servers | Добавление сервера в проект
*ProjectsAPI* | [**AddStorageToProject**](docs/ProjectsAPI.md#addstoragetoproject) | **Post** /api/v1/projects/{project_id}/resources/buckets | Добавление хранилища в проект
*ProjectsAPI* | [**CreateProject**](docs/ProjectsAPI.md#createproject) | **Post** /api/v1/projects | Создание проекта
*ProjectsAPI* | [**DeleteProject**](docs/ProjectsAPI.md#deleteproject) | **Delete** /api/v1/projects/{project_id} | Удаление проекта
*ProjectsAPI* | [**GetAccountBalancers**](docs/ProjectsAPI.md#getaccountbalancers) | **Get** /api/v1/projects/resources/balancers | Получение списка всех балансировщиков на аккаунте
*ProjectsAPI* | [**GetAccountClusters**](docs/ProjectsAPI.md#getaccountclusters) | **Get** /api/v1/projects/resources/clusters | Получение списка всех кластеров на аккаунте
*ProjectsAPI* | [**GetAccountDatabases**](docs/ProjectsAPI.md#getaccountdatabases) | **Get** /api/v1/projects/resources/databases | Получение списка всех баз данных на аккаунте
*ProjectsAPI* | [**GetAccountDedicatedServers**](docs/ProjectsAPI.md#getaccountdedicatedservers) | **Get** /api/v1/projects/resources/dedicated | Получение списка всех выделенных серверов на аккаунте
*ProjectsAPI* | [**GetAccountServers**](docs/ProjectsAPI.md#getaccountservers) | **Get** /api/v1/projects/resources/servers | Получение списка всех серверов на аккаунте
*ProjectsAPI* | [**GetAccountStorages**](docs/ProjectsAPI.md#getaccountstorages) | **Get** /api/v1/projects/resources/buckets | Получение списка всех хранилищ на аккаунте
*ProjectsAPI* | [**GetAllProjectResources**](docs/ProjectsAPI.md#getallprojectresources) | **Get** /api/v1/projects/{project_id}/resources | Получение всех ресурсов проекта
*ProjectsAPI* | [**GetProject**](docs/ProjectsAPI.md#getproject) | **Get** /api/v1/projects/{project_id} | Получение проекта по ID
*ProjectsAPI* | [**GetProjectBalancers**](docs/ProjectsAPI.md#getprojectbalancers) | **Get** /api/v1/projects/{project_id}/resources/balancers | Получение списка балансировщиков проекта
*ProjectsAPI* | [**GetProjectClusters**](docs/ProjectsAPI.md#getprojectclusters) | **Get** /api/v1/projects/{project_id}/resources/clusters | Получение списка кластеров проекта
*ProjectsAPI* | [**GetProjectDatabases**](docs/ProjectsAPI.md#getprojectdatabases) | **Get** /api/v1/projects/{project_id}/resources/databases | Получение списка баз данных проекта
*ProjectsAPI* | [**GetProjectDedicatedServers**](docs/ProjectsAPI.md#getprojectdedicatedservers) | **Get** /api/v1/projects/{project_id}/resources/dedicated | Получение списка выделенных серверов проекта
*ProjectsAPI* | [**GetProjectServers**](docs/ProjectsAPI.md#getprojectservers) | **Get** /api/v1/projects/{project_id}/resources/servers | Получение списка серверов проекта
*ProjectsAPI* | [**GetProjectStorages**](docs/ProjectsAPI.md#getprojectstorages) | **Get** /api/v1/projects/{project_id}/resources/buckets | Получение списка хранилищ проекта
*ProjectsAPI* | [**GetProjects**](docs/ProjectsAPI.md#getprojects) | **Get** /api/v1/projects | Получение списка проектов
*ProjectsAPI* | [**TransferResourceToAnotherProject**](docs/ProjectsAPI.md#transferresourcetoanotherproject) | **Put** /api/v1/projects/{project_id}/resources/transfer | Перенести ресурс в другой проект
*ProjectsAPI* | [**UpdateProject**](docs/ProjectsAPI.md#updateproject) | **Put** /api/v1/projects/{project_id} | Изменение проекта
*S3API* | [**AddStorageSubdomainCertificate**](docs/S3API.md#addstoragesubdomaincertificate) | **Post** /api/v1/storages/certificates/generate | Добавление сертификата для поддомена хранилища
*S3API* | [**AddStorageSubdomains**](docs/S3API.md#addstoragesubdomains) | **Post** /api/v1/storages/buckets/{bucket_id}/subdomains | Добавление поддоменов для хранилища
*S3API* | [**CopyStorageFile**](docs/S3API.md#copystoragefile) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/copy | Копирование файла/директории в хранилище
*S3API* | [**CreateFolderInStorage**](docs/S3API.md#createfolderinstorage) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/mkdir | Создание директории в хранилище
*S3API* | [**CreateStorage**](docs/S3API.md#createstorage) | **Post** /api/v1/storages/buckets | Создание хранилища
*S3API* | [**DeleteStorage**](docs/S3API.md#deletestorage) | **Delete** /api/v1/storages/buckets/{bucket_id} | Удаление хранилища на аккаунте
*S3API* | [**DeleteStorageFile**](docs/S3API.md#deletestoragefile) | **Delete** /api/v1/storages/buckets/{bucket_id}/object-manager/remove | Удаление файла/директории в хранилище
*S3API* | [**DeleteStorageSubdomains**](docs/S3API.md#deletestoragesubdomains) | **Delete** /api/v1/storages/buckets/{bucket_id}/subdomains | Удаление поддоменов хранилища
*S3API* | [**GetStorageFilesList**](docs/S3API.md#getstoragefileslist) | **Get** /api/v1/storages/buckets/{bucket_id}/object-manager/list | Получение списка файлов в хранилище по префиксу
*S3API* | [**GetStorageSubdomains**](docs/S3API.md#getstoragesubdomains) | **Get** /api/v1/storages/buckets/{bucket_id}/subdomains | Получение списка поддоменов хранилища
*S3API* | [**GetStorageTransferStatus**](docs/S3API.md#getstoragetransferstatus) | **Get** /api/v1/storages/buckets/{bucket_id}/transfer-status | Получение статуса переноса хранилища от стороннего S3 в Timeweb Cloud
*S3API* | [**GetStorageUsers**](docs/S3API.md#getstorageusers) | **Get** /api/v1/storages/users | Получение списка пользователей хранилищ аккаунта
*S3API* | [**GetStorages**](docs/S3API.md#getstorages) | **Get** /api/v1/storages/buckets | Получение списка хранилищ аккаунта
*S3API* | [**GetStoragesPresets**](docs/S3API.md#getstoragespresets) | **Get** /api/v1/presets/storages | Получение списка тарифов для хранилищ
*S3API* | [**RenameStorageFile**](docs/S3API.md#renamestoragefile) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/rename | Переименование файла/директории в хранилище
*S3API* | [**TransferStorage**](docs/S3API.md#transferstorage) | **Post** /api/v1/storages/transfer | Перенос хранилища от стороннего провайдера S3 в Timeweb Cloud
*S3API* | [**UpdateStorage**](docs/S3API.md#updatestorage) | **Patch** /api/v1/storages/buckets/{bucket_id} | Изменение хранилища на аккаунте
*S3API* | [**UpdateStorageUser**](docs/S3API.md#updatestorageuser) | **Patch** /api/v1/storages/users/{user_id} | Изменение пароля пользователя-администратора хранилища
*S3API* | [**UploadFileToStorage**](docs/S3API.md#uploadfiletostorage) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/upload | Загрузка файлов в хранилище
*SSHAPI* | [**AddKeyToServer**](docs/SSHAPI.md#addkeytoserver) | **Post** /api/v1/servers/{server_id}/ssh-keys | Добавление SSH-ключей на сервер
*SSHAPI* | [**CreateKey**](docs/SSHAPI.md#createkey) | **Post** /api/v1/ssh-keys | Создание SSH-ключа
*SSHAPI* | [**DeleteKey**](docs/SSHAPI.md#deletekey) | **Delete** /api/v1/ssh-keys/{ssh_key_id} | Удаление SSH-ключа по ID
*SSHAPI* | [**DeleteKeyFromServer**](docs/SSHAPI.md#deletekeyfromserver) | **Delete** /api/v1/servers/{server_id}/ssh-keys/{ssh_key_id} | Удаление SSH-ключей с сервера
*SSHAPI* | [**GetKey**](docs/SSHAPI.md#getkey) | **Get** /api/v1/ssh-keys/{ssh_key_id} | Получение SSH-ключа по ID
*SSHAPI* | [**GetKeys**](docs/SSHAPI.md#getkeys) | **Get** /api/v1/ssh-keys | Получение списка SSH-ключей
*SSHAPI* | [**UpdateKey**](docs/SSHAPI.md#updatekey) | **Patch** /api/v1/ssh-keys/{ssh_key_id} | Изменение SSH-ключа по ID
*ServersAPI* | [**AddServerIP**](docs/ServersAPI.md#addserverip) | **Post** /api/v1/servers/{server_id}/ips | Добавление IP-адреса сервера
*ServersAPI* | [**CloneServer**](docs/ServersAPI.md#cloneserver) | **Post** /api/v1/servers/{server_id}/clone | Клонирование сервера
*ServersAPI* | [**CreateServer**](docs/ServersAPI.md#createserver) | **Post** /api/v1/servers | Создание сервера
*ServersAPI* | [**CreateServerDisk**](docs/ServersAPI.md#createserverdisk) | **Post** /api/v1/servers/{server_id}/disks | Создание диска сервера
*ServersAPI* | [**CreateServerDiskBackup**](docs/ServersAPI.md#createserverdiskbackup) | **Post** /api/v1/servers/{server_id}/disks/{disk_id}/backups | Создание бэкапа диска сервера
*ServersAPI* | [**DeleteServer**](docs/ServersAPI.md#deleteserver) | **Delete** /api/v1/servers/{server_id} | Удаление сервера
*ServersAPI* | [**DeleteServerDisk**](docs/ServersAPI.md#deleteserverdisk) | **Delete** /api/v1/servers/{server_id}/disks/{disk_id} | Удаление диска сервера
*ServersAPI* | [**DeleteServerDiskBackup**](docs/ServersAPI.md#deleteserverdiskbackup) | **Delete** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Удаление бэкапа диска сервера
*ServersAPI* | [**DeleteServerIP**](docs/ServersAPI.md#deleteserverip) | **Delete** /api/v1/servers/{server_id}/ips | Удаление IP-адреса сервера
*ServersAPI* | [**GetConfigurators**](docs/ServersAPI.md#getconfigurators) | **Get** /api/v1/configurator/servers | Получение списка конфигураторов серверов
*ServersAPI* | [**GetOsList**](docs/ServersAPI.md#getoslist) | **Get** /api/v1/os/servers | Получение списка операционных систем
*ServersAPI* | [**GetServer**](docs/ServersAPI.md#getserver) | **Get** /api/v1/servers/{server_id} | Получение сервера
*ServersAPI* | [**GetServerDisk**](docs/ServersAPI.md#getserverdisk) | **Get** /api/v1/servers/{server_id}/disks/{disk_id} | Получение диска сервера
*ServersAPI* | [**GetServerDiskAutoBackupSettings**](docs/ServersAPI.md#getserverdiskautobackupsettings) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/auto-backups | Получить настройки автобэкапов диска сервера
*ServersAPI* | [**GetServerDiskBackup**](docs/ServersAPI.md#getserverdiskbackup) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Получение бэкапа диска сервера
*ServersAPI* | [**GetServerDiskBackups**](docs/ServersAPI.md#getserverdiskbackups) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/backups | Получение списка бэкапов диска сервера
*ServersAPI* | [**GetServerDisks**](docs/ServersAPI.md#getserverdisks) | **Get** /api/v1/servers/{server_id}/disks | Получение списка дисков сервера
*ServersAPI* | [**GetServerIPs**](docs/ServersAPI.md#getserverips) | **Get** /api/v1/servers/{server_id}/ips | Получение списка IP-адресов сервера
*ServersAPI* | [**GetServerLogs**](docs/ServersAPI.md#getserverlogs) | **Get** /api/v1/servers/{server_id}/logs | Получение списка логов сервера
*ServersAPI* | [**GetServerStatistics**](docs/ServersAPI.md#getserverstatistics) | **Get** /api/v1/servers/{server_id}/statistics | Получение статистики сервера
*ServersAPI* | [**GetServers**](docs/ServersAPI.md#getservers) | **Get** /api/v1/servers | Получение списка серверов
*ServersAPI* | [**GetServersPresets**](docs/ServersAPI.md#getserverspresets) | **Get** /api/v1/presets/servers | Получение списка тарифов серверов
*ServersAPI* | [**GetSoftware**](docs/ServersAPI.md#getsoftware) | **Get** /api/v1/software/servers | Получение списка ПО из маркетплейса
*ServersAPI* | [**HardShutdownServer**](docs/ServersAPI.md#hardshutdownserver) | **Post** /api/v1/servers/{server_id}/hard-shutdown | Принудительное выключение сервера
*ServersAPI* | [**ImageUnmountAndServerReload**](docs/ServersAPI.md#imageunmountandserverreload) | **Post** /api/v1/servers/{server_id}/image-unmount | Отмонтирование ISO образа и перезагрузка сервера
*ServersAPI* | [**PerformActionOnBackup**](docs/ServersAPI.md#performactiononbackup) | **Post** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id}/action | Выполнение действия над бэкапом диска сервера
*ServersAPI* | [**PerformActionOnServer**](docs/ServersAPI.md#performactiononserver) | **Post** /api/v1/servers/{server_id}/action | Выполнение действия над сервером
*ServersAPI* | [**RebootServer**](docs/ServersAPI.md#rebootserver) | **Post** /api/v1/servers/{server_id}/reboot | Перезагрузка сервера
*ServersAPI* | [**ResetServerPassword**](docs/ServersAPI.md#resetserverpassword) | **Post** /api/v1/servers/{server_id}/reset-password | Сброс пароля сервера
*ServersAPI* | [**ShutdownServer**](docs/ServersAPI.md#shutdownserver) | **Post** /api/v1/servers/{server_id}/shutdown | Выключение сервера
*ServersAPI* | [**StartServer**](docs/ServersAPI.md#startserver) | **Post** /api/v1/servers/{server_id}/start | Запуск сервера
*ServersAPI* | [**UpdateServer**](docs/ServersAPI.md#updateserver) | **Patch** /api/v1/servers/{server_id} | Изменение сервера
*ServersAPI* | [**UpdateServerDisk**](docs/ServersAPI.md#updateserverdisk) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id} | Изменение параметров диска сервера
*ServersAPI* | [**UpdateServerDiskAutoBackupSettings**](docs/ServersAPI.md#updateserverdiskautobackupsettings) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id}/auto-backups | Изменение настроек автобэкапов диска сервера
*ServersAPI* | [**UpdateServerDiskBackup**](docs/ServersAPI.md#updateserverdiskbackup) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Изменение бэкапа диска сервера
*ServersAPI* | [**UpdateServerIP**](docs/ServersAPI.md#updateserverip) | **Patch** /api/v1/servers/{server_id}/ips | Изменение IP-адреса сервера
*ServersAPI* | [**UpdateServerNAT**](docs/ServersAPI.md#updateservernat) | **Patch** /api/v1/servers/{server_id}/local-networks/nat-mode | Изменение правил маршрутизации трафика сервера (NAT)
*ServersAPI* | [**UpdateServerOSBootMode**](docs/ServersAPI.md#updateserverosbootmode) | **Post** /api/v1/servers/{server_id}/boot-mode | Выбор типа загрузки операционной системы сервера
*VPCAPI* | [**CreateVPC**](docs/VPCAPI.md#createvpc) | **Post** /api/v2/vpcs | Создание VPC
*VPCAPI* | [**DeleteVPC**](docs/VPCAPI.md#deletevpc) | **Delete** /api/v1/vpcs/{vpc_id} | Удаление VPC по ID сети
*VPCAPI* | [**GetVPC**](docs/VPCAPI.md#getvpc) | **Get** /api/v2/vpcs/{vpc_id} | Получение VPC
*VPCAPI* | [**GetVPCPorts**](docs/VPCAPI.md#getvpcports) | **Get** /api/v1/vpcs/{vpc_id}/ports | Получение списка портов для VPC
*VPCAPI* | [**GetVPCServices**](docs/VPCAPI.md#getvpcservices) | **Get** /api/v2/vpcs/{vpc_id}/services | Получение списка сервисов в VPC
*VPCAPI* | [**GetVPCs**](docs/VPCAPI.md#getvpcs) | **Get** /api/v2/vpcs | Получение списка VPCs
*VPCAPI* | [**UpdateVPCs**](docs/VPCAPI.md#updatevpcs) | **Patch** /api/v2/vpcs/{vpc_id} | Изменение VPC по ID сети


## Documentation For Models

 - [AddBalancerToProject200Response](docs/AddBalancerToProject200Response.md)
 - [AddBalancerToProjectRequest](docs/AddBalancerToProjectRequest.md)
 - [AddBitbucket](docs/AddBitbucket.md)
 - [AddClusterToProjectRequest](docs/AddClusterToProjectRequest.md)
 - [AddCountries](docs/AddCountries.md)
 - [AddCountriesToAllowedList201Response](docs/AddCountriesToAllowedList201Response.md)
 - [AddCountriesToAllowedListRequest](docs/AddCountriesToAllowedListRequest.md)
 - [AddDatabaseToProjectRequest](docs/AddDatabaseToProjectRequest.md)
 - [AddDedicatedServerToProjectRequest](docs/AddDedicatedServerToProjectRequest.md)
 - [AddGit](docs/AddGit.md)
 - [AddGithub](docs/AddGithub.md)
 - [AddGitlab](docs/AddGitlab.md)
 - [AddIPsToAllowedList201Response](docs/AddIPsToAllowedList201Response.md)
 - [AddIPsToAllowedListRequest](docs/AddIPsToAllowedListRequest.md)
 - [AddIPsToBalancerRequest](docs/AddIPsToBalancerRequest.md)
 - [AddIps](docs/AddIps.md)
 - [AddKeyToServerRequest](docs/AddKeyToServerRequest.md)
 - [AddProvider201Response](docs/AddProvider201Response.md)
 - [AddServerIP201Response](docs/AddServerIP201Response.md)
 - [AddServerIPRequest](docs/AddServerIPRequest.md)
 - [AddServerToProjectRequest](docs/AddServerToProjectRequest.md)
 - [AddStorageSubdomainCertificateRequest](docs/AddStorageSubdomainCertificateRequest.md)
 - [AddStorageSubdomains200Response](docs/AddStorageSubdomains200Response.md)
 - [AddStorageSubdomainsRequest](docs/AddStorageSubdomainsRequest.md)
 - [AddStorageToProjectRequest](docs/AddStorageToProjectRequest.md)
 - [AddSubdomain201Response](docs/AddSubdomain201Response.md)
 - [AddedSubdomain](docs/AddedSubdomain.md)
 - [ApiKey](docs/ApiKey.md)
 - [App](docs/App.md)
 - [AppConfiguration](docs/AppConfiguration.md)
 - [AppDiskStatus](docs/AppDiskStatus.md)
 - [AppDomainsInner](docs/AppDomainsInner.md)
 - [AppProvider](docs/AppProvider.md)
 - [AppsPresets](docs/AppsPresets.md)
 - [AppsPresetsBackendPresetsInner](docs/AppsPresetsBackendPresetsInner.md)
 - [AppsPresetsFrontendPresetsInner](docs/AppsPresetsFrontendPresetsInner.md)
 - [AutoBackup](docs/AutoBackup.md)
 - [AutoReplyIsDisabled](docs/AutoReplyIsDisabled.md)
 - [AutoReplyIsEnabled](docs/AutoReplyIsEnabled.md)
 - [AvailabilityZone](docs/AvailabilityZone.md)
 - [AvailableFrameworks](docs/AvailableFrameworks.md)
 - [AvailableFrameworksBackendFrameworksInner](docs/AvailableFrameworksBackendFrameworksInner.md)
 - [AvailableFrameworksFrontendFrameworksInner](docs/AvailableFrameworksFrontendFrameworksInner.md)
 - [Backup](docs/Backup.md)
 - [Balancer](docs/Balancer.md)
 - [BaseError](docs/BaseError.md)
 - [BindFloatingIp](docs/BindFloatingIp.md)
 - [Bonus](docs/Bonus.md)
 - [Branch](docs/Branch.md)
 - [Bucket](docs/Bucket.md)
 - [BucketDiskStats](docs/BucketDiskStats.md)
 - [BucketUser](docs/BucketUser.md)
 - [CheckDomain200Response](docs/CheckDomain200Response.md)
 - [ClusterEdit](docs/ClusterEdit.md)
 - [ClusterIn](docs/ClusterIn.md)
 - [ClusterOut](docs/ClusterOut.md)
 - [ClusterResponse](docs/ClusterResponse.md)
 - [Clusterk8s](docs/Clusterk8s.md)
 - [ClustersResponse](docs/ClustersResponse.md)
 - [Commit](docs/Commit.md)
 - [ComponentsSchemasBaseError](docs/ComponentsSchemasBaseError.md)
 - [ConfigParameters](docs/ConfigParameters.md)
 - [CopyStorageFileRequest](docs/CopyStorageFileRequest.md)
 - [CreateAdmin](docs/CreateAdmin.md)
 - [CreateApiKey](docs/CreateApiKey.md)
 - [CreateApp](docs/CreateApp.md)
 - [CreateApp201Response](docs/CreateApp201Response.md)
 - [CreateBalancer](docs/CreateBalancer.md)
 - [CreateBalancer200Response](docs/CreateBalancer200Response.md)
 - [CreateBalancerRule200Response](docs/CreateBalancerRule200Response.md)
 - [CreateCluster](docs/CreateCluster.md)
 - [CreateClusterAdmin](docs/CreateClusterAdmin.md)
 - [CreateClusterInstance](docs/CreateClusterInstance.md)
 - [CreateDatabase201Response](docs/CreateDatabase201Response.md)
 - [CreateDatabaseBackup201Response](docs/CreateDatabaseBackup201Response.md)
 - [CreateDatabaseBackup409Response](docs/CreateDatabaseBackup409Response.md)
 - [CreateDatabaseCluster201Response](docs/CreateDatabaseCluster201Response.md)
 - [CreateDatabaseInstance201Response](docs/CreateDatabaseInstance201Response.md)
 - [CreateDatabaseUser201Response](docs/CreateDatabaseUser201Response.md)
 - [CreateDb](docs/CreateDb.md)
 - [CreateDbAutoBackups](docs/CreateDbAutoBackups.md)
 - [CreateDedicatedServer](docs/CreateDedicatedServer.md)
 - [CreateDedicatedServer201Response](docs/CreateDedicatedServer201Response.md)
 - [CreateDeploy201Response](docs/CreateDeploy201Response.md)
 - [CreateDeployRequest](docs/CreateDeployRequest.md)
 - [CreateDns](docs/CreateDns.md)
 - [CreateDomainDNSRecord201Response](docs/CreateDomainDNSRecord201Response.md)
 - [CreateDomainMailbox201Response](docs/CreateDomainMailbox201Response.md)
 - [CreateDomainMailboxRequest](docs/CreateDomainMailboxRequest.md)
 - [CreateDomainRequest201Response](docs/CreateDomainRequest201Response.md)
 - [CreateFloatingIp](docs/CreateFloatingIp.md)
 - [CreateFloatingIp201Response](docs/CreateFloatingIp201Response.md)
 - [CreateFolderInStorageRequest](docs/CreateFolderInStorageRequest.md)
 - [CreateInstance](docs/CreateInstance.md)
 - [CreateKey201Response](docs/CreateKey201Response.md)
 - [CreateKeyRequest](docs/CreateKeyRequest.md)
 - [CreateMultipleDomainMailboxes201Response](docs/CreateMultipleDomainMailboxes201Response.md)
 - [CreateMultipleDomainMailboxesRequest](docs/CreateMultipleDomainMailboxesRequest.md)
 - [CreateMultipleDomainMailboxesRequestMailboxesInner](docs/CreateMultipleDomainMailboxesRequestMailboxesInner.md)
 - [CreateNetworkDrive](docs/CreateNetworkDrive.md)
 - [CreateNetworkDrive201Response](docs/CreateNetworkDrive201Response.md)
 - [CreateProject](docs/CreateProject.md)
 - [CreateProject201Response](docs/CreateProject201Response.md)
 - [CreateRule](docs/CreateRule.md)
 - [CreateServer](docs/CreateServer.md)
 - [CreateServer201Response](docs/CreateServer201Response.md)
 - [CreateServerConfiguration](docs/CreateServerConfiguration.md)
 - [CreateServerDisk201Response](docs/CreateServerDisk201Response.md)
 - [CreateServerDiskBackup201Response](docs/CreateServerDiskBackup201Response.md)
 - [CreateServerDiskBackupRequest](docs/CreateServerDiskBackupRequest.md)
 - [CreateServerDiskRequest](docs/CreateServerDiskRequest.md)
 - [CreateStorage201Response](docs/CreateStorage201Response.md)
 - [CreateStorageRequest](docs/CreateStorageRequest.md)
 - [CreateToken201Response](docs/CreateToken201Response.md)
 - [CreateVPC201Response](docs/CreateVPC201Response.md)
 - [CreateVpc](docs/CreateVpc.md)
 - [CreatedApiKey](docs/CreatedApiKey.md)
 - [DatabaseAdmin](docs/DatabaseAdmin.md)
 - [DatabaseAdminInstancesInner](docs/DatabaseAdminInstancesInner.md)
 - [DatabaseCluster](docs/DatabaseCluster.md)
 - [DatabaseClusterDiskStats](docs/DatabaseClusterDiskStats.md)
 - [DatabaseClusterNetworksInner](docs/DatabaseClusterNetworksInner.md)
 - [DatabaseClusterNetworksInnerIpsInner](docs/DatabaseClusterNetworksInnerIpsInner.md)
 - [DatabaseInstance](docs/DatabaseInstance.md)
 - [DatabaseType](docs/DatabaseType.md)
 - [DatabaseTypeRequirements](docs/DatabaseTypeRequirements.md)
 - [Db](docs/Db.md)
 - [DbDiskStats](docs/DbDiskStats.md)
 - [DbType](docs/DbType.md)
 - [DedicatedServer](docs/DedicatedServer.md)
 - [DedicatedServerAdditionalService](docs/DedicatedServerAdditionalService.md)
 - [DedicatedServerPreset](docs/DedicatedServerPreset.md)
 - [DedicatedServerPresetCpu](docs/DedicatedServerPresetCpu.md)
 - [DedicatedServerPresetDisk](docs/DedicatedServerPresetDisk.md)
 - [DedicatedServerPresetMemory](docs/DedicatedServerPresetMemory.md)
 - [DeleteBalancer200Response](docs/DeleteBalancer200Response.md)
 - [DeleteCluster200Response](docs/DeleteCluster200Response.md)
 - [DeleteCountriesFromAllowedList200Response](docs/DeleteCountriesFromAllowedList200Response.md)
 - [DeleteCountriesFromAllowedListRequest](docs/DeleteCountriesFromAllowedListRequest.md)
 - [DeleteDatabase200Response](docs/DeleteDatabase200Response.md)
 - [DeleteDatabaseCluster200Response](docs/DeleteDatabaseCluster200Response.md)
 - [DeleteIPsFromAllowedList200Response](docs/DeleteIPsFromAllowedList200Response.md)
 - [DeleteIPsFromAllowedListRequest](docs/DeleteIPsFromAllowedListRequest.md)
 - [DeleteServer200Response](docs/DeleteServer200Response.md)
 - [DeleteServerIPRequest](docs/DeleteServerIPRequest.md)
 - [DeleteServiceResponse](docs/DeleteServiceResponse.md)
 - [DeleteStorage200Response](docs/DeleteStorage200Response.md)
 - [DeleteStorageFileRequest](docs/DeleteStorageFileRequest.md)
 - [Deploy](docs/Deploy.md)
 - [DeploySettingsInner](docs/DeploySettingsInner.md)
 - [DeployStatus](docs/DeployStatus.md)
 - [DnsRecord](docs/DnsRecord.md)
 - [DnsRecordData](docs/DnsRecordData.md)
 - [Domain](docs/Domain.md)
 - [DomainAllowedBuyPeriodsInner](docs/DomainAllowedBuyPeriodsInner.md)
 - [DomainInfo](docs/DomainInfo.md)
 - [DomainNameServer](docs/DomainNameServer.md)
 - [DomainNameServerItemsInner](docs/DomainNameServerItemsInner.md)
 - [DomainPaymentPeriod](docs/DomainPaymentPeriod.md)
 - [DomainPrimeType](docs/DomainPrimeType.md)
 - [DomainProlong](docs/DomainProlong.md)
 - [DomainRegister](docs/DomainRegister.md)
 - [DomainRequest](docs/DomainRequest.md)
 - [DomainTransfer](docs/DomainTransfer.md)
 - [EditApiKey](docs/EditApiKey.md)
 - [Finances](docs/Finances.md)
 - [FirewallGroup](docs/FirewallGroup.md)
 - [FirewallGroupInAPI](docs/FirewallGroupInAPI.md)
 - [FirewallGroupOutResponse](docs/FirewallGroupOutResponse.md)
 - [FirewallGroupResource](docs/FirewallGroupResource.md)
 - [FirewallGroupResourceOutResponse](docs/FirewallGroupResourceOutResponse.md)
 - [FirewallGroupResourcesOutResponse](docs/FirewallGroupResourcesOutResponse.md)
 - [FirewallGroupsOutResponse](docs/FirewallGroupsOutResponse.md)
 - [FirewallRule](docs/FirewallRule.md)
 - [FirewallRuleDirection](docs/FirewallRuleDirection.md)
 - [FirewallRuleInAPI](docs/FirewallRuleInAPI.md)
 - [FirewallRuleOutResponse](docs/FirewallRuleOutResponse.md)
 - [FirewallRuleProtocol](docs/FirewallRuleProtocol.md)
 - [FirewallRulesOutResponse](docs/FirewallRulesOutResponse.md)
 - [FloatingIp](docs/FloatingIp.md)
 - [ForwardingIncomingIsDisabled](docs/ForwardingIncomingIsDisabled.md)
 - [ForwardingIncomingIsEnabled](docs/ForwardingIncomingIsEnabled.md)
 - [ForwardingOutgoingIsDisabled](docs/ForwardingOutgoingIsDisabled.md)
 - [ForwardingOutgoingIsEnabled](docs/ForwardingOutgoingIsEnabled.md)
 - [Frameworks](docs/Frameworks.md)
 - [Free](docs/Free.md)
 - [GetAccountStatus200Response](docs/GetAccountStatus200Response.md)
 - [GetAllProjectResources200Response](docs/GetAllProjectResources200Response.md)
 - [GetAppDeploys200Response](docs/GetAppDeploys200Response.md)
 - [GetAppLogs200Response](docs/GetAppLogs200Response.md)
 - [GetApps200Response](docs/GetApps200Response.md)
 - [GetAuthAccessSettings200Response](docs/GetAuthAccessSettings200Response.md)
 - [GetAuthAccessSettings200ResponseWhiteList](docs/GetAuthAccessSettings200ResponseWhiteList.md)
 - [GetBalancerIPs200Response](docs/GetBalancerIPs200Response.md)
 - [GetBalancerRules200Response](docs/GetBalancerRules200Response.md)
 - [GetBalancers200Response](docs/GetBalancers200Response.md)
 - [GetBalancersPresets200Response](docs/GetBalancersPresets200Response.md)
 - [GetBranches200Response](docs/GetBranches200Response.md)
 - [GetCommits200Response](docs/GetCommits200Response.md)
 - [GetConfigurators200Response](docs/GetConfigurators200Response.md)
 - [GetCountries200Response](docs/GetCountries200Response.md)
 - [GetDatabaseAutoBackupsSettings200Response](docs/GetDatabaseAutoBackupsSettings200Response.md)
 - [GetDatabaseBackups200Response](docs/GetDatabaseBackups200Response.md)
 - [GetDatabaseClusterTypes200Response](docs/GetDatabaseClusterTypes200Response.md)
 - [GetDatabaseClusters200Response](docs/GetDatabaseClusters200Response.md)
 - [GetDatabaseInstances200Response](docs/GetDatabaseInstances200Response.md)
 - [GetDatabaseUsers200Response](docs/GetDatabaseUsers200Response.md)
 - [GetDatabases200Response](docs/GetDatabases200Response.md)
 - [GetDatabasesPresets200Response](docs/GetDatabasesPresets200Response.md)
 - [GetDedicatedServerPresetAdditionalServices200Response](docs/GetDedicatedServerPresetAdditionalServices200Response.md)
 - [GetDedicatedServers200Response](docs/GetDedicatedServers200Response.md)
 - [GetDedicatedServersPresets200Response](docs/GetDedicatedServersPresets200Response.md)
 - [GetDeployLogs200Response](docs/GetDeployLogs200Response.md)
 - [GetDeploySettings200Response](docs/GetDeploySettings200Response.md)
 - [GetDomain200Response](docs/GetDomain200Response.md)
 - [GetDomainDNSRecords200Response](docs/GetDomainDNSRecords200Response.md)
 - [GetDomainMailInfo200Response](docs/GetDomainMailInfo200Response.md)
 - [GetDomainNameServers200Response](docs/GetDomainNameServers200Response.md)
 - [GetDomainRequests200Response](docs/GetDomainRequests200Response.md)
 - [GetDomains200Response](docs/GetDomains200Response.md)
 - [GetFinances200Response](docs/GetFinances200Response.md)
 - [GetFinances400Response](docs/GetFinances400Response.md)
 - [GetFinances401Response](docs/GetFinances401Response.md)
 - [GetFinances403Response](docs/GetFinances403Response.md)
 - [GetFinances429Response](docs/GetFinances429Response.md)
 - [GetFinances500Response](docs/GetFinances500Response.md)
 - [GetFloatingIps200Response](docs/GetFloatingIps200Response.md)
 - [GetImage404Response](docs/GetImage404Response.md)
 - [GetKey200Response](docs/GetKey200Response.md)
 - [GetKeys200Response](docs/GetKeys200Response.md)
 - [GetLocations200Response](docs/GetLocations200Response.md)
 - [GetMailQuota200Response](docs/GetMailQuota200Response.md)
 - [GetMailboxes200Response](docs/GetMailboxes200Response.md)
 - [GetNetworkDrives200Response](docs/GetNetworkDrives200Response.md)
 - [GetNetworkDrivesAvailableResources200Response](docs/GetNetworkDrivesAvailableResources200Response.md)
 - [GetNetworkDrivesPresets200Response](docs/GetNetworkDrivesPresets200Response.md)
 - [GetNotificationSettings200Response](docs/GetNotificationSettings200Response.md)
 - [GetOsList200Response](docs/GetOsList200Response.md)
 - [GetProjectBalancers200Response](docs/GetProjectBalancers200Response.md)
 - [GetProjectClusters200Response](docs/GetProjectClusters200Response.md)
 - [GetProjectDatabases200Response](docs/GetProjectDatabases200Response.md)
 - [GetProjectDedicatedServers200Response](docs/GetProjectDedicatedServers200Response.md)
 - [GetProjectServers200Response](docs/GetProjectServers200Response.md)
 - [GetProjectStorages200Response](docs/GetProjectStorages200Response.md)
 - [GetProjects200Response](docs/GetProjects200Response.md)
 - [GetProviders200Response](docs/GetProviders200Response.md)
 - [GetRepositories200Response](docs/GetRepositories200Response.md)
 - [GetServerDiskAutoBackupSettings200Response](docs/GetServerDiskAutoBackupSettings200Response.md)
 - [GetServerDiskBackup200Response](docs/GetServerDiskBackup200Response.md)
 - [GetServerDiskBackups200Response](docs/GetServerDiskBackups200Response.md)
 - [GetServerDisks200Response](docs/GetServerDisks200Response.md)
 - [GetServerIPs200Response](docs/GetServerIPs200Response.md)
 - [GetServerLogs200Response](docs/GetServerLogs200Response.md)
 - [GetServerStatistics200Response](docs/GetServerStatistics200Response.md)
 - [GetServerStatistics200ResponseCpuInner](docs/GetServerStatistics200ResponseCpuInner.md)
 - [GetServerStatistics200ResponseDiskInner](docs/GetServerStatistics200ResponseDiskInner.md)
 - [GetServerStatistics200ResponseNetworkTrafficInner](docs/GetServerStatistics200ResponseNetworkTrafficInner.md)
 - [GetServerStatistics200ResponseRamInner](docs/GetServerStatistics200ResponseRamInner.md)
 - [GetServers200Response](docs/GetServers200Response.md)
 - [GetServersPresets200Response](docs/GetServersPresets200Response.md)
 - [GetSoftware200Response](docs/GetSoftware200Response.md)
 - [GetStorageFilesList200Response](docs/GetStorageFilesList200Response.md)
 - [GetStorageSubdomains200Response](docs/GetStorageSubdomains200Response.md)
 - [GetStorageTransferStatus200Response](docs/GetStorageTransferStatus200Response.md)
 - [GetStorageUsers200Response](docs/GetStorageUsers200Response.md)
 - [GetStoragesPresets200Response](docs/GetStoragesPresets200Response.md)
 - [GetTLD200Response](docs/GetTLD200Response.md)
 - [GetTLDs200Response](docs/GetTLDs200Response.md)
 - [GetTokens200Response](docs/GetTokens200Response.md)
 - [GetVPCPorts200Response](docs/GetVPCPorts200Response.md)
 - [GetVPCServices200Response](docs/GetVPCServices200Response.md)
 - [GetVPCs200Response](docs/GetVPCs200Response.md)
 - [Image](docs/Image.md)
 - [ImageDownload](docs/ImageDownload.md)
 - [ImageDownloadResponse](docs/ImageDownloadResponse.md)
 - [ImageDownloadsResponse](docs/ImageDownloadsResponse.md)
 - [ImageInAPI](docs/ImageInAPI.md)
 - [ImageOutResponse](docs/ImageOutResponse.md)
 - [ImageStatus](docs/ImageStatus.md)
 - [ImageUpdateAPI](docs/ImageUpdateAPI.md)
 - [ImageUrlAuth](docs/ImageUrlAuth.md)
 - [ImageUrlIn](docs/ImageUrlIn.md)
 - [ImagesOutResponse](docs/ImagesOutResponse.md)
 - [Invoice](docs/Invoice.md)
 - [K8SVersionsResponse](docs/K8SVersionsResponse.md)
 - [Location](docs/Location.md)
 - [LocationDto](docs/LocationDto.md)
 - [Mailbox](docs/Mailbox.md)
 - [MailboxAutoReply](docs/MailboxAutoReply.md)
 - [MailboxForwardingIncoming](docs/MailboxForwardingIncoming.md)
 - [MailboxForwardingOutgoing](docs/MailboxForwardingOutgoing.md)
 - [MailboxSpamFilter](docs/MailboxSpamFilter.md)
 - [MasterPresetOutApi](docs/MasterPresetOutApi.md)
 - [Meta](docs/Meta.md)
 - [MountNetworkDrive](docs/MountNetworkDrive.md)
 - [Network](docs/Network.md)
 - [NetworkDrive](docs/NetworkDrive.md)
 - [NetworkDriveAvailableResource](docs/NetworkDriveAvailableResource.md)
 - [NetworkDrivePreset](docs/NetworkDrivePreset.md)
 - [NetworkDrivePresetRead](docs/NetworkDrivePresetRead.md)
 - [NetworkDrivePresetWrite](docs/NetworkDrivePresetWrite.md)
 - [NetworkDriveServiceListInner](docs/NetworkDriveServiceListInner.md)
 - [NetworkDriversResponse](docs/NetworkDriversResponse.md)
 - [NodeCount](docs/NodeCount.md)
 - [NodeGroupIn](docs/NodeGroupIn.md)
 - [NodeGroupInConfiguration](docs/NodeGroupInConfiguration.md)
 - [NodeGroupOut](docs/NodeGroupOut.md)
 - [NodeGroupResponse](docs/NodeGroupResponse.md)
 - [NodeGroupsResponse](docs/NodeGroupsResponse.md)
 - [NodeOut](docs/NodeOut.md)
 - [NodesResponse](docs/NodesResponse.md)
 - [NotificationSetting](docs/NotificationSetting.md)
 - [NotificationSettingChannel](docs/NotificationSettingChannel.md)
 - [NotificationSettingChannels](docs/NotificationSettingChannels.md)
 - [NotificationSettingType](docs/NotificationSettingType.md)
 - [OS](docs/OS.md)
 - [PerformActionOnBackupRequest](docs/PerformActionOnBackupRequest.md)
 - [PerformActionOnServerRequest](docs/PerformActionOnServerRequest.md)
 - [Policy](docs/Policy.md)
 - [PresetsBalancer](docs/PresetsBalancer.md)
 - [PresetsDbs](docs/PresetsDbs.md)
 - [PresetsResponse](docs/PresetsResponse.md)
 - [PresetsStorage](docs/PresetsStorage.md)
 - [Project](docs/Project.md)
 - [ProjectResource](docs/ProjectResource.md)
 - [Provider](docs/Provider.md)
 - [Providers](docs/Providers.md)
 - [Quota](docs/Quota.md)
 - [RefreshApiKey](docs/RefreshApiKey.md)
 - [RemoveCountries](docs/RemoveCountries.md)
 - [RemoveIps](docs/RemoveIps.md)
 - [RenameStorageFileRequest](docs/RenameStorageFileRequest.md)
 - [Repository](docs/Repository.md)
 - [Resource](docs/Resource.md)
 - [ResourceTransfer](docs/ResourceTransfer.md)
 - [ResourceType](docs/ResourceType.md)
 - [Resources](docs/Resources.md)
 - [ResourcesResponse](docs/ResourcesResponse.md)
 - [Rule](docs/Rule.md)
 - [S3Object](docs/S3Object.md)
 - [S3ObjectOwner](docs/S3ObjectOwner.md)
 - [S3Subdomain](docs/S3Subdomain.md)
 - [SchemasBaseError](docs/SchemasBaseError.md)
 - [SchemasMeta](docs/SchemasMeta.md)
 - [ServerBackup](docs/ServerBackup.md)
 - [ServerDisk](docs/ServerDisk.md)
 - [ServerIp](docs/ServerIp.md)
 - [ServerLog](docs/ServerLog.md)
 - [ServersConfigurator](docs/ServersConfigurator.md)
 - [ServersConfiguratorRequirements](docs/ServersConfiguratorRequirements.md)
 - [ServersOs](docs/ServersOs.md)
 - [ServersOsRequirements](docs/ServersOsRequirements.md)
 - [ServersPreset](docs/ServersPreset.md)
 - [ServersSoftware](docs/ServersSoftware.md)
 - [ServersSoftwareRequirements](docs/ServersSoftwareRequirements.md)
 - [SetLabels](docs/SetLabels.md)
 - [SettingCondition](docs/SettingCondition.md)
 - [SpamFilterIsDisabled](docs/SpamFilterIsDisabled.md)
 - [SpamFilterIsEnabled](docs/SpamFilterIsEnabled.md)
 - [SshKey](docs/SshKey.md)
 - [SshKeyUsedByInner](docs/SshKeyUsedByInner.md)
 - [Status](docs/Status.md)
 - [StatusCompanyInfo](docs/StatusCompanyInfo.md)
 - [Subdomain](docs/Subdomain.md)
 - [TopLevelDomain](docs/TopLevelDomain.md)
 - [TopLevelDomainAllowedBuyPeriodsInner](docs/TopLevelDomainAllowedBuyPeriodsInner.md)
 - [TransferStatus](docs/TransferStatus.md)
 - [TransferStatusErrorsInner](docs/TransferStatusErrorsInner.md)
 - [TransferStorageRequest](docs/TransferStorageRequest.md)
 - [URLType](docs/URLType.md)
 - [UpdateAdmin](docs/UpdateAdmin.md)
 - [UpdateAppSettings200Response](docs/UpdateAppSettings200Response.md)
 - [UpdateAuthRestrictionsByCountriesRequest](docs/UpdateAuthRestrictionsByCountriesRequest.md)
 - [UpdateBalancer](docs/UpdateBalancer.md)
 - [UpdateCluster](docs/UpdateCluster.md)
 - [UpdateDb](docs/UpdateDb.md)
 - [UpdateDedicatedServerRequest](docs/UpdateDedicatedServerRequest.md)
 - [UpdateDomain](docs/UpdateDomain.md)
 - [UpdateDomainAutoProlongation200Response](docs/UpdateDomainAutoProlongation200Response.md)
 - [UpdateDomainMailInfoRequest](docs/UpdateDomainMailInfoRequest.md)
 - [UpdateDomainNameServers](docs/UpdateDomainNameServers.md)
 - [UpdateDomainNameServersNameServersInner](docs/UpdateDomainNameServersNameServersInner.md)
 - [UpdateFloatingIp](docs/UpdateFloatingIp.md)
 - [UpdateInstance](docs/UpdateInstance.md)
 - [UpdateKeyRequest](docs/UpdateKeyRequest.md)
 - [UpdateMailQuotaRequest](docs/UpdateMailQuotaRequest.md)
 - [UpdateMailbox](docs/UpdateMailbox.md)
 - [UpdateNetworkDrive](docs/UpdateNetworkDrive.md)
 - [UpdateNotificationSettingsRequest](docs/UpdateNotificationSettingsRequest.md)
 - [UpdateNotificationSettingsRequestSettingsInner](docs/UpdateNotificationSettingsRequestSettingsInner.md)
 - [UpdateNotificationSettingsRequestSettingsInnerChannels](docs/UpdateNotificationSettingsRequestSettingsInnerChannels.md)
 - [UpdateProject](docs/UpdateProject.md)
 - [UpdateRule](docs/UpdateRule.md)
 - [UpdateServer](docs/UpdateServer.md)
 - [UpdateServerConfigurator](docs/UpdateServerConfigurator.md)
 - [UpdateServerDiskBackupRequest](docs/UpdateServerDiskBackupRequest.md)
 - [UpdateServerDiskRequest](docs/UpdateServerDiskRequest.md)
 - [UpdateServerIPRequest](docs/UpdateServerIPRequest.md)
 - [UpdateServerNATRequest](docs/UpdateServerNATRequest.md)
 - [UpdateServerOSBootModeRequest](docs/UpdateServerOSBootModeRequest.md)
 - [UpdateStorageRequest](docs/UpdateStorageRequest.md)
 - [UpdateStorageUser200Response](docs/UpdateStorageUser200Response.md)
 - [UpdateStorageUserRequest](docs/UpdateStorageUserRequest.md)
 - [UpdateToken200Response](docs/UpdateToken200Response.md)
 - [UpdateVpc](docs/UpdateVpc.md)
 - [UpdeteSettings](docs/UpdeteSettings.md)
 - [UploadSuccessful](docs/UploadSuccessful.md)
 - [UploadSuccessfulResponse](docs/UploadSuccessfulResponse.md)
 - [UrlStatus](docs/UrlStatus.md)
 - [Use](docs/Use.md)
 - [Vds](docs/Vds.md)
 - [VdsDisksInner](docs/VdsDisksInner.md)
 - [VdsImage](docs/VdsImage.md)
 - [VdsNetworksInner](docs/VdsNetworksInner.md)
 - [VdsNetworksInnerIpsInner](docs/VdsNetworksInnerIpsInner.md)
 - [VdsOs](docs/VdsOs.md)
 - [VdsSoftware](docs/VdsSoftware.md)
 - [Vpc](docs/Vpc.md)
 - [VpcPort](docs/VpcPort.md)
 - [VpcPortService](docs/VpcPortService.md)
 - [VpcService](docs/VpcService.md)
 - [WorkerPresetOutApi](docs/WorkerPresetOutApi.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@timeweb.cloud

