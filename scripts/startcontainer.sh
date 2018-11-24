NAME=${1}
RUN_PARAMS=${2}

function isNotRunning() {
  echo $(docker ps --format "{{.Names}}" | grep -E "^${NAME}$" | wc -l)
}

if [ $(isNotRunning) -eq 0 ]; then
        for CID in $(docker ps -a --format "{{.Names}}" | awk '$1 == "'${NAME}'" {print $1}'); do
                docker stop ${CID}
                docker rm ${CID}
        done
        CMD="docker run -d --name ${NAME} ${RUN_PARAMS}"
        echo "Starting container ${IMAGE} with command '${CMD}'"
        eval ${CMD}
fi
