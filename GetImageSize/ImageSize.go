package main

import (
	"fmt"
	//	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

/*
//裁剪图片
func ICropFaceImage(jpgBuffer []byte, x0, y0, x1, y1 int) *protocol.IImage {
	buffer := bytes.NewBuffer(jpgBuffer)
	m, _, err := image.Decode(buffer)
	if err != nil {
		return nil
	}
	rgbImg := m.(*image.YCbCr)
	subImg := rgbImg.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
	subBuffer := new(bytes.Buffer)
	err = jpeg.Encode(subBuffer, subImg, nil)
	if err != nil {
		return nil
	}
	return &protocol.IImage{3,
		int32(x1 - x0),
		int32(y1 - y0),
		subBuffer.Bytes(),
	}
}

func HandlerImage() {
	data, err := ioutil.ReadFile("d:/1.jpg")
	m, _, _ := image.Decode(data) // 图片文件解码
	rgbImg := m.(*image.YCbCr)
	subImg := rgbImg.SubImage(image.Rect(0, 0, 200, 200)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
}*/

//打开文件
func OpenFile() error {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "1.jpg")
	if nil != err {
		return err
	}
	defer file.Close()
	buffer, _ := ioutil.ReadAll(file)
	fmt.Println(buffer)

	return nil
}

//获取当前工作路径
func Get_Current() error {
	dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	fmt.Println(dirctory)
	return nil
}

//获取图片高度与宽度(image/jpeg)
func GetJpg_WidthHeight() {
	file, err := os.Open("d:/1.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	img, err := jpeg.DecodeConfig(file) //解码
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(img.Height)
	fmt.Println(img.Width)
}

//获取文件夹中文件信息
func ReadFolderInfo() {
	dir_list, e := ioutil.ReadDir("d:/test")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name())
		fmt.Println(v.Name(), "的权限是:", v.Mode())
		fmt.Println(v.Name(), "文件大小:", v.Size())
		fmt.Println(v.Name(), "创建时间", v.ModTime())
		fmt.Println(v.Name(), "系统信息", v.Sys())
		if v.IsDir() == true {
			fmt.Println(v.Name(), "是目录")

		}
	}
}

//读取文件信息
func ReadFileInfo() {
	data, err := ioutil.ReadFile("d:/1.jpg")
	if err != nil {
		fmt.Println("read error")
		os.Exit(1)
	}
	fmt.Println(string(data))
}

//保存文件信息
func SaveFileInfo() error {
	err := ioutil.WriteFile("d:/22.jpg", []byte("123456"), 0666)
	if err != nil {
		return err
	}
	return nil
}

//创建图片存储路径
func CreateStoreImagePath() error {
	//获取当前工作路径
	dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	dirctory = dirctory + "\\" + "test"
	//判断图片文件夹是否存在
	_, err = os.Stat(dirctory)
	if os.IsExist(err) {
		return nil
	}
	//不存在，则创建目录
	err = os.MkdirAll(dirctory, 0666)
	if err != nil {
		return err
	}
	return nil
}

//遍历删除目录下的所有文件
func DeleteFileDir(dirctory string) error {
	fileNames := make([]string, 0) //存放文件
	dirNames := make([]string, 0)  //存放空目录
	err := filepath.Walk(dirctory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirNames = append(dirNames, path)
		} else {
			fileNames = append(fileNames, path)
		}
		return err
	})
	if err != nil {
		return err
	}

	//删除所有文件夹下的文件
	for i := len(fileNames) - 1; i >= 0; i-- {
		err = os.Remove(fileNames[i])
		if err != nil {
			return err
		}
	}
	//删除所有空文件夹
	for j := len(dirNames) - 1; j >= 0; j-- {
		err = os.Remove(dirNames[j])
		if err != nil {
			return err
		}
	}
	return nil
}

//ioutil.ReadAll函数
func ReadAllFunc() {
	reader := strings.NewReader("hello word widuu") //返回*strings.Reader
	fmt.Println(reflect.TypeOf(reader))
	data, _ := ioutil.ReadAll(reader)
	fmt.Println(string(data))
}

//创建临时文件
func TempDir() {
	dir, err := ioutil.TempDir("D:/test", "tmp")
	if err != nil {
		fmt.Println("创建临时目录失败")
		return
	}
	fmt.Println(dir) //返回的是D:\test\tmp846626247 就是前边的prefix+随机数

}

//
func TempFile() {
	file, error := ioutil.TempFile("D:/test", "tmp")
	defer file.Close()
	if error != nil {
		fmt.Println("创建文件失败")
		return
	}
	file.WriteString("Hello word") //利用file指针的WriteString()详情见os.WriteString()
	filedata, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(filedata))
}

//
func main() {
	var str string = "1234"
	fmt.Println(str)
	fmt.Println([]byte(str))
	/*
			//获取图片的高度与宽度
			GetJpg_WidthHeight()
			//读取文件夹信息
			ReadFolderInfo()
		//读取文件信息
		ReadFileInfo()
		//保存文件信息
		/*err := SaveFileInfo()
			if err != nil {
				fmt.Println("创建文件夹失败!")
			}
			//创建文件夹
			err = CreateStoreImagePath()
			if err != nil {
				fmt.Println("创建文件夹失败!")
			}
			//遍历某个路径下的所有文件
			err = DeleteFileDir("d:/test")
			if err != nil {
				fmt.Println("创建文件夹失败!")
			}


		ReadAllFunc()

		TempDir()

		TempFile()*/
}

/*
import (
	"fmt"
	"io/ioutil"
)

func main() {
	imagebuffer, errcode := ioutil.ReadFile("d:/1.jpg")
	if errcode != nil {
		fmt.Println("read image file fail")
		return
	}
	fmt.Println(imagebuffer)

	//获取到高度与宽度的十六进制
	var height, width [2]int
	for i := 0; i < len(imagebuffer); i++ {
		if imagebuffer[i] == 255 && imagebuffer[i+1] == 192 {
			height[0] = int(imagebuffer[i+5])
			height[1] = int(imagebuffer[i+6])
			width[0] = int(imagebuffer[i+7])
			width[1] = int(imagebuffer[i+8])
		}
	}

	var str1 [4]int
	var k int = 0
	for j := 0; j < 2; j++ {
		hex, length := toHex(int(height[j]))
		for i := 0; i < length; i++ {
			if hex[i] >= 10 {
				//fmt.Printf("%c", 'A'+hex[i]-10)
				str1[k] = 'A' + hex[i] - 10
			} else {
				//fmt.Print(hex[i])
				str1[k] = hex[i]
			}
			k++
		}
	}
	fmt.Print(str1)

	fmt.Print(width)
	var t int = 0
	var str2 [4]int
	for j := 0; j < 2; j++ {
		hex1, length1 := toHex(int(width[j]))
		for i := 0; i < length1; i++ {
			if hex1[i] >= 10 {
				fmt.Printf("%c", 'A'+hex1[i]-10)
				str2[t] = 'A' + hex1[i] - 10
			} else {
				fmt.Print(hex1[i])
				str2[t] = hex1[i]
			}
		}
		t++
	}
	fmt.Print(str2)

}

//将十六进制转换成十进制

//将十进制转换成十六进制
func toHex(ten int) (hex []int, length int) {
	m := 0
	hex = make([]int, 0)
	length = 0
	if ten < 16 {
		hex = append(hex, 0)
		hex = append(hex, ten)
		length += 2
		return
	}
	for {
		m = ten / 16
		ten = ten % 16
		if m == 0 {
			hex = append(hex, ten)
			length++
			break
		}
		hex = append(hex, m)
		length++
	}
	return
}
*/
