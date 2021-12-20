# GoKalkan [WIP]

<!-- <img src="assets/logo.png" width="200px" align='right'/> -->

⭐ Star on GitHub — it motivates me a lot!

## Описание

GoKalkan - это библиотека-обертка над KalkanCrypt для Golang.

### KalkanCrypt

KalkanCrypt - это набор библиотек для шифрования, дешифрования данных.

Основные методы KalkanCrypt реализованы в `libkalkancryptwr-64`. Это файл доступными методами 
для подписания файлов, текста используя ЭЦП. Подробнее про PKI можно почитать [здесь](wiki/README.md).

## Доступный функционал

```go
// Kalkan - интерфейс с методами KalkanCrypt
type Kalkan interface {
	Init() error
	LoadKeyStore(password, containerPath string) error
	SignXML(data string) (string, error)
	VerifyXML(xml string) (string, error)
	VerifyData(data string) (*VerifiedData, error)
	X509ExportCertificateFromStore() (string, error)
	GetLastErrorString() string
	Close() error
}
```

Не все доступные методы пока были реализованы. Для знакомства со всеми функциями перейти [сюда](cpp/KalkanCrypt.h).

## Запуск

Чтобы использовать библиотеку требуется провести подготовку:

#### 1. Обратиться в [pki.gov.kz](https://pki.gov.kz/developers/) чтобы получить SDK

SDK представляет собой набор библиотек для Java и C.

#### 2. Установить в доверенные сертификаты

Сертификаты будут лежать по пути `SDK/C/Linux/ca-certs/Ubuntu`. Будут два типа сертфикатов - `production` и `test`.

В папке будут скрипты для установки сертификатов, понадобится sudo права.

#### 3. Скопировать `libkalkancryptwr-64.so` и `libkalkancryptwr-64.so.1.1.0` в /usr/lib/

Файлы лежат в директории `SDK/C/Linux/C`. Команда для копирования:

```sh
sudo cp -f libkalkancryptwr-64.so libkalkancryptwr-64.so.1.1.0 /usr/lib/
```

#### 4. Скопировать `kalkancrypt`  в `/opt/`

`kalkancrypt` - представляет набор из общих библиотек и состоит из файлов расширения `.so`.

Скопируйте папку `SDK/C/Linux/libs_for_linux/kalkancrypt` в `/opt/`

```sh
sudo cp -r kalkancrypt /opt/
```

#### 5. LD_LIBRARY_PATH

При обращении к GoKalkan убедитесь что экспортирована переменная окружения

```sh
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/kalkancrypt/:/opt/kalkancrypt/lib/engines
```

Это переменная нужна для динамического обращения к библиотеке KalkanCrypt.


## License

The MIT License (MIT) 2021 - [Abylaikhan Zulbukharov](https://github.com/Zulbukharov).

Please have a look at the [LICENSE.md](https://github.com/Zulbukharov/kalkancrypt-wrapper/blob/master/LICENSE.md) for more details.

