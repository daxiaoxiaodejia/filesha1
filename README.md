# filesha1
生成文件夹里的文件sha1值

# 简单示例
```golang
	fileSha1, err := filesha1.NewFileSha1(`{"root":"file","outputFile":"1.txt","exclude":["nihao","hahahah","hahahah/"]}`)
	if err != nil {
		log.Debug(err.Error())
	}
	fileSha1.HandleFilelist()
```

# 过滤规则

* 只忽略excludeLable目录，不忽略excludeLable文件
excludeLable/

* 忽略excludeLable文件和excludeLable目录
excludeLable

* excludeLable支持正则表达式
