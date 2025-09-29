# Etapa 1: Compilación
FROM golang:1.23-alpine AS builder

# Instalar dependencias necesarias
RUN apk add --no-cache git

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de módulo Go
COPY go.mod go.sum* ./

# Descargar dependencias (si las hay)
RUN go mod download || true

# Copiar el código fuente
COPY . .

# Compilar la aplicación
# CGO_ENABLED=0 para crear un binario estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hero-api .

# Etapa 2: Imagen final (mucho más pequeña)
FROM alpine:latest

# Instalar certificados SSL (necesarios para HTTPS)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el binario compilado desde la etapa de builder
COPY --from=builder /app/hero-api .

# Exponer el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./hero-api"]