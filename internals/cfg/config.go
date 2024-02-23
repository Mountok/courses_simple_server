package cfg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct {
	Port   string
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()              //создаем экземпляр нашего ридера для Env
	v.SetEnvPrefix("COURSE_SERV") //префикс, все переменные нашего сервера должны теперь стартовать с COURSE_SERV для того, чтобы не смешиваться с системными
	v.SetDefault("PORT", "8080")  //ставим умолчальные настройки
	v.SetDefault("DBNAME", "course_service_db")  // название базы данных
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "mount7890")
	v.SetDefault("DBHOST", "")
	v.SetDefault("DBPORT", "5432")
	v.AutomaticEnv() //собираем наши переменные с системных

	var cfg Cfg

	err := v.Unmarshal(&cfg) // закидываем переменные в cfg после анмаршалинга
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBConnetcUrl() string { //маленький метод для сборки строки соединения с БД
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
}
