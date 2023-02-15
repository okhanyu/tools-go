package filetools

//
//import (
//	"fmt"
//	"github.com/beevik/etree"
//)
//
//func Read() {
//	message := `
//		<Person>
//			<FullName>Grace R. Emlin</FullName>
//			<Company>Example Inc.</Company>
//			<Email where="home">
//				<Addr where='work'>gre@example.com</Addr>
//			</Email>
//			<Email >
//				<Addr>gre@work.com</Addr>
//			</Email>
//			<Group>
//				<Value>Friends</Value>
//				<Value>Squash</Value>
//			</Group>
//			<City>Hanga Roa</City>
//			<State>Easter Island</State>
//		</Person>
//	`
//
//	doc := etree.NewDocument()
//	if err := doc.ReadFromString(message); err != nil {
//		//return returnInfo{GetPresetByTokenErr, "read xml failed."}
//	}
//	root := doc.SelectElement("Person")
//	if root == nil {
//		//return returnInfo{GetPresetByTokenErr, "read xml failed."}
//	}
//	addr := root.FindElements("./Email/Addr")
//	if addr == nil {
//		//return returnInfo{GetPresetByTokenErr, "read xml failed."}
//	}
//
//	for _, res := range addr {
//		fmt.Println(res.SelectAttr("where").Value)
//		//fmt.Println(res.FindElement("./").Text())
//	}
//
//}
