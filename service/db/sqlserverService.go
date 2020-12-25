package db

import (
	"database/sql"
	"fmt"
	"github.com/Casper-Mars/dbTool/pojo"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xormplus/xorm"
	"strconv"
	"strings"
)

type MsService struct {
}

// 获取数据库的所有表信息
func (receiver MsService) GetAllTableInfo(username string, password string, ipPort string, dbName string) ([]pojo.TableInfo, error) {
	base, err := receiver.getUrlWithDataBase(username, password, ipPort, dbName)
	if err != nil {
		return nil, err
	}
	mssql, err := xorm.NewEngine(xorm.MSSQL_DRIVER, base)
	if err != nil {
		return nil, err
	}

	tableNames, tableComment, err := receiver.getTableList(mssql)
	if err != nil {
		return nil, err
	}
	info := make([]pojo.TableInfo, len(tableNames))
	for i, k := range tableNames {
		var table pojo.TableInfo
		table.Comment = tableComment[i]
		table.TableName = k
		err := receiver.descTable(mssql, &table)
		if err != nil {
			return nil, err
		}
		table.Comment = tableComment[i]
		info[i] = table
	}

	return info, nil
}

func (receiver MsService) GetAllDBs(username string, password string, ipPort string) ([]string, error) {
	url, err := receiver.getUrl(username, password, ipPort)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mssql", url)
	if err != nil {
		return nil, err
	}
	result, err := db.Query("select name from sys.sysdatabases")
	if err != nil {
		return nil, err
	}
	dbNameList := make([]string, 0)
	var dbName string
	for result.Next() {
		err = result.Scan(&dbName)
		if err != nil {
			return nil, err
		}
		dbNameList = append(dbNameList, dbName)
	}
	return dbNameList, nil
}

// 获取纯粹的数据库链接
func (receiver MsService) getUrl(username string, password string, ipPort string) (string, error) {
	split := strings.Split(ipPort, ":")
	if len(split) != 2 {
		return "", fmt.Errorf("请按照格式输入")
	}
	port, err := strconv.Atoi(split[1])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", split[0], username, password, port), nil
}

// 获取有指定数据库的数据库链接
func (receiver MsService) getUrlWithDataBase(username string, password string, ipPort string, dbName string) (string, error) {
	url, err := receiver.getUrl(username, password, ipPort)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s;database=%s", url, dbName), nil
}

// 获取所有的表名称
func (receiver MsService) getTableList(engine *xorm.Engine) ([]string, []string, error) {
	getAllTableSql := `SELECT DISTINCT
						d.name as name,
						f.value as comment
						FROM
						syscolumns a
						LEFT JOIN systypes b ON a.xusertype= b.xusertype
						INNER JOIN sysobjects d ON a.id= d.id
						AND d.xtype= 'U'
						AND d.name<> 'dtproperties'
						LEFT JOIN syscomments e ON a.cdefault= e.id
						LEFT JOIN sys.extended_properties g ON a.id= g.major_id
						AND a.colid= g.minor_id
						LEFT JOIN sys.extended_properties f ON d.id= f.major_id
						AND f.minor_id= 0 `
	queryString, err := engine.QueryString(getAllTableSql)
	if err != nil {
		return nil, nil, err
	}
	names := make([]string, len(queryString))
	comment := make([]string, len(queryString))
	for i, k := range queryString {
		names[i] = k["name"]
		comment[i] = k["comment"]
	}
	return names, comment, nil
}

// 获取表的结构
func (receiver MsService) descTable(engine *xorm.Engine, table *pojo.TableInfo) error {
	descTableSql := `SELECT 
						a.name name,
						(case when (SELECT count(*) FROM sysobjects  
						WHERE (name in (SELECT name FROM sysindexes  
						WHERE (id = a.id) AND (indid in  
						(SELECT indid FROM sysindexkeys  
						WHERE (id = a.id) AND (colid in  
						(SELECT colid FROM syscolumns WHERE (id = a.id) AND (name = a.name)))))))  
						AND (xtype = 'PK'))>0 then 1 else 0 end) is_primary,
						b.name type,
						COLUMNPROPERTY(a.id,a.name,'PRECISION') as len, 
						(case when a.isnullable=1 then 1 else 0 end) is_empty, 
						isnull(g.[value], ' ') AS comment
						FROM  syscolumns a 
						left join systypes b on a.xtype=b.xusertype  
						inner join sysobjects d on a.id=d.id and d.xtype='U' and d.name<>'dtproperties' 
						left join syscomments e on a.cdefault=e.id  
						left join sys.extended_properties g on a.id=g.major_id AND a.colid=g.minor_id
						left join sys.extended_properties f on d.id=f.class and f.minor_id=0
						where b.name is not null and d.name='%s'
						order by a.id,a.colorder`
	queryInterface, err := engine.QueryInterface(fmt.Sprintf(descTableSql, table.TableName))
	if err != nil {
		return err
	}
	cols := make([]pojo.ColInfo, len(queryInterface))
	for i, k := range queryInterface {
		col := receiver.descCol(k)
		cols[i] = col
	}
	table.Cols = cols
	return nil
}

// 解析出列的数据
func (receiver MsService) descCol(info map[string]interface{}) pojo.ColInfo {
	var col pojo.ColInfo
	{
	}
	col.ColName = info["name"].(string)
	col.Len = int(info["len"].(int64))
	col.ColType = info["type"].(string)
	col.Comment = info["comment"].(string)
	col.IsEmpty = info["is_empty"].(int64) == 1
	col.IsPrimary = info["is_primary"].(int64) == 1
	return col
}
