# Downloads image
FROM golang:1.16.3
# Defines the working directory in the container
WORKDIR /app/src
# Copy all files into the directory
COPY . .
# Run command
RUN go build -o items-api .
#Expose
EXPOSE 8082
#Command
CMD ["./items-api"]