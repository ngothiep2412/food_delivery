APP_NAME=food-delivery

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME} #stop container ${APP_NAME}
#docker rmi $(docker images -qa -f 'dangling=true')


docker run -d --name ${APP_NAME} \
-- network my-net \
-e VIRTUAL_HOST="g0.5.custohub.com" \
-e LETSENCRYPT_HOST="g0.5.custohub.com" \
-e LETSENCRYPT_EMAIL="ngothiep@custohub.com" \
-e DBConnectionStr=""
-e S3BucketName=""
-e S3Region=""
-e S3AccessKey=""
-e S3SecretKey=""
-e S3Domain=""
-e SYSTEM_SECRET=""
-p 8080:8080 \
${APP_NAME}