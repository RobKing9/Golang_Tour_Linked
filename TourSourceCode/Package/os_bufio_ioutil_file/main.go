package main

import (
	"fmt"
	"io"
	"os"
)

/*
读（r）：4
写（w）：2
执行（x）：1
其中 0 后面第一个位置处的 7 代表 4 + 2 + 1，即 rwx 权限，意思就是：文件拥有者可以读写并执行该文件
7 后面的 6 代表 4 + 2 + 0，即 rw- 权限，意思就是：文件拥有者所在的组里的其他成员可对文件进行读写操作
6 后面的 4 代表 4 + 0 + 0，即 r-- 权限，意思就是：除了上面提到的用户，其余用户只能对该文件进行读操作
*/

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
	//read.ReadFile1()
	//read.ReadFile2()
	//read.ReadFile3()
	//write.WriteFile1()
	//write.WriteFile2()
	//write.WriteFile3()
	//Copy()
	//downloadfile.DownloadFile1("http://kr.shanghai-jiuxin.com/file/mm/20211129/jvflowxl13e.jpg", "./image/image3.jpg")
	//downloadfile.DownloadFile1("https://video19.ifeng.com/video09/2022/01/30/p6893405658201198592-102-121302.mp4?reqtype=tsl&vid=065bd6ad-03f5-4adf-811a-43f28bc9b229&uid=0djwHS&from=v_Free&pver=vHTML5Player_v2.0.0&sver=&se=&cat=&ptype=&platform=pc&sourceType=h5&dt=1643618634486&gid=WX9U4X6vTgr2&sign=0faf58593d3c3e490195589db9ff6725&tm=1643618634486&vkey=4pl2W4nMPYkYiFDr4y0MfhUeI83yFQXbiN%2FgixOtNwp0LawDXLldSM8rt8rHyGHaZJm8AdAMLdhdSmb0vw0XnKrvwuYHOZVmvogWQ5piCKUZhEXZHAj0%2BNNwkN8GMBU4Pu5WmMo1AQQtnv3sb8KHd4R%2ByPwYWZEUw%2Fr1X5Ch9zfHVfclvNiasc8ksUySCS8brYG2vmi17n7HEEzu8HHWfJ6eDYHv%2Bf7qEczXADBZSibCMFncoDH1zXsRTT9TykbVN4tNcDE3lIV1hDqEj3g9OGS7Z2pLtac2OovD%2BtVmujI%3D", "./image/image4.mp4")
}

func Copy() {
	////创建目录
	//os.MkdirAll("./copy", 0744)
	////写入文件
	//ioutil.WriteFile("./copy/w.txt", []byte("你好 世界"), 0666)
	////读取文件
	//buf, err := ioutil.ReadFile("./copy/w.txt")
	//if err != nil {
	//	fmt.Println("read file failed, err:", err.Error())
	//}
	//fmt.Println(string(buf))
	// 打开文件
	read, err := os.Open("./copy/w.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err.Error())
	}
	defer read.Close()
	//打开或创建写入的文件
	write, err := os.OpenFile("./copy/wcopy.txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("open witefile failed, err:", err.Error())
	}
	defer write.Close()
	_, err = io.Copy(read, write)
	if err != nil {
		fmt.Println("copy failed, err:", err.Error())
	}
}
