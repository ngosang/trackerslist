# auto-update for linux
This tool will download the best trackers list, and auto update the config file of aria2 "~/.aria2/aria2.conf".

The default url is "https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_best.txt", you can change it by argument.

这个小工具用于下载最佳tracker列表，并且自动更新aria2配置文件"~/.aria2/aria2.conf"。

默认列表文件的下载地址是："https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_best.txt"，用户可以用参数改变下载地址。

#### Usage（用法）:
```
    //update default best list，下载默认最佳列表。
    auto-update
    //update special list，下载指定列表
    auto-update [url of list]
```
