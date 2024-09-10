package model

import (
	"fmt"
	"github.com/tidwall/gjson"
	"go-syntax-train/config"
	"go-syntax-train/model"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 模型映射表
func TestAutoMigrate(t *testing.T) {
	db := config.ConnectTest()
	err := db.AutoMigrate(&model.Message{})
	if err != nil {
		return
	}
}

func TestBatchCreate(t *testing.T) {
	start := time.Now().UnixMilli()
	db := config.ConnectTest()
	res := sync.WaitGroup{}
	for i := 1; i <= 151; i++ {
		res.Add(1)
		go func(i int) {
			message := model.NewMessage("我是消息"+strconv.Itoa(i), 1)
			db.Create(&message)
			res.Add(-1)
		}(i)
	}
	// 阻塞式等待
	res.Wait()
	end := time.Now().UnixMilli()
	log.Println("批量插入数据耗时：", end-start, " ms")
}

func TestTime(t *testing.T) {
	res := time.Now().Add(48 * time.Hour).UnixMilli()
	fmt.Println(res)
}

func TestJson(t *testing.T) {
	response := "{\n    \"code\": 0,\n    \"status_code\": 200,\n    \"message\": \"\",\n    \"detail\": \"\",\n    \"data\": {\n        \"info\": {\n            \"lesson_id\": 1850472,\n            \"org_lesson_id\": 2329946,\n            \"live_class_id\": 0,\n            \"live_lesson_id\": 0,\n            \"vd_room_id\": 0,\n            \"schedule_id\": \"\",\n            \"ai_courseware_id\": \"\",\n            \"ai_release_version\": 0,\n            \"course_type\": 18,\n            \"is_new_step\": 1,\n            \"class_id\": 1905567,\n            \"class_name\": \"预习视频\",\n            \"lesson_name\": \"第1讲-3.0课件\",\n            \"index\": 1,\n            \"index_name\": \"第1讲\",\n            \"course_status\": 2,\n            \"course_class_name\": \"2024-暑期-魔法好课测试\",\n            \"course_type_name\": \"魔法好课\",\n            \"show_time\": \"08月26日 16:00-16:30\",\n            \"duty\": 0,\n            \"duty_name\": \"未考勤\",\n            \"duty_icon\": \"https://ms-public.oss-cn-zhangjiakou.aliyuncs.com/student_app/class/unattendance.png\",\n            \"campus_name\": \"\",\n            \"room_name\": \"\",\n            \"grade_name\": \"\",\n            \"class_hour\": 0,\n            \"subject\": \"语文\",\n            \"subject_name\": \"\",\n            \"subject_name_color\": \"\",\n            \"subject_back_color\": \"\",\n            \"show_class_hour\": \"\",\n            \"grade\": \"二年级\",\n            \"base_class_id\": 123495841\n        },\n        \"list\": [\n            {\n                \"name\": \"学习任务\",\n                \"task\": [\n                    {\n                        \"type\": \"preview\",\n                        \"type_show_name\": \"导学视频\",\n                        \"available\": 1,\n                        \"status\": 0,\n                        \"status_name\": \"去完成\",\n                        \"preview_video\": \"\",\n                        \"preview_video_type\": 4,\n                        \"preview_video_paper\": \"\",\n                        \"preview_paper_id\": \"w_cr3g3ts8c9bneh8af0q0\",\n                        \"icon\": \"https://ms-public.oss-cn-zhangjiakou.aliyuncs.com/student_app/learning_tasks/prepare.png\"\n                    },\n                    {\n                        \"type\": \"classroom\",\n                        \"type_show_name\": \"课堂学习\",\n                        \"available\": 0,\n                        \"status\": 2,\n                        \"status_name\": \"已完成\",\n                        \"tip\": \"课堂学习已结束\",\n                        \"icon\": \"https://ms-public.oss-cn-zhangjiakou.aliyuncs.com/student_app/learning_tasks/learn.png\"\n                    },\n                    {\n                        \"type\": \"homework\",\n                        \"type_show_name\": \"完成加油站\",\n                        \"available\": 1,\n                        \"status\": 0,\n                        \"status_name\": \"去提交\",\n                        \"assignment\": \"w_cog8nrfqt4ic2rps9a60\",\n                        \"assignment_vthree_Id\": \"\",\n                        \"assignment_type\": 1,\n                        \"tip\": \"\",\n                        \"icon\": \"https://ms-public.oss-cn-zhangjiakou.aliyuncs.com/student_app/learning_tasks/practice.png\"\n                    },\n                    {\n                        \"type\": \"exclusive_materials\",\n                        \"type_show_name\": \"专题课\",\n                        \"available\": 1,\n                        \"status\": 1,\n                        \"status_name\": \"去查看\",\n                        \"tip\": \"\",\n                        \"icon\": \"https://ms-public.oss-cn-zhangjiakou.aliyuncs.com/student_app/learning_tasks/strong.png\"\n                    }\n                ]\n            }\n        ]\n    },\n    \"meta\": {\n        \"code\": 0,\n        \"msg\": \"\",\n        \"timestamp\": 0\n    },\n    \"trace_id\": \"dayu_1753a3d55d63ac5ba798c4470908a3be\",\n    \"jaeger_trace_id\": \"47978de5db4ec75\"\n}"
	res := gjson.Get(response, "meta.code").String()
	fmt.Println(res)
}
