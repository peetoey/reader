package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "userid.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Open file error: ", err.Error())
	}

	// fmt.Printf("%#v\n", xlFile.Sheets)
	// a := xlFile.Sheets[0]
	// b := a.Rows[1]
	// c := b.Cells[1]
	// fmt.Printf("%#v\n", c.Value)

	var result string
	for rowNum, row := range xlFile.Sheets[0].Rows {
		for cellNum, cell := range row.Cells {
			if rowNum != 0 && cellNum == 1 && cell.String() != "" {
				userID := cell.String()
				bankCode := userID[len(userID)-3:]
				channel := userID[:3]
				result = result + fmt.Sprintf("INSERT [dbo].[MW_TBL_ONLINE_USERID] ([BANKCODE], [CHANNEL], [USERID], [PPB2C_F], [PPINB_F]) VALUES (N'%s', N'%s', N'%s', 0, 1)\n", bankCode, channel, userID)
			}
		}
	}

	// write the whole body at once
	byteResult := []byte(result)
	err = ioutil.WriteFile("result.sql", byteResult, 0644)
	if err != nil {
		panic(err)
	}
}
