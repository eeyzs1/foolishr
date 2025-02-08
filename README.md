# foolishr


package mypkg 

var PublicVar int = 42  // 可被外部包访问 
var privateVar int = 10  // 仅限本包使用 

func PublicFunc() {}     // 可被外部调用 
func privateFunc() {}    // 仅限本包调用 
main包是程序入口，但无法被其他包导入