package fm

import "testing"

func TestIsFileNameType(t *testing.T) {

	var (
		imageFile = []string{
			"1.png",
			"2.jpeg",
			"3.PNG",
			"3.JPEG",
		}
		wordFile = []string{
			"wads.doc",
			"woer.docx",
		}
		excelFile = []string{
			"eadfa.xls",
			"eadfa.xlsx",
		}
	)

	for _, f := range imageFile {
		if IsFileNameType(f, TypePng, TypeJpg) != true {
			t.Errorf("%s应该是一张图片，但是结果否", f)
		}
		if IsFileNameType(f, TypeDoc, TypeDocx) != false {
			t.Errorf("%s应该不是word文件，但是结果是", f)
		}
		if IsFileNameType(f, TypeXls, TypeXlsx) != false {
			t.Errorf("%s应该不是excel文件，但是结果是", f)
		}
	}

	for _, f := range wordFile {
		if IsFileNameType(f, TypePng, TypeJpg) != false {
			t.Errorf("%s应该不是一张图片，但是结果是", f)
		}
		if IsFileNameType(f, TypeDoc, TypeDocx) != true {
			t.Errorf("%s应该是word文件，但是结果否", f)
		}
		if IsFileNameType(f, TypeXls, TypeXlsx) != false {
			t.Errorf("%s应该不是excel文件，但是结果是", f)
		}
	}

	for _, f := range excelFile {
		if IsFileNameType(f, TypePng, TypeJpg) != false {
			t.Errorf("%s应该不是一张图片，但是结果是", f)
		}
		if IsFileNameType(f, TypeDoc, TypeDocx) != false {
			t.Errorf("%s应该不是word文件，但是结果是", f)
		}
		if IsFileNameType(f, TypeXls, TypeXlsx) != true {
			t.Errorf("%s应该是excel文件，但是结果否", f)
		}
	}

}
