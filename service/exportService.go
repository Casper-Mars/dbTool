package service

type ExportService interface {
	Export(ipPort string, username string, password string, dbNames string, storeLocation string)
}
