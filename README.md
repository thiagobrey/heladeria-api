# Heladería API

API RESTful para la gestión de usuarios, pedidos, gustos, cantidades y administradores de una heladería.

---

## Tabla de contenidos

- [Rutas principales](#rutas-principales)
- [Importar colección Postman](#importar-colección-postman)
- [Configuración y conexión a la base de datos](#configuración-y-conexión-a-la-base-de-datos)
- [Levantar entorno local](#levantar-entorno-local)
- [Build del proyecto](#build-del-proyecto)
- [Variables de entorno](#variables-de-entorno)

---

## Rutas principales

Todas las rutas requieren autenticación por middleware salvo `/admin/login`.

### Usuarios (`/user`)

- `GET /user/:id` - Obtener usuario por ID
- `GET /user/all` - Listar todos los usuarios
- `POST /user` - Crear usuario
- `PUT /user/:id` - Actualizar usuario
- `DELETE /user/:id` - Eliminar usuario

### Pedidos (`/pedido`)

- `GET /pedido/:user_id` - Obtener pedidos por ID de usuario
- `GET /pedido/all` - Listar todos los pedidos
- `POST /pedido` - Crear pedido
- `PUT /pedido/:ID` - Actualizar pedido
- `DELETE /pedido/:id` - Eliminar pedido

### Gustos (`/taste`)

- `GET /taste/:id` - Obtener gusto por ID
- `GET /taste` - Listar todos los gustos
- `POST /taste` - Crear gusto
- `PUT /taste/:id` - Actualizar gusto
- `DELETE /taste/:id` - Eliminar gusto

### Cantidades (`/cantidad`)

- `GET /cantidad/:id` - Obtener cantidad por ID
- `GET /cantidad` - Listar todas las cantidades
- `POST /cantidad` - Crear cantidad
- `PUT /cantidad/:id` - Actualizar cantidad
- `DELETE /cantidad/:id` - Eliminar cantidad

### Administradores (`/admin`)

- `POST /admin` - Crear administrador
- `GET /admin/:id` - Obtener administrador por ID
- `PUT /admin/:id` - Actualizar administrador
- `DELETE /admin/:id` - Eliminar administrador
- `POST /admin/login` - Login de administrador (no requiere autenticación)

---

## Importar colección Postman

1. Abre Postman.
2. Haz clic en "Import".
3. Selecciona el archivo `Heladería.postman_collection.json` incluido en este repositorio.
4. Cambia la variable `url_heladeria` a `http://localhost:8080` para pruebas locales.
5. Utiliza las rutas y ejemplos incluidos para probar la API.

---

## Configuración y conexión a la base de datos

La API utiliza MySQL y GORM para la conexión y migración automática de modelos.

**Configuración por defecto:**

- Host: `localhost`
- Puerto: `3306`
- Usuario: `root`
- Contraseña: `301205`
- Base de datos: `heladeria`

**DSN utilizado:**

```
root:notSecureChangeMe@tcp(localhost:3307)/heladeria?charset=utf8mb4&parseTime=True&loc=Local
```

Puedes modificar estos valores en el archivo de conexión `database/db.go` según tu entorno.

---

## Levantar entorno local

1. Instala Go (versión 1.18+ recomendada).
2. Instala MySQL y crea la base de datos `heladeria` en el puerto 3307.
3. Clona el repositorio y entra en la carpeta.
4. Configura las variables de entorno en `.env` (ver sección abajo).
5. Ejecuta:
    ```
    go run main.go
    ```
    El servidor se levantará en `http://localhost:8080`.

---

## Build del proyecto

Para compilar el binario ejecuta:

```
go build -o heladeria-api main.go
```

Esto generará el ejecutable `heladeria-api` en el directorio actual.

---

## Variables de entorno

El archivo `.env` debe contener:

```
USER_ADMIN=admin@admin.com
USER_PASS=admin123
```

Estas credenciales se utilizan para el usuario administrador inicial y para login en `/admin/login`.

---

## Notas

- Todas las rutas (excepto login) requieren autenticación JWT.
- El middleware de autenticación utiliza la base de datos para validar el token.
- La migración de modelos es automática al iniciar el proyecto.
- Puedes modificar el DSN de la base de datos en `database/db.go` si tu entorno es diferente.

---

¿Dudas o sugerencias? ¡Abre un issue!
