package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"gitlab.test.com/common/snowflake"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Active struct {
	ID          snowflake.SFID        `gorm:"primaryKey;not null;column:f_id;comment:id" json:"id"`
	DeviceID    string                `gorm:"type:varchar(40);column:f_device_id;index:i_active_device_id;not null;comment:设备标识" json:"device_id"`
	AppID       snowflake.SFID        `gorm:"type:bigint;column:f_app_id;index:i_active_app_id;not null;comment:应用id" json:"app_id"`
	AppVersion  string                `gorm:"type:varchar(20);column:f_app_version;not null;default:'';comment:应用版本" json:"app_version"`
	WebmasterID snowflake.SFID        `gorm:"type:bigint;column:f_webmaster_id;index:i_active_webmaster_id;not null;comment:站长id" json:"webmaster_id"`
	Source      int32                 `gorm:"type:smallint;index:i_active_source;not null;column:f_source;comment:来源(1:H5, 2:安卓)" json:"source"`
	Period      int32                 `gorm:"type:smallint;index:i_active_period;not null;column:f_period;comment:所属时段(0-2点,2-4点...22-24点)" json:"period"`
	Hour        int32                 `gorm:"type:smallint;index:i_active_hour;not null;column:f_hour;comment:小时 如:2022020202" json:"hour"`
	Day         int32                 `gorm:"type:smallint;index:i_active_day;not null;column:f_day;comment:天 如:20220202" json:"day"`
	Week        int32                 `gorm:"type:smallint;index:i_active_week;not null;column:f_week;comment:某一年第几周 如:202223" json:"week"`
	Month       int32                 `gorm:"type:smallint;index:i_active_month;not null;column:f_month;comment:月 如:202202" json:"month"`
	Year        int32                 `gorm:"type:smallint;index:i_active_year;not null;column:f_year;comment:年 如:2022" json:"year"`
	CreatedAt   int64                 `gorm:"index:i_active_created_at;type:bigint;not null;column:f_created_at;comment:启动时间" json:"created_at"`
	UpdatedAt   int64                 `gorm:"type:bigint;not null;column:f_updated_at;comment:更新时间" json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `gorm:"type:bigint;not null;default:0;column:f_deleted_at;comment:删除时间" json:"-"`
}

func (*Active) TableName() string {
	return "t_active"
}

const dsnLayout = "host=%s port=%d  user=%s dbname=%s sslmode=%s  password=%s TimeZone=Asia/Shanghai"

func db() *gorm.DB {
	dsn := fmt.Sprintf(dsnLayout, "192.168.2.26", 5432, "postgres", "analysis-dev", "disable", "xiaoluo.618")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var (
	Period     = []int64{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22}
	Devices    = []string{"3s4dsd4s4ww5w5w5", "35456q54q54fq", "as4a64f6q54qq4fq56", "af4qa654fq65e4f4", "a24a54s65a4sa654f", "a24sa54sa545a45", "s4df65a654w654", "sjlshdjsskksjkjk"}
	Webmasters = []int64{1595665043859443712, 1603223244645601280, 1595954363241598976, 1603671552941887488, 1602193338780356608, 32156464565545, 6656556545545, 3213465456, 212134654654}
	Apps       = []int64{1595665043859454655, 1603223298898601280, 1595954363241521320, 1603671552941865445, 1602193338780302454, 316545654456654, 654654654552252, 320241654655, 105456454}
	Sources    = []int64{1, 2}
)

func TestFakeData(t *testing.T) {
	db := db()
	now := time.Now()
	_, week := now.ISOWeek()
	for i := 0; i < 100; i++ {
		active := &Active{}
		err := faker.FakeData(&active)
		if err != nil {
			t.Fatal(err)
			return
		}

		active.ID = snowflake.Generate()
		active.DeviceID = Devices[rand.Intn(len(Devices))]
		active.WebmasterID = snowflake.SFID(Webmasters[rand.Intn(len(Webmasters))])
		active.AppID = snowflake.SFID(Apps[rand.Intn(len(Apps))])
		active.Source = int32(Sources[rand.Intn(len(Sources))])
		active.Period = GetPeriod(now)
		active.Hour = int32(now.Hour())
		active.Day = int32(now.Day())
		active.Week = int32(week)
		active.Month = int32(now.Month())
		active.Year = int32(now.Year())
		active.AppVersion = "v1.0.1"
		active.CreatedAt = 0
		active.UpdatedAt = 0
		active.DeletedAt = 0
		t.Log(active)
		if err := db.Create(active).Error; err != nil {
			t.Fatal(err)
			return
		}
	}
}

func GetPeriod(t time.Time) int32 {
	hour := t.Hour()
	if (hour & 1) == 1 {
		return int32(hour - 1)
	}
	return int32(hour)
}
