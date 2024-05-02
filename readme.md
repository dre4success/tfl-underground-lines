### Running the app locally:
- Clone the repo
- Have [go](https://go.dev/) installed
- run in the cloned repo directory:
```bash
go mod download
```
- after you can start the app:
```bash
go run main.go
```
The app should be available on:
`http://localhost:6060`

Another option is to pull the docker image:
```bash
docker pull dre4success/tfl-app:latest
```
Once pulled, you can run with this command:
```bash
docker run -p 6060:6060 dre4success/tfl-app
```
The app is available on `http://localhost:6060`.

https://github.com/dre4success/switchcraft-interview/assets/26462670/fb29bcde-e238-4ddd-9780-3bb8c934ea02