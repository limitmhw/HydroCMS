package controllers

//识别档案号，对档案号进行正则判别，如果符合正则表达式，则进行提取
import (
	"fmt"
	// "github.com/nfnt/resize"
	// "image"
	// "image/draw"
	// "image/jpeg"
	// "image/png"
	// "io/ioutil"
	// "log"
	// "math/rand"
	// "os"
	"path"
	"regexp"
	// "strconv"
	"strings"
	// "time"
)

func Record(filenameWithSuffix string) (Suffix, number, name, pronumber, projieduan, proleixing, prozhuanye string) {
	FileSuffix := path.Ext(filenameWithSuffix) //只留下后缀名
	LengthSuffix := len([]rune(FileSuffix))
	Suffix = SubString(FileSuffix, 1, LengthSuffix-1)
	fmt.Println("扩展名", Suffix)

	var filenameOnly string
	var FileNumber string
	var FileName string
	//	var projieduan string
	//	var prozhuanye string
	//	var proleixing string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, FileSuffix) //只留下文件名，无后缀
	fmt.Println("文件全名：", filenameOnly)                                //filenameOnly= mai
	//这个测试一个字符串是否符合一个表达式。
	//    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//    fmt.Println(match)
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。
	r, _ := regexp.Compile(`[[:upper:]]{2}[0-9]+[[:upper:]\.0-9]+[-][0-9]+[-][0-9]+[\p{Han} \(\)\/~]`)
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	fmt.Println(r.MatchString(filenameOnly))
	if r.MatchString(filenameOnly) { //如果符合正则表达式
		lengthname := len([]rune(filenameOnly))
		// 查找连续2个的大写字母
		reg := regexp.MustCompile(`[[:upper:]]{2}`)
		fmt.Printf("大写字母%q\n", reg.FindAllString(filenameOnly, -1))
		// ["H" "G"]
		blankloc := UnicodeIndex(filenameOnly, " ") // 查找空格这个字符的位置
		if blankloc == 0 {                          //如果没有空格                                                   //如果没有空格，则用正则表达式获取编号
			re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
			loc := re.FindStringIndex(filenameOnly)
			if loc != nil { //如果有编号
				FileNumber = SubString(filenameOnly, 0, loc[0])
				fmt.Println("文件编号：", FileNumber)
				FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
				fmt.Println("文件名：", FileName)
			} else { //如果没有编号
				FileNumber = filenameOnly
				fmt.Println("文件编号：", FileNumber)
				FileName = filenameOnly
				fmt.Println("文件名：", filenameOnly)
			}
		} else { //如果有空格
			re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
			loc := re.FindStringIndex(filenameOnly)
			if loc != nil { //如果有编号
				FileNumber = SubString(filenameOnly, 0, loc[0])
				fmt.Println("文件编号：", FileNumber)
				FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
				fmt.Println("文件名：", FileName)
			} else { //如果没有编号
				FileNumber = filenameOnly
				fmt.Println("文件编号：", FileNumber)
				FileName = filenameOnly
				fmt.Println("文件名：", filenameOnly)
			}
		}
		//这里继续提取项目号-阶段-成果类别-专业
		//首先判断有无.号，如果有，则是旧图号
		dianhaoloc := UnicodeIndex(FileNumber, ".")
		//	fmt.Println("第一个“.”位置：", dianhaoloc) //
		if dianhaoloc != 0 { //如果是旧编号
			//项目编号
			fmt.Println("项目编号：", SubString(FileNumber, 0, dianhaoloc-1))
			//阶段
			projieduan = SubString(FileNumber, dianhaoloc-1, 1)
			switch projieduan {
			case "A":
				fmt.Println("规划阶段：" + projieduan)
			case "B":
				fmt.Println("项目建议书阶段：" + projieduan)
			case "C":
				fmt.Println("可行性阶段：" + projieduan)
			case "D":
				fmt.Println("初步设计阶段：" + projieduan)
			case "E":
				fmt.Println("招标设计阶段：" + projieduan)
			case "F":
				fmt.Println("施工图设计阶段：" + projieduan)
			case "G":
				fmt.Println("竣工图阶段：" + projieduan)
			case "L":
				fmt.Println("专题：" + projieduan)
			}
			//专业
			prozhuanye = SubString(FileNumber, dianhaoloc+1, 1)
			switch prozhuanye {
			case "1":
				fmt.Println("综合：" + prozhuanye)
			case "2":
				fmt.Println("规划：" + prozhuanye)
			case "3":
				fmt.Println("测量：" + prozhuanye)
			case "4":
				fmt.Println("地质：" + prozhuanye)
			case "5":
				fmt.Println("水工：" + prozhuanye)
			case "6":
				fmt.Println("建筑：" + prozhuanye)
			case "7":
				fmt.Println("机电：" + prozhuanye)
			case "8":
				fmt.Println("资环：" + prozhuanye)
			case "9":
				fmt.Println("施工：" + prozhuanye)
			}
			//二级专业代码
		} else { //新编号
			jianhaoloc := UnicodeIndex(FileNumber, "-")
			//项目编号
			fmt.Println("项目编号：", SubString(FileNumber, 0, jianhaoloc-2))
			//阶段
			projieduan = SubString(FileNumber, jianhaoloc-2, 1)
			switch projieduan {
			case "A":
				fmt.Println("规划阶段：" + projieduan)
			case "B":
				fmt.Println("项目建议书阶段：" + projieduan)
			case "C":
				fmt.Println("可行性阶段：" + projieduan)
			case "D":
				fmt.Println("初步设计阶段：" + projieduan)
			case "E":
				fmt.Println("招标设计阶段：" + projieduan)
			case "F":
				fmt.Println("施工图设计阶段：" + projieduan)
			case "G":
				fmt.Println("竣工图阶段：" + projieduan)
			case "L":
				fmt.Println("专题：" + projieduan)
			}
			//文件类型
			proleixing = SubString(FileNumber, jianhaoloc-1, 1)
			switch proleixing {
			case "B":
				fmt.Println("技术报告：" + proleixing)
			case "D":
				fmt.Println("设计大纲：" + proleixing)
			case "G":
				fmt.Println("设计/修改通知单：" + proleixing)
			case "T":
				fmt.Println("工程图纸：" + proleixing)
			case "J":
				fmt.Println("计算书：" + proleixing)
			}
			//专业
			prozhuanye = SubString(FileNumber, jianhaoloc+1, 1)
			switch prozhuanye {
			case "1":
				fmt.Println("综合：" + prozhuanye)
			case "2":
				fmt.Println("规划：" + prozhuanye)
			case "3":
				fmt.Println("测量：" + prozhuanye)
			case "4":
				fmt.Println("地质：" + prozhuanye)
			case "5":
				fmt.Println("水工：" + prozhuanye)
			case "6":
				fmt.Println("建筑：" + prozhuanye)
			case "7":
				fmt.Println("机电：" + prozhuanye)
			case "8":
				fmt.Println("资环：" + prozhuanye)
			case "9":
				fmt.Println("施工：" + prozhuanye)
			}
			//二级专业代码
		}
	}
	return Suffix, FileNumber, FileName, pronumber, projieduan, proleixing, prozhuanye
}

//	fmt.Println(FileNumber)
//	rs := []rune("SL8888CT-500-88 泵站厂房布置图")
//	lengthname := len([]rune(filenameOnly))
//	if end == 0 { //如果没有空格，则用正则表达式获取编号
//		re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
//		loc := re.FindStringIndex(filenameOnly)
//		fmt.Println("loc=", filenameWithSuffix[loc[0]:loc[1]]) //国
//		fmt.Println("loc[0]=", loc[0])                         //loc[0]= 20
//		if loc != nil {
//			end = loc[0]
//			FileNumber = SubString(filenameOnly, 0, end)
//			fmt.Println("文件编号：", FileNumber)
//			// 查找连续的汉字——改成查找第一个汉字？
//			re = regexp.MustCompile(`[\p{Han}]`)
//			//			lochan := re.FindStringIndex(filenameOnly)
//			//			fmt.Println("第一个汉字位置：", lochan[0])
//			FileName = SubString(filenameOnly, lochan[0], lengthname-lochan[0])
//			fmt.Println("文件名称：", FileName)

//			re, _ = regexp.Compile("[-]")
//			loc = re.FindStringIndex(FileNumber)
//			if loc != nil {
//				end = loc[0]
//			} //7
//			fmt.Println("第一个“-”位置：", loc) //[7 8]

//			loc1 := re.FindAllStringIndex(FileNumber, -1) //[[7 8] [11 12]]
//			if loc1 != nil {
//				fmt.Println("取多个-位置的第一个", loc1[0][0])                        //11
//				fmt.Println("多个的-位置", re.FindAllStringIndex(FileNumber, -1)) //-1表示所有，1表示显示一个，2表示2个
//			}
//			//n换行
//			//查找连续的数字
//			re, _ := regexp.Compile(`[0-9]+`)
//			fmt.Println("连续的数字:", re.FindAllString(filenameOnly, -1))
//			//loc := re.FindStringIndex(filenameOnly)
//			//查找.和-之间的字符
//			re, _ = regexp.Compile(`\..*-`) //  .AT-500-
//			fmt.Println(".和-之间：", re.FindAllString(filenameOnly, -1))
//			//查找第二个字母之前
//			fulleFilename3 := SubString(FileNumber, 0, end-2) //SL888
//			fmt.Println("项目编号：", fulleFilename3)
//			fulleFilename3 = SubString(FileNumber, end-2, 1) //F
//			fulleFilename3 = SubString(FileNumber, end-1, 1) //T
//			fulleFilename3 = SubString(FileNumber, end+1, 1) //5
//			//fmt.Println(fulleFilename3)
//			fulleFilename3 = SubString(FileNumber, end+1, 3) //500
//			fmt.Println("专业中间编号：" + fulleFilename3)
//			fulleFilename3 = SubString(FileNumber, end+5, 3) //500
//			fmt.Println("专业编号：" + fulleFilename3)
//			end = end - 1
//		} else {
//			FileNumber = filenameOnly
//			end = -1
//		}
//	} else { //如果有空格
//		FileNumber = SubString(filenameOnly, 0, end) //这里不能用fullfilename，因为前面赋值后当做了int类型
//	}
//	end = end + 1
//	fulleFilename2 := SubString(filenameOnly, end, lengthname) //这里不能用fullfilename，因为前面赋值后当做了int类型
//	fmt.Println(FileNumber)