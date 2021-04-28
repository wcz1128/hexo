set -e
if [ `whoami` != "hippo" ];then
echo "!!!! only hippo can run me"
exit 1
fi
hexo g && hexo d
git add .
git commit -a -m "update"
git push
#代码保存github
