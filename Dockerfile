# Usa la imagen base oficial de Go para compilar nuestro código
FROM golang:1.20 as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente de la aplicación en el contenedor
COPY . .

# Compila la aplicación
RUN go build -o main .

# Usa la imagen base de Alpine para una imagen más pequeña
FROM alpine:latest

WORKDIR /root/

# Copia el ejecutable desde la etapa de construcción
COPY --from=builder /app/main .

# Expone el puerto 8080
EXPOSE 8080

# Ejecuta el ejecutable
CMD ["./main"]
