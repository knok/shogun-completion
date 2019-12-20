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

# shogun.go

このソースはgo言語によって記述されたshogunコマンド実装です。

# 使い方

    $ go build shogun.go
    $ ./shogun 徳川家光
    徳川家光 (the 3rd Shogun)
    $ ./shogun kamakura 源実朝
    源実朝 (the 3rd Shogun)

# バグ・仕様・課題

* UTF-8限定であり、localeを見ない
* 日本語出力しかない
* 室町幕府に対応していない

# license

Copyright 2013, 2015, 2019 NOKUBI Takatsugu <knok@daionet.gr.jp>, mattn

Copying and distribution of this file, with or without modification,
are permitted in any medium without royalty provided the copyright
notice and this notice are preserved.  This file is offered as-is,

# fish実装 (by homulerさん)

* https://gist.github.com/homuler/3d0842205cf121ad40f4797a6c0c492a
* [本人のコメント](https://qiita.com/knok/items/3a08663895bb3a3d6fd0#comment-30a3a31581227bdfecf1)
* [ツイート](https://twitter.com/eulerdora/status/1207719145556459520)

# ChangeLog

* 2019/12/28
    * 未定義将軍に対するエラー処理の追加 (by mattnさん, PR #1)
* 2015/04/01
    * 鎌倉幕府対応
    * shogunコマンド実装
* 2013/04/01
    * 初期実装
