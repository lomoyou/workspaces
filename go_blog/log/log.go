package log

//import go_logger "github.com/phachon/go-logger"
//
//var Mylog *go_logger.Logger
//
//func Setup() {
//	Mylog = go_logger.NewLogger()
//
//	//ser console out
//	//Mylog.Detach("console")
//	//console := &go_logger.ConsoleConfig{
//	//	Color: true,
//	//	JsonFormat: false,
//	//	Format: "",
//	//}
//	//
//	//Mylog.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, console)
//
//	//set logfile out
//	Mylog.Detach("file")
//	fileconfig := &go_logger.FileConfig{
//		Filename: "mybloglog.txt",
//		LevelFileName: nil,
//		MaxSize: 1024*1024,
//		MaxLine: 1000*100,
//		DateSlice: "d",
//		JsonFormat: false,
//		Format: "",
//	}
//
//	Mylog.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileconfig)
//}
//
//func Infof(format string, a ...interface{}) {
//	Mylog.Infof(format,a...)
//}
//func Info(format string) {
//	Mylog.Info(format)
//}
//
//
//func Errorf(format string, a ...interface{}) {
//	Mylog.Errorf(format,a...)
//}
//func Error(format string) {
//	Mylog.Error(format)
//}
//
//func Warnf(format string, a ...interface{}) {
//	Mylog.Warningf(format,a...)
//}
//func Warn(format string) {
//	Mylog.Warning(format)
//}
//
//func Debugf(format string, a ...interface{}) {
//	Mylog.Debugf(format,a...)
//}
//
//func Debug(format string) {
//	Mylog.Debug(format)
//}
