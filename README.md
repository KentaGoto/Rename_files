# Rename_files 
File name batch replacement tool.  

## Usage  
```
$ go run main.go <DIR> <RENAME_LIST.txt>
```
or for Windows 
```
$ Rename_files.exe <DIR> <RENAME_LIST.txt>
```

## Description  
Rename files in the specified directory based on the tab-delimited rename list.  

**Example:**
```
before1.docx<TAB>after1.docx
before2.docx<TAB>after2.docx
before3.docx<TAB>after3.docx
```

## Requirements
- The rename list is UTF8 (without BOM) encording

## License
MIT

## Author  
Kenta Goto
