# Project Name
> ftp

``` bash
go mod tidy
go mod vendor
go build
```

## Docker

Depois de instalado versão do `Docker` (windows, linux ou mac), irá executar o comando `docker --version` para verificar se o mesmo foi instalado.

Acessar a pasta raiz do projeto e executar o comando `docker-compose up -d`

Quando quiser parar o processo executar o comando `docker-compose down`

## Testes

```bash
go test -v -cover ./...
go test ./... -coverprofile fmtcoverage.html fmt
go test ./... -coverprofile cover.out
```

## Sonarqube

Acessar o Sonarqube pela url `http://localhost:9000/`(http://localhost:9000/), com usuário `admin` e senha `admin`

Antes de iniciar o scanner, precisar instalar as ferramentas `sonar-scanner`.

```bash
# Instalar a ferramenta sonar-scanner (Linux)
apt-get update
apt-get install unzip wget nodejs
mkdir /downloads/sonarqube -p
cd /downloads/sonarqube
wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.5.0.2216-linux.zip
unzip sonar-scanner-cli-4.5.0.2216-linux.zip
mv ssonar-scanner-cli-4.5.0.2216-linux /opt/sonar-scanner
```

```bash
# Criar aquivo para automatizar a configuração de variáveis de ambiente
# Adicionar o comando sonar-scanner à variável PATH
vi /etc/profile.d/sonar-scanner.sh
```

Conteúdo do arquivo sonar-scanner.sh

```bash
#/bin/bash
export PATH="$PATH:/opt/sonar-scanner/bin"
```

```bash
# Use o comando de origem para adicionar o comando do sonar-scanner à variável PATH
source /etc/profile.d/sonar-scanner.sh
# Verificar a variável PATH foi atualizada
env | grep PATH
# Verificar a versão do sonar-scanner instalada
sonar-scanner -v
```

```bash
# Executar o sonar-scanner
sonar-scanner
```