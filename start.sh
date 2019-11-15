docker build -t qk-notes-app -f Dockerfile . &&
docker run -d --network="host" --env-file config/development.env -p 3000:3000 qk-notes-app