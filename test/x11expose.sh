HOST_OS=$(uname -s)

echo "Exposing X11"

if [ $HOST_OS == "Linux" ]; then
    xhost +local:root
elif [ $HOST_OS == "Darwin" ]; then
    xhost + 127.0.0.1
else
    echo "unknown host os"
    exit
fi
