# Admin Backend with Golang

### docker-compose

- 8000:3000 mean internal called with port  3000 external called with port 8000
- .:/app mean if files in go-admin-backend folder has changed it will move to container folder called app (likes workdir)
- depends on mean will start somethin in depends on command first and then right after that it will start command that depends on in it