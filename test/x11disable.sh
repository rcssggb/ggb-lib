HOST_OS=$(uname -s)

echo "Disabling X11 exposure"

if [ $HOST_OS == "Linux" ]; then
    xhost -local:root
elif [ $HOST_OS == "Darwin" ]; then
    xhost - 127.0.0.1
fi
