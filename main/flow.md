## original

http://lists.w3.org/Archives/Public/ietf-http-wg/2013JulSep/1135.html

The thing is that we don't have to emit toggle off first.
After all current header set is processed,
the entries in the reference set and not emitted and not common header are to be removed.
To handle eviction of common header,
emit 2 indexed representation for it just before the removal.


1. For each name/value pair in the current header set:

ヘッダセット内の name/value ペアについて

1.1. If name/value pair is present in the header table,
and the corresponding entry in the header table is NOT in the reference set,
add the entry to the reference set and encode it as indexed representation.
Mark the entry "emitted".

もし、 name/value ペアがヘッダテーブルにあり、
そのエントリがリファレンスセットには無かったら、
エントリをリファレンスセットに追加して、
index representation でエンコードする。
"emitted" とマークする


1.2. If name/value pair is present in the header table,
and the corresponding entry in the header table is in the reference set:

もし、 name/value ペアがヘッダテーブルにあり、
そのエントリがリファレンスセットにもあったら。

1.2.1. If the entry is marked as "common-header",
then this is the 2nd occurrence of the same indexed representation.
To encode this name/value pair,
we have to encode 4 indexed representation.
2 for the 1st one (which was the name/value pair processed in 1.2.3.),
and the another 2 for the current name/value pair.
Unmark the entry "common-header" and mark it "emitted".

もしエントリが "common-header" としてマークされていたら、
これは同じ indexed representation の二回目のチェックである。
この name/value をエンコードするには、 4 つの indexed representation を
エンコードしないといけない。
2 つは、最初のもの(1.2.3 で処理された name/value ペア)
残りの 2 つは現在の name/value ペアのため。
エントリの "common-header" を外して、 "emitted" にする。

1.2.2. If the entry is marked as "emitted",
then this is also the occurrences of the same indexed representation.
But this time, we just encode 2 indexed representations.

エントリが "emitted" とマークされたら、
これも同じ indexed representation への処理である。
しかしこの場合、 2 つの indexed representation だけですむ。

1.2.3. Otherwise, just mark the entry "common-header"
and not encode it at the moment.

これ以外では、エントリは "common-header" としてマークして
この時点ではエンコードしない。


1.3. If name/value pair is not present in the header table,
encoder encodes name/value pair as literal representation.
On eviction or substitution, If the entry to be removed is
in the reference set and marked as "common-header", encode
it as 2 indexed representations before the removal. On
removal, it is removed from the reference set.

もし、 name/value ペアがヘッダテーブルに無かったら、
エンコーダーは、 name/value を literal representation としてエンコードする。
eviction か substitution では、もし削除されるエントリがリファレンスセットにあり、
"common-header" とマークされていたら、
削除する前に、 2 つの indexed representation でエンコードする。
消す時は、リファレンスセットからも消す。


2. For each entry in the reference set:
if the entry is in the reference set
but is neither marked as "emitted" nor "common-header",
remove it from the reference set and encode its index as indexed representation.

リファレンスセットの個々のエントリについて、
もしエントリがリファレンスセットにあるが、
"emitted" でも "common-header" でもマークされてなかったら
リファレンスセットから消して、そのインデックスを
indexed representation としてエンコードする。


3. After all current header set is processed, unmark all entries in
the header table.

全部のヘッダセットを処理したら、
ヘッダテーブルの全エントリからマークを外す。
