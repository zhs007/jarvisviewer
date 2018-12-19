docker build -t jarvisviewer .

if [ ! -d "logs" ]; then
    mkdir logs
fi

if [ ! -d "dat" ]; then
    mkdir dat
fi