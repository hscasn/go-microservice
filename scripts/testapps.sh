ROOTPATH="$(dirname $(readlink -f ${0}))/.."
CI=0
if [ "${1:='-'}" = "ci" ]; then
    CI=1
fi

for APP in $(ls apps); do
    APPPATH="${ROOTPATH}/apps/${APP}"
    if [ ! -d ${APPPATH} ]; then
        continue
    fi
    cd ${APPPATH}
    if [ ${CI} -gt 0 ]; then
        make test-ci
    else
        make test
    fi
    TESTRESULT=$?
    cd ${ROOTPATH}
    if [ ! ${TESTRESULT} -eq 0 ]; then
        (>&2 echo "Tests failed for application ${APPPATH}")
        exit ${TESTRESULT}
    fi
done