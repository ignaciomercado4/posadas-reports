# Usamos una imagen base ligera de Go
FROM golang:1.23.3-alpine AS builder

# Establecemos el directorio de trabajo
WORKDIR /app

# Instalamos dependencias necesarias
RUN apk add --no-cache git

# Copiamos el archivo go.mod y go.sum al contenedor
COPY go.mod go.sum ./

# Descargamos las dependencias del proyecto
RUN go mod download

# Copiamos el resto del código fuente al contenedor
COPY . .

# Construimos la aplicación
RUN go build -o main .

# Imagen final para ejecutar la aplicación
FROM alpine:latest

# Instalamos el cliente de PostgreSQL para conexiones
RUN apk --no-cache add postgresql-client

# Creamos un directorio para la app
WORKDIR /root/

# Copiamos el binario de la fase de construcción
COPY --from=builder /app/main .

# Exponemos el puerto en el que escucha Gin (por defecto 8080)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
