docker container stop jarvistviewer
docker container rm jarvistviewer
docker run -d --name jarvistviewer -v $PWD/cfg:/home/jarvistviewer/cfg -v $PWD/logs:/home/jarvistviewer/logs -v $PWD/dat:/home/jarvistviewer/dat jarvistviewer