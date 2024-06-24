# 阴阳神玉

![](https://img.shields.io/github/languages/top/CuteReimu/YinYangJade "语言")
[![](https://img.shields.io/github/actions/workflow/status/CuteReimu/YinYangJade/golangci-lint.yml?branch=master)](https://github.com/CuteReimu/YinYangJade/actions/workflows/golangci-lint.yml "代码分析")
[![](https://img.shields.io/github/contributors/CuteReimu/YinYangJade)](https://github.com/CuteReimu/YinYangJade/graphs/contributors "贡献者")
[![](https://img.shields.io/github/license/CuteReimu/YinYangJade)](https://github.com/CuteReimu/YinYangJade/blob/master/LICENSE "许可协议")

这是以下几个机器人项目的合集，基于 [onebot-11](https://github.com/botuniverse/onebot-11) 接口编写。

[![](https://img.shields.io/badge/tfcc-red?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADgAAAA4BAMAAABaqCYtAAAAGFBMVEVHcExfJkHiqcppUniibJO6irDw1eWyO1akOSYiAAAAAXRSTlMAQObYZgAAAmhJREFUOMttlT1vGzEMhlXb6VzhWngNBOjmCFKQtcXx4DVQxT9ww+2GXejvlx+6jzjhZPvB+5IiJdqYLQ4FAN7NV3Eol1rr1Xa/PzOHcK8c1r5+ZogFGv314Omc9yjwbK19eWAUHiUt031ZTkMgM9vthdMj3Yz7ue/VmOD1ytCuwrmfJ9VyTUpfViHH6ivCJetB2NwvvuezZNWCvyuc9DSXmgCGaO2PnWujCCMi9X8I3c51bhUhNQqGAeDKvodpTz2lLZBSGi9cr1uYSidSDhSAPBz3Iec8eRLGmABhr3QKsQyDwCID2VFqFblGioT+3TwttiLsZ4dDoP6QMhPMbto6NM++DNyfKPCURdpqZRiktQILkhnHUs9f7XvC7EwZM6fSlAwHnQpBby6AApUuKZuSYHZLuXnyICltUGUFXKbtkaYSVyXB+73Bia41tUdgx9V681bv/+YFUkOjwNCN6ClnrWOD/Miiwi6OrAy13hsEuNE8YpB6RPnH1rvm9HDjjncdsUi3hY7yzdpK7aX2lbfYsVBmwvCd4fkVc0YcbpQqUBAEQHQC7c8R6eIEGyy7shJGUZ6kkVwnF8m6joWsfDZGK2/HIyF9TomVnu5Q66U2zbZyQI5p6CxbkDKIkGxzJngMK+u0GhaOXKwxp7gxqSbSWyBb/8yPJYbNVI4Ba0rybfMV0oRyvQTGFFs2TSinVFfy1TvedE3o20440otbGCQVunVNJX50mm51XffQkb7zrxtzu/2W+IZItotmxN3uOwltQRvDf9jXJ5roAh+ZMSDzFoSYH1d1GXGJ/HnJP0FjX/8/PJXyEf0HxMMjDV0zUjUAAAAASUVORK5CYII=) 东方Project沙包聚集地机器人](tfcc)\
[![](https://img.shields.io/badge/hkbot-292944?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAD8AAACEBAMAAADVfO5aAAAAGFBMVEUGBw0AAAAAAAH///9BSmCkmqbk09hvboBKvLyOAAAAA3RSTlP8An1aKiFpAAAELUlEQVRYw7WY3W6jPBCGB9ZpTglXALkFAj31p5LmNNIuzSkS3ewpUun29r+Z8Rhs4yRardZqU4qfvPNr4wR290cGfwjk+/QukBcAJf4t+DUC4DzfAEijAM2nclFGgBznIREgjQAAm9G8MwMtEi6QgbpqkMsvkXABgNdabufwOqOOwLZ6F2EEemvMCeGj6sFeW7EFyEFV1VkU0FpjbLjAU1WN4gNs3ypjYwHIwgwUm6p65pS4wEAmJLjtUDUBgDFUjZ6j+I5yqQuwC4e5RvBKcj7wIrd4aMR7D8hQtNFztxToBGfNAYaqnttkl6kBw3CBAigwJ+/CLwDm8V0FQOIAgKnpl4YlE40PoNvn5AEw/ncP+IY3Ss+HxnUyfwy8rIA6VKgcgGr7vAL0EgWl5eAlioBxyQP137tXLALOTv9twmoC13dZ009hP9Cd57lfCu4Or6OoTRtcsnseQC2nPSBDr7Cp7WA82QUNgzZkKBSo/XXBfdxMbcvzvyvr0ZwoXjdVM1yvXfd2rRhwE4VbylMVDF5FFgDQmxCoaUMToIDNGVskGCPGIQDG3ZCX/ujRhgHM3vAaAuimAJz6ehvaoLVsgIKTMP4OJX4sAL25PoXAL0gMQHsDjuttYJUDmwkD5OsshsDLfSDjQt0Fvj8ChjjwLEDx74CDBfxW+NkNKyCoodT11xQFemxr0xqjmn3wGw1M8VAqBjS8cga5ikRRM/AhVxHgwMCTaUlI18D7DBxmBa9YfRT4eAS83Ae8ljPANwHKW8CHADuv7e16k0S5gFNOTpQSNJ3XphOGFh+pFskMHL1MiWBvgXLH25qt1uU0WK2SgYKCiSy+RrYgOhztIbJ2atnEMnK7gPUW1ctDrYCThhRWRG33SYDuYvbft6B3Zact4Nj9kC368+1qdJqfGlRpAfXZjfCl5528/Zron/kcVcC2uxzFzjKWsxzGeexwnN1p7zyJRgjo0GyroaXf+Qlpt0GWGHESf+hFp1EABVR7YhUfAPjE+YsRmPChgxcukBsfJuD30lNJtd6ZNqdMsAVoTydM/BogGxdQ6AEJ0F//VIxPrC25QB5MFCQCiXcqRuI4gvFAg5sI56kHwAKtyZYOztW5ASYSoESoEGBCkQK9xgDqPEoS+wExACtGFuhVK3zOQgxoJyOAOjEAy2QEgM8AIUAuGA8on15H2YO5CWGigui1Qm7C41LQRUzhpK0AH0MChQyMhyKgVgoILG4CASsFSjVbIQ/U6kNQxvOkwh5EAZ4yOUB7acRJ0jY5gLCj6OLE80YgCpj5ySxyvfooR5/NlnmlYRcJs7UGliD8KCTHtwBnfi6VD0ztev8IimV2zKWUoZPmra6BCFDuXAM+kGYWSCJAvt+zCnUv3Prc/RDIBUhvfnIXIHkElDeBjNQB7nw5QECR3vv2oPzrLyjWwP8sVtBU587EMAAAAABJRU5ErkJggg==) 空洞骑士speedrun推送小助手](hkbot)\
[![](https://img.shields.io/badge/maplebot-yellow?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADgAAAA4BAMAAABaqCYtAAAAJ1BMVEVHcEz8nQzUs5aPNwMBtxwCA8TRiSySYzIC+NL97tP82D7IbB9nPiLjlKXkAAAACXRSTlMA/v0f/vy+bP2QnEi9AAACOElEQVQ4y82Uv2vbQBSATQX22iNqwNYUlVJMNMgVONBN4A71f5BMSQsNaRbTQkMydiixtxjs6nAXtTIJ18kZAvZp8lakP6rvvZP104XSqQ+Mz/fd9+7u+e5qtf86HOeP6OCEsd2P21m3JRjE222s3gwA6e42qr0/8lFk+q5dgZ2JcIkxd6ciNkc9ZfYYK6sdHtByIFxWUrWzUZ9gS4igrHYngfATKEqqdnbUVzMCA9XPq915Mh9B4bZ2CnvU8zDQc2p97rsJbCn1dVamTEygrk/tqphQlqk5Uc3qs1StzwM36VMUvzfqFYqUTVIsqP76lPbxBP8OgoaikuYeDTCrB8U0aJE/E4hy/9VjgC9HPjMkQV9mVNzd4pQ9YAC/yzCmUFQEX3AjC/zhKxKmWIgbgG+wCZ3SbHOIsbkfxxFBW8EwjoioGFtAW8JPILAxz0UYy5XAjV6h+IsXYhZHC4EL6iCEpKmKbVjC6gcWwQAIfRv7G44MI7k6xvqdyDCiTApaOMqKpDqCXUOuMdN6A5egy4U6Rtq5XI5NHi75ctyewShutm/kQ3IUtHNr38B9rq3IwuUs7j8/DDZnqOHtMcp4P6Pce83hcXrAHvFKPE3hi8My876m8HL4r7CS1svB+qFXhMNJtqDGM15I7PHhILtml6W83jR3Pxvtkll4iz5wL8XXQ/68cO21T1DTawqo7aldeDEc7eCdadLhM08vbMcuYEfTnAsKB5rV189WH5BsbdvLadfsv364fwMSwLUJ9DUjjAAAAABJRU5ErkJggg==) GMSR群机器人](maplebot)\
[![](https://img.shields.io/badge/fengsheng-223922?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADgAAAA4BAMAAABaqCYtAAAAD1BMVEXqvYvktYS8lXFdPS6LaVBusrW8AAACm0lEQVQ4y31Ti5WkMAyzM1OAw1BACBSw4BQAhP5rOtkJs3Pvbpf5viiSZcWh+MtD/6ycv4GFfwYHTT+CErSwyP/Ba1Yd1zD9DS4TarFU1TIXXT9BAWECJi/VddacPsFQci0oxU8tWb+e/AFyTeDswsRatrofHRQXvVYedWNmquWq5w3iR8YtaHqVjYmeZVE875oSCp/z9liJIs2Hbs83yBx0qjrrvIsQzYvu1Qo0psz7rE0K6EsrWt2JGjiWoB0k+NGCv+WczAzJuL4MKnMChjap6VzJmHw5po9EKKmSRj2GBQrGlBFYLQWgMKxmrBONrYzA+fbSc0kC1S/zc5mfS8BEK5qwARYRYcF7I3/YZGkpbJ4QQtD1qeeJfrHM5p4oRbBPiNGYuHamBy9QxnEe9NhtIzkYc14Oiw9RBd2wpWbrxjrJOYslzTg0VMXJUajgyOM8c3RDEILscO1sqVO2HK6jQdxk88mYn9bAUs/d/4iwf2ebno6VhXdfFHG3YTKi2F5WxN+YNjOmzbavLcDvzYzYjnQjMwpIVzZQzGv0LzcNKr9BZ1pDYsrdK3fQdaymabSbYiOMn72pxR5FHELOFpo9vJsWjMSWhIwAwxuEgN9RsrNjAm3KN5gMsGvS/HD+fiZ+33vv7s61PdGm2y+++UGIwxAEnzDEKYFptwsvawXNLCHYJwzD4WBrzBtnWgIOhAzcyWs2WbGzoAHgAPA54VSS+M00pkckwycIyZvpmQ8mi8EdEI+30vts/htIoYHsXLNqZHIyQBT0VmJPqJ8jdbD16SP3PQE+clgxprQ+mT+5bEmnlrtPn12om2kOIi6qhYejtEP7zh0eojGbrt0UufvxuGLcbZUldln5MCRe0/36XDK1mW6j3FfR5x8z7W6hKetuAQAAAABJRU5ErkJggg==) 风声群机器人](fengsheng)

## 开始

本项目只含有业务逻辑，不负责QQ机器人的连接与认证、收发消息等功能。

在使用本项目之前，你应该首先自行搭建一个支持 [onebot-11](https://github.com/botuniverse/onebot-11) 接口的QQ机器人。例如：

- [NapCat](https://github.com/NapNeko/NapCatQQ) 基于NTQQ的无头Bot框架
- [OpenShamrock](https://github.com/whitechi73/OpenShamrock) 基于 Lsposed(Non-Riru) 实现 Kritor 标准的 QQ 机器人框架
- [Lagrange](https://github.com/LagrangeDev/Lagrange.Core) 一个基于纯C#的NTQQ协议实现，源自Konata.Core
- [LiteLoaderQQNT](https://github.com/LiteLoaderQQNT/LiteLoaderQQNT) QQNT 插件加载器
- [Gensokyo](https://github.com/Hoshinonyaruko/Gensokyo) 基于qq官方api开发的符合onebot标准的golang实现，轻量、原生跨平台
- [LLOneBot](https://github.com/LLOneBot/LLOneBot) LiteLoaderQQNT插件，使你的NTQQ支持OneBot11协议进行QQ机器人开发

> [!IMPORTANT]
> 本项目是基于onebot的正向ws接口，因此你需要开启对应机器人项目的ws监听。
>
> 本项目处理消息的格式是消息段数组，因此你需要将onobot中的`event.message_format`配置为`array`。

## 编译

```shell
go build -o YinYangJade
```

## 运行

第一次运行会生成配置文件`config.yaml`，请根据实际情况修改配置文件后重新运行。

```yml
# OneBot的ws的host
host: localhost

# OneBot的ws的port
port: 8080

# 你的机器人的QQ号
qq: 123456789

# 对应OneBot的accessToken
verifykey: ABCDEFGHIJK

# 自动退出除了以下群之外的所有群，为空则是不启用此功能
check_qq_groups:
  - 123456789
  - 987654321
```
