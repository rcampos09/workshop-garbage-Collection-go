# Usa una imagen base de Go
FROM golang:1.20-alpine

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos go.mod y go.sum
COPY go.mod . 

# Descargar las dependencias
RUN go mod download

# Copia el archivo Go al contenedor
COPY example2.go .  

# Compila el archivo Go
RUN go build -o example2 . 

# Expone el puerto 8080
EXPOSE 8080

# Expone el puerto 6060 para pprof
EXPOSE 6060

# Ejecuta el binario compilado
CMD ["./example2"]  