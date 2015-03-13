# shogun-completion

このスクリプトは、徳川将軍の名前をbashで補完できるように
するためのものです。

# 使い方

    $ source shogun-completion
    $ shogun 徳川[TAB][TAB]
    徳川家慶  徳川家康  徳川家重  徳川家定  徳川慶喜
    徳川家継  徳川家綱  徳川家斉  徳川家茂  徳川綱吉
    徳川家光  徳川家治  徳川家宣  徳川吉宗  徳川秀忠
    $ shogun kamakura [TAB][TAB]
    惟康親王  源実朝    源頼朝    宗尊親王  藤原頼嗣
    久明親王  源頼家    守邦親王  藤原頼経

# バグ・仕様・課題

* UTF-8限定であり、localeを見ない
* 日本語出力しかない
* 誰が何代目であるかがわからない
* 室町幕府に対応していない
* shogunコマンド自体は実在していない
