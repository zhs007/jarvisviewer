docker container stop jarvisviewer
docker container rm jarvisviewer
docker run -d --name jarvisviewer \
    -v $PWD/cfg:/home/jarvisviewer/cfg \
    -v $PWD/logs:/home/jarvisviewer/logs \
    -v $PWD/dat:/home/jarvisviewer/dat \
    -p 7710:7710 \
    -p 7788:7788 \
    jarvisviewer