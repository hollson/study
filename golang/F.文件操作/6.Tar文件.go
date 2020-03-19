package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)
//tar 是一种打包格式，但不对文件进行压缩，所以打包后的文档一般远远大于 zip 和 tar.gz，因为不需要压缩的原因，所以打包的速度是非常快的，打包时 CPU 占用率也很低。
//
//tar 的目的在于方便文件的管理，比如在我们的生活中，有很多小物品分散在房间的各个角落，为了方便整洁可以将这些零散的物品整理进一个箱子中，而 tar 的功能就类似这样。
//
//创建 tar 归档文件与创建 .zip 归档文件非常类似，主要不同点在于我们将所有数据都写入相同的 writer 中，并且在写入文件的数据之前必须写入完整的头部，而非仅仅是一个文件名。
//
//tar 打包实现原理如下：
//
//创建一个文件 x.tar，然后向 x.tar 写入 tar 头部信息；
//打开要被 tar 的文件，向 x.tar 写入头部信息，然后向 x.tar 写入文件信息；
//当有多个文件需要被 tar 时，重复第二步直到所有文件都被写入到 x.tar 中；
//关闭 x.tar，完成打包。



func main() {
	f, err := os.Create("./output.tar") //创建一个 tar 文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	tw := tar.NewWriter(f)
	defer tw.Close()

	fileinfo, err := os.Stat("./main.exe") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err)
	}

	err = tw.WriteHeader(hdr) //写入头文件信息
	if err != nil {
		fmt.Println(err)
	}

	f1, err := os.Open("./main.exe")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) //将main.exe文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}


func main2() {
	f, err := os.Open("output.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer f.Close()
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("123" + fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}