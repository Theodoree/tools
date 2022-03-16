package _00_699



/*
635. 设计日志存储系统

你将获得多条日志，每条日志都有唯一的 id 和 timestamp，timestamp 是形如 Year:Month:Day:Hour:Minute:Second 的字符串，例如 2017:01:01:23:59:59，所有值域都是零填充的十进制数。

设计一个日志存储系统实现如下功能：

void Put(int id, string timestamp)：给定日志的 id 和 timestamp，将这个日志存入你的存储系统中。



int[] Retrieve(String start, String end, String granularity)：返回在给定时间区间内的所有日志的 id。start 、 end 和 timestamp 的格式相同，granularity 表示考虑的时间级。比如，start = "2017:01:01:23:59:59", end = "2017:01:02:23:59:59", granularity = "Day" 代表区间 2017 年 1 月 1 日到 2017 年 1 月 2 日。





样例 1 ：

put(1, "2017:01:01:23:59:59");
put(2, "2017:01:01:22:59:59");
put(3, "2016:01:01:00:00:00");
retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Year"); // 返回值 [1,2,3]，返回从 2016 年到 2017 年所有的日志。
retrieve("2016:01:01:01:01:01","2017:01:01:23:00:00","Hour"); // 返回值 [1,2], 返回从 2016:01:01:01 到 2017:01:01:23 区间内的日志，日志 3 不在区间内。
*/

type LogSystem struct {
    arr []*log
}

type log struct {
    id        int
    TimeStamp string
}

func Constructor() LogSystem {
    return LogSystem{}
}

func (this *LogSystem) Put(id int, timestamp string) {

    f := &log{
        id:        id,
        TimeStamp: timestamp,
    }
    this.arr = append(this.arr, f)
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {

    var idx int
    switch gra {
    case "Year":
        idx = 4
    case "Month":
        idx = 7
    case "Day":
        idx = 10
    case "Hour":
        idx = 14
    case "Minute":
        idx = 16
    case "Second":
        idx = 19
    }

    var result []int
    for _, v := range this.arr {

        t := v.TimeStamp[:idx]
        if t >= s[:idx] && t <=e[:idx]{
            result = append(result,v.id)
        }

    }

    return result
}
