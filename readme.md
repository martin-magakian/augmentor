#build image
docker build -t martinmagakian/augmentor .
docker run -it --add-host redis_db:192.168.99.100 -p 3000:3000 martinmagakian/augmentor

#run tools
cd vendor/beeswax/tools/augmentor/augmentor_requests_generator/augmentor
./augmentor_requests_generator ../augmentor_sample_request_1.txt http://192.168.99.100:3000/ --path-to-responses-file responses.txt --header-secret some-secret --log-level debug