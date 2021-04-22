set -e
hexo g && hexo d
git add .
git commit -a -m "update"
git push
#代码保存github
