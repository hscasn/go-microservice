ROOTPATH="$(dirname $(readlink -f ${0}))/.."

MODIFIED_APPS=($(git log --name-only --oneline -1 | sed 1d | grep -e '^apps/' | cut -d "/" -f2 | uniq))
MODIFIED_PKG=($(git log --name-only --oneline -1 | sed 1d | grep -e '^pkg/' | cut -d "/" -f2 | uniq))

APPS_TO_DEPLOY=()
for APP in $(ls apps); do
    if [ ${#MODIFIED_PKG[@]} -gt 0 ]; then
		APPS_TO_DEPLOY+=("${APP}")
        continue
    fi
	for MOD_APP in ${MODIFIED_APPS[@]}; do
		if [ ${MOD_APP} = ${APP} ]; then
			APPS_TO_DEPLOY+=("${APP}")
			continue
		fi
	done
done

echo "================================================================================"
echo "Modified apps (${#MODIFIED_APPS[@]}): $(echo ${MODIFIED_APPS[@]} | tr '\n' ' ')"
echo "Modified packages (${#MODIFIED_PKG[@]}): $(echo ${MODIFIED_PKG[@]} | tr '\n' ' ')"
echo "Apps to deploy (${#APPS_TO_DEPLOY[@]}): $(echo ${APPS_TO_DEPLOY[@]} | tr '\n' ' ')"
echo "================================================================================"

for APP in ${APPS_TO_DEPLOY[@]}; do
	APPPATH="${ROOTPATH}/apps/${APP}"
	cd ${APPPATH}
	make deploy
	cd ${ROOTPATH}
done