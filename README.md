How to use groupcache

使用的groupcache的顺序

1.首先需要先建立httppool，建立的时候可以设定defaultBasePath和defaultReplicas。前者是默认基础路径，在URL中排在端口号后。后者是默认的副本数，所有提供的地址都会经过一定处理生成一个哈希，副本数代表一个地址存在副本数的哈希地址后（此处用到哈希一致性的知识，不同哈希地址处存同一个地址，便于找到最近的哈希地址取值）。

2.创建一个group，同时附带getter，getter为当group在mainCache,hotCache以及other peers出都无法获得所需数据时，获取数据的途径（一般是从数据库等数据源处获取，也可以设定为再次添加数据，作为set data的方法）。

3.通过创建的group变量Get数据。
    1.Get数据的流程是先从当前Group的mainCache和hotCache中获取，如果有则改变groupStatus的状态并返回该值。
    2.如果没有则通过loadGroup去找。loadGroup对于查询同样数据的请求只通过第一个，其余通过waitGroup阻塞，在通过的请求取到值的时候同时返回。
    3.通过loadGroup的请求首先重新在mainCache和hotCache再查找一遍，然后通过哈希一致性等规则找到哈希地址最近的一个结点地址，然后通过http发送请求获取相应数据，如果收到则更新groupStatus的状态，一定几率更新本地hotCache,并返回数据，如果没有，则从本地查找（创建group时附带的getter方法）。如果没有错误则更新groupStatus的状态,存入mainCache，并返回数据。


