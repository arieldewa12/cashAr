package controllers

import (
	"bytes"
	"cashAr/models"
	"encoding/csv"
	"fmt"
	"io"
	"logistix/controllers/responses"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BulkUploadController struct {
	beego.Controller
	cashInOrmer  models.CashInOrmer
	cashOutOrmer models.CashOutOrmer
}

func (ths BulkUploadController) Prepare() {
	ormer := orm.NewOrm()
	ths.cashInOrmer = models.NewCashInOrmer(ormer)
	ths.cashOutOrmer = models.NewCashOutOrmer(ormer)
}

// @router / [post]
func (ths *BulkUploadController) BulkUploadCash() error {

	file, _, err := ths.Ctx.Request.FormFile("csv")
	if err != nil {
		return responses.Error(err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return responses.Error(err)
	}

	fileReader := bytes.NewReader(buf.Bytes())

	reader := csv.NewReader(fileReader)
	var bulkUploadCash []models.BulkUploadCash

	isValid, err := validateHeader(reader, []string{"Tanggal Transaksi", "Keterangan", "Cabang", "Jumlah"})
	if err != nil {
		return err
	}
	if isValid {
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				if perr, ok := err.(*csv.ParseError); ok {
					return fmt.Errorf("[BulkEditMappingLocation] Error Parsing CSV Cause: %v", perr)
				}
				return fmt.Errorf("[BulkEditMappingLocation] Error Parsing CSV Cause: %v", err)
			}
			bulkUploadCash = append(bulkUploadCash, models.BulkUploadCash{
				Date:        line[0],
				Description: line[1],
				Branch:      line[2],
				Amount:      line[3],
			})
		}

	} else {
		return responses.Error(err)
	}

	var bulkUploadCashIn []models.CashIn
	var bulkUploadCashOut []models.CashOut

	for _, value := range bulkUploadCash {
		if strings.Contains("CR", value.Amount) {
			amount, err := strconv.ParseFloat(value.Amount, 64)
			if err != nil {
				return err
			}
			bulkUploadCashIn = append(bulkUploadCashIn, models.CashIn{
				Date:        value.Date,
				Description: value.Description,
				Amount:      amount,
			})
		} else {
			amount, err := strconv.ParseFloat(value.Amount, 64)
			if err != nil {
				return err
			}
			bulkUploadCashOut = append(bulkUploadCashOut, models.CashOut{
				TaxInvoiceDate: value.Date,
				Document:       value.Description,
				Amount:         amount,
			})
		}
	}

	err = ths.cashInOrmer.AddCashInFromCSV(bulkUploadCashIn)
	if err != nil {
		return fmt.Errorf("[AddCashInFromCSV] Error Add CashIn from CSV Cause : %v", err)
	}

	err = ths.cashInOrmer.AddCashOutFromCSV(bulkUploadCashOut)
	if err != nil {
		return fmt.Errorf("[AddCashOutFromCSV] Error Add CashOut from CSV/.h. Cause : %v", err)
	}

	return err
}

func validateHeader(reader *csv.Reader, validHeader []string) (bool, error) {
	for i := 0; i > 6; i++ {
		_, err := reader.Read()
		if err != nil { //read unnecessary line
			return false, err
		}
	}

	header, err := reader.Read()
	if err != nil { //read header
		return false, err
	}

	if len(header) != len(validHeader) {
		return false, fmt.Errorf("Invalid CSV Header: %v", header)
	}
	for i, v := range validHeader {
		if v != header[i] {
			return false, fmt.Errorf("Invalid CSV Header: %v", header)
		}
	}

	return true, nil
}
