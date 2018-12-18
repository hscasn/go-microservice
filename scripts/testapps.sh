ROOTPATH="$(dirname $(readlink -f ${0}))/.."

for APP in $(ls apps); do
    APPPATH="${ROOTPATH}/apps/${APP}"
    if [ ! -d ${APPPATH} ]; then
        continue
    fi
    cd ${APPPATH}
    make test
    TESTRESULT=$?
    cd ${ROOTPATH}
    if [ ! ${TESTRESULT} -eq 0 ]; then
        (>&2 echo "Tests failed for application ${APPPATH}")
        exit ${TESTRESULT}
    fi
done