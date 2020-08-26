# cuentas_contables_crud

El API cuentas_contables_crud, proporciona interfaces para la manipulación(CRUD) de los datos almacenados en una base de datos no relacional **MongoDB** (comprobantes, tipos de comprobantes, árbol de cuentas contables).  
Esta API representa la capa de datos del subsistema de contabilidad, el cual, a su vez, hacer parte de el sistema de gestión financiero KRONOS.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [MongoDB]()
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# Ejemplo que se debe actualizar acorde al proyecto
CUENTAS_CONTABLES_CRUD_DB_USER = [descripción]
CUENTAS_CONTABLES_CRUD_DB_PASS = [descripción]
CUENTAS_CONTABLES_CRUD_DB_HOST = [descripción]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con CUENTAS_CONTABLES_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/cuentas_contables_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/cuentas_contables_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
CUENTAS_CONTABLES_HTTP_PORT=8080 CUENTAS_CONTABLES_DB_HOST=127.0.0.1:27017 CUENTAS_CONTABLES_SOME_VARIABLE=some_value bee run
```


### Ejecución Dockerfile
```shell
# docker build --tag=cuentas_contables_crud . --no-cache
# docker run -p 80:80 cuentas_contables_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/cuentas_contables_crud

#2. Moverse a la carpeta del repositorio
cd cuentas_contables_crud

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```
## Estado CI
| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cuentas_contables_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/cuentas_contables_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cuentas_contables_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/cuentas_contables_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cuentas_contables_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/cuentas_contables_crud) |


## Licencia
This file is part of cuentas_contables_crud

cuentas_contables_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

cuentas_contables_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with necesidades_crud. If not, see https://www.gnu.org/licenses/.
