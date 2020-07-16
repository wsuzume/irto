# ansible のインストール
pip3 install ansible

# ansible の設定ファイルのダウンロード
# https://qiita.com/ponsuke0531/items/1e0ab0d6845ec93a0dc0
mkdir ansible
cd ansible \
&& git init \
&& git config core.sparsecheckout true \
&& git remote add origin https://github.com/wsuzume/irto.git \
&& echo ansible > .git/info/sparse-checkout \
&& git pull origin master
