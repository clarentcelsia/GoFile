package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	// "github.com/jung-kurt/gofpdf"
	"github.com/gin-gonic/gin"
	"github.com/go-pdf/fpdf"
)

type (
	ConsumerInfo struct {
		No           string
		Area         string
		GoLiveDate   string
		ContractNo   string
		InvoiceNo    string
		ConsumerName string
		DealerCode   string
		DealerName   string
		Network      string
	}

	DealerInfo struct {
		SignerName  string
		SignerTitle string
		Name        string
		Address     string
	}

	RequestGeneratePdfLegalisasi struct {
		StartDate      string           `json:"StartDate"`
		EndDate        string           `json:"EndDate"`
		LOB            string           `json:"LOB"`
		Cabang         string           `json:"Cabang"`
		ListSupplierId []ListSuplierIds `json:"ListSupplierId"`
		Name           string           `json:"-"`
		JobsPosisition string           `json:"-"`
	}

	ListSuplierIds struct {
		SupplierId string `json:"SupplierId"`
	}
)

func fetchData() ([]DealerInfo, []ConsumerInfo, error) {
	var dealers = []DealerInfo{
		{
			SignerName:  "SIGNER NAME",
			SignerTitle: "SIGNER TITLE",
			Name:        "SIGNER'S DEALER NAME",
			Address:     "SIGNER'S DEALER ADDRESS",
		},
		{
			SignerName:  "SIGNER NAME 2",
			SignerTitle: "SIGNER TITLE 2",
			Name:        "SIGNER'S DEALER NAME 2 njdcnscndsNSNCSNCDSKJCNDSCDJCNSD CKDNCDNCSCNSCNDJSC CNSCNSNCDNCDJC DCDNCSC SCKCNL DVVN  FUF JIFFIOFF NFDHFJSDNDNIGFBLHFJFIFFDJ ",
			Address:     "SIGNER'S DEALER ADDRESS 2",
		},
		{
			SignerName:  "SIGNER NAME 3",
			SignerTitle: "SIGNER TITLE 3",
			Name:        "SIGNER'S DEALER NAME 3",
			Address:     "SIGNER'S DEALER ADDRESS 3",
		},
		{
			SignerName:  "SIGNER NAME 4",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 4",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		}, {
			SignerName:  "SIGNER NAME 5",
			SignerTitle: "SIGNER TITLE 5",
			Name:        "SIGNER'S DEALER NAME 5",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		},
		{
			SignerName:  "SIGNER NAME 9",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 9",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		}, {
			SignerName:  "SIGNER NAME 7",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 7",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		}, {
			SignerName:  "SIGNER NAME 8",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 8",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		},
		{
			SignerName:  "SIGNER NAME 6",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 6",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		},
		{
			SignerName:  "SIGNER NAME 10",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 10",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		},
		{
			SignerName:  "SIGNER NAME 11",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 11",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		},
		{
			SignerName:  "SIGNER NAME 4",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 12",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		},
		{
			SignerName:  "SIGNER NAME 4",
			SignerTitle: "SIGNER TITLE 4",
			Name:        "SIGNER'S DEALER NAME 113",
			Address:     "SIGNER'S DEALER ADDRESS 4",
		},
		// {
		// 	SignerName:  "SIGNER NAME 4",
		// 	SignerTitle: "SIGNER TITLE 4",
		// 	Name:        "SIGNER'S DEALER NAME 14",
		// 	Address:     "SIGNER'S DEALER ADDRESS 4",
		// },
		// {
		// 	SignerName:  "SIGNER NAME 4",
		// 	SignerTitle: "SIGNER TITLE 4",
		// 	Name:        "SIGNER'S DEALER NAME 15",
		// 	Address:     "SIGNER'S DEALER ADDRESS 4",
		// }, {
		// 	SignerName:  "SIGNER NAME 4",
		// 	SignerTitle: "SIGNER TITLE 4",
		// 	Name:        "SIGNER'S DEALER NAME 16",
		// 	Address:     "SIGNER'S DEALER ADDRESS 4",
		// }, {
		// 	SignerName:  "SIGNER NAME 4",
		// 	SignerTitle: "SIGNER TITLE 4",
		// 	Name:        "SIGNER'S DEALER NAME 18",
		// 	Address:     "SIGNER'S DEALER ADDRESS 4",
		// },
		// {
		// 	SignerName:  "SIGNER NAME 4",
		// 	SignerTitle: "SIGNER TITLE 4",
		// 	Name:        "SIGNER'S DEALER NAME 117",
		// 	Address:     "SIGNER'S DEALER ADDRESS 4",
		// }, {
		// 	SignerName:  "SIGNER NAME 4",
		// 	SignerTitle: "SIGNER TITLE 4",
		// 	Name:        "SIGNER'S DEALER NAME 119",
		// 	Address:     "SIGNER'S DEALER ADDRESS 4",
		// },
	}

	var consumers = []ConsumerInfo{
		{
			No:           "1",
			Area:         "STSI",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182 - OFFICE JAKARTA",
		},
		{
			No:           "2",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "3",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		},
		{
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182",
		}, {
			No:           "4",
			Area:         "STSI KALIMANTAN TANJUNG BARAT",
			GoLiveDate:   "13/04/24",
			ContractNo:   "CTRTFOEN223J",
			InvoiceNo:    "INVONO223LL20",
			ConsumerName: "HARY TANOE SRIWIJIAYA SJDSANACKSACNCCD",
			DealerCode:   "SB00165",
			DealerName:   "PT RIDWAN KAMIL",
			Network:      "182 - AREA OFFICE JAKARTA",
		},
	}
	return dealers, consumers, nil
}

func getRequest() RequestGeneratePdfLegalisasi {
	return RequestGeneratePdfLegalisasi{
		StartDate: "12 January 2024",
		EndDate:   "20 January 2024",
		LOB:       "MP",
		Cabang:    "Jakarta",
		ListSupplierId: []ListSuplierIds{
			{
				SupplierId: "123",
			},
			{
				SupplierId: "124",
			},
		},
		Name:           "ReqName",
		JobsPosisition: "Managing Directur",
	}
}

func GetFormattedDate() string {
	now := time.Now()
	// day := now.Weekday().String()
	dayNames := map[time.Weekday]string{
		time.Monday:    "Senin",
		time.Tuesday:   "Selasa",
		time.Wednesday: "Rabu",
		time.Thursday:  "Kamis",
		time.Friday:    "Jumat",
		time.Saturday:  "Sabtu",
		time.Sunday:    "Minggu",
	}
	// month := now.Month().String()
	monthNames := map[time.Month]string{
		time.January:   "Januari",
		time.February:  "Februari",
		time.March:     "Maret",
		time.April:     "April",
		time.May:       "Mei",
		time.June:      "Juni",
		time.July:      "Juli",
		time.August:    "Agustus",
		time.September: "September",
		time.October:   "Oktober",
		time.November:  "November",
		time.December:  "Desember",
	}

	return "Pada hari ini " + dayNames[now.Weekday()] + ", tanggal " + now.Format("02") + " bulan " + monthNames[now.Month()] + " tahun " + now.Format("2006")
}

func GetLegalDocPDF(c *gin.Context) {
	pdfbytes, err := GenerateLegalDocument2()
	if err != nil {
		c.JSON(400, err)
		return
	}
	pdfStr := base64.StdEncoding.EncodeToString(pdfbytes)
	c.JSON(200, pdfStr)
}

func getContent(i int, v ConsumerInfo) []string {
	var contents []string
	contents = append(contents, strconv.Itoa(i+1))
	contents = append(contents, v.Area)
	contents = append(contents, v.GoLiveDate)
	contents = append(contents, v.ContractNo)
	contents = append(contents, v.InvoiceNo)
	contents = append(contents, v.ConsumerName)
	contents = append(contents, v.DealerCode)
	contents = append(contents, v.DealerName)
	contents = append(contents, v.Network)
	return contents
}

func generateContents(data []ConsumerInfo) [][]string {
	var contents_ [][]string
	for i, v := range data {
		contents := getContent(i, v)
		contents_ = append(contents_, contents)
	}
	return contents_
}

func isCenterValue(headers, centered_value []string, i int) bool {
	for _, data := range centered_value {
		if headers[i] == data {
			return true
		}
	}
	return false
}

func GenerateLegalDocument2() ([]byte, error) {
	//-----------
	notes1 := `PIC yang menandatangani Berita Acara Legalisasi Dokumen <i>Digital Invoice</i> adalah orang yang telah terdaftar dan ` +
		`ditunjuk secara resmi oleh Dealer berdasarkan Surat Penunjukan PIC <i>Digital Invoice.</i>`
	notes2 := `Untuk Dealer dengan proses tagihan sentralisasi, maka Berita Acara Legalisasi Dokumen <i>Digital Invoice</i> dapat dibuat dalam satu dokumen dengan menjelaskan tagihan yang diberikan untuk outlet Dealer yang bersangkutan. `

	req := getRequest()

	cleanedJobsPositions := strings.Replace(req.JobsPosisition, "Group Akses", "", -1)
	cleanedJobsPositions = strings.Replace(cleanedJobsPositions, "NMC", "", -1)
	cleanedJobsPositions = strings.TrimSpace(cleanedJobsPositions)

	dealers, consumers, err := fetchData()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	//-----------
	const (
		lspLineY   = 2.
		llineBreak = 1.
		lmarginX   = 10.
		lmarginY   = 10.
		lenterY    = 8.
		ltabX      = 5.
		tbheight   = 5.
	)

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(lmarginX, lmarginY, lmarginX)
	pagewidth, pageheight := pdf.GetPageSize()
	_, _, _, mbottom := pdf.GetMargins()
	areaW := pagewidth - (2 * lmarginX)

	pdf.SetFont("Arial", "B", 12)
	_, lineHt := pdf.GetFontSize()

	pdf.SetFooterFunc(func() {
		pdf.SetY(-mbottom)
		pdf.SetFont("Arial", "I", 8.5)
		pdf.SetTextColor(128, 128, 128)
		pdf.WriteAligned(areaW, lineHt, "Paraf", "R")
	})
	pdf.AddPage()

	title := "BERITA ACARA LEGALISASI DOKUMEN"
	ititle := "DIGITAL INVOICE"
	title_width := pdf.GetStringWidth(title) + pdf.GetStringWidth(ititle)
	centerX0 := (areaW - title_width) / 2
	pdf.SetX(centerX0 + 8)

	html := pdf.HTMLBasicNew()
	html.Write(
		(lineHt),
		fmt.Sprintf("%s <b><i>%s<i></b>", title, ititle),
	)
	pdf.Ln(-1)
	pdf.SetY(pdf.GetY() + llineBreak)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())

	formattedDate := GetFormattedDate()
	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(lmarginX, pdf.GetY()+(1.7*lenterY))
	pdf.WriteAligned(areaW, lineHt, formattedDate+", bertempat di _____________,yang bertanda tangan di bawah ini:", "L")

	pdf.SetXY(lmarginX, pdf.GetY()+lenterY-2)
	pdf.Write(lineHt, "Nama")
	pdf.SetX(pdf.GetX() + (3.7 * ltabX))
	Xcolonpos := pdf.GetX()
	pdf.Write(lineHt, " : ")
	pdf.SetX(pdf.GetX())
	pdf.MultiCell(areaW-pdf.GetX(), lineHt, req.Name, "", "L", false)
	Ypos := pdf.GetY()

	pdf.SetXY(lmarginX, Ypos+llineBreak)
	pdf.Write(lineHt, "Jabatan")
	pdf.SetX(Xcolonpos)
	pdf.Write(lineHt, " : ")
	pdf.SetX(pdf.GetX())
	pdf.MultiCell(areaW-pdf.GetX(), lineHt, cleanedJobsPositions, "", "L", false)

	pdf.SetY(pdf.GetY() + (2 * lspLineY))
	displayedDealers := make(map[string]bool)
	for _, dealer := range dealers {
		if !displayedDealers[dealer.Name] {
			pdf.SetXY(lmarginX, pdf.GetY())
			pdf.Write(lineHt, "Nama Dealer")
			pdf.SetX(Xcolonpos)
			pdf.Write(lineHt, " : ")
			pdf.SetX(pdf.GetX())
			pdf.MultiCell(areaW-pdf.GetX(), lineHt, dealer.Name, "", "L", false)

			pdf.SetXY(lmarginX, pdf.GetY()+1)
			pdf.Write(lineHt, "Alamat Dealer")
			pdf.SetX(Xcolonpos)
			pdf.Write(lineHt, " : ")
			pdf.SetX(pdf.GetX())
			pdf.MultiCell(areaW-pdf.GetX(), lineHt, dealer.Address, "", "L", false)
			// pdf.SetY(pdf.GetY() + (1.7 * lspLineY)) ///
			pdf.Ln(1.7 * lspLineY)

			displayedDealers[dealer.Name] = true
		}
	}

	pdf.SetXY(lmarginX, pdf.GetY()+lenterY-2)
	detail := `menerangkan bahwa pengalihan Dokumen <i>Digital Invoice</i> yang dibuat di atas kertas ke dalam media lainnya/secara digital kepada PT Bussan Auto Finance telah dilakukan sesuai dengan aslinya. <br><br>` +
		`Detail masing-masing Dokumen <i>Digital Invoice</i> dilampirkan pada lembar terpisah yang menjadi satu kesatuan dengan Berita Acara Legalisasi Dokumen <i>Digital Invoice</i> ini. <br><br>` +
		`Demikian Berita Acara Legalisasi Dokumen <i>Digital Invoice</i> ini dibuat dengan sebenar-benarnya dan untuk dipergunakan sebagaimana mestinya.`
	html.Write(lineHt, detail)

	pdf.SetXY(lmarginX, pdf.GetY()+(1.5*lenterY))
	if pdf.GetY()+(3.5*lenterY) > pageheight-mbottom {
		pdf.AddPage()
	}
	pdf.Line(lmarginX, pdf.GetY(), pdf.GetX()+50, pdf.GetY())
	pdf.Ln(3.5 * lenterY)
	pdf.Line(lmarginX, pdf.GetY(), pdf.GetX()+40, pdf.GetY())

	pdf.SetFontSize(9)
	pdf.SetXY(lmarginX, pdf.GetY()+(1.7*lenterY))
	pdf.Write(lineHt, "Catatan:")
	pdf.Ln(-1)
	pdf.Write(lineHt, "1. ")
	pdf.SetX(pdf.GetX())
	pdf.SetLeftMargin(pdf.GetX())

	html.Write(lineHt, notes1)
	pdf.Ln(-1)
	pdf.SetX(lmarginX)
	pdf.Write(lineHt, "2. ")
	pdf.SetX(pdf.GetX())
	html.Write(lineHt, notes2)

	pdf.SetLeftMargin(lmarginX)
	//---------------------------------------------
	headers := []string{"No", "Area", "Tanggal Go-live", "No. Kontrak", "No. Invoice (Pelunasan)", "Nama Konsumen", "Kode Dealer", "Nama Dealer", "Network"}
	centered_value := []string{"No", "Tanggal Go-live", "No. Kontrak", "No. Invoice (Pelunasan)", "Kode Dealer"}

	// pdf.SetFooterFunc(func() {
	// 	pdf.SetY(-mbottom)
	// 	pdf.SetFont("Arial", "I", 8.5)
	// 	pdf.SetTextColor(128, 128, 128)
	// 	pdf.WriteAligned(areaW, lineHt, "Paraf", "R")
	// })

	pdf.AddPage()

	pdf.SetMargins(lmarginX, lmarginY, lmarginX)
	pdf.SetFont("Arial", "B", 12)
	pdf.WriteAligned(areaW, lineHt, "LAMPIRAN", "C")
	pdf.Ln(-1)
	attachment_width := pdf.GetStringWidth("Daftar Konsumen yang telah menggunakan Proses Digital Invoice")
	centerX0 = (areaW - attachment_width) / 2
	pdf.SetXY(centerX0+8, pdf.GetY()+lspLineY)
	html.Write(lineHt, "<u><b>Daftar Konsumen yang telah menggunakan Proses <i>Digital Invoice</i></b></u>")

	pdf.SetXY(lmarginX, (pdf.GetY() + lenterY + 1.5))
	pdf.SetFont("Arial", "B", 9)

	widths := calculateColumnWidths(pdf, headers, consumers, areaW)
	header_ht := 0.
	for col := 0; col < len(headers); col++ {
		lines := pdf.SplitText(headers[col], widths[col])
		h := float64(len(lines)) * pdf.PointConvert(12)
		if h > header_ht {
			header_ht = h
		}
	}

	x, y := pdf.GetXY()
	for col := 0; col < len(headers); col++ {
		// pdf.Rect(x, y, widths[col], header_ht-0.2, "")
		pdf.MultiCell(widths[col], header_ht/float64(len(pdf.SplitText(headers[col], widths[col]))), headers[col], "1", "C", false)

		// pdf.MultiCell(widths[col], lineHt+1.4, headers[col], "", "CM", false)
		x += widths[col]
		pdf.SetXY(x, y)
	}

	pdf.SetY(pdf.GetY() + header_ht)
	pdf.SetFont("Arial", "", 7)
	height := 0.
	consumers_list := generateContents(consumers)
	for row := 0; row < len(consumers_list); row++ {
		list := consumers_list[row]
		currX, currY := pdf.GetXY()
		x := currX
		for cols := 0; cols < len(headers); cols++ {
			rowdata := list[cols]

			lines := pdf.SplitText(rowdata, widths[cols])
			h := float64(len(lines)) * pdf.PointConvert(12)
			if h > height {
				height = h
			}

			// lines := pdf.SplitLines([]byte(rowdata), widths[cols])
			// h := float64(len(lines))*(lineHt) + 2*float64(len(lines))
			// if h > height {
			// 	height = h
			// }
		}

		if pdf.GetY()+height > pageheight-mbottom {
			pdf.AddPage()
			currY = pdf.GetY()
		}

		for cols := 0; cols < len(headers); cols++ {
			rowdata := list[cols]
			if isCenterValue(headers, centered_value, cols) {
				pdf.MultiCell(widths[cols], height/float64(len(pdf.SplitText(rowdata, widths[cols]))), rowdata, "1", "CM", false)
			} else {
				pdf.MultiCell(widths[cols], height/float64(len(pdf.SplitText(rowdata, widths[cols]))), rowdata, "1", "L", false)
			}
			x += widths[cols]
			pdf.SetXY(x, currY)
		}
		pdf.SetXY(currX, currY+height)
		height = 0
	}

	// pdf.SetFooterFunc(func() {
	// 	pdf.SetY(-mbottom)
	// 	pdf.SetFont("Arial", "I", 8.5)
	// 	pdf.WriteAligned(areaW, lineHt, "Paraf", "R")
	// })

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to buffer: %v", err)
	}

	return buf.Bytes(), nil

}

func GenerateLegalDocument() ([]byte, error) {
	req := getRequest()
	dealers, consumers, err := fetchData()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	title := "BERITA ACARA LEGALISASI DOKUMEN DIGITAL INVOICE"
	pdf.SetFont("Arial", "B", 12)
	pdf.SetY(20)
	width := pdf.GetStringWidth(title) + 6
	pdf.SetX((210 - width) / 2)
	pdf.CellFormat(width, 10, title, "", 0, "C", false, 0, "")
	pdf.Ln(12)
	pdf.Line(10, 30, 200, 30)

	pdf.SetFont("Arial", "", 10)
	pdf.Ln(10)
	formattedDate := GetFormattedDate()
	pdf.MultiCell(0, 7, formattedDate+", bertempat di _____________,yang bertanda tangan di", "", "L", false)
	pdf.MultiCell(0, 2, "bawah ini:", "", "L", false)
	pdf.Ln(3)

	cleanedJobsPositions := strings.Replace(req.JobsPosisition, "Group Akses", "", -1)
	cleanedJobsPositions = strings.Replace(cleanedJobsPositions, "NMC", "", -1)
	cleanedJobsPositions = strings.TrimSpace(cleanedJobsPositions)

	pdf.Cell(40, 8, "Nama")
	pdf.Cell(0, 8, ": "+req.Name)
	pdf.Ln(5)
	pdf.Cell(40, 8, "Jabatan")
	pdf.Cell(0, 8, ": "+cleanedJobsPositions)
	pdf.Ln(10)

	displayedDealers := make(map[string]bool)

	for _, dealer := range dealers {
		if !displayedDealers[dealer.Name] {
			pdf.Cell(40, 8, "Nama Dealer")
			pdf.Cell(0, 8, ": "+dealer.Name)
			pdf.Ln(-1)
			pdf.Cell(40, 8, "Alamat Dealer")
			// pdf.Cell(5, 8, ":")
			pdf.MultiCell(0, 5, ": "+dealer.Address, "", "L", false)
			pdf.Ln(-1)

			displayedDealers[dealer.Name] = true
		}
	}

	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, "menerangkan bahwa pengalihan Dokumen ")
	pdf.SetFont("Arial", "I", 10)
	pdf.Write(5, "Digital Invoice")
	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, " yang dibuat di atas kertas ke dalam media lainnya/secara digital kepada PT Bussan Auto Finance telah dilakukan sesuai dengan aslinya.")

	pdf.Ln(7)

	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, "Detail masing-masing Dokumen ")
	pdf.SetFont("Arial", "I", 10)
	pdf.Write(5, "Digital Invoice")
	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, " dilampirkan pada lembar terpisah yang menjadi satu kesatuan dengan Berita Acara Legalisasi Dokumen Digital Invoice ini.")

	pdf.Ln(7)

	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, "Demikian Berita Acara Legalisasi Dokumen ")
	pdf.SetFont("Arial", "I", 10)
	pdf.Write(5, "Digital Invoice")
	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, " ini dibuat dengan sebenar-benarnya dan untuk dipergunakan sebagaimana mestinya.")

	pdf.Ln(7)

	pdf.Ln(15)

	pdf.Cell(0, 8, "___________________________________")
	pdf.Ln(5)
	pdf.Ln(15)

	pdf.Cell(0, 8, "_______________________")
	pdf.Ln(15)

	pdf.Cell(0, 8, "Catatan:")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 9)

	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, "1. "+"PIC yang menandatangani Berita Acara Legalisasi Dokumen ")
	pdf.SetFont("Arial", "I", 10)
	pdf.Write(5, "Digital Invoice")
	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, " adalah orang yang telah terdaftar dan")
	pdf.Writef(5, "\t\t\t\t%s", "ditunjuk secara resmi oleh Dealer berdasarkan Surat Penunjukan PIC Digital Invoice.")

	pdf.Ln(7)

	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, "2. "+"Untuk Dealer dengan proses tagihan sentralisasi, maka Berita Acara Legalisasi Dokumen ")
	pdf.SetFont("Arial", "I", 10)
	pdf.Write(5, "Digital Invoice")
	pdf.SetFont("Arial", "", 10)
	pdf.Write(5, " dapat dibuat")
	pdf.Writef(5, "\t\t\t\t%s", "dalam satu dokumen dengan menjelaskan tagihan yang diberikan untuk outlet Dealer yang bersangkutan.")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 8, "LAMPIRAN", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0, 8, "Daftar Konsumen yang telah menggunakan Proses Digital Invoice", "", 1, "C", false, 0, "")
	pdf.Ln(8)
	pdf.Line(10, 30, 200, 30)

	pdf.SetFont("Arial", "B", 9)
	headers := []string{"No", "Area", "Tanggal Go-live", "No. Kontrak", "No. Invoice (Pelunasan)", "Nama Konsumen", "Kode Dealer", "Nama Dealer", "Network"}
	widths := calculateColumnWidths(pdf, headers, consumers, 190.0)

	renderRow(pdf, headers, widths, true)

	pdf.SetFont("Arial", "", 7)
	for _, row := range consumers {
		renderRow(pdf, []string{
			row.No, row.Area, row.GoLiveDate, row.ContractNo, row.InvoiceNo,
			row.ConsumerName, row.DealerCode, row.DealerName, row.Network,
		}, widths, false)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to buffer: %v", err)
	}

	return buf.Bytes(), nil
}

func calculateColumnWidths(pdf *fpdf.Fpdf, headers []string, consumers []ConsumerInfo, maxWidth float64) []float64 {
	widths := make([]float64, len(headers))
	for i := range widths {
		widths[i] = pdf.GetStringWidth(headers[i]) + 6
	}

	for _, consumer := range consumers {
		data := []string{
			consumer.No, consumer.Area, consumer.GoLiveDate, consumer.ContractNo, consumer.InvoiceNo,
			consumer.ConsumerName, consumer.DealerCode, consumer.DealerName, consumer.Network,
		}
		for i, value := range data {
			if i < len(widths) {
				lineWidth := pdf.GetStringWidth(value) + 6
				if lineWidth > widths[i] {
					widths[i] = lineWidth
				}
			}
		}
	}

	totalWidth := 0.0
	for _, w := range widths {
		totalWidth += w
	}

	if totalWidth > maxWidth {
		scale := maxWidth / totalWidth
		for i := range widths {
			widths[i] *= scale
		}
	}

	return widths
}

func renderRow(pdf *fpdf.Fpdf, data []string, widths []float64, isHeader bool) {
	if len(data) != len(widths) {
		return
	}

	maxHeight := 0.0
	for i, value := range data {
		lines := pdf.SplitText(value, widths[i])
		height := float64(len(lines)) * pdf.PointConvert(12)
		if height > maxHeight {
			maxHeight = height
		}
	}

	// if pdf.GetY()+maxHeight > pageheight-mbottom {
	// 	pdf.AddPage()
	// 	currY = pdf.GetY()
	// }

	x, y := pdf.GetXY()
	for i, value := range data {
		pdf.SetXY(x, y)
		if isHeader {
			pdf.SetFont("Arial", "B", 7)
			pdf.MultiCell(widths[i], maxHeight/float64(len(pdf.SplitText(value, widths[i]))), value, "1", "C", false)
		} else {
			pdf.SetFont("Arial", "", 7)
			pdf.MultiCell(widths[i], maxHeight/float64(len(pdf.SplitText(value, widths[i]))), value, "1", "L", false)
		}
		x += widths[i]
	}
}
