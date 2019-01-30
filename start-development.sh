docker build -t chlorine .
docker run -p 8080:8080 --rm --env-file ./.env --name chlorine-development chlorine