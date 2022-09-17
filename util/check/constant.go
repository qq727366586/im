package check

/**
S	密码中必须存在特殊字符、大小写字母和数字			Ljy#2020
A	对特殊字符、大写字母、小写字母和数字至少存在3种    ljy#2020
B	对特殊字符、大写字母、小写字母和数字至少存在2种	ljy2020
C	对特殊字符、大写字母、小写字母和数字至少存在1种	ljy
D	不存在特殊字符、大小写字母和数字。	/、\
————————————————
*/
const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)
