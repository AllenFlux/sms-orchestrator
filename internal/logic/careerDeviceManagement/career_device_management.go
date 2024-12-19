package careerDeviceManagement

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"sms_backend/api/v1/career"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"strconv"
)

func New() *sCareerDeviceManagement {
	return &sCareerDeviceManagement{}
}

func init() {
	service.RegisterCareerDeviceManagement(New())
}

const RedisPrefixSmsTraceTask = "RedisPrefixSmsTraceTask"

type sCareerDeviceManagement struct{}

func (s *sCareerDeviceManagement) DeviceRegister(ctx context.Context, req *career.RegisterReq) (res *career.RegisterRes, err error) {
	raw := entity.DeviceList{
		Number:       req.PhoneNumber,
		DeviceNumber: req.DeviceNumber,
		ActiveTime:   req.ActiveTime,
	}

	if count, err := dao.DeviceList.Ctx(ctx).Where("device_number = ?", raw.DeviceNumber).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if count > 0 {
		return nil, errors.New("设备已注册")
	}
	var rowId int64
	if rowId, err = dao.DeviceList.Ctx(ctx).Data(raw).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("注册数据库错误")
	}
	res = &career.RegisterRes{
		ID: rowId,
	}
	return
}

type FileData struct {
	TargetPhoneNumber []string `json:"target_phone_number"`
	Content           []string `json:"content"`
	DeviceNumber      []string `json:"device_number"`
}

func (s *sCareerDeviceManagement) FetchTasks(ctx context.Context, req *career.FetchTaskReq) (res *career.FetchTaskRes, err error) {
	// todo 限制下device获取任务的次数 每太设备最多可以获取1条任务 上一条任务如果没有提交发送报告则不能再次获取
	var device entity.DeviceList
	c := 0
	if err = dao.DeviceList.Ctx(ctx).Where("device_number = ?", req.DeviceNumber).ScanAndCount(&device, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("获取group id失败")
	}
	if c == 0 {
		return nil, errors.New("未查询到device信息")
	}

	if device.GroupId == 0 {
		return nil, errors.New("这台Device目前没有被分配到任何Group")
	}

	// 优先处理对话接口传递的任务
	if c, err := g.Redis().LLen(ctx, req.DeviceNumber); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("从redis中获取对话任务Len错误 请优先修复")
	} else if c > 0 {
		g.Log().Info(ctx, "正在处理对话优先任务 ")
		if messageData, err := g.Redis().LPop(ctx, req.DeviceNumber); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("LPop 从List中获取任务失败 请优先修复")
		} else {
			var subMessageData *sms.SubPostConversationRecordData
			if err = messageData.Scan(&subMessageData); err != nil {
				return nil, errors.New("从redis中获取的数据映射错误 请优先修复")
			}
			res = &career.FetchTaskRes{
				TargetPhoneNumber: subMessageData.TargetPhoneNumber,
				Content:           subMessageData.Content,
				DeviceNumber:      subMessageData.DeviceNumber,
				Interval:          "0",
				TaskId:            subMessageData.TaskID,
				StartAt:           gtime.Now(),
			}
			return res, nil
		}

	} else {
		g.Log().Info(ctx, "无对话任务可以处理 开始处理文件任务")
	}

	var jobs []*entity.SmsMissionReport
	// 任务状态，1-待发送 2-发送中 3-已发送 4-已撤销
	if err = dao.SmsMissionReport.Ctx(ctx).Where("group_id = ?", device.GroupId).Where("task_status = ?", 1).WhereOr("task_status = ?", 2).Limit(1).Scan(&jobs); err != nil {
		return nil, errors.New("查询数据库Mission Report失败")
	}

	if len(jobs) == 0 {
		return nil, errors.New("目前设备无可执行任务List")
	}

	var content []byte
	// 确定需要更新的任务report 条目
	ii := 0

	// Get File
	// todo 循环可以去掉
	for _, job := range jobs {
		g.Log().Infof(ctx, "读取文件中 %s", job.FileName)
		g.Log().Infof(ctx, "fetch task <<<<< filename===%s", job.FileName)
		if v, err := g.Redis().Do(ctx, "GET", job.FileName); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("Redis Get Error ")
		} else {
			content = gconv.Bytes(v)
			g.Log().Infof(ctx, "fetch content by redis <<<<< filename===%s", string(content))
		}
		//ii = i
		if len(content) == 0 {
			g.Log().Info(ctx, "从 缓存 中获取的 content 长度为0")
			// 从文件中重新加载的数据属于新的task list【造成这种情况的原因是程序重启导致的偏差】
			if content, err = ioutil.ReadFile(consts.TaskFilePath + "/" + job.FileName); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("存储的读取文件失败")
			}
			if job.TaskStatus == 2 {
				//   从db中获取已经被发送的任务
				var record []*entity.SmsMissionRecord
				if err = dao.SmsMissionRecord.Ctx(ctx).Where("sub_task_id = ?", job.Id).Scan(record); err != nil {
					g.Log().Error(ctx, err)
					return nil, errors.New("查询SmsMissionRecord错误 程序健壮性错误")
				}
				if len(record) == 0 {
					//	有设备领了这个task的任务没有回报 重复发送任务
				} else {
					//	还原文件content 直接添加，因为device num 写入 mos 是顺序数组 所以只需要根据长度就可以还原未完成的任务
					var tpayload FileData
					err = json.Unmarshal(content, &tpayload)
					if err != nil {
						g.Log().Error(ctx, err)
						return nil, errors.New("还原cache时解析错误")
					}
					length := len(record)
					for i := 0; i < length; i++ {
						tpayload.DeviceNumber = append(tpayload.DeviceNumber, "1")
					}
					content, err = ffjson.Marshal(tpayload)
					if err != nil {
						return nil, errors.New("tpayload 文件格式转换错误")
					}

				}

			}
		} else {
			// go out loop
			// fetch base64
			//g.Log().Infof(ctx, "content ------- : %s", string(content))
			//g.Log().Infof(ctx, "content 000 : %s", string(content[0]))
			//g.Log().Infof(ctx, "content : %v", content)
			g.Log().Info(ctx, "go out loop")
			break
		}

	}

	g.Log().Infof(ctx, "ii = %d", ii)

	if len(content) == 0 {
		return nil, errors.New("最终获取的 content 长度为0 说明无可执行任务块 属于异常错误")
	}
	// Now let's unmarshall the data into `payload`
	var payload FileData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		g.Log().Error(ctx, err)
		g.Log().Infof(ctx, "content : %s", string(content))
		return nil, errors.New("文件解析json错误")
	}

	if len(payload.Content) <= len(payload.DeviceNumber) {
		// 当前文件无可执行任务 需要更新挑选机制 遇到这种状况的原因是有device领取了任务没有即使回报 末尾添加发送数量来限制这种情况
		// 发放的任务可和数据库记录可能会有时间差别 确认下数据库数据
		if jobs[ii].TaskStatus != 3 {
			g.Log().Info(ctx, "数据库数据未即使更新 关闭此条任务窗口")
			if _, err = dao.SmsMissionReport.Ctx(ctx).Data("task_status = ?", 3).Where("id = ?", jobs[ii].Id).Update(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("SmsMissionReport 数据库更新检查错误")
			}
		}
		return nil, errors.New("文件已经被执行完 无任务可以返回")
	}

	i := len(payload.DeviceNumber)
	g.Log().Infof(ctx, "length : %d", i)
	res = &career.FetchTaskRes{
		TargetPhoneNumber: payload.TargetPhoneNumber[i],
		Content:           payload.Content[i],
		DeviceNumber:      req.DeviceNumber,
		Interval:          jobs[ii].IntervalTime,
		TaskId:            jobs[ii].Id,
		StartAt:           jobs[ii].StartTime,
	}

	payload.DeviceNumber = append(payload.DeviceNumber, req.DeviceNumber)

	g.Log().Infof(ctx, "before   ==== > %s", payload)
	g.Log().Infof(ctx, "payload.DeviceNumber = %s", payload.DeviceNumber)
	// 将结构体的格式，转为json字符串的格式。这里用的到库包是"github.com/pquerna/ffjson/ffjson"
	data, err := json.Marshal(&payload)
	if err != nil {
		return nil, errors.New("文件格式转换错误")
	}
	//g.Log().Infof(ctx, "data ---- > %s", string(data))
	// 更新后的任务块条目 将json格式的数据写入mos
	if _, err = g.Redis().Do(ctx, "SET", jobs[ii].FileName, string(data)); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("redis 写入文件错误")
	}
	// 判断剩余短信数量
	sq := jobs[ii].SurplusQuantity - 1
	if sq < 0 {
		return nil, errors.New("剩余短信数量不能为小于0的数 请检查程序逻辑")
	}
	// 判断report的sent status 如果 == 1 需要 更新为 == 2
	// 判断文件中最后一条任务 如果是最后一条可以更新DB报告状态为 更新report的sent status 已完成 -> 已完成的report不会被再次挑选出来

	db := g.DB()
	var tx gdb.TX
	if tx, err = db.Begin(ctx); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("开启事务操作失败")
	}
	g.Log().Info(ctx, "开启事务操作")

	if len(payload.DeviceNumber) == len(payload.Content) {
		if _, err = tx.Model("sms_mission_report").Ctx(ctx).Data(g.Map{"task_status": 3, "surplus_quantity": sq}).Where("id = ?", jobs[ii].Id).Update(); err != nil {
			g.Log().Error(ctx, err)
			if err = tx.Rollback(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("rollback Error")
			}
			return nil, errors.New("更新SmsMissionReport状态错误")
		}
	} else {
		if _, err = tx.Model("sms_mission_report").Ctx(ctx).Data(g.Map{"task_status": 2, "surplus_quantity": sq}).Where("id = ?", jobs[ii].Id).Update(); err != nil {
			g.Log().Error(ctx, err)
			if err = tx.Rollback(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("rollback Error")
			}
			return nil, errors.New("更新SmsMissionReport状态错误 : task_status 2")
		}
	}

	// 添加记录到任务追踪
	var traceTask entity.SmsTraceTask
	c = 0
	if err := dao.SmsTraceTask.Ctx(ctx).Where("target_number = ?", res.TargetPhoneNumber).Where("device_number = ?", res.DeviceNumber).ScanAndCount(&traceTask, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询SmsTraceTask错误")
	}
	if c != 0 {
		_, err = tx.Model("sms_trace_task").Ctx(ctx).Data("task_id = ?", res.TaskId).Update()

		if err != nil {
			g.Log().Error(ctx, err)
			if err = tx.Rollback(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("rollback Error")
			}
			return nil, errors.New("更新记录到任务追踪 错误")
		}
	} else {
		_, err = tx.Model("sms_trace_task").Ctx(ctx).Data(entity.SmsTraceTask{
			TargetNumber: res.TargetPhoneNumber,
			DeviceNumber: res.DeviceNumber,
			TaskId:       res.TaskId,
		}).Insert()

		if err != nil {
			g.Log().Error(ctx, err)
			if err = tx.Rollback(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("rollback Error")
			}
			return nil, errors.New("添加记录到任务追踪 错误")
		}
	}

	// 存储任务追踪到redis
	if _, err = g.Redis().Do(ctx, "SET", MakeNameRedisPrefixSmsTraceTask(res.DeviceNumber, res.TargetPhoneNumber), res.TaskId); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("redis 写入 RedisPrefixSmsTraceTask 文件错误")
	}
	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Commit Error ")
	}
	return
}

func MakeNameRedisPrefixSmsTraceTask(deviceNumber, targetPhoneNumber string) string {
	return RedisPrefixSmsTraceTask + deviceNumber + targetPhoneNumber
}

type SentStatus int

const (
	UnknownSentStatus SentStatus = iota
	SentSuccess
	SentFailure
	EndSentStatus
)

func (s SentStatus) isValid() bool {
	if s > UnknownSentStatus && s < EndSentStatus {
		return true
	}
	return false
}

func (s SentStatus) Value() SentStatus {
	if s == 1 {
		return SentSuccess
	}
	if s == 2 {
		return SentFailure
	}
	return UnknownSentStatus
}

func GenHash(taskName, taskId, targetPhoneNumber, deviceNumber, content, sendStatus, aa, aaId, projectName, projectId, st, SentOrReceive string) string {
	// SentOrReceive : 1是为了char log表 表示此短信是发送的 2表示接收
	hashData := taskName + taskId + targetPhoneNumber + deviceNumber + content + sendStatus + aa + aaId + projectName + projectId + st + SentOrReceive
	hashByte := sha256.Sum256([]byte(hashData))
	rowHash := hex.EncodeToString(hashByte[:])
	return rowHash
}

func (s *sCareerDeviceManagement) ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (res *career.ReportTaskResultRes, err error) {
	// todo 更新device list中 device 的状态
	var mission entity.SmsMissionReport
	// Get TaskName by task ID
	c := 0
	if err = dao.SmsMissionReport.Ctx(ctx).Where("id = ?", req.TaskId).ScanAndCount(&mission, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询SmsMissionReport by id错误")
	}
	if c == 0 {
		return nil, errors.New("非法提交 不存在的任务id")
	}
	// 生成hash串
	// 对整行数据进行hash加研

	rowHash := GenHash(mission.TaskName, strconv.Itoa(req.TaskId), req.TargetPhoneNumber, req.DeviceNumber, req.Content, strconv.Itoa(req.SendStatus), mission.AssociatedAccount, strconv.Itoa(mission.AssociatedAccountId), mission.ProjectName, strconv.Itoa(mission.ProjectId), req.StartTime.String(), "1")
	g.Log().Infof(ctx, "row hash -> %s", rowHash)

	if c, err := dao.SmsMissionRecord.Ctx(ctx).Where("row_hash = ?", rowHash).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询数据库 SmsMissionRecord 错误 row_hash")
	} else if c > 0 {
		return nil, errors.New("数据记录已存在 请勿重复提交")
	}

	data := entity.SmsMissionRecord{
		TaskName:            mission.TaskName,
		SubTaskId:           req.TaskId,
		TargetPhoneNumber:   req.TargetPhoneNumber,
		DeviceNumber:        req.DeviceNumber,
		SmsTopic:            "短信无topic 有只有个文件名比较合理",
		SmsContent:          req.Content,
		SmsStatus:           strconv.Itoa(req.SendStatus),
		AssociatedAccount:   mission.AssociatedAccount,
		AssociatedAccountId: mission.AssociatedAccountId,
		ProjectName:         mission.ProjectName,
		ProjectId:           mission.ProjectId,
		StartTime:           req.StartTime,
		RowHash:             rowHash,
	}

	// 校验提交的SentStatus
	if SentStatus(req.SendStatus).isValid() == false {
		return nil, errors.New("SendStatus验证错误 请提交合法的SendStatus值")
	}
	status := SentStatus(req.SendStatus).Value()

	// 生成DB sms_chart_log中的内容
	charLog := entity.SmsChartLog{
		TaskId:            mission.Id,
		ProjectName:       mission.ProjectName,
		ProjectId:         mission.ProjectId,
		TargetPhoneNumber: req.TargetPhoneNumber,
		DeviceNumber:      req.DeviceNumber,
		SmsTopic:          "todo 短信内容没有topic",
		SmsContent:        req.Content,
		SmsStatus:         req.SendStatus,
		// 这个api接收的数据都属于发送结果 所以在log表中表示的状态都是 1
		SentOrReceive:       1,
		AssociatedAccount:   mission.AssociatedAccount,
		AssociatedAccountId: mission.AssociatedAccountId,
		RowHash:             rowHash,
	}

	var rawId int64
	if rawId, err = dao.SmsMissionRecord.Ctx(ctx).Data(data).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("SmsMissionRecord InsertAndGetId error")
	}
	res = &career.ReportTaskResultRes{ID: rawId}
	db := g.DB()
	var tx gdb.TX
	if tx, err = db.Begin(ctx); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("事务操作开启失败")
	}

	// 更新db中 report的数量信息
	if SentSuccess == status {
		if _, err = tx.Model("sms_mission_report").Ctx(ctx).Data(g.Map{"quantity_sent": mission.QuantitySent + 1, "sent_success_quantity": mission.SentSuccessQuantity + 1}).Where("id = ?", mission.Id).Update(); err != nil {
			g.Log().Error(ctx, err)
			if err = tx.Rollback(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("Rollback Error ")
			}
			return nil, errors.New("发送成功情况下 ： 更新DB SmsMissionReport 错误")
		}
	} else if SentFailure == status {
		if _, err = tx.Model("sms_mission_report").Ctx(ctx).Data(g.Map{"quantity_sent": mission.QuantitySent + 1, "sent_fail_quantity": mission.SentFailQuantity + 1}).Where("id = ?", mission.Id).Update(); err != nil {
			g.Log().Error(ctx, err)
			if err = tx.Rollback(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("Rollback Error ")
			}
			return nil, errors.New("发送失败情况下 ： 更新DB SmsMissionReport 错误")
		}
	} else {
		return nil, errors.New("未知的发送状态 ，请检查验证逻辑是否成功")
	}

	// 更新 sms chart log 表
	if _, err = tx.Model("sms_chart_log").Ctx(ctx).Data(charLog).Insert(); err != nil {
		g.Log().Error(ctx, err)
		if err = tx.Rollback(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("Rollback Error ")
		}
		return nil, errors.New("更新chart log表失败")
	}
	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Commit Error ")
	}
	return
}

func (s *sCareerDeviceManagement) ReportReceiveContent(ctx context.Context, req *career.ReportReceiveContentReq) (res *career.ReportReceiveContentRes, err error) {
	// todo 使用redis记录unread 数量
	var taskId int
	// get task id
	if v, err := g.Redis().Get(ctx, MakeNameRedisPrefixSmsTraceTask(req.DeviceNumber, req.TargetPhoneNumber)); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("从redis中获取数据错误")
	} else {
		taskId = gconv.Int(v)
		g.Log().Info(ctx, "taskId = ", taskId)
		if taskId == 0 {
			t_c := 0
			var t_trace_data entity.SmsTraceTask
			if err = dao.SmsTraceTask.Ctx(ctx).Where("device_number=?", req.DeviceNumber).Where("target_number=?", req.TargetPhoneNumber).ScanAndCount(&t_trace_data, &t_c, false); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New(" SmsTraceTask 数据库查询错误")
			}
			if t_c == 0 {
				return nil, errors.New("t_c == 0 没有查询到追踪任务 请确定是平台发出的任务")
			}
			taskId = t_trace_data.TaskId
		}
	}

	// Get Task Report Info
	var report entity.SmsMissionReport
	c := 0
	if err = dao.SmsMissionReport.Ctx(ctx).Where("id = ?", taskId).ScanAndCount(&report, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询report错误")
	}
	if c == 0 {
		return nil, errors.New("report地址为nil 未映射到相关记录")
	}
	// 生成hash防止重复上报
	rowHash := GenHash(report.TaskName, strconv.Itoa(taskId), req.TargetPhoneNumber, req.DeviceNumber, req.SmsContent, "1", report.AssociatedAccount, strconv.Itoa(report.AssociatedAccountId), report.ProjectName, strconv.Itoa(report.ProjectId), req.StartTime.String(), "2")
	g.Log().Infof(ctx, "Receive Data row hash -> %s", rowHash)
	data := entity.SmsChartLog{
		TaskId:              report.Id,
		ProjectName:         report.ProjectName,
		ProjectId:           report.ProjectId,
		TargetPhoneNumber:   req.TargetPhoneNumber,
		DeviceNumber:        req.DeviceNumber,
		SmsTopic:            "todo 接收的短信没有什么topic",
		SmsContent:          req.SmsContent,
		SmsStatus:           1, //设备都已经将短信接收到了 所以状态一定成功
		AssociatedAccount:   report.AssociatedAccount,
		SentOrReceive:       2, // 2表示接收的信息
		AssociatedAccountId: report.AssociatedAccountId,
		RowHash:             rowHash,
	}
	// 查一下数据库 看是否有相同日志插入
	if c, err := dao.SmsChartLog.Ctx(ctx).Where("row_hash = ?", rowHash).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询数据库SmsChartLog错误")
	} else if c > 0 {
		return nil, errors.New("当前提交的记录有重复 请检查设备是否重复提交")
	}
	// DB Save
	var rowId int64
	if rowId, err = dao.SmsChartLog.Ctx(ctx).Data(data).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("SmsChartLog 数据库插入错误")
	}
	res = &career.ReportReceiveContentRes{ID: rowId}
	return res, nil
}
