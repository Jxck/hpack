Encoder
-------

0 KeepSet は空で初期化する
keep set is initialized as empty set.

1 reference set から不要なものを消す

reference set の各エントリについて
それが、 header set 内にあるかを調べる。
もしなければ indexed でエンコードして
reference set から消す。

For each entry in the reference set,
check that it is present in the current header set.
If it is not, encode it as indexed representation
and remove it from the reference set.


2 header set から送る必要の無いものを消す

reference set の各エントリについて
それが、 header set 内にあるかを調べる。
もしあったら、 "common-header" としてマークし
header set から一致する name/value のペアを消す。
もし、複数の name/value が一致したら
そのうちの一つだけを、 header set から消す。

For each entry in the reference set,
check that it is present in the current header set.
If it is present, mark the entry as "common-header"
and remove the matching name/value pair from current header set
(if multiple name/value pairs are matched,
 only one of them is removed from the current header set).


3 残りの処理

header set 内の、残りの name/value ペアをエンコードする
各ペアについて:

Encode the rest of name/value pair in current header set.
For each name/value pair:

3.1

もし、 name/value ペアが header table にあり、
それが、まだ reference set には無かったら
reference set にエントリを追加して
indexed としてエンコードし
"emmitted" としてマークする。

If name/value pair is present in the header table,
and the corresponding entry in the header table is NOT in the reference set,
add the entry to the reference set
and encode it as indexed representation.
Mark the entry "emitted".

3.2

もし、 name/value ペアが header table にあり、
それが、すでに reference set にあった場合。:

もし、エントリが "common-header" とマークされていたら
それは、同じ indexed に対する二回目のチェックである。
この name/value ペアをエンコードするために
4 つの indexed が必要になる。
つは、最初の一つのため (step 2 で消したもの)
残りの 2 つは、現在の name/value ペアのため。
エントリの "common-header" を消して、 "emitted" にする。
もし、 "emitted" にマークされていたら
それもまた、同じ indexed に対する二回目の表現であるが、
この場合、エンコードは 2 つの indexed だけで済む。

If name/value pair is present in the header table,
and the corresponding entry in the header table is in the reference set:
If the entry is marked as "common-header",
then this is the 2nd occurrence of the same indexed representation.
To encode this name/value pair,
we have to encode 4 indexed representation.
2 for the 1st one (which was removed in step 2),
and the another 2 for the current name/value pair.
Unmark the entry "common-header" and mark it "emitted".

If the entry is marked as "emitted",
then this is also the occurrences of the same indexed representation.
But this time, we just encode 2 indexed representation.

3.3

それ以外の name/value ペアは literal でエンコードする。
削除か変更の場合、消されたエントリが reference set にあったら、
reference set からは消す。
消されたエントリが "common-header" だったら、 keep set に加える。

Otherwise, encoder encodes name/value pair as literal representation.
On eviction or substitution, if the removed entry is in the reference set,
it is removed from the reference set.
If the removed entry is marked as "common-header", add it to keep set.

4

keep set 内の各エントリの name/value について
3.1 から 3.3 までの同じプロセスを実行する。
keep set はイテレーションの中で更新される。

For each name/value of entry in keep set,
do the same processing described in 3.1 thourgh 3.3.
keep set may be updated in the iteration.


5

全ての header set が処理されたら
header table の全てのエントリからマークを外す。

After all current header set is processed,
unmark all entries in the header table.

Decoder
-------

基本的には、エンコーダの言う通り動くだけ。
しかし、 common header を適切に消すためには
header table にあるエントリが消されるとき
削除や更新のために
もしエントリが reference set にあって
まだそれが emit されていなかったら
消すときに emit する。

Decoder generally just performs what the encoder emitted.
But to handle common header gracefully with eviction,
when the entry in the header table is removed from the header table
due to the eviction or substitution,
if the entry is in the reference set
and it is not emitted in the current header processing,
emit the entry on the removal.
