# Conexión básica a Base de Datos
## Archivos
* ### `main.go`
Contiene las funciones necesarias para ejecutar el servidor.
* ### `connection.go`
Establece la conexión con la base de datos. La función `loadDotEnv(env string)` lee las variables de entorno con las credenciales y datos necesarios para esta conexión.
* ### `create.go`
La función `create()` muestra el formulario, mientras que `insert()` recpeciona las peticiones POST del mismo para realizar las inserciones en la base de datos.
* ### `index.go`
Muestra los registros de la base de datos y los muestra en pantalla. Hace uso de la función `connectionDB()` de `connection.go`.
* ### `templates/*`
Dentro se encuentran los archivos .html para la interfaz.