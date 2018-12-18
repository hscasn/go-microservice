MODIFIED_APPS=$(git log --name-only --oneline -1 | sed 1d | grep -e '^apps/' | cut -d "/" -f2 | uniq)

echo "================================================================================"
echo "Modified apps: $(echo ${MODIFIED_APPS} | tr '\n' ' ')"
echo "================================================================================"

ROOTPATH="$(dirname $(readlink -f ${0}))/.."
for APP in ${MODIFIED_APPS}; do
	APPPATH="${ROOTPATH}/apps/${APP}"
	cd ${APPPATH}
	make deploy
	cd ${ROOTPATH}
done