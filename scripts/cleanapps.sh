ROOTPATH="$(dirname $(readlink -f ${0}))/.."

for APP in $(ls apps); do
    APPPATH="${ROOTPATH}/apps/${APP}"
    if [ ! -d ${APPPATH} ]; then
        continue
    fi
    cd ${APPPATH}
    make clean
    cd ${ROOTPATH}
done