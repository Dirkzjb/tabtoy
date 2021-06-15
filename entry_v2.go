package main

import (
	"flag"
	"github.com/Dirkzjb/tabtoy/build"
	"github.com/Dirkzjb/tabtoy/v2"
	"github.com/Dirkzjb/tabtoy/v2/i18n"
	"github.com/Dirkzjb/tabtoy/v2/printer"
	"io/ioutil"
	"os"
	"path"
)

// v2特有
var (
	paramProtoVersion = flag.Int("protover", 3, "output .proto file version, 2 or 3")

	paramLuaEnumIntValue = flag.Bool("luaenumintvalue", false, "use int type in lua enum value")
	paramLuaTabHeader    = flag.String("luatabheader", "", "output string to lua tab header")

	paramGenCSharpBinarySerializeCode = flag.Bool("cs_gensercode", true, "generate c# binary serialize code, default is true")
)

func V2Entry() {
	g := printer.NewGlobals()

	if !i18n.SetLanguage(*paramLanguage) {
		log.Infof("language not support: %s", *paramLanguage)
		os.Exit(1)
	}

	g.Version = build.Version

	/*for _, v := range flag.Args() {
		g.InputFileList = append(g.InputFileList, v)
	}*/

	g.ParaMode = *paramPara
	g.CacheDir = *paramCacheDir
	g.UseCache = *paramUseCache
	g.CombineStructName = *paramCombineStructName
	g.ProtoVersion = *paramProtoVersion
	g.LuaEnumIntValue = *paramLuaEnumIntValue
	g.LuaTabHeader = *paramLuaTabHeader
	g.GenCSSerailizeCode = *paramGenCSharpBinarySerializeCode
	g.PackageName = *paramPackageName

	if *paramFileInputDir != "" {
		// 读取当前目录中的所有文件和子目录
		files, err := ioutil.ReadDir(*paramFileInputDir)
		if err != nil {
			panic(err)
		}

		g.InputFileList = append(g.InputFileList, "Globals.xlsx")
		for _, file := range files {
			fileName := file.Name()
			if fileName == "Globals.xlsx" {
				continue
			}

			fileExt := path.Ext(fileName)
			if fileExt != ".xlsx" {
				continue
			}

			pre := fileName[0:2]
			if pre == "~$" {
				continue
			}

			g.InputFileList = append(g.InputFileList, fileName)
		}
	}

	if *paramProtoOut != "" {
		g.AddOutputType("proto", *paramProtoOut)
	}

	if *paramPbtOut != "" {
		g.AddOutputType("pbt", *paramPbtOut)
	}

	if *paramJsonOut != "" {
		g.AddOutputType("json", *paramJsonOut)
	}

	if *paramLuaOut != "" {
		g.AddOutputType("lua", *paramLuaOut)
	}

	if *paramCSharpOut != "" {
		g.AddOutputType("cs", *paramCSharpOut)
	}

	if *paramGoOut != "" {
		g.AddOutputType("go", *paramGoOut)
	}

	if *paramCppOut != "" {
		g.AddOutputType("cpp", *paramCppOut)
	}

	if *paramBinaryOut != "" {
		g.AddOutputType("bin", *paramBinaryOut)
	}

	if *paramTypeOut != "" {
		g.AddOutputType("type", *paramTypeOut)
	}

	if *paramModifyList != "" {
		g.AddOutputType("modlist", *paramModifyList)
	}

	if !v2.Run(g) {
		os.Exit(1)
	}
}
