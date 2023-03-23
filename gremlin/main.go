package main

import (
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/driver"
	"log"
	"regexp"
)

func main() {
	sql := "4||||insert"
	sqlMatcherOne, err := regexp.Compile("^[0-9]+\\|\\|\\|\\|")
	if err != nil {
		fmt.Println("error")
	}
	// 部分SQL以[0-9]+||||的格式开头，后面紧跟一个正常的SQL
	if ok := sqlMatcherOne.MatchString("4||||insert"); ok {
		sql = sqlMatcherOne.ReplaceAllString(sql, "")
	}
}

func gremlinConfigFunc() func(settings *gremlingo.DriverRemoteConnectionSettings) {
	return func(settings *gremlingo.DriverRemoteConnectionSettings) {
		settings.AuthInfo = &gremlingo.AuthInfo{
			Username: "hetu",
			Password: "Hetu@2022",
		}
		settings.InitialConcurrentConnections = 1
	}
}

// queryProperties 查询traversal的所有properties，并根据id分组后返回
func queryProperties(traversal *gremlingo.GraphTraversal) (error, map[string]map[string]string) {
	resultMap := make(map[string]map[string]string)
	traversalBackup := traversal.Clone()
	resultSet, err := traversalBackup.ValueMap(true).GetResultSet()
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	all, err := resultSet.All()
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	for _, result := range all {
		rawData := result.GetInterface()
		data := rawData.(map[interface{}]interface{})
		if id, ok := data["id"]; ok {
			groupProperties := make(map[string]string)
			for k, v := range data {
				groupProperties[fmt.Sprintf("%v", k)] = fmt.Sprintf("%v", v)
			}
			resultMap[fmt.Sprintf("%v", id)] = groupProperties
		}
	}
	return nil, resultMap
}

func query() {
	// create a graph instance named driverRemoteConnection
	driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection("ws://43.138.47.95:10001/gremlin", gremlinConfigFunc())
	if err != nil {
		log.Fatalf("error create gremlin go connection caused by : %s", err)
		return
	}
	// create a TraversalSource provide gremlin context
	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	var id = "ieg_ngr_ods|ods_ngr_function_login_playerlogin_di"

	traversal := g.V(id).BothE().Path().Unfold()
	err, _ = queryProperties(traversal)

	list, err := traversal.Clone().GetResultSet()

	//g = r.V(vertexId).InE(edgeLabel).Path().Unfold()
	result, _ := list.All()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the result
	for _, r := range result {
		fmt.Printf("%v", r)
		//r.getI
		//data := rawData.(map[interface{}]interface{})
		//for k, v := range data {
		//	fmt.Printf("%v = %v\n", k, v)
		//}
	}

	driverRemoteConnection.Close()
	return
}
